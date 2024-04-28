package transactions

import "gorm.io/gorm"

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
