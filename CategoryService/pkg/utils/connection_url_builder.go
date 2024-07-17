package utils

import (
	"fmt"
	"library-management/CategoryService/pkg/configs"
	"net/url"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var urlB string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgresSQL connection.
		urlB = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			url.QueryEscape("Asia/Jakarta"))
		println(os.Getenv("DB_NAME"))
	case "redis":
		// URL for Redis connection.
		urlB = os.Getenv("REDIS_ADDRESS")
	case "fiber":
		// URL for Fiber connection.
		urlB = fmt.Sprintf(
			"%s:%d",
			"0.0.0.0",
			configs.Config.Apps.HttpPort,
		)
	case "grpc":
		// URL for Redis connection.
		urlB = fmt.Sprintf(
			"tcp:%d",
			os.Getenv("GRPC_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return urlB, nil
}
