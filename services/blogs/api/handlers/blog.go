package handlers

import (
	"blogs/internal/services"

	"github.com/gofiber/fiber/v2"
)

type BlogHandler struct {
	service *services.BlogService
}

func NewBlogHandler(service services.BlogService) *BlogHandler {
	return &BlogHandler{service: &service}
}

func (bh *BlogHandler) CreateBlog(ctx *fiber.Ctx) error {

	return ctx.SendString("Test")
}
