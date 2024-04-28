package v1

import (
	"bwastartup/src/auth"
	"bwastartup/src/controller"
	"bwastartup/src/helper/mysql"
	"bwastartup/src/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup) {
	userRepository := user.NewRepository(mysql.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userController := controller.NewUserController(userService, authService)

	// Per route middleware, you can add as many as you desire.
	//api.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	api.POST("/users", userController.RegisterUser)
	api.POST("/sessions", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatar", userController.UploadAvatar)
}
