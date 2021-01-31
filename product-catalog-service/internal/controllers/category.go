package controllers

import (
	"fmt"

	"github.com/easyms/archv1/pc-service/internal/db"
	"github.com/easyms/archv1/pc-service/internal/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// CategoryController enumerates the category controller methods
type CategoryController struct {
	model *db.Connection
}

// NewController create a new category controller
func NewCategoryController(m *db.Connection) *CategoryController {

	return &CategoryController{
		model: m,
	}
}

// CreateCategory creates an new category
func (cat *CategoryController) CreateCategory(c *gin.Context) {

	var category models.Category
	c.BindJSON(&category)
	c.JSON(200, category)
}

// GetAllCategories retrieve all categories
func (cat *CategoryController) GetAllCategories(c *gin.Context) {

}

// GetCategoryByName retreive category by name
func (cat *CategoryController) GetCategoryByName(c *gin.Context) {

	name := c.Param("name")
	collection := cat.model.Use("category")

	result := models.Category{}

	err := collection.Find(bson.M{"name": name}).One(&result)

	if err != nil {
		c.JSON(404, fmt.Sprintf("No category found for name: %s", name))
	}

	c.JSON(200, result)
}

// GetAllSubcategories retrevies all subcategories of a category
func (cat *CategoryController) GetAllSubcategories(c *gin.Context) {
}

// UpdateCategory update category
func (cat *CategoryController) UpdateCategory(c *gin.Context) {

}

// DeleteCategory delete category
func (cat *CategoryController) DeleteCategory(c *gin.Context) {

}
