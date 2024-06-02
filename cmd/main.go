package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"skillfactory_task_31.3.1/internal/api"
	"skillfactory_task_31.3.1/internal/config"
	"skillfactory_task_31.3.1/internal/repository"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Panic().Err(err).Msg("failed to init config")
	}

	fmt.Println(cfg)

	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse log level")
	}

	logger := zerolog.New(os.Stdout).Level(logLevel).With().Timestamp().Logger()

	ctx := context.Background()

	dbLog := logger.With().Str("module", "database").Logger()

	storPsg, err := repository.NewStoragePG(ctx, cfg.PgConnString, dbLog)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to init postgres db")
	}

	storMongo, err := repository.NewStorageMongo(ctx, cfg.MongoConnString, dbLog)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to init mongo db")
	}

	repositories := repository.NewRepository(ctx, storPsg, storMongo, dbLog)

	apiLog := logger.With().Str("module", "api").Logger()
	fmt.Println(cfg.APIPort)

	APIConfig := &api.Opts{
		Addr:       fmt.Sprintf("localhost:%v", cfg.APIPort),
		Log:        apiLog,
		Repository: repositories,
	}

	server := api.NewAPI(APIConfig)

	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to start api server")
		}
	}()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGABRT)

	logger.Info().Msg("awaiting signal")

	sig := <-sigs

	log.Info().Str("signal", sig.String()).Msg("signal received")

	server.Stop(context.Background())
	repositories.StopMongo(ctx)
	repositories.StopPG()

	logger.Info().Msg("exiting")

}
