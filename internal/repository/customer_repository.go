package repository

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, customerId gocql.UUID) (bool, error)
}

type customerRepository struct {
	session *gocql.Session
}

func NewCustomerRepository(session *gocql.Session) CustomerRepository {
	return &customerRepository{session: session}
}

func (r *customerRepository) CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error) {
	return &models.Customer{}, nil
}

func (r *customerRepository) DeleteCustomer(ctx context.Context, customerId gocql.UUID) (bool, error) {
	return false, nil
}
