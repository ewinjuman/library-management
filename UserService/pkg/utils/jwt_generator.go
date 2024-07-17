package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Tokens struct to describe tokens object.
type Tokens struct {
	Access string `json:"access"`
}

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewTokens(id string, metaData TokenMetadata) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(id, metaData)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access: accessToken,
	}, nil
}

func generateNewAccessToken(id string, metaData TokenMetadata) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET")

	// Set expires minutes count for secret key from resource file.
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["userId"] = metaData.UserID
	claims["expires"] = expirationTime.Unix()
	claims["exp"] = expirationTime.Unix()
	claims["meta_data"] = metaData

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
