package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Cache struct {
	pool *pgxpool.Pool
}

func NewCache(pgxPool *pgxpool.Pool) Cache {
	return Cache{pool: pgxPool}
}

func (c *Cache) AddOrder() {

}