package user

import "bwastartup/domain/helper"

var userRepository Repository //var userRepository *repository
var userService Service       //var userService *service

func init() {
	helper.ConnectDatabase()
	userRepository = NewRepository(helper.DB)
	userService = NewService(userRepository)
}
