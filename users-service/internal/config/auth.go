package config

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtVerifyMiddlware(c *fiber.Ctx) error {
	var header string = c.Get("Authorization")
	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "No token provided",
		})
	}
	tokenString := strings.Split(header, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	return c.Next()
}
