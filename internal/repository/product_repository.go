package repository

import (
	"context"
	"fmt"

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
	query := `INSERT INTO eccomerce.products (id, name, description, price, stock, category, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the query
	err := r.session.Query(query, product.ID, product.Name, product.Description, product.Price, product.Stock, product.Category, product.CreatedAt, product.UpdatedAt).Exec()
	if err != nil {
		return nil, fmt.Errorf("failed to insert product: %v", err)
	}

	// Return the created product
	return &product, nil
}

func (r *productRepository) GetProduct(ctx context.Context, category string, productId gocql.UUID) (*models.Product, error) {
	var product models.Product
	query := `SELECT id, name, description, price, stock, category, created_at, updated_at FROM eccomerce.products WHERE id = ? AND category = ?`

	// Execute the query and scan the result into the product struct
	if err := r.session.Query(query, productId, category).Consistency(gocql.One).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.Category,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		// Handle the error appropriately
		if err == gocql.ErrNotFound {
			return nil, fmt.Errorf("product with id %s not found: %w", productId, err)
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return &product, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, category string, productId gocql.UUID) (bool, error) {
	query := `DELETE FROM eccomerce.products WHERE category = ? AND id = ?`
	err := r.session.Query(query, category, productId).Exec()
	if err != nil {
		return false, fmt.Errorf("failed to delete product: %v", err)
	}
	return true, nil
}

func (r *productRepository) ListProducts(ctx context.Context, limit int, paging_state []byte, category string) (*[]models.Product, []byte, error) {
	return nil, nil, nil
}

func (r *productRepository) UpdateProducts(ctx context.Context, product models.Product, productId gocql.UUID, category string) (*models.Product, error) {
	query := `UPDATE eccomerce.products SET price = ?, stock = ?, name = ?, description = ?, updated_at = ? WHERE id = ? AND category = ?`

	// Execute the update query
	if err := r.session.Query(query, product.Price, product.Stock, product.Name, product.Description, product.UpdatedAt, productId, category).Exec(); err != nil {
		return nil, fmt.Errorf("failed to update product with id %s: %w", productId, err)
	}

	return &product, nil
}

// /products/{productId}/category/{categoryName}
