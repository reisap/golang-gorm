package dto

import (
	"database/sql"
)

type CampaignImages struct {
	ID         int            `db:"id"`
	CampaignID sql.NullInt64  `db:"campaign_id"`
	FileName   sql.NullString `db:"file_name"`
	IsPrimary  sql.NullInt64  `db:"is_primary"`
	CreatedAt  sql.NullTime   `db:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at"`
}
