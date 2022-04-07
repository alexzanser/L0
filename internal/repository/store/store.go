package store

import (
	"github.com/alexzanser/L0.git/internal/domain"
)


func (s domain.Store) getOrder(orderID string) domain.Order {
	s.Lock()
	defer s.Unlock()


}