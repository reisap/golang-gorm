package dto

import (
	"time"
)

const TableNameTransactions_table = "transactions"

// Transactions_table mapped from table <transactions>
type Transactions_table struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CampaignID int       `gorm:"column:campaign_id" json:"campaign_id"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
	Amount     int       `gorm:"column:amount" json:"amount"`
	Status     string    `gorm:"column:status type:varchar(25)" json:"status"`
	Code       string    `gorm:"column:code type:varchar(255)" json:"code"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Transactions_table's table name
func (*Transactions_table) TableName() string {
	return TableNameTransactions_table
}
