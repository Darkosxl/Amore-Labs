package middleware

import (
	"net/http"
	"os"

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
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		frontendURL := os.Getenv("FRONTEND_URL")

		// Allow configured frontend or localhost for dev, otherwise fall back to strict check
		if origin == frontendURL || origin == "http://localhost:5173" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
            // Default safe fallback if origin doesn't match
             c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
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

