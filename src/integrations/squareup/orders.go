package squareup

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"goose/src/integrations/squareup/dto"
	"goose/src/modules/orders/api/v1/models"
)

// Order represents the structure of the order to be created
type Order struct {
	IdempotencyKey string    `json:"idempotency_key"`
	Order          OrderBody `json:"order"`
}

func (order *Order) fromModel(m models.Order) {

	order.IdempotencyKey = uuid.NewString()
	order.Order.LocationID = "L64JNY26EYBXF"

	// Map LineItems
	for _, item := range m.Items {
		orderItem := OrderLineItem{
			Quantity:       item.Quantity,
			BasePriceMoney: Money{Amount: item.UnitPrice, Currency: "USD"},
			Name:           item.Name,
		}

		// Map Modifiers
		for _, modifier := range item.Modifiers {
			orderItem.Modifiers = append(orderItem.Modifiers, Modifier{
				BasePriceMoney: Money{Amount: modifier.UnitPrice, Currency: "USD"},
				Quantity:       modifier.Quantity,
				Name:           modifier.Name,
			})
		}

		order.Order.LineItems = append(order.Order.LineItems, orderItem)
	}

	// Map Discounts
	//for _, discount := range m.Totals.Discounts {
	//	orderDiscount := OrderLineItemDiscount{
	//		AmountMoney:  Money{Amount: discount.Amount, Currency: "USD"}, // Assuming discounts are in USD
	//		AppliedMoney: Money{Amount: discount.Amount, Currency: "USD"}, // Assuming applied money is the same as amount
	//	}
	//	order.Order.Discounts = append(order.Order.Discounts, orderDiscount)
	//}

	// Map Taxes
	//for _, tax := range m.Totals.Tax {
	//	orderTax := OrderLineItemTax{
	//		Name:       tax.Name,
	//		Percentage: fmt.Sprintf("%d", tax.Percentage), // Assuming tax percentage is an integer
	//	}
	//	order.Order.Taxes = append(order.Order.Taxes, orderTax)
	//}

	// Map ServiceCharges
	//for _, serviceCharge := range m.Totals.ServiceCharge {
	//	orderServiceCharge := OrderServiceCharge{
	//		AmountMoney: Money{Amount: serviceCharge.Amount, Currency: "USD"}, // Assuming service charges are in USD
	//		Type:        serviceCharge.Type,                                   // Assuming Type field exists in models.ServiceCharge
	//	}
	//	order.Order.ServiceCharges = append(order.Order.ServiceCharges, orderServiceCharge)
	//}

	// Map other fields
	order.Order.PricingOptions = OrderPricingOptions{
		AutoApplyDiscounts: true, // Adjust accordingly based on your requirements
		AutoApplyTaxes:     true, // Adjust accordingly based on your requirements
	}
	//order.Order.ReferenceID = m.ID.Hex() // Assuming ID in models.Order is a primitive.ObjectID
	order.Order.State = "OPEN" // Assuming order is always open when created

	//Need to set other fields as well

	order.Order.Table = m.Table

	//	Mapping Sources
	order.Order.Source.Name = m.Table
}

type OrderSource struct {
	Name string `json:"name"`
}

type OrderBody struct {
	LocationID string `json:"location_id"`
	//CustomerID     string                  `json:"customer_id"`
	Discounts      []OrderLineItemDiscount `json:"discounts"`
	LineItems      []OrderLineItem         `json:"line_items"`
	PricingOptions OrderPricingOptions     `json:"pricing_options"`
	ReferenceID    string                  `json:"reference_id"`
	ServiceCharges []OrderServiceCharge    `json:"service_charges"`
	Source         OrderSource             `json:"source"`
	State          string                  `json:"state"`
	Taxes          []OrderLineItemTax      `json:"taxes"`
	Table          string                  `json:"ticket_name"`
}

type OrderLineItemTax struct {
}

type OrderServiceCharge struct {
}

type OrderPricingOptions struct {
	AutoApplyDiscounts bool `json:"auto_apply_discounts"`
	AutoApplyTaxes     bool `json:"auto_apply_taxes"`
}

type OrderLineItemDiscount struct {
	AmountMoney  Money `json:"amount_money"`
	AppliedMoney Money `json:"applied_money"`
}

type OrderLineItem struct {
	Quantity       string     `json:"quantity"`
	BasePriceMoney Money      `json:"base_price_money"`
	Name           string     `json:"name"`
	Modifiers      []Modifier `json:"modifiers,omitempty"`
}

type Modifier struct {
	BasePriceMoney Money  `json:"base_price_money"`
	Quantity       string `json:"quantity"`
	Name           string `json:"name"`
}

type Money struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

func CreateOrder(payload models.Order) string {
	var order Order
	order.fromModel(payload)
	requestBody, _ := json.Marshal(order)
	resp := dto.CreateSquareOrderRes{}
	getClient().R().
		SetBody(requestBody).
		SetResult(&resp).
		Post("https://connect.squareupsandbox.com/v2/orders")
	SqureUpOrderID := resp.Order.ID
	return SqureUpOrderID
}

func GetOrderByTable(LocationID string, tableID string) models.Order {

	query := map[string]any{
		"return_entries": false,
		"query": map[string]any{
			"filter": map[string]any{
				"source_filter": map[string]any{
					"source_names": []string{
						tableID,
					},
				},
			},
		},
		"location_ids": []string{
			LocationID,
		},
		"limit": 1,
	}
	resp := dto.GetSquareOrderByTableRes{}
	getClient().R().
		SetResult(&resp).
		SetBody(query).
		Post("/orders/search")

	if len(resp.Orders) == 0 {
		panic(fiber.NewError(fiber.StatusNotFound, "No Orders Found for the table provided"))
	}
	return resp.Orders[0].ToOrder()
}

func GetOrderById(orderID string) models.Order {
	resp := dto.CreateSquareOrderRes{}
	getClient().R().
		SetResult(&resp).
		SetError(&resp).
		Get("/orders/" + orderID)
	if len(resp.Errors) > 0 {
		panic(fiber.NewError(fiber.StatusNotFound, "No Order Found for the ID provided"))
	}
	return resp.Order.ToOrder()
}
