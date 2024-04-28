package main

import (
	"bwastartup/domain/auth"
	"bwastartup/domain/handler"
	"bwastartup/domain/helper"
	"bwastartup/domain/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	helper.ConnectDatabase()
	helper.AutoMigrateDB()

	userRepository := user.NewRepository(helper.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error running API in port", err)
	}
}
