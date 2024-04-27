package user

import (
	"bwastartup/app/database"
	"bwastartup/app/entity"
	"bwastartup/app/model"
)

type Repository interface {
	Save(user entity.User) (entity.User, error)
}

func Save(user entity.User) (entity.User, error) {
	u := model.Use(database.DB)
	userDB := u.User
	err := userDB.Create(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
