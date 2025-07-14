package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// IsPasswordValid checks if the provided password matches the hashed password
func IsPasswordValid(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
