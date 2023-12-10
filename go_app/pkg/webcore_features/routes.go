package webcore_features

import (
	"github.com/elanticrypt0/gasonline/pkg/access"
	"github.com/elanticrypt0/gasonline/pkg/webcore"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(gas *webcore.GasonlineApp) {

	// setup
	setup := gas.Fiber.Group("/setup")

	setup.Get("/", func(c *fiber.Ctx) error {
		return Setup(c, gas)
	})

	// app monitor
	setup.Get("/monitor", monitor.New(monitor.Config{Title: gas.App.Config.App_name + " Monitor Page"}))

	//status
	setup.Get("/status", func(c *fiber.Ctx) error {
		return Status(c)
	})

	// seeder
	if gas.App.Config.App_setup_enabled {
		setup.Get("/seed", func(c *fiber.Ctx) error {
			return Seed(c, gas)
		})
		setup.Get("/seed/:table_name", func(c *fiber.Ctx) error {
			return Seed(c, gas)
		})
	}

	// access pkg
	var accessConfig access.AccessConfig
	access.LoadConfig(&accessConfig)
	if accessConfig.IsEnabled {
		AccessRoutesSetup(gas, &accessConfig)
	}

}

func AccessRoutesSetup(gas *webcore.GasonlineApp, config *access.AccessConfig) {
	access := gas.Fiber.Group(config.BaseURL)

	access.Get("/login", func(c *fiber.Ctx) error {
		// todo
		return c.SendString("login")
	})

	access.Get("/logout", func(c *fiber.Ctx) error {
		// todo
		return c.SendString("logout")
	})

	access.Post("/create", func(c *fiber.Ctx) error {
		// todo
		return c.SendString("Create")
	})

	access.Put("/update", func(c *fiber.Ctx) error {
		// todo
		return c.SendString("update")
	})

	access.Get("/delete", func(c *fiber.Ctx) error {
		// todo
		return c.SendString("delete")
	})
}
