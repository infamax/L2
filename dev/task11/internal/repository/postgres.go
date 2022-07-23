package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresDB struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, connString string) (*postgresDB, error) {
	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return &postgresDB{
		pool: pool,
	}, nil
}
