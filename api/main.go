package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)
func main() {
	// logger := log.New()
	// logger.SetLevel(log.Deb)

	dsn := "postgres://user_go:8956_go@localhost:5442/order" +
	"?sslmode=disable"

	pool, err := pkg
}