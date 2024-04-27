package main

import (
	"bwastartup/domain/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCacheChannel()

	insert_user := entity.User{Name: "test"}
	user.Save(insert_user)
}

*/

func main() {
	dsn := "root:secret@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database!")
	}
	fmt.Println("Database connected")
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Rambo"
	userInput.Occupation = "pejuang"
	userInput.Password = "secret"
	userInput.Email = "rambol@test.com"
	userService.RegisterUser(userInput)

	//user := user.User{
	//
	//	Name:           "jarwo",
	//	Occupation:     "it",
	//	Email:          "jarwo@test.com",
	//	PasswordHash:   "test",
	//	AvatarFileName: "test.png",
	//	Role:           "admin",
	//	Token:          "test",
	//	CreatedAt:      time.Time{},
	//	UpdatedAt:      time.Time{},
	//}
	//userRepository.Save(user)
}
