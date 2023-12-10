package webcore

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const logFormat string = "${time} - [${ip}]:${port} ${status} - ${method} ${path}\n"

func LogOnFile(fiber *fiber.App) {
	// Custom File Writer
	file, err := os.OpenFile("./_logs/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	fiber.Use(logger.New(logger.Config{
		Format: logFormat,
	}))
}

func LogOn(fiber *fiber.App) {
	fiber.Use(logger.New(logger.Config{
		Format: logFormat,
	}))
}
