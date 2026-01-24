package api

import (
	"github.com/stripe/stripe-go/v84"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"github.com/stripe/stripe-go/v84/checkout/session"
)

var AmoreLabsPrices = map[string]string{
	"voice_ai_italy_inbound": "price_1Ssl0sLOfzIWfxHPQmMX4JbI",
	"voice_ai_italy_outbound": "price_1StBmVLOfzIWfxHP7VXZsMbB",
	"voice_ai_italy_outbound-inbound": "price_1StBn4LOfzIWfxHPzOS1DXq8",
	"rinova_ai": "price_1StBnwLOfzIWfxHPdy0Ke7mH",
}

func CreateCheckoutSession(c *gin.Context) {
  product, ok := AmoreLabsPrices[c.Param("product")]
  if !ok {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Product doesn't exist"})
    return
  }
  
  stripe.Key = os.Getenv("STRIPE_API_KEY")
  domain := os.Getenv("STRIPE_DOMAIN")
  params := &stripe.CheckoutSessionParams{
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      &stripe.CheckoutSessionLineItemParams{
        // Provide the exact Price ID (for example, price_1234) of the product you want to sell
        Price: stripe.String(product),
        Quantity: stripe.Int64(1),
      },
    },
    Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
    SuccessURL: stripe.String(domain + "/success"),
    CancelURL: stripe.String(domain + "/failure"),
    AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
    CreateCustomers: %stripe.CheckoutSessionParams.CustomerCreation.ALWAYS,
    // Provide the Customer ID (for example, cus_1234) for an existing customer to associate it with this session
    // Customer: "cus_RnhPlBnbBbXapY",
  }

  s, err := session.New(params)

  if err != nil {
    log.Printf("session.New: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout session"})
    return
  }

  c.Redirect(http.StatusSeeOther, s.URL)
}