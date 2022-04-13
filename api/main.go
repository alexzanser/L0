package main

import (
	"fmt"
	"log"
	repo "github.com/alexzanser/L0.git/internal/repository"
	sub "github.com/alexzanser/L0.git/internal/subscribe"
	"net/http"
	"github.com/alexzanser/L0.git/pkg/postgres"
	"github.com/go-chi/chi"
	"github.com/alexzanser/L0.git/internal/handlers"
)

const (
	clusterID = "test-cluster"
	clientID = "client-222"
	connstr = "postgres://user_go:8956_go@db:5432/orders"
)

func main() {
	pool, err := postgres.NewPool(connstr)
	if err != nil {
		log.Fatal(fmt.Errorf("Can`t create new pool %v", err))
	}
	repo := repo.NewRepository(pool)
	repo.RestoreCache()
	if err != nil {
		log.Fatal(fmt.Errorf("Can`t restore data from cache %v", err))
	}
	sc, err := sub.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(fmt.Errorf("Error during connection %w", err))
	}
	defer sc.Close()
	sub, err := sub.Subscribe(sc, repo)
	if err != nil {
		log.Fatal(fmt.Errorf("Error during subscription %w", err))
	}
	defer sub.Unsubscribe()

	r := chi.NewRouter()
	ordersHandler := handlers.NewOrders(*repo)
	r.Mount("/", ordersHandler.Routes())

	http.ListenAndServe(":8080", r)
}
