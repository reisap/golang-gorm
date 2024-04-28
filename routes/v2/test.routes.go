package v2

import "github.com/gin-gonic/gin"

func TestRoutes(api *gin.RouterGroup) {
	api.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test api",
		})
	})
}
