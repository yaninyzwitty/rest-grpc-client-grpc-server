package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/cmd/client/models"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/helpers"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
)

type ProductController struct {
	client pb.ProductServiceClient
}

func NewProductController(client pb.ProductServiceClient) *ProductController {
	return &ProductController{client: client}
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	var ctx = r.Context()
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createProductReq := &pb.CreateProductRequest{
		Name:        product.Name,
		Description: product.Description,
		Price:       float64(product.Price),
		Stock:       int32(product.Stock),
		Category:    product.Category,
	}

	createdProduct, err := c.client.CreateProduct(ctx, createProductReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create product: %v", err), http.StatusBadRequest)
		return

	}

	productIDInUUID, err := gocql.ParseUUID(createdProduct.Product.Id)
	if err != nil {
		http.Error(w, "failed to parse id into uuid", http.StatusBadRequest)
		return

	}

	createdAt := helpers.ProtoToTime(createdProduct.Product.CreatedAt)
	updatedAt := helpers.ProtoToTime(createdProduct.Product.UpdatedAt)

	createProductInJson := &models.Product{
		ID:          productIDInUUID,
		Name:        createProductReq.Name,
		Description: createProductReq.Description,
		Price:       uint32(createdProduct.Product.Price),
		Stock:       uint32(createdProduct.Product.Stock),
		Category:    createProductReq.Category,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	if err := helpers.ConvertStructToJson(w, http.StatusCreated, createProductInJson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

}
func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	var productId = chi.URLParam(r, "id")
	var category = r.URL.Query().Get("category")
	if productId == "" || category == "" {
		http.Error(w, "Both category and product id are missing", http.StatusBadRequest)
		return

	}
	var ctx = r.Context()

	productReq := &pb.GetProductRequest{
		Category:  category,
		ProductId: productId,
	}
	product, err := c.client.GetProduct(ctx, productReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	productIDInUUID, err := gocql.ParseUUID(product.Product.Id)
	if err != nil {
		http.Error(w, "failed to parse id into uuid", http.StatusBadRequest)
		return

	}

	createdAt := helpers.ProtoToTime(product.Product.CreatedAt)
	updatedAt := helpers.ProtoToTime(product.Product.UpdatedAt)

	productInJson := models.Product{
		ID:          productIDInUUID,
		Name:        product.Product.Name,
		Description: product.Product.Description,
		Price:       uint32(product.Product.Price),
		Stock:       uint32(product.Product.Stock),
		Category:    product.Product.Category,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	if err := helpers.ConvertStructToJson(w, http.StatusOK, productInJson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

}
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var productId = chi.URLParam(r, "id")
	var category = r.URL.Query().Get("category")
	if productId == "" || category == "" {
		http.Error(w, "Both category and product id are missing", http.StatusBadRequest)
		return

	}
	var ctx = r.Context()

	deletedProductRes, err := c.client.DeleteProduct(ctx, &pb.DeleteProductRequest{
		Category:  category,
		ProductId: productId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	if deletedProductRes.Success {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, deletedProductRes.Message, http.StatusInternalServerError)

}
func (c *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	category := r.URL.Query().Get("category")
	limitStr := r.URL.Query().Get("limit")
	pagingState := r.URL.Query().Get("paging_state")

	// Validate category
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	// Set default limit if not provided
	if limitStr == "" {
		limitStr = "10"
	}

	// Parse limit as an integer
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid limit value: %v", err), http.StatusBadRequest)
		return
	}

	// Create context from request
	ctx := r.Context()

	// Call the ListProducts RPC
	productsResponse, err := c.client.ListProducts(ctx, &pb.ListProductsRequest{
		Limit:       int32(limit),
		PagingState: []byte(pagingState),
		Category:    category,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list products: %v", err), http.StatusInternalServerError)
		return
	}

	// Initialize an array of products for JSON response
	var productsInJson []models.Product

	// Loop through each product in the response and convert to your model
	for _, product := range productsResponse.Products {
		// Parse UUID for the product ID
		productIDInUUID, err := gocql.ParseUUID(product.Id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse product ID into UUID: %v", err), http.StatusBadRequest)
			return
		}

		// Convert timestamps
		createdAt := helpers.ProtoToTime(product.CreatedAt)
		updatedAt := helpers.ProtoToTime(product.UpdatedAt)

		// Map to your product model
		productsInJson = append(productsInJson, models.Product{
			ID:        productIDInUUID,
			Name:      product.Name,
			Category:  product.Category,
			Price:     uint32(product.Price),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	// Send the products list as a JSON response
	if err := helpers.ConvertStructToJson(w, http.StatusOK, productsInJson); err != nil {
		http.Error(w, fmt.Sprintf("Failed to serialize products: %v", err), http.StatusInternalServerError)
		return
	}
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL
	productID := chi.URLParam(r, "id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// Decode the request body to get the updated product details
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Convert timestamps from the product model
	createdAtProto := helpers.TimeToProto(product.CreatedAt)
	updatedAtProto := helpers.TimeToProto(product.UpdatedAt)

	// Prepare the gRPC request for updating the product
	updateProductReq := &pb.UpdateProductRequest{
		ProductId: productID,
		Product: &pb.Product{
			Id:        productID, // Keep the original product ID
			Name:      product.Name,
			Category:  product.Category,
			Price:     float64(product.Price),
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		},
	}

	// Get the context from the request
	ctx := r.Context()

	// Call the gRPC client's UpdateProduct method
	updatedProduct, err := c.client.UpdateProduct(ctx, updateProductReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update product: %v", err), http.StatusInternalServerError)
		return
	}

	if updatedProduct.Success {

		if err := helpers.ConvertStructToJson(w, http.StatusOK, "Update product succcesfully"); err != nil {
			http.Error(w, fmt.Sprintf("Failed to serialize response: %v", err), http.StatusInternalServerError)
			return
		}
	}
	return

}
