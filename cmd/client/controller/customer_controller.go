package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	defer r.Body.Close()
	var ctx = r.Context()

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	if customer.Name == "" || customer.Email == "" {
		http.Error(w, "Both 'Name' and 'Email' fields are required", http.StatusBadRequest)
		return
	}

	createUserReq := &pb.CreateCustomerRequest{
		Name:  customer.Name,
		Email: customer.Email,
	}

	createdCustomer, err := c.client.CreateCustomer(ctx, createUserReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create customer: %v", err), http.StatusBadRequest)
		return
	}

	createdCustomerID, err := gocql.ParseUUID(createdCustomer.Customer.Id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse UUID: %v", err), http.StatusInternalServerError)
		return

	}

	createdAt := helpers.ProtoToTime(createdCustomer.Customer.CreatedAt)
	updatedAt := helpers.ProtoToTime(createdCustomer.Customer.UpdatedAt)

	createCustomerInjSON := &models.Customer{
		ID:        createdCustomerID,
		Name:      createdCustomer.Customer.Name,
		Email:     createdCustomer.Customer.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// marshal created customer to json and return it

	err = helpers.ConvertStructToJson(w, http.StatusCreated, createCustomerInjSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")
	var ctx = r.Context()

	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	deleteCustomerReq := &pb.DeleteCustomerRequest{
		CustomerId: id,
	}
	deletedCustomerResp, err := c.client.DeleteCustomer(ctx, deleteCustomerReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete customer: %v", err), http.StatusInternalServerError)
		return
	}
	if deletedCustomerResp.Success {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, deletedCustomerResp.Message, http.StatusInternalServerError)

}
func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")
	var ctx = r.Context()

	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return

	}

	getCustomerReq := &pb.GetCustomerRequest{
		CustomerId: id,
	}
	customer, err := c.client.GetCustomer(ctx, getCustomerReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get customer: %v", err), http.StatusInternalServerError)
		return
	}

	customerId, err := gocql.ParseUUID(customer.Customer.Id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse id to uuid: %v", err), http.StatusInternalServerError)
		return
	}
	createdAt := helpers.ProtoToTime(customer.Customer.CreatedAt)
	updatedAt := helpers.ProtoToTime(customer.Customer.UpdatedAt)

	customerInJson := models.Customer{
		ID:        customerId,
		Name:      customer.Customer.Name,
		Email:     customer.Customer.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if err := helpers.ConvertStructToJson(w, http.StatusOK, customerInJson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
