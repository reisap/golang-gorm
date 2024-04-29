package user

import (
	abstractRepo "bwastartup/src/helper/repository"
	"bwastartup/src/user/dto"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user dto.User) (dto.User, error)
	FindByEmail(email string) (dto.User, error)
	DeleteUserByEmail(email string) (dto.User, error)

	FindUserById(id int) (dto.User, error)
	UpdateUser(user dto.User) (dto.User, error)
}

// digunakan untuk initialize
type repository struct {
	db               *gorm.DB
	userAbstractRepo abstractRepo.AbstractRepository[dto.User]
}

func NewRepository(db *gorm.DB) *repository { //*repository
	model := new(dto.User)
	return &repository{
		db: db,
		userAbstractRepo: abstractRepo.AbstractRepository[dto.User]{
			DB:     db,
			Entity: model,
		},
	}
}
func (r *repository) Create(user dto.User) (dto.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}
func (r *repository) FindByEmail(email string) (dto.User, error) {
	var user dto.User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) DeleteUserByEmail(email string) (dto.User, error) {
	//biasanya untuk hal ini bukan delete tapi update berupa flag
	//hanya sekedar test agar proses testing bisa lebih real
	var user dto.User
	err := r.db.Where("email=?", email).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindUserById(id int) (dto.User, error) {
	var user dto.User
	err := r.db.Where("id=?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) UpdateUser(user dto.User) (dto.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
