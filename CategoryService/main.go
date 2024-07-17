package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
	"library-management/CategoryService/app/proto"
	"library-management/CategoryService/pkg/configs"
	"library-management/CategoryService/pkg/middleware"
	"library-management/CategoryService/pkg/routes"
	"library-management/CategoryService/pkg/utils"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private_libs.sh routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	myFigure := figure.NewColorFigure("Category Service", "", "green", true)
	myFigure.Print()

	go func() {
		proto.StartGrpcServer()
	}()
	// Start server (with or without graceful shutdown).
	//if configs.Config.Apps.Mode == "local" {
	//	utils.StartServer(app)
	//} else {
	utils.StartServerWithGracefulShutdown(app)
	//}
}
