package handlers

import (
	"log"
	"net/http"
)

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{logger: l}
}

func (ph *ProductHandler) GetAll(rw http.ResponseWriter, req *http.Request) {

}

func (ph *ProductHandler) GetOne(rw http.ResponseWriter, req *http.Request) {

}

func (ph *ProductHandler) Post(rw http.ResponseWriter, req *http.Request) {

}

func (ph *ProductHandler) Delete(rw http.ResponseWriter, req *http.Request) {

}

func (ph *ProductHandler) Update(rw http.ResponseWriter, req *http.Request) {

}
