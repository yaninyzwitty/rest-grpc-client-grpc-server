package main

import (
	"log/slog"
	"os"

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

}
