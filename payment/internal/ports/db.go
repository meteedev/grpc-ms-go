package ports

import "github.com/meteedev/grpc-ms-go/payment/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Payment, error)
	Save(payment *domain.Payment) error
}
