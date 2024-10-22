package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/yaninyzwitty/golang-rest-grpc-proj/internal/database"
	"github.com/yaninyzwitty/golang-rest-grpc-proj/pkg"
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
	defer db.Close()

}
