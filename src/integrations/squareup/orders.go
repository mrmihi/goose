package squareup

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// Order represents the structure of the order to be created
type Order struct {
	IdempotencyKey string    `json:"idempotency_key"`
	Order          OrderBody `json:"order"`
}

// OrderBody represents the body of the order
type OrderBody struct {
	LocationID     string                  `json:"location_id"`
	CustomerID     string                  `json:"customer_id"`
	Discounts      []OrderLineItemDiscount `json:"discounts"`
	LineItems      []OrderItem             `json:"line_items"`
	PricingOptions OrderPricingOptions
	ReferenceID    string `json:"reference_id"`
	ServiceCharges []OrderServiceCharge
	State          string `json:"state"`
	Taxes          []OrderLineItemTax
	Table          string `json:"ticket_name"`
}

type OrderLineItemTax struct {
}

type OrderServiceCharge struct {
}

type OrderPricingOptions struct {
}

type OrderLineItemDiscount struct {
}

// OrderItem represents an item in the order
type OrderItem struct {
	Quantity       string     `json:"quantity"`
	BasePriceMoney Money      `json:"base_price_money"`
	Name           string     `json:"name"`
	Modifiers      []Modifier `json:"modifiers,omitempty"`
}

// Modifier represents a modifier for an item
type Modifier struct {
	BasePriceMoney Money  `json:"base_price_money"`
	Quantity       string `json:"quantity"`
	Name           string `json:"name"`
}

// Money represents the amount and currency
type Money struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

func CreateOrder() {
	// Create a new order instance
	order := Order{
		IdempotencyKey: "unique-idempotency-key",
		Order: OrderBody{
			LocationID: "L64JNY26EYBXF",
			LineItems: []OrderItem{
				{
					Quantity: "1",
					BasePriceMoney: Money{
						Amount:   1200,
						Currency: "USD",
					},
					Name: "Burger",
					Modifiers: []Modifier{
						{
							BasePriceMoney: Money{
								Amount:   50,
								Currency: "USD",
							},
							Quantity: "1",
							Name:     "extra cheese",
						},
					},
				},
			},
		},
	}

	// Convert order instance to JSON
	requestBody, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error marshaling order:", err)
		return
	}
	fmt.Println(requestBody)

	// Create a Resty Client
	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		SetAuthToken("EAAAl0Yxj-imj10VhSAU88v_bc3YLb9Zsl1jDgUd7h8hJ1jnRXMdP4SFyWLNoxOG").
		Post("https://connect.squareupsandbox.com/v2/orders")
	fmt.Println(resp)
}
