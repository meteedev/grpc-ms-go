package ports

import (
	"github.com/meteedev/grpc-ms-go/order/internal/domain"
)

type PaymentPort type interface{
	Charge(*domain.Order)
}