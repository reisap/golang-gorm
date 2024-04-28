package helper

import (
	"bwastartup/domain/campaigns"
	"bwastartup/domain/campaigns_images"
	"bwastartup/domain/transactions"
	"bwastartup/domain/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "root:secret@tcp(db:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	//jika dijalankan di luar docker
	//dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database!")
	}
	fmt.Println("Database Connected")

}

func AutoMigrateDB() {
	DB.AutoMigrate(user.User{}, campaigns.Campaign{}, campaigns_images.CampaignImage{}, transactions.Transactions_table{})
}
