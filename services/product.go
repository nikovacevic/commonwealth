package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/nikovacevic/commonwealth/models"
)

// ProductService is a Service that operates on Products.
type ProductService struct {
	db *sql.DB
}

// ProductNotFound is an error to be used when a Product is not found.
type ProductNotFound error

// NewProduct creates and returns a new instance of the ProductService.
func NewProduct(db *sql.DB) *ProductService {
	return &ProductService{db}
}

// ByID queries for a Product by ID, returning the Product if found or a
// ProductNotFound error if the Product is not found.
func (ps *ProductService) ByID(id uint64) (*models.Product, error) {
	product := &models.Product{ID: id}
	ctx := context.Background()
	row := ps.db.QueryRowContext(
		ctx,
		`SELECT   p.cost,
			  p.description,
			  p.is_active,
			  p.name,
			  p.price
		FROM      products AS p
		WHERE     p.id = $1;`,
		id,
	)
	err := row.Scan(
		&(product.Cost),
		&(product.Description),
		&(product.IsActive),
		&(product.Name),
		&(product.Price),
	)
	if err != nil {
		return nil, *new(ProductNotFound)
	}
	return product, nil
}

// Create attempts to INSERT the given Product into the database.
func (ps *ProductService) Create(product *models.Product) (*models.Product, error) {
	// TODO validate

	ctx := context.Background()
	stmt, err := ps.db.PrepareContext(
		ctx,
		`INSERT INTO products (
			cost,
			description,
			is_active,
			name,
			price
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.ExecContext(
		ctx,
		product.Cost,
		product.Description,
		product.IsActive,
		product.Name,
		product.Price,
	)
	if err != nil {
		log.Fatal(err)
	}

	// TODO scan INSERTED ID into product.ID

	return product, nil
}

// Update attempts to UPDATE the given Product record by ID.
func (ps *ProductService) Update(product *models.Product) (*models.Product, error) {
	// TODO validate

	ctx := context.Background()
	stmt, err := ps.db.PrepareContext(
		ctx,
		`UPDATE products
		 SET    cost = $1,
			description = $2,
			is_active = $3,
			name = $4,
			price = $5
		 WHERE   id = $6`,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.ExecContext(
		ctx,
		product.Cost,
		product.Description,
		product.IsActive,
		product.Name,
		product.Price,
		product.ID,
	)
	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}
