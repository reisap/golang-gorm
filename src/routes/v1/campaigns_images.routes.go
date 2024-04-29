package v1

import (
	"bwastartup/src/campaigns_images"
	"bwastartup/src/helper/mysql"

	"github.com/gin-gonic/gin"
)

func CampaignsImagesRoutes(api *gin.RouterGroup) {
	repository := campaigns_images.NewCampaignImagesRepository(mysql.DB)
	service := campaigns_images.NewCampaignImagesService(repository)
	controller := campaigns_images.NewCampaignImagesController(service)

	api.GET("/campaign/images", controller.ListPaging)
}
