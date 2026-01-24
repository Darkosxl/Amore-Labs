package main

import (
	h "amorelabs/backend/api"
	auth "amorelabs/backend/auth"
	mw "amorelabs/backend/middleware"
	"github.com/gin-gonic/gin"
	
)

func main() {
	r := gin.Default()
	r.Use(mw.CORSMiddleware())
	authorized := r.Group("/v1")
	// CORS MIDDLEWARE
	

	// HEALTH
	r.GET("/health", h.HealthHandler)

	// AUTH
	r.POST("/auth/login", auth.LoginHandler)
	
	r.POST("/callback", auth.CallbackHandler)
	
	r.POST("/auth/verify-masterkey", auth.VerifyMasterKeyHandler)
	// PROTECTED ROUTES
	authorized.Use(mw.AuthMiddleware())
	authorized.POST("/admin_console", auth.AdminConsoleHandler)
	authorized.GET("/me", auth.Me)
	authorized.POST("/billing/:product", h.CreateCheckoutSession)
	
	r.Run(":8173")
}