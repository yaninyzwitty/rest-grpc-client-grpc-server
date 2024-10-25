package services

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	DeleteOrder(ctx context.Context, orderId gocql.UUID) (bool, error)
	GetOrder(ctx context.Context, orderId gocql.UUID) (*models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order, orderId gocql.UUID) (*models.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	return s.repo.CreateOrder(ctx, order)
}

func (s *orderService) DeleteOrder(ctx context.Context, orderId gocql.UUID) (bool, error) {
	return s.repo.DeleteOrder(ctx, orderId)
}

func (s *orderService) UpdateOrder(ctx context.Context, order models.Order, orderId gocql.UUID) (*models.Order, error) {
	return s.repo.UpdateOrder(ctx, order, orderId)
}
func (s *orderService) GetOrder(ctx context.Context, orderId gocql.UUID) (*models.Order, error) {
	return s.repo.GetOrder(ctx, orderId)
}
