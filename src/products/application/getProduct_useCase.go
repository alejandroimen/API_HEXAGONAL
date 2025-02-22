package application

import (
	"github.com/alejandroimen/API_HEXAGONAL/src/products/domain/entities"
	"github.com/alejandroimen/API_HEXAGONAL/src/products/domain/repository"
)

type GetProducts struct {
	repo repository.ProductRepository
}

func NewGetProducts(repo repository.ProductRepository) *GetProducts {
	return &GetProducts{repo: repo}
}

func (gp *GetProducts) Run() ([]entities.Product, error) {
	products, err := gp.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
