package user

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserSaveRepository(t *testing.T) {

	userModel := User{
		Name:           "rambo",
		Occupation:     "anggota",
		Email:          "rambo@rambo.com",
		PasswordHash:   "hashsah",
		AvatarFileName: "",
		Role:           "",
		Token:          "",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	userRepository.Save(userModel)
}

func TestUserFindByEmail(t *testing.T) {
	//success
	userEmail := "rambo@rambo.com"
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
