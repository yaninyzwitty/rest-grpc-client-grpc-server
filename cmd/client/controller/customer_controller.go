package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/cmd/client/models"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
)

type CustomerController struct {
	client pb.CustomerServiceClient
}

func NewCustomerController(client pb.CustomerServiceClient) *CustomerController {
	return &CustomerController{client: client}
}

func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	var ctx = r.Context()
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, http.StatusText(500), http.StatusBadRequest)
		return

	}

	createCustomer, err := c.client

}
func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
}
func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
}
