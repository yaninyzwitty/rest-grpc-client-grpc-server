package services

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/repository"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, customerId gocql.UUID) (bool, error)
	GetCustomer(ctx context.Context, customerID gocql.UUID) (*models.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) CreateCustomer(ctx context.Context, customer models.Customer) (*models.Customer, error) {
	return &models.Customer{}, nil
}
func (s *customerService) DeleteCustomer(ctx context.Context, customerID gocql.UUID) (bool, error) {
	return false, nil
}
func (s *customerService) GetCustomer(ctx context.Context, customerID gocql.UUID) (*models.Customer, error) {
	return nil, nil
}
