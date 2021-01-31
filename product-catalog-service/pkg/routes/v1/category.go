package v1

import (
	"github.com/easyms/archv1/pc-service/internal/controllers"
	"github.com/easyms/archv1/pc-service/internal/db"
	"github.com/gin-gonic/gin"
)

// ApplyCategoryRoutes adds category routes to the root group
func ApplyCategoryRoutes(r *gin.RouterGroup) {

	database := db.NewConnection("127.0.0.1", "product-catalog")
	controller := controllers.NewCategoryController(database)

	category := r.Group("/categories")
	{
		category.GET("/", controller.GetAllCategories)
		category.GET("/:name", controller.GetCategoryByName)
		category.GET("/:name/subcategories", controller.GetAllSubcategories)
		category.POST("/", controller.CreateCategory)
		category.DELETE("/:id", controller.DeleteCategory)
		category.PUT("/:id", controller.UpdateCategory)
	}
}
