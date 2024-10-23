package controller

import (
	"context"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/services"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
)

type OrderController struct {
	service services.OrderService
	pb.UnimplementedOrderServiceServer
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (c *OrderController) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return nil, nil
}

func (c *OrderController) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	return nil, nil
}

func (c *OrderController) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	return nil, nil
}

func (c *OrderController) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return nil, nil
}
