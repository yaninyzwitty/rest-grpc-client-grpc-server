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

type OrderController struct {
	client pb.OrderServiceClient
}

func NewOrderController(client pb.OrderServiceClient) *OrderController {
	return &OrderController{client: client}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	var ctx = r.Context()
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	if order.ProductID.String() == "" || order.Quantity == 0 || order.CustomerID.String() == "" {
		http.Error(w, "Product Id, Quantity and Customer id are required", http.StatusBadRequest)
	}

	createOrderRequest := &pb.CreateOrderRequest{
		ProductId:  order.ProductID.String(),
		Quantity:   int32(order.Quantity),
		CustomerId: order.CustomerID.String(),
	}
	createdOrder, err := c.client.CreateOrder(ctx, createOrderRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create an order: %v ", err), http.StatusInternalServerError)
		return
	}

	createdOrderUUID, err := gocql.ParseUUID(createdOrder.Order.CustomerId)
	if err != nil {
		http.Error(w, "failed to parse id to uuid", http.StatusInternalServerError)
		return
	}
	productUUID, err := gocql.ParseUUID(createdOrder.Order.ProductId)
	if err != nil {
		http.Error(w, "failed to parse id to uuid", http.StatusInternalServerError)
		return
	}
	customerUUID, err := gocql.ParseUUID(createdOrder.Order.CustomerId)
	if err != nil {
		http.Error(w, "failed to parse id to uuid", http.StatusInternalServerError)
		return
	}
	createdAt := helpers.ProtoToTime(createdOrder.Order.CreatedAt)
	updatedAt := helpers.ProtoToTime(createdOrder.Order.UpdatedAt)

	createOrderInjSON := models.Order{
		ID:         createdOrderUUID,
		ProductID:  productUUID,
		Quantity:   uint32(createdOrder.Order.Quantity),
		CustomerID: customerUUID,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
	if err := helpers.ConvertStructToJson(w, http.StatusCreated, createOrderInjSON); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (c *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")
	var ctx = r.Context()

	if id == "" {
		http.Error(w, "Order id is required", http.StatusBadRequest)
		return
	}

	deleteOrderReq := &pb.DeleteOrderRequest{
		OrderId: id,
	}
	deletedOrderResp, err := c.client.DeleteOrder(ctx, deleteOrderReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete order: %v", err), http.StatusInternalServerError)
		return
	}

	if deletedOrderResp.Success {
		w.WriteHeader(http.StatusNoContent) // 204 No Content
		return
	}
	http.Error(w, "Failed to delete order: order not found", http.StatusNotFound)

}
func (c *OrderController) Getorder(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")
	var ctx = r.Context()

	if id == "" {
		http.Error(w, "Order id is required", http.StatusBadRequest)
		return

	}

	getOrderReq := &pb.GetOrderRequest{
		OrderId: id,
	}

	order, err := c.client.Getorder(ctx, getOrderReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderInUUID, err := gocql.ParseUUID(order.Order.Id)
	if err != nil {
		http.Error(w, "failed to parse id to uuid", http.StatusBadRequest)
		return
	}
	productInUUID, err := gocql.ParseUUID(order.Order.ProductId)
	if err != nil {
		http.Error(w, "failed to parse id to uuid", http.StatusBadRequest)
		return
	}
	customerInUUID, err := gocql.ParseUUID(order.Order.CustomerId)
	if err != nil {
		http.Error(w, "failed to parse id to uuid", http.StatusBadRequest)
		return
	}

	createdAt := helpers.ProtoToTime(order.Order.CreatedAt)
	updatedAt := helpers.ProtoToTime(order.Order.UpdatedAt)

	orderInJson := models.Order{
		ID:         orderInUUID,
		ProductID:  productInUUID,
		Quantity:   uint32(order.Order.Quantity),
		CustomerID: customerInUUID,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
	if err := helpers.ConvertStructToJson(w, http.StatusOK, orderInJson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (c *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	var ctx = r.Context()
	var id = chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Missing order id", http.StatusBadRequest)
		return

	}

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	createdAtInGRPCFormat := helpers.TimeToProto(order.CreatedAt)
	updatedAtInGRPCFormat := helpers.TimeToProto(order.UpdatedAt)

	updatedOrderReq := &pb.UpdateOrderRequest{
		OrderId: id,
		Order: &pb.Order{
			Id:         id,
			ProductId:  order.ProductID.String(),
			Quantity:   int32(order.Quantity),
			CustomerId: order.CustomerID.String(),
			CreatedAt:  createdAtInGRPCFormat,
			UpdatedAt:  updatedAtInGRPCFormat,
		},
	}

	updatedOrder, err := c.client.UpdateOrder(ctx, updatedOrderReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update order: %v", err), http.StatusInternalServerError)
		return
	}

	if updatedOrder.Success {
		err = helpers.ConvertStructToJson(w, http.StatusOK, updatedOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
