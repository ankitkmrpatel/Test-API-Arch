package main

import (
	"log"

	"test-api-arch/config"
	"test-api-arch/internal/delivery/http"
	"test-api-arch/pkg/db"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := db.Connect(cfg.Database)
	defer dbConn.Close()

	server := http.NewServer(cfg, dbConn)
	log.Println("Starting API server on port:", cfg.Port)
	server.Start()
}
