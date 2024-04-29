package transactions

import (
	abstractRepo "bwastartup/src/helper/repository"
	"bwastartup/src/transactions/dto"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db           *gorm.DB
	AbstractRepo abstractRepo.AbstractRepository[dto.Transactions_table]
}

func NewTransactionsRepository(db *gorm.DB) Repository {
	model := new(dto.Transactions_table)
	return &repository{
		db: db,
		AbstractRepo: abstractRepo.AbstractRepository[dto.Transactions_table]{
			DB:     db,
			Entity: model,
		},
	}
}
