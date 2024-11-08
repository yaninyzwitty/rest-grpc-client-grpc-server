package repository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, customerId gocql.UUID) (bool, error)
	GetCustomer(ctx context.Context, customerID gocql.UUID) (*models.Customer, error)
}

type customerRepository struct {
	session *gocql.Session
}

func NewCustomerRepository(session *gocql.Session) CustomerRepository {
	return &customerRepository{session: session}
}

func (r *customerRepository) CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error) {
	query := `INSERT INTO eccomerce.customers (id, name, email, created_at, updated_at) VALUES(?, ?, ?, ?, ?)`

	err := r.session.Query(query, customer.ID, customer.Name, customer.Email, customer.CreatedAt, customer.UpdatedAt).Exec()
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}

	return &customer, nil
}

func (r *customerRepository) DeleteCustomer(ctx context.Context, customerId gocql.UUID) (bool, error) {
	query := `DELETE FROM eccomerce.customers WHERE id = ?`
	err := r.session.Query(query, customerId).Exec()
	if err != nil {
		return false, fmt.Errorf("failed to delete customer: %v", err)
	}
	return true, nil
}
func (r *customerRepository) GetCustomer(ctx context.Context, customerId gocql.UUID) (*models.Customer, error) {
	var customer models.Customer

	query := `SELECT id, name, email, created_at, updated_at FROM eccomerce.customers WHERE id = ?`

	// Execute the query and scan the results into the customer struct
	err := r.session.Query(query, customerId).WithContext(ctx).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil // No customer found with the given ID
		}
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	return &customer, nil
}
