package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID       string   `json:"user_id"`
	Email        string   `json:"email"`
	Role         string   `json:"role"`
	Entitlements []string `json:"entitlements"`
	SessionID    string   `json:"session_id"` // WorkOS session ID for logout
	jwt.RegisteredClaims
}

func GenerateToken(userid string, email string, role string, entitlements []string, sessionID string) (string, error) {
	claims := Claims{
		UserID:       userid,
		Email:        email,
		Role:         role,
		Entitlements: entitlements,
		SessionID:    sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "amorelabs",
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenString.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(token string) (*Claims, error)  {
	tokenString, err := jwt.ParseWithClaims(token, &Claims{}, 
		func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenString.Claims.(*Claims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token claims")
	}
}