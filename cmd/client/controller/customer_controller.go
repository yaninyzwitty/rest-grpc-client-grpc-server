package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/cmd/client/models"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/helpers"
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
	if customer.Name == "" || customer.Email == "" {
		http.Error(w, "Name and email cant be missing", http.StatusBadRequest)
		return
	}

	createUserReq := &pb.CreateCustomerRequest{
		Name:  customer.Name,
		Email: customer.Email,
	}

	createdCustomer, err := c.client.CreateCustomer(ctx, createUserReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdCustomerID, err := gocql.ParseUUID(createdCustomer.Customer.Id)
	if err != nil {
		http.Error(w, err.Error()+"error parsing uuid", http.StatusInternalServerError)
		return

	}

	createdAt := helpers.ProtoToTime(createdCustomer.Customer.CreatedAt)
	updatedAt := helpers.ProtoToTime(createdCustomer.Customer.UpdatedAt)

	createCustomerInjSON := models.Customer{
		ID:        createdCustomerID,
		Name:      createdCustomer.Customer.Name,
		Email:     createdCustomer.Customer.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// marshal created customer to json and return it

}
func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
}
func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
}
