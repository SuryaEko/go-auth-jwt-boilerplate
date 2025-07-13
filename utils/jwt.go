package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Load the secret key from environment variable
var secretKey = os.Getenv("JWT_SECRET")

// JWTClaims defines the structure for JWT claims
type JWTClaims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for the given userId
func GenerateJWT(userID uint) (string, error) {
	exp_time, err := time.ParseDuration(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		return "", err
	}

	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-auth-jwt-boilerplate",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp_time)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken and parses the JWT token and returns the claims
func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err // Return error if token parsing fails
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
