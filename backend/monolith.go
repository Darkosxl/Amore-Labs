package main

import (
	h "amorelabs/backend/api"
	auth "amorelabs/backend/auth"
	mw "amorelabs/backend/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/workos/workos-go/v6/pkg/usermanagement"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	usermanagement.SetAPIKey(os.Getenv("WORKOS_API_KEY"))

	r := gin.Default()
	r.Use(mw.CORSMiddleware())
	authorized := r.Group("/v1")
	// CORS MIDDLEWARE
	

	// HEALTH
	r.GET("/health", h.HealthHandler)

	// WEBHOOKS (unprotected)
	r.POST("/webhooks/stripe", h.StripeWebhook)

	// AUTH
	r.GET("/auth/login", auth.LoginHandler)
	
	r.GET("auth/callback", auth.CallbackHandler)
	r.GET("/auth/logout", auth.LogoutHandler)
	
	r.POST("/auth/verify-masterkey", auth.VerifyMasterKeyHandler)
	
	authorized.Use(mw.AuthMiddleware())
	authorized.GET("/me", auth.Me)
	authorized.POST("/billing/:product", h.CreateCheckoutSession)
	
	
	authorized.GET("/subscriptions", h.GetSubscriptions)
	authorized.GET("/entitlements", h.GetEntitlements)
	
	r.Run(":8173")
}