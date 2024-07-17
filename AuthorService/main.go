package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
	"library-management/AuthorService/app/proto"
	"library-management/AuthorService/pkg/configs"
	"library-management/AuthorService/pkg/middleware"
	"library-management/AuthorService/pkg/routes"
	"library-management/AuthorService/pkg/utils"
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

	myFigure := figure.NewColorFigure("Authors Service", "", "green", true)
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
