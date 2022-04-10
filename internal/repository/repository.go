package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	Storage
	Cache
}

func NewRepository(pgxPool *pgxpool.Pool) *Repository {
	return &Repository{
		Storage: NewStorage(),
		Cache: NewCache(pgxPool),
	}
}
