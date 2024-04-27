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
	userEmail := "reisap@mail.com"
	result, err := userRepository.FindByEmail(userEmail)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	//if err != nil {
	//	fmt.Println("Error find this email")
	//}
	//fmt.Println(result)
}
