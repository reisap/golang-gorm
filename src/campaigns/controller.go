package campaigns

type campaignsController struct {
	campaignService Service
}

func NewCampaignsController(campaignService Service) *campaignsController {
	return &campaignsController{campaignService: campaignService}
}
