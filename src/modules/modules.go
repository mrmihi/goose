package modules

import (
	"goose/src/middleware"
	"goose/src/modules/auth"
	"goose/src/modules/system"
	"goose/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	modules := fiber.New()

	modules.Mount("/auth", auth.New())

	modules.Mount("/system", system.New())

	modules.All("/*", middleware.Protect)

	modules.Mount("/users", users.New())

	return modules
}
