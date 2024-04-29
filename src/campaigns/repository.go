package campaigns

import (
	"bwastartup/src/campaigns/dto"
	abstractRepo "bwastartup/src/helper/repository"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db           *gorm.DB
	AbstractRepo abstractRepo.AbstractRepository[dto.Campaign]
}

func NewCampaignsRepository(db *gorm.DB) *repository {
	model := new(dto.Campaign)
	return &repository{
		db: db,
		AbstractRepo: abstractRepo.AbstractRepository[dto.Campaign]{
			DB:     db,
			Entity: model,
		},
	}
}
