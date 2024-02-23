package orders

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/modules/orders/api/v1"
)

func New() *fiber.App {
	orders := fiber.New()
	orders.Mount("/v1", v1.New())
	return orders
}
