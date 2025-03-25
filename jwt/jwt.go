package jwt

import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

// sample secret key
var jwtKey = []byte("sample_secret_key")

func CreateJWTKey(secret string) {
	jwtKey = []byte(secret)
}

// GenerateToken creates a JWT token with custom claims
func GenerateToken(claims map[string]any, expiration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiration)
	claims["exp"] = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString(jwtKey)
}

// ValidateToken checks the validity of the JWT token
func ValidateToken(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
