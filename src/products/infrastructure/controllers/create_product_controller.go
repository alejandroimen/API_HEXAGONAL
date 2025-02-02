package controllers

import (
	"log"

	"github.com/alejandroimen/API_HEXAGONAL/products/application"
	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	createProduct *application.CreateProduct
}

func NewCreateProductController(createProduct *application.CreateProduct) *CreateProductController {
	return &CreateProductController{createProduct: createProduct}
}

func (c *CreateProductController) Handle(ctx *gin.Context) {
	log.Println("Received request to create a product")

	// Estructura para decodificar el JSON de la solicitud
	var request struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	// Decodificar el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decoding request body: %v", err)
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}
	log.Printf("Creating product: Name=%s, Price=%f", request.Name, request.Price)

	// Ejecutar el caso de uso para crear el producto
	if err := c.createProduct.Run(request.Name, request.Price); err != nil {
		log.Printf("Error creating product: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	log.Println("Product created successfully")
	ctx.JSON(201, gin.H{"message": "product created successfully"})
}
