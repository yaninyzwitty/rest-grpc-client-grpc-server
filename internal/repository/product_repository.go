package repository

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product models.Product) (*models.Product, error)
	GetProduct(ctx context.Context, category string, productId gocql.UUID) (*models.Product, error)
	DeleteProduct(ctx context.Context, category string, productId gocql.UUID) (bool, error)
	ListProducts(ctx context.Context, limit int, paging_state []byte, category string) (*[]models.Product, []byte, error)
	UpdateProducts(ctx context.Context, product models.Product, productId gocql.UUID, category string) (*models.Product, error)
}

type productRepository struct {
	session *gocql.Session
}

func NewProductRepository(session *gocql.Session) ProductRepository {
	return &productRepository{session: session}
}

func (r *productRepository) CreateProduct(ctx context.Context, product models.Product) (*models.Product, error) {
	return &models.Product{}, nil //not implemented
}
func (r *productRepository) GetProduct(ctx context.Context, category string, productId gocql.UUID) (*models.Product, error) {
	return &models.Product{}, nil // not implemented
}
func (r *productRepository) DeleteProduct(ctx context.Context, category string, productId gocql.UUID) (bool, error) {
	return false, nil // not implemented
}

func (r *productRepository) ListProducts(ctx context.Context, limit int, paging_state []byte, category string) (*[]models.Product, []byte, error) {
	return nil, nil, nil
}

func (r *productRepository) UpdateProducts(ctx context.Context, product models.Product, productId gocql.UUID, category string) (*models.Product, error) {
	return nil, nil
}
