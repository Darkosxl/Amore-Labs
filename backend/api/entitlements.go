package api

import (
	mw "amorelabs/backend/middleware"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/customer"
	"github.com/stripe/stripe-go/v84/subscription"
)

// SubscriptionResponse represents a subscription for API response
type SubscriptionResponse struct {
	ID               string   `json:"id"`
	Status           string   `json:"status"`
	ProductName      string   `json:"product_name"`
	PriceID          string   `json:"price_id"`
	CurrentPeriodEnd int64    `json:"current_period_end"`
	CancelAtPeriodEnd bool    `json:"cancel_at_period_end"`
}

// EntitlementsResponse represents entitlements from WorkOS access token
type EntitlementsResponse struct {
	Entitlements map[string]interface{} `json:"entitlements"`
	UserID       string                  `json:"user_id"`
	Email        string                  `json:"email"`
}

// GetSubscriptions retrieves active subscriptions for the user
func GetSubscriptions(c *gin.Context) {
	// Get user claims from middleware
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

	stripe.Key = os.Getenv("STRIPE_API_KEY")

	// Search for Stripe customer by email
	customerParams := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: fmt.Sprintf("email:'%s'", claims.Email),
		},
	}
	customerResults := customer.Search(customerParams)
	
	var stripeCustomerID string
	if customerResults.Next() {
		stripeCustomerID = customerResults.Customer().ID
	} else {
		// No customer found
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"subscriptions": []SubscriptionResponse{},
			"message": "No Stripe customer found for this user",
		})
		return
	}

	// Get subscriptions for this customer
	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(stripeCustomerID),
	}
	params.AddExpand("data.items.data.price.product")

	i := subscription.List(params)
	var subscriptions []SubscriptionResponse

	// Reverse map price IDs to product names
	priceToProduct := make(map[string]string)
	for productName, priceID := range AmoreLabsPrices {
		priceToProduct[priceID] = productName
	}

	for i.Next() {
		sub := i.Subscription()
		
		// Get the first line item (assuming one product per subscription)
		if len(sub.Items.Data) > 0 {
			priceID := sub.Items.Data[0].Price.ID
			productName := priceToProduct[priceID]
			if productName == "" {
				productName = "Unknown Product"
			}

			// Calculate current period end
			// In Stripe v84, we can use TrialEnd or calculate from billing cycle
			var currentPeriodEnd int64
			if sub.TrialEnd > 0 {
				currentPeriodEnd = sub.TrialEnd
			} else {
				// Use created date + billing interval as approximation
				// A better approach would be to look at the latest invoice
				currentPeriodEnd = sub.Created + (30 * 24 * 60 * 60) // 30 days default
			}

			subscriptions = append(subscriptions, SubscriptionResponse{
				ID:               sub.ID,
				Status:           string(sub.Status),
				ProductName:      productName,
				PriceID:          priceID,
				CurrentPeriodEnd: currentPeriodEnd,
				CancelAtPeriodEnd: sub.CancelAtPeriodEnd,
			})
		}
	}

	if err := i.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Failed to list subscriptions", "details": err.Error()})
		return
	}

	log.Printf("Found %d subscriptions for user %s (customer: %s)", len(subscriptions), claims.Email, stripeCustomerID)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"subscriptions": subscriptions,
		"user_id": claims.UserID,
		"stripe_customer_id": stripeCustomerID,
	})
}

// GetEntitlements extracts entitlements from the user's WorkOS access token
func GetEntitlements(c *gin.Context) {
	// Get user claims from middleware
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

	// Get the raw access token from the Authorization header or cookie
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		// Try to get from cookie
		accessToken, _ = c.Cookie("access_token")
	} else {
		// Remove "Bearer " prefix if present
		if len(accessToken) > 7 && accessToken[:7] == "Bearer " {
			accessToken = accessToken[7:]
		}
	}

	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "error": "No access token found"})
		return
	}

	// Parse the JWT without verification to extract the entitlements claim
	// (we already verified it in the middleware)
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Failed to parse token", "details": err.Error()})
		return
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Invalid token claims"})
		return
	}

	// Extract entitlements if they exist
	var entitlements map[string]interface{}
	if ent, exists := mapClaims["entitlements"]; exists {
		// Handle entitlements as either map or JSON string
		switch v := ent.(type) {
		case map[string]interface{}:
			entitlements = v
		case string:
			// Try to parse as JSON string
			if err := json.Unmarshal([]byte(v), &entitlements); err != nil {
				entitlements = make(map[string]interface{})
				entitlements["raw"] = v
			}
		default:
			entitlements = make(map[string]interface{})
			entitlements["raw"] = fmt.Sprintf("%v", v)
		}
	} else {
		entitlements = make(map[string]interface{})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"entitlements": entitlements,
		"user_id": claims.UserID,
		"email": claims.Email,
	})
}
