package handlers

import (
	"log"
	"net/http"
)

type CategoriesHandler struct {
	logger *log.Logger
}

func NewCategoriesHandler(l *log.Logger) *CategoriesHandler {
	return &CategoriesHandler{logger: l}
}

func (ph *CategoriesHandler) GetAll(rw http.ResponseWriter, req *http.Request) {

}

func (ph *CategoriesHandler) GetOne(rw http.ResponseWriter, req *http.Request) {

}
