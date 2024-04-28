package user

import "gorm.io/gorm"

type Repository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	DeleteUserByEmail(email string) (User, error)

	FindUserById(id int) (User, error)
	UpdateUser(user User) (User, error)
}

// digunakan untuk initialize
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository { //*repository
	return &repository{db: db}
}
func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) DeleteUserByEmail(email string) (User, error) {
	//biasanya untuk hal ini bukan delete tapi update berupa flag
	//hanya sekedar test agar proses testing bisa lebih real
	var user User
	err := r.db.Where("email=?", email).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindUserById(id int) (User, error) {
	var user User
	err := r.db.Where("id=?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}