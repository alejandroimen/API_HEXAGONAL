package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/API_HEXAGONAL/src/products/application"
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

	var request struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decoding request body: %v", err)
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	log.Printf("Creating product: Name=%s, Price=%f", request.Name, request.Price)

	if err := c.createProduct.Run(request.Name, request.Price); err != nil {
		log.Printf("Error creating product: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Product created successfully")
	ctx.JSON(201, gin.H{"message": "product created successfully"})
}

// Controlador para Short Polling
func (c *CreateProductController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (c *CreateProductController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
