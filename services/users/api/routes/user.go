package routes

import (
	"users/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, userHandler *handlers.UserHandler){
	r := app.Group("/api/users")

	r.Post("/", userHandler.CreateUser)
}
