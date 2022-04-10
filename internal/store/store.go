package store

import "encoding/json"

type Order struct {
	ID	string `json:"order_uid"`
	Info interface{}
}

type Storage struct {
	Orders []Order
}

func New() *Storage {
	return &Storage{
		Orders : make([]Order, 0),
	}
}

func (s *Storage) Save(data []byte) {
	order := &Order{}
	json.Unmarshal(data, order)
	order.Info = data
	s.Orders = append(s.Orders, *order)
}

type Saver interface {
	Save(data []byte)
}
