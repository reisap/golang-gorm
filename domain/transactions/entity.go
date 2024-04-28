package transactions

import (
	"time"
)

const TableNameTransactions_table = "transactions"

// Transactions_table mapped from table <transactions>
type Transactions_table struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CampaignID int32     `gorm:"column:campaign_id" json:"campaign_id"`
	UserID     int32     `gorm:"column:user_id" json:"user_id"`
	Amount     int32     `gorm:"column:amount" json:"amount"`
	Status     string    `gorm:"column:status" json:"status"`
	Code       string    `gorm:"column:code" json:"code"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Transactions_table's table name
func (*Transactions_table) TableName() string {
	return TableNameTransactions_table
}
