package repository

import "github.com/alejandroimen/API_HEXAGONAL/products/domain/entities"

type ProductRepository interface {
	Save(product entities.Product) error
	FindByID(id int) (*entities.Product, error)
	FindAll() ([]entities.Product, error)
	Update(product entities.Product) error
	Delete(id int) error
}
