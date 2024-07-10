package main

import (
	_ "gopkg.in/yaml.v2"
	"log"
	"net"
	"net/http"
	"time"
	"tradeMetricsCollector/internal/config"
	"tradeMetricsCollector/pkg/client/postgreSQL"
)

func main() {
	cfg := config.LoadConfiguration("config.yaml")
	db, err := postgreSQL.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	mux := http.NewServeMux()
	start(mux, cfg)
}

func start(router *http.ServeMux, cfg *config.Config) {
	log.Println("Start the application...")

	listener, err := net.Listen(cfg.Server.Network, cfg.Server.Address)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
	}

	log.Printf("Server is listening port %s\n", cfg.Server.Address)
	log.Fatal(server.Serve(listener))
}
