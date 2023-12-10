package webcore

import (
	"github.com/elanticrypt0/gasonline/pkg/access"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MiddlewareSetup(gas *GasonlineApp) {
	// CORS
	gas.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: gas.App.Config.App_CORS_origins,
		AllowHeaders: gas.App.Config.App_CORS_headers,
	}))

	//  Recover from error
	gas.Fiber.Use(recover.New())

	// use this middleware to check if the user is authenticated
	gas.Fiber.Use(access.AccessWithAuthenticatedUser)

	// LoggerOnFile(gas.Fiber)
	LogOn(gas.Fiber)

}
