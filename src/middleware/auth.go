package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goose/src/modules/restaurants/api/v1/repository"
)

func RestaurantProtect(ctx *fiber.Ctx) error {
	token := ctx.Get("X-API-Key")
	if token == "" {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Missing API Key"))
	}
	restaurant := repository.GetInstance().FindOne(primitive.M{"apikey": token})
	if restaurant == nil {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid API Key"))
	}
	ctx.Locals("restaurant", restaurant)
	return ctx.Next()
}
