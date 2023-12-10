package webcore

import (
	"fmt"

	"github.com/elanticrypt0/go4it"
	"github.com/gofiber/fiber/v2"
)

type GasonlineApp struct {
	App   *go4it.App
	Fiber *fiber.App
}

func (gas *GasonlineApp) PrintAppInfo() {
	fmt.Printf("Starting app: %s v%s", gas.App.Config.App_name, gas.App.Config.App_version)
}

func (gas *GasonlineApp) GetAppUrl() string {
	return fmt.Sprintf("%s:%d", gas.App.Config.App_server_host, gas.App.Config.App_server_port)
}

func (gas *GasonlineApp) GetPortAsStr() string {
	return fmt.Sprintf("%d", gas.App.Config.App_server_port)
}
