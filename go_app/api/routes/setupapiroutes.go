package routes

import (
	"github.com/elanticrypt0/gasonline/pkg/webcore"
)

func SetupApiRoutes(gas *webcore.GasonlineApp) {
	api := gas.Fiber.Group("/api")
	// categories
	categoriesRoutes(gas, api)
}
