package repository

import (
	"fmt"
)

type Storage struct {
	Orders map[string][]byte
}

func NewStorage() Storage {
	return Storage{
		Orders : make(map[string][]byte, 0),
	}
}

func (s *Storage) Store(orderID string, data []byte) error {
	if _, ok := s.Orders[orderID]; ok {
		return fmt.Errorf("Order already exists")
	}
	s.Orders[orderID] = data
	return nil
}

func (s *Storage) GetOrder(orderID string) ([]byte, error) {
	if val, ok := s.Orders[orderID]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("Order with id %s not found", orderID)
}
