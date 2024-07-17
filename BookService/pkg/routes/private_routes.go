package routes

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "library-management/BookService/app/handlers/http"
	"library-management/BookService/pkg/middleware"
)

// PrivateRoutes func for describe group of private_libs.sh routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	v1 := a.Group("/v1")
	v1.Post("/books", middleware.JWTProtectedG, httpHandler.Create)
	v1.Get("/books", middleware.JWTProtectedG, httpHandler.Get)
	v1.Get("/books/:id", middleware.JWTProtectedG, httpHandler.GetBook)
	v1.Post("/books/borrow", middleware.JWTProtectedG, httpHandler.Borrow)
	v1.Post("/books/return", middleware.JWTProtectedG, httpHandler.Return)
	v1.Get("/books/borrow", middleware.JWTProtectedG, httpHandler.GetBorrow)
}
