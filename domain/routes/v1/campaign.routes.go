package v1

import (
	"bwastartup/domain/campaigns"
	"bwastartup/domain/helper/mysql"
	"github.com/gin-gonic/gin"
)

func CampaignRoutes(api *gin.RouterGroup) {
	repository := campaigns.NewCampaignsRepository(mysql.DB)
	service := campaigns.NewCampaignsService(repository)
	controller := campaigns.NewCampaignsController(service)

	api.GET("/campaign", controller.ListPaging)
}
