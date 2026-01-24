package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler provides a simple status check for the API.
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}


