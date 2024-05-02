package ports

import "github.com/meteedev/grpc-ms-go/payment/internal/application/core/domain"

type APIPort interface {
	Charge(payment domain.Payment) (domain.Payment, error)
}
