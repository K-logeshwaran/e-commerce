package handlers

import (
	"log"
	"net/http"
)

type CustomerHandler struct {
	logger *log.Logger
}

func NewCustomerHandler(l *log.Logger) *CustomerHandler {
	return &CustomerHandler{logger: l}
}

func (ph *CustomerHandler) GetAll(rw http.ResponseWriter, req *http.Request) {

}

func (ph *CustomerHandler) GetOne(rw http.ResponseWriter, req *http.Request) {

}

func (ph *CustomerHandler) Post(rw http.ResponseWriter, req *http.Request) {

}

func (ph *CustomerHandler) Delete(rw http.ResponseWriter, req *http.Request) {

}

func (ph *CustomerHandler) Update(rw http.ResponseWriter, req *http.Request) {

}
