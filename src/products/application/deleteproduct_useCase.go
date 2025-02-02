package application

import (
	"fmt"

	"github.com/alejandroimen/API_HEXAGONAL/products/domain/repository"
)

type DeleteProduct struct {
	repo repository.ProductRepository
}

func NewDeleteProduct(repo repository.ProductRepository) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (dp *DeleteProduct) Run(id int) error {
	// Verificar si el producto existe
	_, err := dp.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("product no encontrado: %w", err)
	}

	// Eliminar el producto del repositorio
	if err := dp.repo.Delete(id); err != nil {
		return fmt.Errorf("error eliminando product: %w", err)
	}

	return nil
}
