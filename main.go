package main

import (
	"log"
	"os"

	"github.com/e-commerce-app/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	logger := log.New(os.Stdout, "E-com-API", log.LstdFlags)
	// Routers
	getRouter := r.Methods(http.MethodGet).Subrouter()
	postRouter := r.Methods(http.MethodPost).Subrouter()
	putRouter := r.Methods(http.MethodPut).Subrouter()
	deleteRouter := r.Methods(http.MethodDelete).Subrouter()

	//handlers
	ph := handlers.NewProductHandler(logger)
	cushand := handlers.NewCustomerHandler(logger)
	ordhand := handlers.NewProductHandler(logger)

	//Routes
	getRouter.HandleFunc("/products", ph.GetAll)
	getRouter.HandleFunc("/products/{id}", ph.GetOne)
	postRouter.HandleFunc("products", ph.Post)
	putRouter.HandleFunc("/products/{id}", ph.Update)
	deleteRouter.HandleFunc("/products/{id}", ph.Delete)

	getRouter.HandleFunc("/customers", cushand.GetAll)
	getRouter.HandleFunc("/customers/{id}", cushand.GetOne)
	postRouter.HandleFunc("customers", cushand.Post)
	putRouter.HandleFunc("/customers/{id}", cushand.Update)
	deleteRouter.HandleFunc("/customers/{id}", cushand.Delete)

	getRouter.HandleFunc("/orders", ordhand.GetAll)
	getRouter.HandleFunc("/orders/{id}", ordhand.GetOne)
	postRouter.HandleFunc("orders", ordhand.Post)
	deleteRouter.HandleFunc("/orders/{id}", ordhand.Delete)

	http.ListenAndServe(":7080", r)
}
