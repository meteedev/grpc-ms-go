package ports

import "github.com/meteedev/grpc-ms-go/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
