package main

import (
	"fmt"
	"log"
	"time"
	"users/api/handlers"
	"users/api/routes"
	"users/internal/repository"
	"users/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var totalRequest = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_total_request_users",
		Help: "Total number of requests",
	},
	[]string{"method", "path"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_status_users",
		Help: "Total number of response status",
	},
	[]string{"status"},
)

var httpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds_users",
		Help: "HTTP request duration in seconds",
	},
	[]string{"method", "path"},
)

func prometheusMiddleware(c *fiber.Ctx) error {
	totalRequest.WithLabelValues(c.Method(), c.Path()).Inc()
	start := time.Now()
	err := c.Next()
	duration := time.Since(start).Seconds()
	responseStatus.WithLabelValues(string(c.Response().StatusCode())).Inc()
	httpDuration.WithLabelValues(c.Method(), c.Path()).Observe(duration)
	log.Println(err)
	return err
}

func init() {
	prometheus.MustRegister(totalRequest)
	prometheus.MustRegister(responseStatus)
	prometheus.MustRegister(httpDuration)
}

func main() {
	app := fiber.New()
	userRepo := repository.NewORMRepository()
	fmt.Println(userRepo.Db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(*userService)
	routes.NewRouter(app, userHandler)
	app.Use(prometheusMiddleware)

	app.Get("/metrics/", adaptor.HTTPHandler(promhttp.Handler()))

	app.Get("/health/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Listen(":8001")
}
