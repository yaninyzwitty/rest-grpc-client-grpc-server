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

type OrderController struct {
	service services.OrderService
	pb.UnimplementedOrderServiceServer
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (c *OrderController) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderID := gocql.TimeUUID()

	productId, err := gocql.ParseUUID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to decode product id into uuid: %v", err))
	}
	customerID, err := gocql.ParseUUID(req.CustomerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to decode customer id into uuid: %v", err))
	}
	order := models.Order{
		ID:         orderID,
		ProductID:  productId,
		CustomerID: customerID,
		Quantity:   uint32(req.Quantity),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createdOrder, err := c.service.CreateOrder(ctx, order)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create order: %v", err))
	}

	createdAt := helpers.TimeToProto(createdOrder.CreatedAt)
	updatedAt := helpers.TimeToProto(createdOrder.UpdatedAt)

	return &pb.CreateOrderResponse{Order: &pb.Order{
		Id:         createdOrder.ID.String(),
		ProductId:  createdOrder.ProductID.String(),
		Quantity:   int32(createdOrder.Quantity),
		CustomerId: createdOrder.CustomerID.String(),
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	},
		Success: true,
	}, nil

}

func (c *OrderController) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	orderId, err := gocql.ParseUUID(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to parse order id into uuid: %v ", err))
	}

	deleted, err := c.service.DeleteOrder(ctx, orderId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete order: %v", err))

	}

	if deleted {
		return &pb.DeleteOrderResponse{
				Success: true,
				Message: "Order deleted succesfully",
			},
			nil
	}

	return nil, status.Errorf(codes.Unknown, fmt.Sprintf("something went wrong: %v", err))
}

func (c *OrderController) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	orderId, err := gocql.ParseUUID(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to parse order id into uuid: %v ", err))
	}
	productId, err := gocql.ParseUUID(req.Order.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to parse product id into uuid: %v ", err))
	}
	customerId, err := gocql.ParseUUID(req.Order.CustomerId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to parse customer id into uuid: %v ", err))
	}
	order := models.Order{
		ID:         orderId,
		ProductID:  productId,
		CustomerID: customerId,
		Quantity:   uint32(req.Order.Quantity),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	_, err = c.service.UpdateOrder(ctx, order, orderId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to update order: %v", err))
	}

	return &pb.UpdateOrderResponse{
		Message: "Order updated succesfully",
		Success: true,
	}, nil

}

func (c *OrderController) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	orderId, err := gocql.ParseUUID(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to parse id to uuid: %v", err))
	}

	order, err := c.service.GetOrder(ctx, orderId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("failed to get order: %v", err))
	}

	createdAt := helpers.TimeToProto(order.CreatedAt)
	updatedAt := helpers.TimeToProto(order.UpdatedAt)

	return &pb.GetOrderResponse{
		Order: &pb.Order{
			Id:         order.ID.String(),
			ProductId:  order.ProductID.String(),
			Quantity:   int32(order.Quantity),
			CustomerId: order.CustomerID.String(),
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		},
	}, nil
}
