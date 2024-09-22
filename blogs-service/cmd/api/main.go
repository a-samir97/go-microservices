package main

import (
	"blogs/api/handlers"
	"blogs/api/routes"
	"blogs/internal/repository"
	"blogs/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	blogRepo := repository.NewSqlRepository()
	blogService := services.NewBlogService(blogRepo)
	blogHandler := handlers.NewBlogHandler(blogService)
	blogRoutes := routes.NewBlogRoutes(app, *blogHandler)
	blogRoutes.CreateRoutes()
	app.Listen(":8000")
}
