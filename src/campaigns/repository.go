package campaigns

import "gorm.io/gorm"

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewCampaignsRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
