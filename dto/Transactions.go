package dto

import (
	"database/sql"
)

type Transactions struct {
	ID         int            `db:"id"`
	CampaignID sql.NullInt64  `db:"campaign_id"`
	UserID     sql.NullInt64  `db:"user_id"`
	Amount     sql.NullInt64  `db:"amount"`
	Status     sql.NullString `db:"status"`
	Code       sql.NullString `db:"code"`
	CreatedAt  sql.NullTime   `db:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at"`
}
