package store

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/alexzanser/L0.git/internal/domain"
)

type Storage struct {
	pool   *pgxpool.Pool
	Orders map[string][]byte
}

func New() *Storage {
	return &Storage{
		Orders : make(map[string][]byte, 0),
	}
}

func (s *Storage) Save(data []byte) error {
	order := &order.Order{}
	err := json.Unmarshal(data, order)
	if err != nil {
		return fmt.Errorf("Cant`t unmarshal to json (invalid data)%v", err)
	}

	if _, ok := s.Orders[order.OrderUid]; ok {
		return fmt.Errorf("Order already exists %v", err)
	}

	s.Orders[order.OrderUid] = data
	return nil
}

func (s *Storage) GetOrder(orderID string) []byte {
	if val, ok := s.Orders[orderID]; ok {
		return val
	}
	return nil
}

type Saver interface {
	Save(data []byte)
}
