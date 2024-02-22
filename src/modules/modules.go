package modules

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/middleware"
	"goose/src/modules/auth"
	"goose/src/modules/restaurants"
	"goose/src/modules/system"
	"goose/src/modules/users"
)

func New() *fiber.App {
	modules := fiber.New()

	modules.Mount("/auth", auth.New())

	modules.Mount("/system", system.New())

	modules.Mount("/users", users.New())

	modules.Mount("/restaurants", restaurants.New())

	modules.All("/*", middleware.Protect)

	return modules
}
