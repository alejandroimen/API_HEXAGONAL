package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alejandroimen/API_HEXAGONAL/src/products/application"
	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	updateProduct *application.UpdateProduct
}

func NewUpdateProductController(updateProduct *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{updateProduct: updateProduct}
}

func (c *UpdateProductController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid product ID"})
		return
	}

	var request struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if err := c.updateProduct.Run(id, request.Name, request.Price); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "product updated successfully"})
}

// Controlador para Short Polling
func (c *UpdateProductController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (c *UpdateProductController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
