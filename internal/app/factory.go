package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var appLocalConfig HTTPConfig

func AppFactory(mode string) *fiber.App {
	appLocalConfig, _ = LoadConfig(context.Background())
	fiberConfig := fiber.Config{
		AppName:       "UploadService",
		BodyLimit:     20 * 1024 * 1024,
		CaseSensitive: true,
		ServerHeader:  "Fiber/Go",
		StrictRouting: true,
	}

	app := fiber.New(fiberConfig)
	app.Use(cors.New())

	app.Use(logger.New())

	return app
}
