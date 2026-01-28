package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		token, err := c.Cookie("access_token")
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid access token"})
			return
		}
		claims, err := ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
	
}

func CORSMiddleware() gin.HandlerFunc {
	allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
	allowedOrigins := strings.Split(allowedOriginsStr, ",")

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if origin is in allowed list
		allowed := origin == "http://localhost:5173" // Always allow localhost for dev
		for _, o := range allowedOrigins {
			if origin == strings.TrimSpace(o) {
				allowed = true
				break
			}
		}

		if allowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

