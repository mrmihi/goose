package auth

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/modules/auth/api/v1"
)

func New() *fiber.App {
	auth := fiber.New()
	auth.Mount("/v1", v1.New())
	return auth
}
