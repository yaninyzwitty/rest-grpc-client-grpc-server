package repository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order) (*models.Order, error)
	DeleteOrder(ctx context.Context, orderId gocql.UUID) (bool, error)
	UpdateOrder(ctx context.Context, order models.Order, orderId gocql.UUID) (*models.Order, error)
	GetOrder(ctx context.Context, orderId gocql.UUID) (*models.Order, error)
}

type orderRepository struct {
	session *gocql.Session
}

func NewOrderRepository(session *gocql.Session) OrderRepository {
	return &orderRepository{session: session}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	query := `INSERT INTO ecommerce.orders (id, product_id, quantity, customer_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)`
	err := r.session.Query(query, order.ID, order.ProductID, order.Quantity, order.CustomerID, order.CreatedAt, order.UpdatedAt).Exec()
	if err != nil {
		return nil, fmt.Errorf("failed to insert data: %v", err)
	}
	return &order, nil
}

func (r *orderRepository) DeleteOrder(ctx context.Context, orderId gocql.UUID) (bool, error) {
	query := `DELETE FROM eccomerce.orders WHERE id = ?`
	err := r.session.Query(query, orderId).Exec()
	if err != nil {
		return false, fmt.Errorf("failed to delete the order: %v", err)
	}
	return true, nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, order models.Order, orderId gocql.UUID) (*models.Order, error) {
	query := `UPDATE eccomerce.orders SET quantity = ?, updated_at = ? WHERE id = ?`
	if err := r.session.Query(query, order.Quantity, order.UpdatedAt, orderId).Exec(); err != nil {
		return nil, fmt.Errorf("failed to update the orders: %v", err)
	}
	return &order, nil
}
func (r *orderRepository) GetOrder(ctx context.Context, orderId gocql.UUID) (*models.Order, error) {
	var order models.Order
	query := `SELECT id, product_id, quantity, customer_id, updated_at, created_at FROM orders WHERE id = ?`

	// Execute the query and scan the result into the order variable
	err := r.session.Query(query, orderId).Consistency(gocql.One).Scan(
		&order.ID,
		&order.ProductID,
		&order.Quantity,
		&order.CustomerID,
		&order.UpdatedAt,
		&order.CreatedAt,
	)

	if err != nil {
		if err == gocql.ErrNotFound {
			// Return nil order and nil error for not found case
			return nil, nil
		}
		// Return an error for any other issues
		return nil, err
	}
	return &order, nil // Return the populated order
}
