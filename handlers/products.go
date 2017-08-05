package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/views"
	"github.com/nikovacevic/money"
)

var createProductView = views.NewView("default", "views/products/create.gohtml")
var productsView = views.NewView("default", "views/products/index.gohtml")
var updateProductView = views.NewView("default", "views/products/update.gohtml")
var viewProductView = views.NewView("default", "views/products/view.gohtml")

// Products serves a listing of all products
func (hdl *Handler) Products(w http.ResponseWriter, r *http.Request) {
	rows, err := hdl.db.Query("SELECT id, description, is_active, name, price FROM products WHERE is_active;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	activeProducts := make([]*models.Product, 0)
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
		activeProducts = append(activeProducts, product)
	}

	rows, err = hdl.db.Query("SELECT id, description, is_active, name, price FROM products WHERE NOT is_active;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	inactiveProducts := make([]*models.Product, 0)
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
		inactiveProducts = append(inactiveProducts, product)
	}

	productsView.Render(w, struct {
		ActiveProducts   []*models.Product
		InactiveProducts []*models.Product
	}{
		activeProducts,
		inactiveProducts,
	})
}

// ViewProduct serves details of the identified product
func (hdl *Handler) ViewProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	product, err := productService.ByID(id)
	if err != nil {
		log.Fatal(err)
	}

	viewProductView.Render(w, struct {
		Product *models.Product
	}{
		product,
	})
}

// CreateProduct serves a form for creating a new product
func (hdl *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	createProductView.Render(w, nil)
}

// PostProduct attempts to create a product from the given request data
func (hdl *Handler) PostProduct(w http.ResponseWriter, r *http.Request) {
	c, err := strconv.ParseFloat(r.FormValue("cost"), 64)
	if err != nil {
		log.Fatal(err)
	}
	cost := money.ToUSD(c)
	description := r.FormValue("description")
	isActive, err := strconv.ParseBool(r.FormValue("is_active"))
	if err != nil {
		// TODO do better
		isActive = false
	}
	name := r.FormValue("name")
	p, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Fatal(err)
	}
	price := money.ToUSD(p)

	product := &models.Product{
		Cost:        cost,
		Description: description,
		IsActive:    isActive,
		Name:        name,
		Price:       price,
	}

	_, err = productService.Create(product)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
	return
}

// UpdateProduct serves a form for editing a product
func (hdl *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	product, err := productService.ByID(id)
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Product *models.Product
	}{
		product,
	}
	updateProductView.Render(w, data)
}

// PatchProduct attempts to update a product from the given request data
func (hdl *Handler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	product, err := productService.ByID(id)
	if err != nil {
		log.Fatal(err)
	}

	cost, err := strconv.ParseFloat(r.FormValue("cost"), 64)
	if err != nil {
		log.Fatal(err)
	}
	product.Cost = money.ToUSD(cost)
	product.Description = r.FormValue("description")
	product.IsActive, err = strconv.ParseBool(r.FormValue("is_active"))
	if err != nil {
		// TODO do better
		product.IsActive = false
	}
	product.Name = r.FormValue("name")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Fatal(err)
	}
	product.Price = money.ToUSD(price)

	_, err = productService.Update(product)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, fmt.Sprintf("/products/%v", product.ID), http.StatusSeeOther)
	return
}
