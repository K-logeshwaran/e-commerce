package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/e-commerce-app/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{logger: l}
}

func (ph *ProductHandler) GetAll(rw http.ResponseWriter, req *http.Request) {
	var products []db.Product
	base, err := sql.Open("postgres", db.ConStr)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	defer base.Close()
	rows, err := base.Query(`SELECT * FROM "products"`)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product db.Product
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
		products = append(products, product)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(products)

}

func (ph *ProductHandler) GetOne(rw http.ResponseWriter, req *http.Request) {
	prams := mux.Vars(req)
	id, _ := strconv.Atoi(prams["id"])

	qs := `SELECT * FROM "products" WHERE id = $1`
	base, err := sql.Open("postgres", db.ConStr)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	defer base.Close()
	rows, err := base.Query(qs, id)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	defer rows.Close()
	var product db.Product

	if rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte("No records found from the given id"))
		return

	}

	rw.Header().Set("Content-Type", "application/json")
	product.ToJson(rw)

}

func (ph *ProductHandler) Post(rw http.ResponseWriter, req *http.Request) {
	qs := `
	INSERT INTO "products" (name, description, price, category_id, image_url, quantity, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`
	var product db.Product

	base, err := sql.Open("postgres", db.ConStr)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	defer base.Close()
	if err := db.Tostruct(&product, req.Body); err != nil {
		ph.logger.Fatalln(err)
	}
	result, err := base.Exec(qs, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL, product.Quantity, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	ra, _ := result.RowsAffected()

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{"message": fmt.Sprintf("Product added Successfully, %d rows affected", ra)})
}

func (ph *ProductHandler) Delete(rw http.ResponseWriter, req *http.Request) {
	prams := mux.Vars(req)
	id, _ := strconv.Atoi(prams["id"])
	qs := `DELETE  FROM "products" WHERE id = $1`

	base, err := sql.Open("postgres", db.ConStr)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	defer base.Close()
	_, err = base.Exec(qs, id)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	//ra,_:=result.RowsAffected()

	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("Delection Success"))
}

func (ph *ProductHandler) Update(rw http.ResponseWriter, req *http.Request) {
	qs := `
	UPDATE "products" SET name = $1, description = $2, price = $3, category_id = $4, image_url = $5, quantity = $6, updated_at = $7
	WHERE id = $8;
	`
	var product db.Product

	base, err := sql.Open("postgres", db.ConStr)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	defer base.Close()

	if err := db.Tostruct(product, req.Body); err != nil {
		ph.logger.Fatalln(err)
	}

	result, err := base.Exec(qs, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL, product.Quantity, time.Now(), product.ID)
	if err != nil {
		ph.logger.Fatalln(err)
	}
	ra, _ := result.RowsAffected()

	if ra == 0 {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Product Not Found"))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{"message": fmt.Sprintf("Product updated Successfully, %d rows affected", ra)})
}
