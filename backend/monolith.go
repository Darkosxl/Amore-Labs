package main

import (
	h "amorelabs/backend/api"
	auth "amorelabs/backend/auth"
	mw "amorelabs/backend/middleware"
	"github.com/gin-gonic/gin"
	
)

func main() {
	r := gin.Default()
	authorized := r.Group("/v1")
	// CORS MIDDLEWARE
	r.Use(mw.CORSMiddleware())

	// HEALTH
	r.GET("/health", h.HealthHandler)

	// AUTH
	r.POST("/auth/login", auth.LoginHandler)
	
	r.POST("/callback", auth.CallbackHandler)

	// PROTECTED ROUTES
	authorized.Use(mw.AuthMiddleware())
	authorized.POST("/admin_console", auth.AdminConsoleHandler)
	authorized.POST("/me", auth.Me)
	authorized.POST("/billing", h.createCheckoutSession)
	r.Run(":8173")
}