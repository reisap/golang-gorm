package main

import (
	"bwastartup/domain/auth"
	"bwastartup/domain/handler"
	"bwastartup/domain/helper"
	"bwastartup/domain/user"
	"github.com/gin-gonic/gin"
)

func main() {
	helper.ConnectDatabase()

	userRepository := user.NewRepository(helper.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run(":8080")
}
