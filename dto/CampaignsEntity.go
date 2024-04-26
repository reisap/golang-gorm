package dto

import (
	"database/sql"
)

type CampaignsEntity struct {
	ID               int            `db:"id"`
	UserID           sql.NullInt64  `db:"user_id"`
	Name             sql.NullString `db:"name"`
	ShortDescription sql.NullString `db:"short_description"`
	Description      sql.NullString `db:"description"`
	Perks            sql.NullString `db:"perks"`
	BackerCount      sql.NullInt64  `db:"backer_count"`
	GoalAmount       sql.NullInt64  `db:"goal_amount"`
	CurrentAmount    sql.NullInt64  `db:"current_amount"`
	Slug             sql.NullString `db:"slug"`
	CreatedAt        sql.NullTime   `db:"created_at"`
	UpdatedAt        sql.NullTime   `db:"updated_at"`
}
