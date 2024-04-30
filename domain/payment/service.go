package payment

type Service interface {
	GetPaymentUrl() (string, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl() (string, error) {
	return "", nil
}
