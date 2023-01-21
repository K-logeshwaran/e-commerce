package handlers

import (
	//"database/sql"
	//"encoding/json"
	"github.com/e-commerce-app/db"
	"log"
	"net/http"
	//"strconv"
)

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{logger: l}
}

var product db.Product

func (ph *ProductHandler) GetAll(rw http.ResponseWriter, req *http.Request) {

}

func (ph *ProductHandler) GetOne(rw http.ResponseWriter, req *http.Request) {
	//id := req.URL.Query().Get("id")

}

func (ph *ProductHandler) Post(rw http.ResponseWriter, req *http.Request) {

}

func (ph *ProductHandler) Delete(rw http.ResponseWriter, req *http.Request) {

}

func (ph *ProductHandler) Update(rw http.ResponseWriter, req *http.Request) {

}
