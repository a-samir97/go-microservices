package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// ServiceConfig holds the configuration for the API Gateway service
type ServiceConfig struct {
	Port int
}

// LoadServiceConfig loads the configuration for the API Gateway service
func LoadServiceConfig() *ServiceConfig {
	port := 8080
	if os.Getenv("PORT") != "" {
		port = 8080
	}

	return &ServiceConfig{
		Port: port,
	}
}

// StartAPIGatewayService starts the API Gateway service
func StartAPIGatewayService() {
	serviceConfig := LoadServiceConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the API Gateway Service!")
	})

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", serviceConfig.Port),
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting API Gateway Service on port %d", serviceConfig.Port)
	log.Fatal(server.ListenAndServe())
}
