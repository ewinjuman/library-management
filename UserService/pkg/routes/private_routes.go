package routes

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "library-management/UserService/app/handlers/http"
	"library-management/UserService/pkg/middleware"
)

// PrivateRoutes func for describe group of private_libs.sh routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	v1 := a.Group("/v1")
	v1.Get("/users", middleware.JWTProtectedG, httpHandler.Get)
	v1.Get("/users/me", middleware.JWTProtectedG, httpHandler.Me)
}
