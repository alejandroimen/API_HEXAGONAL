package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alejandroimen/API_HEXAGONAL/src/products/application"
	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	deleteProduct *application.DeleteProduct
}

func NewDeleteProductController(deleteProduct *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{deleteProduct: deleteProduct}
}

func (c *DeleteProductController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid product ID"})
		return
	}

	if err := c.deleteProduct.Run(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "product deleted successfully"})
}

// Controlador para Short Polling
func (c *DeleteProductController) ShortPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (c *DeleteProductController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}
