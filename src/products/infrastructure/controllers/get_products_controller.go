package controllers

import (
	"log"

	"github.com/alejandroimen/API_HEXAGONAL/products/application"
	"github.com/gin-gonic/gin"
)

type GetProductsController struct {
	getProducts *application.GetProducts
}

func NewGetProductsController(getProducts *application.GetProducts) *GetProductsController {
	return &GetProductsController{getProducts: getProducts}
}

func (c *GetProductsController) Handle(ctx *gin.Context) {
	log.Println("Petición para listar todos los productos, recibido")
	// Obtener los productos
	products, err := c.getProducts.Run()
	if err != nil {
		log.Printf("Error buscando products: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Devolver los productos en formato JSON
	log.Printf("Retornando %d products", len(products))
	ctx.JSON(200, products)
}
