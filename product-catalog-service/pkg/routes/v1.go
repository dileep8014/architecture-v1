package routes

import (
	apiV1 "github.com/easyms/archv1/pc-service/pkg/routes/v1"
	"github.com/gin-gonic/gin"
)

// Ping route tests connection to server
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK"})
}

// ApplyRoutes applies v1 routes to the router
func ApplyRoutes(r *gin.RouterGroup) {

	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		apiV1.ApplyProductRoutes(v1)
		apiV1.ApplyCategoryRoutes(v1)

	}
}
