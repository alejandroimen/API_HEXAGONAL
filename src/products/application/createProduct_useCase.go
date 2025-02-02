package application

import (
	"fmt"

	"github.com/alejandroimen/API_HEXAGONAL/products/domain/entities"
	"github.com/alejandroimen/API_HEXAGONAL/products/domain/repository"
)

type CreateProduct struct {
	repo repository.ProductRepository
}

func NewCreateProduct(repo repository.ProductRepository) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Run(name string, price float64) error {
	product := entities.Product{Name: name, Price: price}
	if err := cp.repo.Save(product); err != nil {
		return fmt.Errorf("error saving product: %w", err)
	}
	return nil
}
