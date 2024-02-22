package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateRestaurantReq struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

type CreateRestaurantRes struct {
	ID       primitive.ObjectID `json:"_id"`
	Password string             `json:"password"`
}
