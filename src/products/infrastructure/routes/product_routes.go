package routes

import (
	"github.com/alejandroimen/API_HEXAGONAL/src/products/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine, createProductController *controllers.CreateProductController, getProductsController *controllers.GetProductsController, updateProductController *controllers.UpdateProductController, deleteProductController *controllers.DeleteProductController) {
	// las rutas
	r.POST("/products", createProductController.Handle)
	r.GET("/products", getProductsController.Handle)
	r.PUT("/products/:id", updateProductController.Handle)
	r.DELETE("/products/:id", deleteProductController.Handle)
}
