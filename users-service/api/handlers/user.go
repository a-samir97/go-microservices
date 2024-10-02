package handlers

import (
	"log"
	"strconv"
	"users/api/dto"
	"users/internal/domain"
	"users/internal/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {

	return &UserHandler{service: service}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userRequest dto.CreateUserRequest

	if err := c.BodyParser(&userRequest); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: hash password
	user := domain.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Avatar:    userRequest.Avatar,
		Password:  string(pass),
	}

	if _, err := uh.service.CreateUser(user); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// parsing with user response struct
	var userResponse dto.UserResponse
	if err := c.BodyParser(&userResponse); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusCreated).JSON(userResponse)
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	var userRequest dto.LoginRequest
	if err := c.BodyParser(&userRequest); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := uh.service.Login(userRequest.Email, userRequest.Password)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user, err := uh.service.FindByID(idInt)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var userResponse dto.UserResponse
	userResponse.FirstName = user.FirstName
	userResponse.LastName = user.LastName
	userResponse.Email = user.Email
	userResponse.Avatar = user.Avatar

	return c.Status(fiber.StatusOK).JSON(userResponse)
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var userRequest dto.CreateUserRequest
	if err := c.BodyParser(&userRequest); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := uh.service.FindByID(idInt)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.Email = userRequest.Email
	user.Avatar = userRequest.Avatar

	if _, err := uh.service.UpdateUser(*user, idInt); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var userResponse dto.UserResponse
	userResponse.FirstName = user.FirstName
	userResponse.LastName = user.LastName
	userResponse.Email = user.Email
	userResponse.Avatar = user.Avatar

	return c.Status(fiber.StatusOK).JSON(userResponse)
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uh.service.DeleteUser(id); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
