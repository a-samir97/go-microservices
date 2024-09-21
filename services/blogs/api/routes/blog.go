package routes

import (
	"blogs/api/handlers"

	"github.com/gofiber/fiber/v2"
)

type BlogRoutes struct {
	handler handlers.BlogHandler
	app     *fiber.App
}

func NewBlogRoutes(app *fiber.App, handler handlers.BlogHandler) *BlogRoutes {
	return &BlogRoutes{handler: handler, app: app}
}

func (br *BlogRoutes) CreateRoutes() {
	r := br.app.Group("/api/users")

	r.Post("/", br.handler.CreateBlog)
}
