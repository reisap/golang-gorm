package user

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement:true"  json:"id"`
	Name           string    `gorm:"column:name" gorm:"type:varchar(255)" json:"name"`
	Occupation     string    `gorm:"column:occupation" gorm:"type:varchar(255)" json:"occupation"`
	Email          string    `gorm:"column:email" gorm:"type:varchar(255)" json:"email"`
	PasswordHash   string    `gorm:"column:password_hash" gorm:"type:varchar(255)" json:"password_hash"`
	AvatarFileName string    `gorm:"column:avatar_file_name" gorm:"type:varchar(255)" json:"avatar_file_name"`
	Role           string    `gorm:"column:role" gorm:"type:varchar(255)" json:"role"`
	Token          string    `gorm:"column:token" gorm:"type:varchar(255)" json:"token"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
