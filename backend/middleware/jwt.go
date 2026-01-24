package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"fmt"
)

func GenerateToken(userid string, email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userid,
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) (jwt.MapClaims, error)  {
	tokenString, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenString.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}