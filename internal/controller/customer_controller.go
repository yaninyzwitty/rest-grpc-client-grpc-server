package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/helpers"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/services"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerController struct {
	service services.CustomerService
	pb.UnimplementedCustomerServiceServer
}

func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (c *CustomerController) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	customerID := gocql.TimeUUID()
	customer := models.Customer{
		ID:        customerID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdCustomer, err := c.service.CreateCustomer(ctx, customer)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create customer: %v", err))
	}

	createdTimeInProto := helpers.TimeToProto(createdCustomer.CreatedAt)
	updatedTimeInProto := helpers.TimeToProto(createdCustomer.UpdatedAt)

	return &pb.CreateCustomerResponse{
		Customer: &pb.Customer{
			Id:        createdCustomer.ID.String(),
			Name:      createdCustomer.Name,
			Email:     createdCustomer.Email,
			CreatedAt: createdTimeInProto,
			UpdatedAt: updatedTimeInProto,
		},
	}, nil

}

func (c *CustomerController) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	customerID, err := gocql.ParseUUID(req.CustomerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to decode id to uuid: %v", err))
	}

	deleted, err := c.service.DeleteCustomer(ctx, customerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete customer: %v", err))
	}
	if deleted {
		return &pb.DeleteCustomerResponse{
			Success: true,
			Message: "Customer was deleted succesfully",
		}, nil
	}

	return nil, status.Errorf(codes.Internal, fmt.Sprintf("something went wrong, item failed to delete: %v", err))

}

func (c *CustomerController) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customerId, err := gocql.ParseUUID(req.CustomerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to decode id to uuid: %v", err))
	}
	customer, err := c.service.GetCustomer(ctx, customerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get customer: %v", err))
	}

	createdTimeInProto := helpers.TimeToProto(customer.CreatedAt)
	updatedTimeInProto := helpers.TimeToProto(customer.UpdatedAt)
	return &pb.GetCustomerResponse{
		Customer: &pb.Customer{
			Id:        customer.ID.String(),
			Name:      customer.Name,
			Email:     customer.Email,
			CreatedAt: createdTimeInProto,
			UpdatedAt: updatedTimeInProto,
		},
	}, nil
}
