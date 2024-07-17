package routes

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "library-management/CategoryService/app/handlers/http"
	"library-management/CategoryService/pkg/middleware"
)

// PrivateRoutes func for describe group of private_libs.sh routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	v1 := a.Group("/v1")
	v1.Post("/categories", middleware.JWTProtectedG, httpHandler.Create)
	v1.Get("/categories", middleware.JWTProtectedG, httpHandler.Get)
	v1.Get("/categories/detail/:id", middleware.JWTProtectedG, httpHandler.GetCategory)
}
