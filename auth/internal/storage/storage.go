package storage

import (
	"auth/internal/config"
	"auth/internal/storage/repository"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func NewStorage(ctx context.Context, config config.DBConfig) (*repository.Storage, error) {
	pool, err := pgxpool.Connect(ctx, config.GetDSN())
	if err != nil {
		log.Fatal("failed to connect to db: ", err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("failed to ping db: ", err)
	}

	return &repository.Storage{
		Pool: pool,
	}, nil
}
