package controller

import (
	"context"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/services"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
)

type ProductController struct {
	service services.ProductService
	pb.UnimplementedProductServiceServer
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return nil, nil
}

func (c *ProductController) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	return nil, nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	return nil, nil
}

func (c *ProductController) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return nil, nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	return nil, nil
}
