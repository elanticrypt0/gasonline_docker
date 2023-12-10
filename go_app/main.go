package main

import (
	"log"

	"github.com/elanticrypt0/gasonline/api"
	"github.com/elanticrypt0/gasonline/pkg/webcore"
	"github.com/elanticrypt0/gasonline/pkg/webcore_features"
	"github.com/elanticrypt0/go4it"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app_config := go4it.NewApp("./config/appconfig")

	gas := webcore.GasonlineApp{
		App: &app_config,
		Fiber: fiber.New(fiber.Config{
			Prefork:               false,
			CaseSensitive:         true,
			StrictRouting:         true,
			ServerHeader:          "Fiber",
			AppName:               app_config.Config.App_name,
			DisableStartupMessage: false,
			PassLocalsToViews:     true,
		}),
	}
	gas.PrintAppInfo()

	// make the connection
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	// middleware
	webcore.MiddlewareSetup(&gas)

	// Routes setup

	// webcore setup routes
	if gas.App.Config.App_setup_enabled {
		webcore_features.SetupRoutes(&gas)
		webcore_features.SetupOnStartup(&gas)
	}

	api.ApiSetup(&gas)

	// static routes
	webcore.SetupStaticRoutes(gas.Fiber)

	portAsStr := gas.GetPortAsStr()

	// go4it.OpenInBrowser("http://" + gas.GetAppUrl())

	log.Fatal(gas.Fiber.Listen(gas.GetAppUrl()), "Server is running on port "+portAsStr)

}
