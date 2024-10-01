package routes

import (
	"users/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, userHandler *handlers.UserHandler) {
	r := app.Group("/api/users")
	// r.Use(config.JwtVerify)

	r.Post("/signup/", userHandler.CreateUser)
	r.Post("/login/", userHandler.Login)
	r.Get("/:id/", userHandler.GetUser)
}
