package user

import "time"

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name           string    `gorm:"column:name" json:"name"`
	Occupation     string    `gorm:"column:occupation" json:"occupation"`
	Email          string    `gorm:"column:email" json:"email"`
	PasswordHash   string    `gorm:"column:password_hash" json:"password_hash"`
	AvatarFileName string    `gorm:"column:avatar_file_name" json:"avatar_file_name"`
	Role           string    `gorm:"column:role" json:"role"`
	Token          string    `gorm:"column:token" json:"token"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}