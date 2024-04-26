package database

import (
	"bwastartup/app/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
	fmt.Println("Database connected")
}

func AutoMigrate() {
	DB.AutoMigrate(entity.User{}, entity.Campaign{}, entity.CampaignImage{}, entity.Transactions_table{})
}
