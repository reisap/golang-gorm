package transactions

type Service interface {
}

type service struct {
	repository Repository
}

func NewTransactionsService(repository Repository) Service {
	return &service{repository}
}
