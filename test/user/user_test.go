package user

import (
	"bwastartup/src/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var userRepository user.Repository //var userRepository *repository
var userService user.Service       //var userService *service

var DbTest *gorm.DB

func init() {
	var err error
	dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	DbTest, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the mysql!")
	}
	fmt.Println("Database Connected")
	userRepository := user.NewRepository(DbTest)
	userService = user.NewService(userRepository)
}
