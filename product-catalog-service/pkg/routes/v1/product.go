package v1

import (
	"github.com/easyms/archv1/pc-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

//
func ApplyProductRoutes(r *gin.RouterGroup) {

	product := r.Group("/product")
	{
		product.GET("/", controllers.GetAllProducts)
		product.GET("/:id", controllers.GetProductByID)
		product.POST("/:id", controllers.CreateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
		product.PUT("/:id", controllers.UpdateProduct)
	}
}
