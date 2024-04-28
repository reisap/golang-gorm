package user

import (
	"bwastartup/domain/helper/mysql"
)

var userRepository Repository //var userRepository *repository
var userService Service       //var userService *service

func init() {
	mysql.ConnectDatabase()
	userRepository = NewRepository(mysql.DB)
	userService = NewService(userRepository)
}
