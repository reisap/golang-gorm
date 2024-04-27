package user

import (
	"testing"
	"time"
)

func TestUserRepository(t *testing.T) {

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
