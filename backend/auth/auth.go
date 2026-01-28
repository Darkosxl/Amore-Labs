package auth

import (
	mw "amorelabs/backend/middleware"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/workos/workos-go/v6/pkg/usermanagement"
)


func LoginHandler(c *gin.Context) {
	authUrl, err := usermanagement.GetAuthorizationURL(usermanagement.GetAuthorizationURLOpts{
		ClientID: os.Getenv("WORKOS_CLIENT_ID"),
		RedirectURI: os.Getenv("API_URL") + "/auth/callback",
		Provider: "authkit",
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, authUrl.String())
	
}

func LogoutHandler(c *gin.Context) {
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	frontendUrl := strings.TrimSpace(allowedOrigins[0])
	returnTo := frontendUrl + "#/signin"

	// Get the access token cookie to extract session ID
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		// No token, just redirect to signin
		c.SetCookie("access_token", "", -1, "/", os.Getenv("COOKIE_DOMAIN"), false, true)
		c.Redirect(http.StatusFound, returnTo)
		return
	}

	// Parse our JWT to get the session ID
	claims, err := mw.ValidateToken(tokenString)
	if err != nil || claims.SessionID == "" {
		// Can't get session ID, just clear cookie and redirect
		c.SetCookie("access_token", "", -1, "/", os.Getenv("COOKIE_DOMAIN"), false, true)
		c.Redirect(http.StatusFound, returnTo)
		return
	}

	// Clear the access token cookie
	c.SetCookie("access_token", "", -1, "/", os.Getenv("COOKIE_DOMAIN"), false, true)

	// Get WorkOS Logout URL with session ID
	logoutUrl, err := usermanagement.GetLogoutURL(usermanagement.GetLogoutURLOpts{
		SessionID: claims.SessionID,
	})

	if err != nil {
		// Fallback if URL gen fails: just redirect to signin
		c.Redirect(http.StatusFound, returnTo)
		return
	}

	// Redirect to WorkOS to end their session
	c.Redirect(http.StatusFound, logoutUrl.String())
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
	
	// Extract session ID from WorkOS access token (it's in the 'sid' claim)
	var sessionID string
	token, _, err := jwt.NewParser().ParseUnverified(user.AccessToken, jwt.MapClaims{})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sid, ok := claims["sid"].(string); ok {
				sessionID = sid
			}
		}
	}

	// Generate access token with empty entitlements and session ID for logout
	accesstoken, err := mw.GenerateToken(user.User.ID, user.User.Email, "user", []string{}, sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Failed to generate token"})
		return
	}
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	frontendUrl := strings.TrimSpace(allowedOrigins[0])
	c.SetCookie("access_token", accesstoken, 86400, "/", os.Getenv("COOKIE_DOMAIN"), false, true)
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
