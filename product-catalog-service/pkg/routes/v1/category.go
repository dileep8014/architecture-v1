package v1

import (
	"github.com/easyms/archv1/pc-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

// ApplyCategoryRoutes adds category routes to the root group
func ApplyCategoryRoutes(r *gin.RouterGroup) {

	category := r.Group("/categories")
	{
		category.GET("/", controllers.GetAllProducts)
		category.GET("/:name", controllers.GetCategoryByName)
		category.GET("/:name/subcategories", controllers.GetAllSubcategories)
		category.POST("/", controllers.CreateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)
		category.PUT("/:id", controllers.UpdateCategory)
	}

}
