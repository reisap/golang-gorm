package user

import (
	"bwastartup/src/user/dto"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServiceRegisterUser(t *testing.T) {

	//register service
	registerUserInput := dto.RegisterUserInput{
		Name:       "akandidelete",
		Occupation: "anggota",
		Email:      "akandidelete@rambo.com",
		Password:   "anggota",
	}

	result, err := userService.RegisterUser(registerUserInput)
	require.NoError(t, err)
	require.NotEmpty(t, result)

}

func TestServiceLoginUser(t *testing.T) {

	//login service success
	loginUserInput := dto.LoginUserInput{
		Email:    "akandidelete@rambo.com",
		Password: "anggota",
	}

	result, err := userService.Login(loginUserInput)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	//login service email not found
	loginEmailNotFound := dto.LoginUserInput{
		Email:    "ngkketemu@mail.com",
		Password: "lucubgt",
	}
	resultNotFound, errNotFound := userService.Login(loginEmailNotFound)
	require.Error(t, errNotFound)
	require.Empty(t, resultNotFound)

}

func TestPaginationUser(t *testing.T) {
	result, err := userService.UserPaging(1, 10)
	b, err := json.Marshal(result)
	fmt.Println(string(b))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
