package campaigns_images

type Service interface {
	FindPagingTbl(limit int, page int) (any, error)
	FindAllTbl() (any, error)
}

type service struct {
	repository *repository
}

func NewCampaignImagesService(repository *repository) *service {
	return &service{repository: repository}
}

func (s *service) FindPagingTbl(limit int, page int) (any, error) {

	list, err := s.repository.abstractRepo.FindPaging(limit, page)
	if err != nil {
		return list, err
	}
	return list, nil

}
func (s *service) FindAllTbl() (any, error) {

	list, err := s.repository.abstractRepo.Find()
	if err != nil {
		return list, err
	}
	return list, nil

}
