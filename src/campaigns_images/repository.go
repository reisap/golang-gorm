package campaigns_images

import (
	"bwastartup/src/campaigns_images/dto"
	abstractRepo "bwastartup/src/helper/repository"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db           *gorm.DB
	abstractRepo abstractRepo.AbstractRepository[dto.CampaignImage]
}

func NewCampaignImagesRepository(db *gorm.DB) Repository {
	model := new(dto.CampaignImage)
	return &repository{
		db: db,
		abstractRepo: abstractRepo.AbstractRepository[dto.CampaignImage]{
			DB:     db,
			Entity: model,
		},
	}
}
