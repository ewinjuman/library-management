package routes

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "library-management/AuthorService/app/handlers/http"
	"library-management/AuthorService/pkg/middleware"
)

// PrivateRoutes func for describe group of private_libs.sh routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	v1 := a.Group("/v1")
	v1.Post("/authors", middleware.JWTProtectedG, httpHandler.Create)
	v1.Get("/authors", middleware.JWTProtectedG, httpHandler.Get)
	v1.Get("/authors/:id", middleware.JWTProtectedG, httpHandler.GetAuthor)
}
