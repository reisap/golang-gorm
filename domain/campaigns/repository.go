package campaigns

import (
	"bwastartup/domain/campaigns/dto"
	abstractRepo "bwastartup/domain/helper/repository"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db           *gorm.DB
	abstractRepo abstractRepo.AbstractRepository[dto.Campaign]
}

func NewCampaignsRepository(db *gorm.DB) *repository {
	model := new(dto.Campaign)
	return &repository{
		db: db,
		abstractRepo: abstractRepo.AbstractRepository[dto.Campaign]{
			DB:     db,
			Entity: model,
		},
	}
}
