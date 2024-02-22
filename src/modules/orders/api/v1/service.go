package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"goose/src/modules/orders/api/v1/dto"
	"goose/src/modules/orders/api/v1/models"
)

func createOrderService(c *fiber.Ctx, payload dto.CreateOrderReq) *dto.CreateOrderRes {
	log.Info("Creating order within system - ")
	//squareup.CreateOrder()

	order := mapDTOToModel(payload)

	insertedID := repository.Create(order.WithDefaults())

	return &dto.CreateOrderRes{
		ID:   insertedID,
		Data: order,
	}
}

func mapDTOToModel(req dto.CreateOrderReq) models.Order {
	// Here you can map the fields from the DTO to the Order model
	// Note: You need to handle the list of items, discounts, and modifiers accordingly
	// You may need to loop through the lists and map each item individually

	order := models.Order{
		OpenedAt: req.OpenedAt,
		IsClosed: req.IsClosed,
		Table:    req.Table,
		// Map other fields as needed
	}

	// Map items
	for _, item := range req.Items {
		orderItem := models.OrderItem{
			Name:      item.Name,
			Comment:   item.Comment,
			UnitPrice: item.UnitPrice,
			Quantity:  item.Quantity,
			Amount:    item.Amount,
			// Map other fields as needed
		}

		// Map discounts for the item
		for _, discount := range item.Discounts {
			orderItem.Discounts = append(orderItem.Discounts, models.Discount{
				Name:         discount.Name,
				IsPercentage: discount.IsPercentage,
				Value:        discount.Value,
				Amount:       discount.Amount,
			})
		}

		// Map modifiers for the item
		for _, modifier := range item.Modifiers {
			orderItem.Modifiers = append(orderItem.Modifiers, models.Modifier{
				Name:      modifier.Name,
				UnitPrice: modifier.UnitPrice,
				Quantity:  modifier.Quantity,
				Amount:    modifier.Amount,
			})
		}

		order.Items = append(order.Items, orderItem)
	}

	return order
}
