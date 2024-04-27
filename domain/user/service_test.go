package user

import (
	"testing"
)

func TestServiceRegisterUser(t *testing.T) {

	//register service
	registerUserInput := RegisterUserInput{
		Name:       "service test",
		Occupation: "anak servis",
		Email:      "servis@test.com",
		Password:   "servis",
	}
	userService.RegisterUser(registerUserInput)

}

func TestServiceLoginUser(t *testing.T) {

	//login service
	loginUserInput := LoginUserInput{
		Email:    "reisap@mail.com",
		Password: "lucubgt",
	}
	userService.Login(loginUserInput)

}
