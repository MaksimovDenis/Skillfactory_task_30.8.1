package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"skillfactory_task_31.3.1/internal/repository"
)

const (
	FileUploadBufferSize       = 512e+6 //512MB for now
	ServerShutdownDefaultDelay = 5 * time.Second
)

type Opts struct {
	Addr       string
	Log        zerolog.Logger
	Repository *repository.Repository
}

type API struct {
	l          zerolog.Logger
	server     *http.Server
	router     *gin.Engine
	repository *repository.Repository
}

func NewAPI(opts *Opts) *API {
	router := gin.Default()

	srvHdl := &http.Server{
		Addr:    opts.Addr,
		Handler: router,
	}

	api := &API{
		l:          opts.Log,
		server:     srvHdl,
		router:     router,
		repository: opts.Repository,
	}

	postgresRouter := router.Group("postgres")
	{
		postgresRouter.POST("/posts", api.createPostPG)
		postgresRouter.GET("/posts", api.getPostsPG)
		postgresRouter.PUT("/posts", api.updatePostPG)
		postgresRouter.DELETE("/posts/:id", api.deletPostByIdPG)
	}
	mongoRouter := router.Group("mongo")
	{
		mongoRouter.POST("/posts", api.createPostMongo)
		mongoRouter.GET("/posts", api.getPostsMongo)
		mongoRouter.PUT("/posts", api.updatePostMongo)
		mongoRouter.DELETE("/posts/:id", api.deletPostByIdMongo)
	}

	return api

}

func (hdl *API) Serve() error {
	if err := hdl.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		hdl.l.Error().Err(err).Msg("failed to start api server")
		return err
	}
	return nil
}

func (hdl *API) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, ServerShutdownDefaultDelay)
	defer cancel()

	if err := hdl.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		hdl.l.Error().Err(err).Msg("failed to stop api server")
	}
}
