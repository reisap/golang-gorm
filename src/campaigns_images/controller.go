package campaigns_images

type campaignImagesController struct {
	campaignImagesService Service
}

func NewCampaignImagesController(campaignImagesService Service) *campaignImagesController {
	return &campaignImagesController{campaignImagesService: campaignImagesService}
}
