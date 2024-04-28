package campaigns_images

import (
	"time"
)

const TableNameCampaignImage = "campaign_images"

// CampaignImage mapped from table <campaign_images>
type CampaignImage struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CampaignID int       `gorm:"column:campaign_id" json:"campaign_id"`
	FileName   string    `gorm:"column:file_name type:varchar(255)" json:"file_name"`
	IsPrimary  int       `gorm:"column:is_primary" json:"is_primary"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName CampaignImage's table name
func (*CampaignImage) TableName() string {
	return TableNameCampaignImage
}
