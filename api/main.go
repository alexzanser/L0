package main

import (
	"fmt"
	"log"
	"time"
	"github.com/alexzanser/L0.git/internal/store"
	sub "github.com/alexzanser/L0.git/internal/subscribe"
)

const (
	clusterID = "test-cluster"
	clientID = "client-222"
)

func main() {
	quit := make(chan struct {})

	sc, err := sub.Connect(clusterID, clientID)
	if err != nil {
		log.Println(fmt.Errorf("Error during connection %w", err))
	}
	defer sc.Close()

	storage := store.New()
	sub, err := sub.Subscribe(sc, storage)
	if err != nil {
		log.Println(fmt.Errorf("Error during subscription %w", err))
	}
	defer sub.Unsubscribe()
	
	
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
