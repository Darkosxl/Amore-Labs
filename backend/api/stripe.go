package api

import (
	mw "amorelabs/backend/middleware"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/checkout/session"
	"github.com/workos/workos-go/v6/pkg/usermanagement"
)

var AmoreLabsPrices = map[string]string{
	"voice_ai_italy_inbound":          "price_1Ssl0sLOfzIWfxHPQmMX4JbI",
	"voice_ai_italy_outbound":         "price_1StBmVLOfzIWfxHP7VXZsMbB",
	"voice_ai_italy_outbound-inbound": "price_1StBn4LOfzIWfxHPzOS1DXq8",
	"rinova_ai":                       "price_1StBnwLOfzIWfxHPdy0Ke7mH",
	"test_product":                    "price_1StqUxLOfzIWfxHPQva2nD7l", // Test product for bscemarslan@gmail.com
}

func CreateCheckoutSession(c *gin.Context) {
	product, ok := AmoreLabsPrices[c.Param("product")]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product doesn't exist"})
		return
	}

	// Get authenticated user from middleware
	claimsCheck, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "error": "Unauthorized"})
		return
	}
	claims, ok := claimsCheck.(*mw.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Invalid claims type"})
		return
	}

	// Get user details from WorkOS
	ctx := context.Background()
	user, err := usermanagement.GetUser(ctx, usermanagement.GetUserOpts{
		User: claims.UserID,
	})
	if err != nil {
		log.Printf("Failed to get user from WorkOS: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user details"})
		return
	}

	stripe.Key = os.Getenv("STRIPE_API_KEY")
	
	// Use test API key for test products
	productName := c.Param("product")
	if productName == "test_product" {
		stripe.Key = os.Getenv("STRIPE_TEST_API_KEY")
	}
	
	// Check if free trial is enabled
	freeTrialEnabled := c.Query("free_trial") == "true"
	
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(product),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(os.Getenv("FRONTEND_URL") + "/#/payment_success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(os.Getenv("FRONTEND_URL") + "/#/admin_console?payment_failed=true"),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
		CustomerEmail: stripe.String(user.Email),
		Metadata: map[string]string{
			"workos_user_id": user.ID,
			"user_email":     user.Email,
		},
	}

	// Add free trial configuration if enabled
	if freeTrialEnabled {
		params.SubscriptionData = &stripe.CheckoutSessionSubscriptionDataParams{
			TrialPeriodDays: stripe.Int64(30), // 1 month trial
		}
		// Ensure payment method is collected during trial
		params.PaymentMethodCollection = stripe.String("always")
		
		log.Printf("Creating checkout session with 30-day free trial for user: %s", user.Email)
	} else {
		log.Printf("Creating immediate checkout session for user: %s", user.Email)
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout session"})
		return
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}