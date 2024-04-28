package v1

import (
	"bwastartup/src/auth"
	"bwastartup/src/handler"
	"bwastartup/src/helper/mysql"
	"bwastartup/src/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup) {
	userRepository := user.NewRepository(mysql.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	// Per route middleware, you can add as many as you desire.
	//api.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)
}
