package controller

import (
	"net/http"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
)

type OrderController struct {
	client pb.OrderServiceClient
}

func NewOrderController(client pb.OrderServiceClient) *OrderController {
	return &OrderController{client: client}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
}
func (c *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}
func (c *OrderController) Getorder(w http.ResponseWriter, r *http.Request) {
}
func (c *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
}
