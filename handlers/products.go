package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/views"
)

var productsView = views.NewView("default", "views/products/index.gohtml")
var productsCreateView = views.NewView("default", "views/products/create.gohtml")

// GETProducts GET /products
func (hdl *Handler) GETProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := hdl.db.Query("SELECT id, description, is_active, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	products := make([]*models.Product, 0)
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Description,
			&product.IsActive,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	productsView.Render(w, struct {
		Products []*models.Product
	}{
		products,
	})
}

// GETProductsCreate GET /products/create
func (hdl *Handler) GETProductsCreate(w http.ResponseWriter, r *http.Request) {
	productsCreateView.Render(w, nil)
}

// POSTProductsCreate POST /products/create
func (hdl *Handler) POSTProductsCreate(w http.ResponseWriter, r *http.Request) {
	cost, err := strconv.ParseFloat(r.FormValue("cost"), 64)
	if err != nil {
		log.Fatal(err)
	}
	description := r.FormValue("description")
	isActive, err := strconv.ParseBool(r.FormValue("is_active"))
	if err != nil {
		// TODO do better
		isActive = false
	}
	name := r.FormValue("name")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Fatal(err)
	}

	product := models.Product{
		Cost:        cost,
		Description: description,
		IsActive:    isActive,
		Name:        name,
		Price:       price,
	}

	ctx := context.Background()
	stmt, err := hdl.db.PrepareContext(ctx, "INSERT INTO products (description, name, price, cost, is_active) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.ExecContext(ctx, product.Description, product.Name, product.Price, product.Cost, product.IsActive)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
	return
}
