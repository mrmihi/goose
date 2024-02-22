package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateRestaurantReq struct {
	Name       string `validate:"required"`
	LocationID string `validate:"required"`
	Email      string `validate:"required,email"`
}

type CreateRestaurantRes struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	APIKey   string             `json:"apikey"`
}
