package controller

import (
	"net/http"

	"google.golang.org/grpc"
)

type OrderController struct {
	conn *grpc.ClientConn
}

func NewOrderController(conn *grpc.ClientConn) *OrderController {
	return &OrderController{conn: conn}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
}
func (c *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}
func (c *OrderController) Getorder(w http.ResponseWriter, r *http.Request) {
}
func (c *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
}
