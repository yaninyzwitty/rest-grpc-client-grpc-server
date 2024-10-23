package controller

import (
	"net/http"

	"google.golang.org/grpc"
)

type ProductController struct {
	conn *grpc.ClientConn
}

func NewProductController(conn *grpc.ClientConn) *ProductController {
	return &ProductController{conn: conn}
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}
