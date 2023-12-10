package webcore

import (
	"github.com/gofiber/fiber/v2"
)

const publicDir = "./public"

func SetupStaticRoutes(app *fiber.App) {

	app.Static("/", publicDir)
	app.Static("/", publicDir+"/assets")
	app.Static("/", publicDir+"/js")
	app.Static("/", publicDir+"/css")
	app.Static("/", publicDir+"/images")

}
