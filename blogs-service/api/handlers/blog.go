package handlers

import (
	"blogs/api/dto"
	"blogs/internal/domain"
	"blogs/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BlogHandler struct {
	service *services.BlogService
}

func NewBlogHandler(service services.BlogService) *BlogHandler {
	return &BlogHandler{service: &service}
}

func (bh *BlogHandler) CreateBlog(ctx *fiber.Ctx) error {
	var blogRequest dto.BlogRequest

	if err := ctx.BodyParser(&blogRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	blog := domain.Blog{
		Title:       blogRequest.Title,
		Description: blogRequest.Description,
		UserId:      blogRequest.UserId,
		Likes:       0,
		Dislikes:    0,
		Claps:       0,
	}

	if _, err := bh.service.Create(blog); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var blogResponse dto.BlogResponse

	if err := ctx.BodyParser(&blogResponse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(blogResponse)
}

func (bh *BlogHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	blog, err := bh.service.GetById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(blog)
}

func (bh *BlogHandler) List(c *fiber.Ctx) error {
	blogs, err := bh.service.List()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(blogs)
}

func (bh *BlogHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = bh.service.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (bh *BlogHandler) LikeBlog(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	blog, err := bh.service.LikeBlog(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(blog)
}

func (bh *BlogHandler) DislikeBlog(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	blog, err := bh.service.DislikeBlog(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(blog)
}

func (bh *BlogHandler) ClapBlog(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	blog, err := bh.service.ClapBlog(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(blog)
}
