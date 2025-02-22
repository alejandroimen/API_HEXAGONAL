package repository

import (
	"database/sql"
	"fmt"

	"github.com/alejandroimen/API_HEXAGONAL/src/products/domain/entities"
)

type ProductRepoMySQL struct {
	db *sql.DB
}

func NewProductRepoMySQL(db *sql.DB) *ProductRepoMySQL {
	return &ProductRepoMySQL{db: db}
}

func (r *ProductRepoMySQL) Save(product entities.Product) error {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	_, err := r.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("error inserting product: %w", err)
	}
	return nil
}

func (r *ProductRepoMySQL) FindByID(id int) (*entities.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var product entities.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return nil, fmt.Errorf("error fetching product: %w", err)
	}
	return &product, nil
}

func (r *ProductRepoMySQL) FindAll() ([]entities.Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepoMySQL) Update(product entities.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}
	return nil
}

func (r *ProductRepoMySQL) Delete(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}
