package user

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServiceRegisterUser(t *testing.T) {

	//register service
	registerUserInput := RegisterUserInput{
		Name:       "akandidelete",
		Occupation: "anggota",
		Email:      "akandidelete@rambo.com",
		Password:   "anggota",
	}
	userService.RegisterUser(registerUserInput)

}

func TestServiceLoginUser(t *testing.T) {

	//login service success
	loginUserInput := LoginUserInput{
		Email:    "akandidelete@rambo.com",
		Password: "anggota",
	}
	result, err := userService.Login(loginUserInput)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	//login service email not found
	loginEmailNotFound := LoginUserInput{
		Email:    "ngkketemu@mail.com",
		Password: "lucubgt",
	}
	resultNotFound, errNotFound := userService.Login(loginEmailNotFound)
	require.Error(t, errNotFound)
	require.Empty(t, resultNotFound)

}
