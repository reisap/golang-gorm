package user

import (
	"bwastartup/src/helper/mysql"
)

var userRepository Repository //var userRepository *repository
var userService Service       //var userService *service

func init() {
	mysql.ConnectDatabase()
	userRepository = NewRepository(mysql.DB)
	userService = NewService(userRepository)
}
