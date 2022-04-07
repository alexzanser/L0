package repository

import (
	"github.com/alexzanser/L0.git/internal/domain"
	"github.com/alexzanser/L0.git/internal/repository/cache"
	"github.com/alexzanser/L0.git/internal/repository/store"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type repo struct {
	*store.Store
	*cache.Cache
	pool   *pgxpool.Pool
	logger logrus.FieldLogger
}

func NewRepository(pgxPool *pgxpool.Pool, logger logrus.FieldLogger) Repository {
	return &repo {
		Store: 	store.New(),
		Cache: cache.New(pgxPool),
		pool:	pgxPool,
		logger: logger,
	}
}

type Repository interface {
	GetOrder(orderID string) (domain.Order, error)
}
