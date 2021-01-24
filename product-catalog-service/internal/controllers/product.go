package controllers

import "github.com/easyms/archv1/pc-service/internal/models"

// ProductController enumerates the product controller methods
type ProductController struct{}

//CreateProduct creates a new product
func (p *ProductController) CreateProduct() (product *models.Product) {
	return &models.Product{}
}
