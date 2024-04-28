package v1

import (
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, limiter *limiter.Limiter) {
	api := router.Group("/api/v1")
	api.Use(tollbooth_gin.LimitHandler(limiter))
	UserRoutes(api)
	CampaignRoutes(api)
}
