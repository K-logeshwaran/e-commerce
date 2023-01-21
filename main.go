package main

import (
	//"fmt"

	//"database/sql"

	"log"
	"os"

	"github.com/gorilla/mux"
	"net/http"
	// _ "github.com/lib/pq"
	"github.com/e-commerce-app/handlers"
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

	//Routes
	getRouter.HandleFunc("/products", ph.GetAll)
	getRouter.HandleFunc("/products/{id}", ph.GetOne)
	postRouter.HandleFunc("products", ph.Post)
	putRouter.HandleFunc("/products/{id}", ph.Update)
	deleteRouter.HandleFunc("/products/{id}", ph.Delete)

	http.ListenAndServe(":7080", r)
}
