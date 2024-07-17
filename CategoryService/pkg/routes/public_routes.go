package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	httpHandler "library-management/CategoryService/app/handlers/http"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")
	v1 := route.Group("/v1")
	v1.Get("/monitor", monitor.New(monitor.Config{Title: "Go-Fiber v2 Metrics Page"}))
	v1.Get("/health", httpHandler.HealthCheck) // register a new user
}
