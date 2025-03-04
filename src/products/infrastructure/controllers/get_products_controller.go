package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alejandroimen/API_HEXAGONAL/src/products/application"
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
		log.Printf("Error buscando productos: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Devolver los productos en formato JSON
	log.Printf("Retornando %d productos", len(products))
	ctx.JSON(200, products)
}

// Controlador para Short Polling
func (c *GetProductsController) ShortPoll(ctx *gin.Context) {
	// Simulación de no tener nuevos datos
	ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (c *GetProductsController) LongPoll(ctx *gin.Context) {
	timeout := time.After(30 * time.Second)
	select {
	case <-timeout:
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
	case newData := <-waitForNewData():
		ctx.JSON(http.StatusOK, gin.H{"data": newData})
	}
}

// Simulación de espera para nuevos datos
func waitForNewData() <-chan string {
	newDataChannel := make(chan string)
	go func() {
		time.Sleep(10 * time.Second) // Simula el tiempo hasta que haya nuevos datos
		newDataChannel <- "Datos nuevos disponibles"
	}()
	return newDataChannel
}
