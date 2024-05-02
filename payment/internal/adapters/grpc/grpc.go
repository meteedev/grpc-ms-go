package grpc

import (
	"context"
	"fmt"

	"github.com/meteedev/grpc-ms-go/payment/internal/application/core/domain"
	"github.com/meteedev/grpc-ms-go/payment/internal/ports"
	grpc_payment "github.com/meteedev/grpc-ms-go/payment/proto/go_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	grpc_payment.UnimplementedPaymentServer
}

func (a Adapter) Stop() {
	a.server.Stop()
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Create(ctx context.Context, request *grpc_payment.CreatePaymentRequest) (*grpc_payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &grpc_payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
