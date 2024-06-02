package repository

import (
	"context"

	"github.com/MaksimovDenis/skillfactory_task_30.8.1/internal/models"
	"github.com/MaksimovDenis/skillfactory_task_30.8.1/internal/repository/postsMongo"
	"github.com/MaksimovDenis/skillfactory_task_30.8.1/internal/repository/postsPG"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface задаёт контракт на работу с БД.
type Posts interface {
	Posts() ([]models.Post, error)           // получение всех публикаций
	AddPost(post models.Post) error          // создание новой публикации
	UpdatePost(post models.UpdatePost) error // обновление публикации
	DeletePost(id int) error                 // удаление публикации по ID
}

type Repository struct {
	PostsPG    Posts
	PostsMongo Posts
	Log        zerolog.Logger
	PostgresDB *pgxpool.Pool
	MongoDB    *mongo.Client
}

func NewRepository(ctx context.Context, postgresDB *pgxpool.Pool, mongoDB *mongo.Client, log zerolog.Logger) *Repository {
	return &Repository{
		PostsPG:    postsPG.NewPostPostgres(postgresDB),
		PostsMongo: postsMongo.NewPostMongo(mongoDB),
		Log:        log,
		PostgresDB: postgresDB,
		MongoDB:    mongoDB,
	}
}
