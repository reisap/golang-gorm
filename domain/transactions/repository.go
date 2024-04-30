package transactions

import (
	abstractRepo "bwastartup/domain/helper/repository"
	"bwastartup/domain/transactions/dto"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db           *gorm.DB
	abstractRepo abstractRepo.AbstractRepository[dto.Transactions_table]
}

func NewRepository(db *gorm.DB) *repository { //*repository
	model := new(dto.Transactions_table)
	return &repository{
		db: db,
		abstractRepo: abstractRepo.AbstractRepository[dto.Transactions_table]{
			DB:     db,
			Entity: model,
		},
	}
}
