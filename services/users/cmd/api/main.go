package main

import (
	"fmt"
	"users/api/handlers"
	"users/api/routes"
	"users/internal/repository"
	"users/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	userRepo := repository.NewSqlRepository()
	fmt.Println(userRepo.Db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(*userService)
	routes.NewRouter(app, userHandler)

	app.Listen(":8001")
}
