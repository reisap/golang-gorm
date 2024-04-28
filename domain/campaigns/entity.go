package campaigns

import (
	"time"
)

const TableNameCampaign = "campaigns"

// Campaign mapped from table <campaigns>
type Campaign struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID           int32     `gorm:"column:user_id" json:"user_id"`
	Name             string    `gorm:"column:name" json:"name"`
	ShortDescription string    `gorm:"column:short_description" json:"short_description"`
	Description      string    `gorm:"column:description" json:"description"`
	Perks            string    `gorm:"column:perks" json:"perks"`
	BackerCount      int32     `gorm:"column:backer_count" json:"backer_count"`
	GoalAmount       int32     `gorm:"column:goal_amount" json:"goal_amount"`
	CurrentAmount    int32     `gorm:"column:current_amount" json:"current_amount"`
	Slug             string    `gorm:"column:slug" json:"slug"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Campaign's table name
func (*Campaign) TableName() string {
	return TableNameCampaign
}
