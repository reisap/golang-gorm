package user

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserSaveRepository(t *testing.T) {

	userModel := User{
		Name:           "akandidelete",
		Occupation:     "anggota",
		Email:          "akandidelete@rambo.com",
		PasswordHash:   "anggota",
		AvatarFileName: "",
		Role:           "",
		Token:          "",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	result, err := userRepository.Save(userModel)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestUserFindByEmail(t *testing.T) {
	//success
	userEmail := "akandidelete@rambo.com"
	result, err := userRepository.FindByEmail(userEmail)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	//not success email not found
	userEmailNotFound := "emailngkada@mail.com"
	resultNotFound, errFound := userRepository.FindByEmail(userEmailNotFound)
	//if errFound != nil {
	//fmt.Println("Error find this email", errFound)
	//}
	//fmt.Println(resultNotFound)
	require.Error(t, errFound)
	require.Empty(t, resultNotFound)

}

func TestUserDeleteByEmail(t *testing.T) {
	userEmail := "akandidelete@rambo.com"
	result, err := userRepository.DeleteUserByEmail(userEmail)
	require.NoError(t, err)
	require.Empty(t, result)
}
