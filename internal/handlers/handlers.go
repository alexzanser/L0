package handlers

import (
	"log"
	"net/http"

	"html/template"

	"github.com/alexzanser/L0.git/internal/repository"
	"github.com/go-chi/chi"
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

const getOrderHTML= `<!DOCTYPE html>
<html>
    <head>
        <title>Orders</title>
    </head>
    <body>
        <h3>Please enter the order ID.</h3>
		<form action="http://localhost:8080/orders/" method="POST" id="myForm">
			<input type="text" name="id">
    		<button type="submit" class="btn" name="Submit"><i class="btn"></i>Submit</button>
		</form>
    </body>
</html>`

const getBackHTML =`<!DOCTYPE html>
<html>
    <head>
        <title>Orders</title>
    </head>
    <body>
		<p>
			{{ .Order}}
		</p>
		<a href="http://localhost:8080/orders/">
			<input font="h1" type="button" value="Get back">
		</a>
    </body>
</html>`


func (o *Orders) Order(w http.ResponseWriter, r*http.Request) {
	log.Printf("handling get task at %s\n", r.URL.Path)

	w.Header().Set("Content-Type", "html;charset=utf-8")
	w.Write([]byte(getOrderHTML))
}



func (o *Orders) getOrder(w http.ResponseWriter, r*http.Request) {
	log.Printf("handling post task at %s\n", r.URL.Path)

	id := r.PostFormValue("id")
	order, err := o.repo.GetOrder(id)
		if err != nil {
		order = []byte(err.Error())
	}
	t := template.New("getBack")
	
	item := struct {
       	Order string
    }{
       	Order: string(order),
    }
	t, _ = t.Parse(getBackHTML)
	t.Execute(w, item)
}
