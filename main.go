package main

import (
	"bwastartup/app/database"
	"bwastartup/app/entity"
	"bwastartup/app/repository/user"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCacheChannel()

	insert_user := entity.User{Name: "test"}
	user.Save(insert_user)
}
