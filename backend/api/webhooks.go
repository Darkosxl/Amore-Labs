//AI coded
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/webhook"
)

// StripeWebhook handles incoming Stripe webhook events
func StripeWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Error reading request body"})
		return
	}

	// Get the Stripe signature from headers
	signatureHeader := c.GetHeader("Stripe-Signature")
	
	// Verify webhook signature if webhook secret is configured
	webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	var event stripe.Event
	
	if webhookSecret != "" {
		// Verify the webhook signature
		event, err = webhook.ConstructEvent(payload, signatureHeader, webhookSecret)
		if err != nil {
			log.Printf("Webhook signature verification failed: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
			return
		}
	} else {
		// In development, if no secret is set, just parse the event
		log.Println("WARNING: No STRIPE_WEBHOOK_SECRET set, skipping signature verification")
		err := json.Unmarshal(payload, &event)
		if err != nil {
			log.Printf("Failed to parse webhook JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
	}

	// Handle the event
	switch event.Type {
	case "customer.subscription.created":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing subscription: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing subscription"})
			return
		}
		handleSubscriptionCreated(&subscription)

	case "customer.subscription.updated":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing subscription: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing subscription"})
			return
		}
		handleSubscriptionUpdated(&subscription)

	case "customer.subscription.deleted":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing subscription: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing subscription"})
			return
		}
		handleSubscriptionDeleted(&subscription)

	case "invoice.payment_failed":
		var invoice stripe.Invoice
		err := json.Unmarshal(event.Data.Raw, &invoice)
		if err != nil {
			log.Printf("Error parsing invoice: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing invoice"})
			return
		}
		handlePaymentFailed(&invoice)

	default:
		log.Printf("Unhandled event type: %s", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func handleSubscriptionCreated(subscription *stripe.Subscription) {
	log.Printf("✅ Subscription created: %s for customer: %s (Status: %s)", 
		subscription.ID, subscription.Customer.ID, subscription.Status)
	
	// Future: You could trigger WorkOS organization updates here
	// For now, WorkOS Stripe integration should handle this automatically
}

func handleSubscriptionUpdated(subscription *stripe.Subscription) {
	log.Printf("🔄 Subscription updated: %s (Status: %s, Cancel at period end: %t)", 
		subscription.ID, subscription.Status, subscription.CancelAtPeriodEnd)
	
	// Log important status changes
	if subscription.Status == "past_due" {
		log.Printf("⚠️  Subscription %s is past due", subscription.ID)
	} else if subscription.Status == "canceled" {
		log.Printf("❌ Subscription %s was canceled", subscription.ID)
	}
}

func handleSubscriptionDeleted(subscription *stripe.Subscription) {
	log.Printf("🗑️  Subscription deleted: %s for customer: %s", 
		subscription.ID, subscription.Customer.ID)
}

func handlePaymentFailed(invoice *stripe.Invoice) {
	log.Printf("💳 Payment failed for customer: %s (Invoice: %s, Amount: %d %s)", 
		invoice.Customer.ID, invoice.ID, invoice.AmountDue, invoice.Currency)
	
	// Future: Send notification emails, update user status, etc.
}

// Helper function to format currency amounts
func formatCurrency(amount int64, currency string) string {
	return fmt.Sprintf("%.2f %s", float64(amount)/100.0, currency)
}
