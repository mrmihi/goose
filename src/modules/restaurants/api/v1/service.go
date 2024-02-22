package v1

import (
	"goose/src/modules/restaurants/api/v1/dto"
	"goose/src/modules/restaurants/api/v1/models"
	"goose/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sethvargo/go-password/password"
)

func createRestaurant(c *fiber.Ctx, payload dto.CreateRestaurantReq) *dto.CreateRestaurantRes {
	log.Info("Creating restaurant within system - ", payload.Email)
	generatedPassword, _ := password.Generate(8, 2, 1, false, false)
	insertedID := repository.Create(models.Restaurant{
		Email:    payload.Email,
		Name:     payload.Name,
		Password: utils.HashStr(generatedPassword),
	}.WithDefaults())
	return &dto.CreateRestaurantRes{
		ID:       insertedID,
		Password: generatedPassword,
	}
}
