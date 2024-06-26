package campaigns_images

import (
	"bwastartup/domain/campaigns_images/dto"
	abstractRepo "bwastartup/domain/helper/repository"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db           *gorm.DB
	abstractRepo abstractRepo.AbstractRepository[dto.CampaignImage]
}

func NewCampaignImagesRepository(db *gorm.DB) *repository {
	model := new(dto.CampaignImage)
	return &repository{
		db: db,
		abstractRepo: abstractRepo.AbstractRepository[dto.CampaignImage]{
			DB:     db,
			Entity: model,
		},
	}
}
