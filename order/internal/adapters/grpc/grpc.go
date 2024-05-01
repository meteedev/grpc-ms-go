package grpc

import (
	"context"
	"github.com/meteedev/grpc-ms-go/order/internal/application/core/domain"
	"github.com/meteedev/grpc-ms-go/proto/go_proto"
)

func (a Adapter) Create(ctx context.Context, request *grpc_order.CreateOrderRequest) (*grpc_order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &grpc_order.CreateOrderResponse{OrderId: result.ID}, nil
}
