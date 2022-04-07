package cache

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Cache struct {
	pool *pgxpool.Pool
}

func New(pgxPool *pgxpool.Pool) *Cache {
	return &Cache{pool: pgxPool}
}