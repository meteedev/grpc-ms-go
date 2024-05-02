package ports

import (
	"github.com/meteedev/grpc-ms-go/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(*domain.Order) error
}
