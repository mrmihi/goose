package v1

import (
	"goose/src/modules/users/api/v1/models"
	"goose/src/utils"
)

var repository = utils.NewRepository[models.User]("users")

func Repository() utils.Repository[models.User] {
	return repository
}
