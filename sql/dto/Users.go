package dto

import (
	"database/sql"
)

type Users struct {
	ID             int            `db:"id"`
	Name           sql.NullString `db:"name"`
	Occupation     sql.NullString `db:"occupation"`
	Email          sql.NullString `db:"email"`
	PasswordHash   sql.NullString `db:"password_hash"`
	AvatarFileName sql.NullString `db:"avatar_file_name"`
	Role           sql.NullString `db:"role"`
	Token          sql.NullString `db:"token"`
	CreatedAt      sql.NullTime   `db:"created_at"`
	UpdatedAt      sql.NullTime   `db:"updated_at"`
}
