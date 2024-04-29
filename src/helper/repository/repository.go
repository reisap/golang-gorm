package repository

import (
	paginator "github.com/dmitryburov/gorm-paginator"
	"gorm.io/gorm"
	"log"
)

type AbstractRepository[model any] struct {
	DB     *gorm.DB
	paging Paging[model]
}

type Paging[model any] struct {
	Model      []*model              `json:"data"`
	Pagination *paginator.Pagination `json:"pagination"`
}

func (global *AbstractRepository[model]) Create(entity *model) error {
	return global.DB.Create(entity).Error

}

func (global *AbstractRepository[model]) Find(entity *model) error {
	return global.DB.Find(entity).Error
}

func (global *AbstractRepository[T]) FindById(entity *T, id any) error {
	return global.DB.Where("id = ?", id).Take(entity).Error
}

func (global *AbstractRepository[T]) Update(entity *T, id int) error {
	return global.DB.Where("id = ?", id).Save(entity).Error
}

func (global *AbstractRepository[T]) Delete(entity *T, id int) error {
	return global.DB.Where("id = ?", id).Delete(entity).Error
}

func (global *AbstractRepository[T]) CountById(id any) (int64, error) {
	var total int64
	err := global.DB.Model(new(T)).Where("id = ?", id).Count(&total).Error
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
