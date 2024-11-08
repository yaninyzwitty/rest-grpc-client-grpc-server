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

type ProductController struct {
	service services.ProductService
	pb.UnimplementedProductServiceServer
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	productId := gocql.TimeUUID()
	product := models.Product{
		ID:          productId,
		Name:        req.Name,
		Description: req.Description,
		Price:       uint32(req.Price) * 100,
		Stock:       uint32(req.Stock),
		Category:    req.Category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdProduct, err := c.service.CreateProduct(ctx, product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create the product: %v", err))
	}

	createdAt := helpers.TimeToProto(createdProduct.CreatedAt)
	updatedAt := helpers.TimeToProto(createdProduct.UpdatedAt)

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          createdProduct.ID.String(),
			Name:        createdProduct.Name,
			Description: createdProduct.Description,
			Price:       float64(createdProduct.Price),
			Stock:       int32(createdProduct.Stock),
			Category:    createdProduct.Category,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		},
	}, nil
}

func (c *ProductController) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	productID, err := gocql.ParseUUID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to parse product id to uuid: %v", err))
	}

	product, err := c.service.GetProduct(ctx, req.Category, productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get product: %v", err))
	}

	createdAt := helpers.TimeToProto(product.CreatedAt)
	updatedAt := helpers.TimeToProto(product.UpdatedAt)

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:          product.ID.String(),
			Name:        product.Name,
			Description: product.Description,
			Price:       float64(product.Price) / 100,
			Stock:       int32(product.Stock),
			Category:    product.Category,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		},
	}, nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	productId, err := gocql.ParseUUID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to parse product id into uuid: %v", err))
	}
	deleted, err := c.service.DeleteProduct(ctx, req.Category, productId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to delete product: %v", err))

	}

	if deleted {

		return &pb.DeleteProductResponse{
			Success: true,
			Message: "Product deleted succesfully",
		}, nil
	}

	return nil, nil

}

// TODO, finish list products
func (c *ProductController) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return nil, nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	productId, err := gocql.ParseUUID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to parse product id to uuid: %v", err))
	}

	product := &models.Product{
		ID:          productId,
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       uint32(req.Product.Price) * 100,
		Stock:       uint32(req.Product.Stock),
		Category:    req.Category,
		UpdatedAt:   time.Now(),
	}

	_, err = c.service.UpdateProducts(ctx, *product, productId, req.Category)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to update product: %v", err))
	}

	return &pb.UpdateProductResponse{
		Success: true,
		Message: "Product updated succesfully",
	}, nil

}
