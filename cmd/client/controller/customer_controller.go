package controller

import (
	"net/http"

	"google.golang.org/grpc"
)

type CustomerController struct {
	conn *grpc.ClientConn
}

func NewCustomerController(conn *grpc.ClientConn) *CustomerController {
	return &CustomerController{conn: conn}
}

func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
