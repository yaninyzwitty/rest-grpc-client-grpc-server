package controller

import (
	"context"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/services"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
)

type CustomerController struct {
	service services.CustomerService
	pb.UnimplementedCustomerServiceServer
}

func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (c *CustomerController) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	return nil, nil
}

func (c *CustomerController) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	return nil, nil
}

func (c *CustomerController) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	return nil, nil
}
