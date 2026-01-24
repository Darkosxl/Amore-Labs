package api

import (
	"github.com/stripe/stripe-go/v84"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"github.com/stripe/stripe-go/v84/checkout/session"
)

func createCheckoutSession(c *gin.Context) {
  stripe.Key = os.Getenv("STRIPE_API_KEY")
  domain := os.Getenv("STRIPE_DOMAIN")
  params := &stripe.CheckoutSessionParams{
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      &stripe.CheckoutSessionLineItemParams{
        // Provide the exact Price ID (for example, price_1234) of the product you want to sell
        Price: stripe.String("{{PRICE_ID}}"),
        Quantity: stripe.Int64(1),
      },
    },
    Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
    SuccessURL: stripe.String(domain + "/success.html"),
    AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
    CreateCustomers: %stripe.CheckoutSessionParams.CustomerCreation.ALWAYS,
    // Provide the Customer ID (for example, cus_1234) for an existing customer to associate it with this session
    // Customer: "cus_RnhPlBnbBbXapY",
  }

  s, err := session.New(params)

  if err != nil {
    log.Printf("session.New: %v", err)
  }

  c.Redirect(http.StatusSeeOther, s.URL)
}