package repository

import (
	"github.com/alexzanser/L0.git/internal/domain"
	"github.com/alexzanser/L0.git/internal/repository/store"
	"github.com/sirupsen/logrus"
)

type repo struct {
	*store.Store
	logger logrus.FieldLogger
}

func NewRepository()

type Repository interface {
	Order() ([]domain.Order, error)
}