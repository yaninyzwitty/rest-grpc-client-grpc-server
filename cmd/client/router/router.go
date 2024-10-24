package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/cmd/client/controller"
)

func NewRouter(productController *controller.ProductController, orderController *controller.OrderController, customerController *controller.CustomerController) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/products", func(r chi.Router) {
		r.Post("/", productController.CreateProduct)
		r.Get("/{id}", productController.GetProduct)
		r.Get("/", productController.ListProducts)
		r.Put("/", productController.UpdateProduct)

	})
	router.Route("/customers", func(r chi.Router) {
		r.Post("/", customerController.CreateCustomer)
		r.Get("/{id}", customerController.GetCustomer)
		r.Delete("/{id}", customerController.DeleteCustomer)
	})

	router.Route("/orders", func(r chi.Router) {
		r.Post("/", orderController.CreateOrder)
		r.Delete("/", orderController.DeleteOrder)
		r.Get("/", orderController.Getorder)
		r.Put("/", orderController.UpdateOrder)
	})

	return router

}