package dto

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement:true"  json:"id"`
	Name           string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Occupation     string    `gorm:"column:occupation;type:varchar(255)" json:"occupation"`
	Email          string    `gorm:"column:email;type:varchar(255) unique" json:"email"`
	PasswordHash   string    `gorm:"column:password_hash;type:varchar(255)" json:"password_hash"`
	AvatarFileName string    `gorm:"column:avatar_file_name;type:varchar(255)" json:"avatar_file_name"`
	Role           string    `gorm:"column:role;type:varchar(255)" json:"role"`
	Token          string    `gorm:"column:token;type:varchar(255)" json:"token"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
