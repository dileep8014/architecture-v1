package main

import (
	"github.com/easyms/archv1/pc-service/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api := router.Group("/api")
	{
		routes.ApplyRoutes(api)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.Run(":8081")
}
