package payment

import (
	"context"
	"fmt"

	"github.com/meteedev/grpc-ms-go/order/internal/application/core/domain"
	grpc_payment "github.com/meteedev/grpc-ms-go/order/proto/go_proto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment grpc_payment.PaymentClient
	
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		fmt.Println("err in init grpc payment adapter ",err.Error())
		return nil, err
	}
	defer conn.Close()

	client := grpc_payment.NewPaymentClient(conn)
	
	fmt.Println("connection state in step new ",conn.GetState().String())
	

	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	
	
	response, err := a.payment.Create(context.Background(),
		&grpc_payment.CreatePaymentRequest{
			UserId:     order.CustomerID,
			OrderId:    order.ID,
			TotalPrice: order.TotalPrice(),
		})

	
	if err != nil{
		fmt.Println("err ",err.Error())
	}
		
	
	fmt.Println("response ",response)
		
	return err
}
