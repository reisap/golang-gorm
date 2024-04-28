package v1

import (
	"bwastartup/domain/auth"
	"bwastartup/domain/handler"
	"bwastartup/domain/helper/mysql"
	"bwastartup/domain/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup) {
	userRepository := user.NewRepository(mysql.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)
}