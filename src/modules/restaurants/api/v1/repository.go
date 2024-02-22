package v1

import (
	"goose/src/modules/restaurants/api/v1/models"
	"goose/src/utils"
)

var repository = utils.NewRepository[models.Restaurant]("restaurants")

func Repository() utils.Repository[models.Restaurant] {
	return repository
}
