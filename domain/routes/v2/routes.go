package v2

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	api := router.Group("/api/v2")
	TestRoutes(api)
}
