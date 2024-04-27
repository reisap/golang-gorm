package user

import (
	"github.com/stretchr/testify/require"
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
	result, err := userService.Login(loginUserInput)
	require.NoError(t, err)
	require.NotEmpty(t, result)

}
