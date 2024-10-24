package controller

import (
	"net/http"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
)

type ProductController struct {
	client pb.ProductServiceClient
}

func NewProductController(client pb.ProductServiceClient) *ProductController {
	return &ProductController{client: client}
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
}
func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
}
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
}
func (c *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) {
}
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {

}
