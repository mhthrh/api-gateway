package service

import (
	"context"
	"github.com/mhthrh/GoNest/model/customer/grpc/customer"
	"google.golang.org/grpc"
	"time"
)

type Grpc struct {
	cnn *grpc.ClientConn
}

func NewGrpc() (*Grpc, error) {
	conn, err := grpc.NewClient("localhost:6985")
	if err != nil {
		return nil, err
	}
	return &Grpc{cnn: conn}, nil
}

func (g Grpc) Register() (*customer.Response, error) {
	pb := customer.NewCustomerServiceClient(g.cnn)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	//a, b := pb.RegisterCustomer(ctx, &customer.Request{
	//	Address: &address.Address{
	//		Street:     "",
	//		City:       nil,
	//		State:      "",
	//		PostalCode: "",
	//		Country:    nil,
	//	},
	//	Customer: &customer.Customer{
	//		CustomerId: "",
	//		IdType:     0,
	//		UserName:   "",
	//		Password:   "",
	//		Email:      "",
	//		Mobile:     "",
	//		FirstName:  "",
	//		MiddleName: "",
	//		LastName:   "",
	//		CreatedAt:  nil,
	//		UpdatedAt:  nil,
	//		Status:     0,
	//		Picture:    nil,
	//		Document:   nil,
	//	},
	//})
	a, b := pb.RegisterCustomer(ctx, nil)
	return a, b
}
