package api

import (
	"github.com/meteedev/grpc-ms-go/payment/internal/application/core/domain"
	"github.com/meteedev/grpc-ms-go/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(&payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
