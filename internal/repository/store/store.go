package store

import (
	"fmt"
	"sync"

	"github.com/alexzanser/L0.git/internal/domain"
)

type Store struct {
	sync.Mutex
	Orders	map[string]domain.Order
}

func New() *Store {
	return &Store{}
}

func (s *Store) GetOrder(orderID string) (domain.Order, error) {
	s.Lock()
	defer s.Unlock()

	order, ok := s.Orders[orderID]

	if ok {
		return order, nil
	} else {
		return domain.Order{}, fmt.Errorf("Order with id=%s not found", orderID)
	}
}

func (s *Store) AddOrder() {
	
}