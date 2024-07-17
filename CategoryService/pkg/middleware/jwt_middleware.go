package middleware

import (
	Session "github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/fiber/v2"
	"library-management/CategoryService/platform/grpc/user"
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
	println(t)
	s := Session.GetSession(c)
	userGrpc := user.NewUserGrpc(s)
	_, err := userGrpc.JwtInterceptor(t)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": err.Error(),
			"status":  "Unauthorized",
			"data":    nil,
		})

	}
	c.Next()

	return nil
}
