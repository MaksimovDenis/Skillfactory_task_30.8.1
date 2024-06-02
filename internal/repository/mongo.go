package repository

import (
	"context"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewStorageMongo(ctx context.Context, mongoConnString string, log zerolog.Logger) (*mongo.Client, error) {
	mongoOpts := options.Client().ApplyURI(mongoConnString)
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init mongo")
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongo")
		return nil, err
	}
	return client, nil
}

func (p *Repository) StopMongo(ctx context.Context) {
	p.MongoDB.Disconnect(ctx)
}
