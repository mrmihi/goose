package v1

import (
	"github.com/gofiber/fiber/v2"
	"goose/src/global"
	"goose/src/modules/orders/api/v1/dto"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreateOrderReq)
	c.BodyParser(payload)
	res := createOrderService(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateOrderRes]{
		Message: "Orders created successfully",
		Data:    res,
	})
}
