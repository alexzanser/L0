package domain

import "sync"

type Store struct {
	sync.Mutex
	Orders	[]Order
}

