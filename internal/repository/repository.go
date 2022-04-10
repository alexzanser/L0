package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

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

func (r *Repository) Save(orderID string, data[]byte) error {
	err := r.Store(orderID, data)
	if err != nil {
		return err
	}
	err = r.CacheOrder(orderID, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RestoreCache() error {
	var s map[string][]byte
	s, err := r.Cache.GetOrders()
	if err != nil {
		return (err)
	}
	r.Storage.Orders = s

	return nil
}