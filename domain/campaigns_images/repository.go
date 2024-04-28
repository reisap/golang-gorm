package campaigns_images

import "gorm.io/gorm"

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewCampaignImagesRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
