package campaigns

import (
	"time"
)

const TableNameCampaign = "campaigns"

// Campaign mapped from table <campaigns>
type Campaign struct {
	ID               int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID           int       `gorm:"column:user_id" json:"user_id"`
	Name             string    `gorm:"column:name" gorm:"type:varchar(255)" json:"name"`
	ShortDescription string    `gorm:"column:short_description" gorm:"type:varchar(255)" json:"short_description"`
	Description      string    `gorm:"column:description" gorm:"type:varchar(255)" json:"description"`
	Perks            string    `gorm:"column:perks" gorm:"type:varchar(255)" json:"perks"`
	BackerCount      int       `gorm:"column:backer_count"  json:"backer_count"`
	GoalAmount       int       `gorm:"column:goal_amount" json:"goal_amount"`
	CurrentAmount    int       `gorm:"column:current_amount" json:"current_amount"`
	Slug             string    `gorm:"column:slug" gorm:"type:varchar(255)" json:"slug"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Campaign's table name
func (*Campaign) TableName() string {
	return TableNameCampaign
}
