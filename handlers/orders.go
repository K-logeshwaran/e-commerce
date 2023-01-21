package handlers

import (
	"log"
	"net/http"
)

type OrdersHandler struct {
	logger *log.Logger
}

func NewOrdersHandler(l *log.Logger) *OrdersHandler {
	return &OrdersHandler{logger: l}
}

func (ph *OrdersHandler) GetAll(rw http.ResponseWriter, req *http.Request) {

}

func (ph *OrdersHandler) GetOne(rw http.ResponseWriter, req *http.Request) {

}

func (ph *OrdersHandler) Post(rw http.ResponseWriter, req *http.Request) {

}

func (ph *OrdersHandler) Delete(rw http.ResponseWriter, req *http.Request) {

}
