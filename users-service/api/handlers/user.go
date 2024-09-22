package handlers

import (
	"users/api/dto"
	"users/internal/domain"
	"users/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {

	return &UserHandler{service: service}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userRequest dto.CreateUserRequest

	if err := c.BodyParser(userRequest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: hash password
	user := domain.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Avatar:    userRequest.Avatar,
		Password:  userRequest.Password,
	}

	if _, err := uh.service.CreateUser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// parsing with user response struct
	var userResponse dto.UserResponse
	if err := c.BodyParser(userResponse); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusCreated).JSON(userResponse)
}
