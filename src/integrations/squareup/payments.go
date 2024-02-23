package squareup

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"goose/src/modules/payments/api/v1/models"
)

func CreatePayment(payload models.Payment) {

	client := resty.New().SetDebug(true)
	fmt.Println("Payment Body", payload)

	_, _ = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetAuthToken("EAAAl0Yxj-imj10VhSAU88v_bc3YLb9Zsl1jDgUd7h8hJ1jnRXMdP4SFyWLNoxOG").
		Post("https://connect.squareupsandbox.com/v2/payments")
}
