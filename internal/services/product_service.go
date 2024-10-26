package services

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/models"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product models.Product) (*models.Product, error)
	GetProduct(ctx context.Context, category string, productID gocql.UUID) (*models.Product, error)
	ListProducts(ctx context.Context, limit int, paging_state []byte, category string) (*[]models.Product, []byte, error)
	DeleteProduct(ctx context.Context, category string, productID gocql.UUID) (bool, error)
	UpdateProducts(ctx context.Context, product models.Product, productId gocql.UUID, category string) (*models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(ctx context.Context, product models.Product) (*models.Product, error) {

	return s.repo.CreateProduct(ctx, product)
}

func (s *productService) GetProduct(ctx context.Context, category string, productID gocql.UUID) (*models.Product, error) {
	return s.repo.GetProduct(ctx, category, productID)
}
func (s *productService) DeleteProduct(ctx context.Context, category string, productID gocql.UUID) (bool, error) {
	return s.repo.DeleteProduct(ctx, category, productID)
}

func (s productService) ListProducts(ctx context.Context, limit int, paging_state []byte, category string) (*[]models.Product, []byte, error) {
	return s.repo.ListProducts(ctx, limit, paging_state, category)
}
func (s *productService) UpdateProducts(ctx context.Context, product models.Product, productId gocql.UUID, category string) (*models.Product, error) {
	return s.repo.UpdateProducts(ctx, product, productId, category)
}
