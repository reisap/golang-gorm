package repository

import (
	paginator "github.com/dmitryburov/gorm-paginator"
	"gorm.io/gorm"
	"log"
)

type AbstractRepository[model any] struct {
	DB     *gorm.DB
	Entity *model
	paging Paging[model]
}

type Paging[model any] struct {
	Model      []*model              `json:"data"`
	Pagination *paginator.Pagination `json:"pagination"`
}

func (global *AbstractRepository[model]) Create(entity *model) error {
	return global.DB.Create(entity).Error

}

func (global *AbstractRepository[model]) Update(entity *model, id int) error {

	return global.DB.Where("id = ?", id).Save(&entity).Error
}

func (global *AbstractRepository[T]) Find() (T, error) {
	var model T
	err := global.DB.Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (global *AbstractRepository[T]) FindById(id any) (T, error) {
	var model T
	err := global.DB.Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (global *AbstractRepository[T]) Delete(id int) error {
	var model T
	return global.DB.Where("id = ?", id).Delete(&model).Error
}

func (global *AbstractRepository[T]) CountById(id any) (int64, error) {
	var total int64
	var model T
	err := global.DB.Model(&model).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (global *AbstractRepository[model]) FindPaging(limit int, page int) (Paging[model], error) {
	var err error

	paging := paginator.Paging{}
	paging.Page = limit
	paging.Limit = page

	global.paging.Pagination, err = paginator.Pages(&paginator.Param{
		DB:     global.DB,
		Paging: &paging,
	}, &global.paging.Model)

	if err != nil {
		log.Fatal("Error get list: ", err.Error())
	}

	return global.paging, nil

}
