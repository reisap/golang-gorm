package campaigns

type Service interface {
}

type service struct {
	repository Repository
}

func NewCampaignsService(repository Repository) *service {
	return &service{repository: repository}
}
