package application

import (
	"fmt"

	_ "github.com/alejandroimen/API_HEXAGONAL/src/products/domain/entities"
	"github.com/alejandroimen/API_HEXAGONAL/src/products/domain/repository"
)

type UpdateProduct struct {
	repo repository.ProductRepository
}

func NewUpdateProduct(repo repository.ProductRepository) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (up *UpdateProduct) Run(id int, name string, price float64) error {
	// Verificar si el producto existe
	product, err := up.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("product no encontrado: %w", err)
	}

	// Actualizar los campos del producto
	product.Name = name
	product.Price = price

	// Guardar los cambios en el repositorio
	if err := up.repo.Update(*product); err != nil {
		return fmt.Errorf("error actualizando product: %w", err)
	}

	return nil
}
