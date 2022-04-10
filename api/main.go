package main

import (
	"fmt"
	"log"
	"time"

	repo "github.com/alexzanser/L0.git/internal/repository"
	sub "github.com/alexzanser/L0.git/internal/subscribe"
	"github.com/alexzanser/L0.git/pkg/postgres"
)

const (
	clusterID = "test-cluster"
	clientID = "client-222"
	connstr = "postgres://user_go:8956_go@localhost:5422/orders"
)

func main() {
	quit := make(chan struct {})

	pool, err := postgres.NewPool(connstr)
	if err != nil {
		log.Fatal(fmt.Errorf("Can`t create new pool %v", err))
	}
	repo := repo.NewRepository(pool)


	sc, err := sub.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(fmt.Errorf("Error during connection %w", err))
	}
	defer sc.Close()
	sub, err := sub.Subscribe(sc, &repo.Storage)
	if err != nil {
		log.Fatal(fmt.Errorf("Error during subscription %w", err))
	}
	defer sub.Unsubscribe()
	
	go func () {
		for {
			for _, order := range repo.Orders {
				fmt.Printf("AAAAAA %s orders: %d\n", string(order), len(repo.Orders))
			}
			time.Sleep(time.Second * 5)
		}
	}()
	<- quit
}
