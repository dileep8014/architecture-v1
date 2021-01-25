package controllers

import (
	"github.com/gin-gonic/gin"
)

// ProductController enumerates the product controller methods
type ProductController struct{}

// CreateProduct creates a new product
func (p *ProductController) CreateProduct(c *gin.Context) error {
	return nil
}

// GetAllProducts retrieves all products
func (p *ProductController) GetAllProducts(c *gin.Context) error {
	return nil
}

// GetProductByID retrieves product by its ID
func (p *ProductController) GetProductByID(c *gin.Context) error {
	return nil
}

// UpdateProduct updates product
func (p *ProductController) UpdateProduct(c *gin.Context) error {
	return nil
}

// DeleteProduct deletes product
func (p *ProductController) DeleteProduct(c *gin.Context) error {
	return nil
}
