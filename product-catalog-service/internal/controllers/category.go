package controllers

import (
	"github.com/gin-gonic/gin"
)

// CategoryController enumerates the category controller methods
type CategoryController struct{}

// CreateCategory creates an new category
func (cat *CategoryController) CreateCategory(c *gin.Context) error {
	return nil
}

// GetAllCategories retrieve all categories
func (cat *CategoryController) GetAllCategories(c *gin.Context) error {
	return nil
}

// GetCategoryByName retreive category by name
func (cat *CategoryController) GetCategoryByName(c *gin.Context) error {
	return nil
}

// GetAllSubcategories retrevies all subcategories of a category
func (cat *CategoryController) GetAllSubcategories(c *gin.Context) error {
	return nil
}

// UpdateCategory update category
func (cat *CategoryController) UpdateCategory(c *gin.Context) error {
	return nil
}

// DeleteCategory delete category
func (cat *CategoryController) DeleteCategory(c *gin.Context) error {
	return nil
}
