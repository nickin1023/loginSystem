package main

import (
	"backend-app/config"
	"backend-app/router"
	"backend-app/utils/handler"
	"log/slog"
	"os"
)

func main() {
	// Load Config
	config, err := config.Init()
	if err != nil {
		slog.Error("Config Loading error:", err)
		os.Exit(1)
	}

	// Establish Database Connection
	db, err := handler.OpenDBConnection(config)
	if err != nil {
		slog.Error("DB connection error:", err)
		os.Exit(1)
	}
	defer db.Close()

	router := router.GetRouter(db)
	router.Run(":8000")
}
