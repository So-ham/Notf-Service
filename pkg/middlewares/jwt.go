package middlewares

import (
	"errors"
	"os"

	"github.com/So-ham/notf-service/internal/entities"
	"github.com/golang-jwt/jwt/v4"
)

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string) (*entities.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entities.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			return nil, errors.New("JWT secret key not found")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entities.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
