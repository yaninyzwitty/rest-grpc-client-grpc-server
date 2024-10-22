package repository

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
)

type OrderRepository interface {
}

type orderRepository struct {
	session *gocql.Session
}

func NewOrderRepository(session *gocql.Session) OrderRepository {
	return &orderRepository{session: session}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	return nil, nil
}

func (r *orderRepository) DeleteOrder(ctx context.Context, orderId gocql.UUID) (bool, error) {
	return false, nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, order models.Order, orderId gocql.UUID) (*models.Order, error) {
	return nil, nil
}
