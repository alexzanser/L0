package handlers

import (
	"log"
	"net/http"

	"github.com/alexzanser/L0.git/internal/repository"
	"github.com/go-chi/chi"
	// "encoding/json"
)
type Orders struct {
	repo repository.Repository
}

func NewOrders (repo repository.Repository) *Orders {
	return &Orders{repo}
}

func (o *Orders) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/orders/", o.Order)
		r.Post("/orders/", o.getOrder)		
	})

	return r
}

const text = `<!DOCTYPE html>
<html>
    <head>
        <title>Orders</title>
		<link rel="stylesheet" type="text/css" href="gay.css">
    </head>
    <body>
        <p>Please enter the order ID.</p>
		<form action="http://localhost:8080/orders/" method="POST">
    		<input type="text" name="id">
		</form>
    </body>
</html>`

func (o *Orders) Order(w http.ResponseWriter, r*http.Request) {
	log.Printf("handling get task at %s\n", r.URL.Path)
	w.Header().Set("Content-Type", "html")
	w.Write([]byte(text))
}

func (o *Orders) getOrder(w http.ResponseWriter, r*http.Request) {
	log.Printf("handling post task at %s\n", r.URL.Path)

	id := r.PostFormValue("id")
	order, err := o.repo.GetOrder(id)
		if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	// js, err := json.Marshal(order)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	w.Write(order)
}