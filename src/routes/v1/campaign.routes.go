package v1

import "github.com/gin-gonic/gin"

func CampaignRoutes(api *gin.RouterGroup) {
	api.GET("/campaign", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Campaign test api",
		})
	})
}
