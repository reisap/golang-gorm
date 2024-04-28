package v1

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	api := router.Group("/api/v1")
	UserRoutes(api)
}
