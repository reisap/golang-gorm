package v1

import (
	"bwastartup/src/campaigns"
	"bwastartup/src/helper/mysql"
	"github.com/gin-gonic/gin"
)

func CampaignRoutes(api *gin.RouterGroup) {
	repository := campaigns.NewCampaignsRepository(mysql.DB)
	service := campaigns.NewCampaignsService(repository)
	controller := campaigns.NewCampaignsController(service)

	api.GET("/campaign", controller.ListPaging)
}
