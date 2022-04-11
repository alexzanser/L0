package handlers

import (
	"github.com/alexzanser/L0.git/internal/repository"
	"github.com/go-chi/chi"
	"net/http"
	"log"
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
		r.Get("/orders/{id}", o.getOrder)
	})

	return r
}

const text = `<!DOCTYPE html>
<html>
    <head>
        <title>Example</title>
		<link rel="stylesheet" type="text/css" href="gay.css">
    </head>
    <body>
        <p>This is an example of a simple HTML page with one paragraph.</p>
		<button onclick="">
		 I am gay
		</button>
    </body>
</html>`

func (o *Orders) getOrder(w http.ResponseWriter, r*http.Request) {
	log.Printf("handling get task at %s\n", r.URL.Path)

	// id := chi.URLParam(r, "id")
	// order, err := o.repo.GetOrder(id)
	// 	if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// }

	// js, err := json.Marshal(order)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	w.Header().Set("Content-Type", "html")
	w.Write([]byte(text))
	// w.Write(js)

}
