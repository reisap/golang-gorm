package user

import (
	"bwastartup/src/user/dto"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input dto.RegisterUserInput) (dto.User, error)
	Login(input dto.LoginUserInput) (dto.User, error)
	IsEmailAvailable(input dto.CheckEmailInput) (bool, error)
	SaveAvatarUser(ID int, fileLocation string) (dto.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input dto.RegisterUserInput) (dto.User, error) {
	user := dto.User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	newUser, err := s.repository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input dto.LoginUserInput) (dto.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) IsEmailAvailable(input dto.CheckEmailInput) (bool, error) {
	email := input.Email
	_, err := s.repository.FindByEmail(email)
	emailRegister := errors.New("email found")
	if err != nil {
		return true, nil
	}

	return false, emailRegister

}

func (s *service) SaveAvatarUser(ID int, fileLocation string) (dto.User, error) {
	user, err := s.repository.FindUserById(ID)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation
	updateUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}
