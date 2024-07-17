package utils

import (
	"encoding/json"
	"fmt"
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/golang-jwt/jwt/v4"
	"library-management/UserService/pkg/repository"
	"net/http"
	"os"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Id       string
	UserID   int
	Expires  int64
	Username string
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func JWTInterceptor(tokenString string) (*TokenMetadata, error) {

	//tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, Error.NewError(http.StatusUnauthorized, repository.FailedStatus, fmt.Sprintf("invalid token: %v", err.Error()))
	}
	if !token.Valid {
		return nil, Error.NewError(http.StatusUnauthorized, repository.FailedStatus, "invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		meta := claims["meta_data"].(map[string]interface{})
		jsonData, _ := json.Marshal(meta)

		// Convert the JSON to a struct
		var structData *TokenMetadata
		json.Unmarshal(jsonData, &structData)
		return structData, nil
	}
	return nil, nil
}
