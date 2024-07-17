package middleware

import (
	Session "github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/fiber/v2"
	"library-management/UserService/pkg/utils"
	"os"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ContextKey:   "jwt", // used in private_libs.sh routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
			"status":  "Unauthorized",
			"data":    nil,
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    fiber.StatusUnauthorized,
		"message": err.Error(),
		"status":  "Unauthorized",
		"data":    nil,
	})
}

func JWTProtectedG(c *fiber.Ctx) error {
	t := c.Get("Authorization")

	s := Session.GetSession(c)
	claim, err := utils.JWTInterceptor(t)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": err.Error(),
			"status":  "Unauthorized",
			"data":    nil,
		})
	}
	s.Put("meta", claim)
	c.Next()

	return nil
}
