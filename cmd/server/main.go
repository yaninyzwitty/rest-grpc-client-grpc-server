package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/controller"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/database"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/repository"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/services"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pb"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pkg"
	"google.golang.org/grpc"
)

func main() {
	var cfg pkg.Config
	file, err := os.Open("config.yaml")
	if err != nil {
		slog.Error("Failed to open the config: ", "error", err)
		os.Exit(1)
	}

	defer file.Close()

	if err := cfg.LoadConfig(file); err != nil {
		slog.Error("Error loading config: ", "error", err)
		os.Exit(1)
	}
	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	databaseCFG := database.DatabaseConfig{
		Path:     cfg.Database.Path,
		Username: cfg.Database.Username,
		Password: cfg.Database.Passsword,
		Timeout:  20 * time.Second,
	}
	db, err := database.NewDatabaseConnection(databaseCFG)
	if err != nil {
		slog.Error("Error connecting to astra", "error", err)
		os.Exit(1)
	}
	slog.Info("Connected to ASTRA succesfully🚀")

	defer db.Close()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		os.Exit(1)
	}
	// set up injects 💉

	// products
	productdRepo := repository.NewProductRepository(db)
	productsService := services.NewProductService(productdRepo)
	productController := controller.NewProductController(productsService)

	// orders
	ordersRepo := repository.NewOrderRepository(db)
	ordersService := services.NewOrderService(ordersRepo)
	orderController := controller.NewOrderController(ordersService)

	// customers
	customersRepo := repository.NewCustomerRepository(db)
	customersService := services.NewCustomerService(customersRepo)
	customersController := controller.NewCustomerController(customersService)

	server := grpc.NewServer()
	pb.RegisterCustomerServiceServer(server, customersController)
	pb.RegisterOrderServiceServer(server, orderController)
	pb.RegisterProductServiceServer(server, productController)

	slog.Info("Server is listening", "address", lis.Addr().String())

	// handle graceful stop, signals etc.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		slog.Info("Received shutdown signal", "signal", sig)
		slog.Info("Shutting down gRPC server...")

		// Gracefully stop the gRPC server
		server.GracefulStop()
		cancel()

		slog.Info("gRPC server has been stopped gracefully")
	}()

	slog.Info("Starting gRPC server", "port", 50051)
	if err := server.Serve(lis); err != nil {
		slog.Error("gRPC server encountered an error while serving", "error", err)
		os.Exit(1)
	}

}
