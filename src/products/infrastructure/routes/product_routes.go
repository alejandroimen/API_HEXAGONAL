package routes

import (
	"github.com/alejandroimen/API_HEXAGONAL/src/products/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine, createProductController *controllers.CreateProductController, getProductsController *controllers.GetProductsController, updateProductController *controllers.UpdateProductController, deleteProductController *controllers.DeleteProductController) {
	// Rutas para productos
	r.POST("/products", createProductController.Handle)
	r.GET("/products", getProductsController.Handle)
	r.PUT("/products/:id", updateProductController.Handle)
	r.DELETE("/products/:id", deleteProductController.Handle) // Ajuste en la ruta para consistencia

	// Nuevas rutas para polling en POST
	r.POST("/products/poll/short", createProductController.ShortPoll)
	r.POST("/products/poll/long", createProductController.LongPoll)

	// Nuevas rutas para polling en DELETE
	r.DELETE("/products/poll/short", deleteProductController.ShortPoll)
	r.DELETE("/products/poll/long", deleteProductController.LongPoll)

	// Nuevas rutas para polling en PUT
	r.PUT("/products/poll/short", updateProductController.ShortPoll)
	r.PUT("/products/poll/long", updateProductController.LongPoll)

	// Nuevas rutas para polling en GET
	r.GET("/products/poll/short", getProductsController.ShortPoll)
	r.GET("/products/poll/long", getProductsController.LongPoll)
}
