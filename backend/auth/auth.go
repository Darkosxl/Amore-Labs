package auth

import (
	"context"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/workos/workos-go/v6/pkg/usermanagement"
	mw "amorelabs/backend/middleware"
)

func LoginHandler(c *gin.Context) {
	authUrl, err := usermanagement.GetAuthorizationURL(usermanagement.GetAuthorizationURLOpts{
		ClientID: os.Getenv("WORKOS_CLIENT_ID"),
		RedirectURI: "http://localhost:8173/callback",
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
	
	accesstoken, err := mw.GenerateToken(user.User.ID, user.User.Email, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "Failed to generate token"})
		return
	}
	c.SetCookie("access_token", accesstoken, 86400, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/admin_console")
	
}

func AdminConsoleHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "success_admin_console"})
}

func Me (c *gin.Context) {
	claims := c.Get("claims").(*Claims)
	return c.JSON(200, map[string]any{
		"email": claims.Email,
		"role": claims.Role
	})
}

func VerifyMasterKeyHandler(c *gin.Context) {
	key := c.PostForm("key")
	if key != os.Getenv("MASTER_KEY") {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "error": "Invalid master key"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
