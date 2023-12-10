package api

import (
	"github.com/elanticrypt0/gasonline/api/routes"
	"github.com/elanticrypt0/gasonline/pkg/webcore"
)

func ApiSetup(gas *webcore.GasonlineApp) {

	gas.App.DB.Primary.AutoMigrate()
	// features routes
	routes.SetupApiRoutes(gas)

}
