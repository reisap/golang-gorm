package campaigns_images

type Service interface {
}

type service struct {
	repository Repository
}

func NewCampaignImagesService(repository Repository) Service {
	return &service{repository: repository}
}
