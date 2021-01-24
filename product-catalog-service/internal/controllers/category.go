package controllers

import "github.com/easyms/archv1/pc-service/internal/models"

// CategoryController enumerates the category controller methods
type CategoryController struct{}

// CreateCategory creates an new category
func (c *CategoryController) CreateCategory() (category *models.Category) {
	return &models.Category{}
}
