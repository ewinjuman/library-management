package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	httpHandler "library-management/UserService/app/handlers/http"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	v1 := a.Group("/v1")
	v1.Get("/monitor", monitor.New(monitor.Config{Title: "Go-Fiber v2 Metrics Page"}))
	v1.Get("/health", httpHandler.HealthCheck) // register a new user

	v1.Post("/register", httpHandler.Register)
	v1.Post("/login", httpHandler.Login)
}
