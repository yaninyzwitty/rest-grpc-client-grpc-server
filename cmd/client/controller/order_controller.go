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

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *OrderController) Getorder(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) error {
	return nil
}
