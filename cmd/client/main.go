package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/cmd/client/controller"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/cmd/client/router"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var cfg pkg.Config
	file, err := os.Open("config.yaml")
	if err != nil {
		slog.Error("Failed to open  config.yaml file")
		os.Exit(1)
	}

	defer file.Close()
	if err := cfg.LoadConfig(file); err != nil {
		slog.Error("Error loading config.yaml", "error", err)
		os.Exit(1)
	}

	address := fmt.Sprintf(":%d", 50051)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to create a grpc conn", "error", err)
		os.Exit(1)

	}
	defer conn.Close()

	orderController := controller.NewOrderController(conn)
	customerController := controller.NewCustomerController(conn)
	productController := controller.NewProductController(conn)
	mux := router.NewRouter(productController, orderController, customerController)
	server := &http.Server{
		Addr:    ":" + fmt.Sprintf("%d", cfg.Server.PORT),
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			return
		}

	}()
	slog.Info("Server is running at port", "port", cfg.Server.PORT)

	quitCH := make(chan os.Signal, 1)
	signal.Notify(quitCH, os.Interrupt)

	<-quitCH

	slog.Info("Received termination signal, shutting down server...")
	shutdownCTX, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(shutdownCTX); err != nil {
		slog.Error("Failed to gracefully shut down server", "error", err)
	}
	slog.Info("Server shutdown successful")

}
