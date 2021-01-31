package v1

import (
	"github.com/easyms/archv1/pc-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

//
func ApplyProductRoutes(r *gin.RouterGroup) {

	controller := new(controllers.ProductController)

	product := r.Group("/product")
	{
		product.GET("/", controller.GetAllProducts)
		product.GET("/:id", controller.GetProductByID)
		product.POST("/:id", controller.CreateProduct)
		product.DELETE("/:id", controller.DeleteProduct)
		product.PUT("/:id", controller.UpdateProduct)
	}
}
