package mysql

import (
	dto2 "bwastartup/src/campaigns/dto"
	dto3 "bwastartup/src/campaigns_images/dto"
	dto4 "bwastartup/src/transactions/dto"
	"bwastartup/src/user/dto"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PORT = os.Getenv("DB_PORT")

	// ===== If using Mysql =====
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	//dsn := "root:secret@tcp(db:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"

	// ===== If using Postgres =====
	// Setup database connection here ...
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	//DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//jika dijalankan di luar docker
	//dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the mysql!")
	}
	fmt.Println("Database Connected")

}

func AutoMigrateDB() {
	err := DB.AutoMigrate(&dto.User{}, &dto2.Campaign{}, &dto3.CampaignImage{}, &dto4.Transactions_table{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
