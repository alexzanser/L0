package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alexzanser/L0.git/internal/store"
	// sub "github.com/alexzanser/L0.git/internal/subscribe"
	"github.com/alexzanser/L0.git/pkg/postgres"
)

const (
	clusterID = "test-cluster"
	clientID = "client-222"
	connstr = "postgres://user_go:8956_go@db:5422/orders"
)

func main() {
	quit := make(chan struct {})

	_, err := postgres.NewPool(connstr)
	if err != nil {
		log.Fatal(fmt.Errorf("Can`t create new pool %v", err))
	}
	storage := store.New()


	// sc, err := sub.Connect(clusterID, clientID)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("Error during connection %w", err))
	// }
	// defer sc.Close()
	// sub, err := sub.Subscribe(sc, storage)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("Error during subscription %w", err))
	// }
	// defer sub.Unsubscribe()
	
	go func () {
		for {
			for _, order := range storage.Orders {
				fmt.Printf("AAAAAA %s orders: %d\n", string(order), len(storage.Orders))
			}
			time.Sleep(time.Second * 5)
		}
	}()
	<- quit
}
