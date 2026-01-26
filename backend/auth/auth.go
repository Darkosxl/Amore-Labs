package auth

import (
	mw "amorelabs/backend/middleware"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/workos/workos-go/v6/pkg/usermanagement"
)


func LoginHandler(c *gin.Context) {
	authUrl, err := usermanagement.GetAuthorizationURL(usermanagement.GetAuthorizationURLOpts{
		ClientID: os.Getenv("WORKOS_CLIENT_ID"),
		RedirectURI: "http://localhost:8173/auth/callback",
		Provider: "authkit",
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, authUrl.String())
	
}

func CallbackHandler(c *gin.Context){
	code := c.Query("code")
	ctx := context.Background()
	user, err := usermanagement.AuthenticateWithCode(ctx,usermanagement.AuthenticateWithCodeOpts{
		ClientID: os.Getenv("WORKOS_CLIENT_ID"),
		Code: code,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	
	// Generate access token with empty entitlements (will be populated from WorkOS token)
	accesstoken, err := mw.GenerateToken(user.User.ID, user.User.Email, "user", []string{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Failed to generate token"})
		return
	}
	frontendUrl := os.Getenv("FRONTEND_URL")
	c.SetCookie("access_token", accesstoken, 86400, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, frontendUrl + "#/admin_console")
	
}


func Me (c *gin.Context) {
	claimscheck, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "error": "Unauthorized"})
		return
	}
	claims, ok := claimscheck.(*mw.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Invalid claims type"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"email": claims.Email,
		"role": claims.Role,
		"userId": claims.UserID,
	})
}

func VerifyMasterKeyHandler(c *gin.Context) {
	key := c.PostForm("key")
	if key != os.Getenv("MASTER_KEY") {
		fmt.Printf("Received: '%s' vs Expected: '%s'\n", key, os.Getenv("MASTER_KEY"))
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "error": "Invalid master key"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
