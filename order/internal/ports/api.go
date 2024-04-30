package ports

import "github.com/meteedev/go-ms-grpc/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
