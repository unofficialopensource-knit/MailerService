package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var appLocalConfig HTTPConfig
var service Service

func AppFactory() *fiber.App {
	appLocalConfig, _ = NewHTTPConfig(context.Background())
	fiberConfig := fiber.Config{
		AppName:       "UploadService",
		BodyLimit:     20 * 1024 * 1024,
		CaseSensitive: true,
		ServerHeader:  "Fiber/Go",
		StrictRouting: true,
	}
	service = *NewService(appLocalConfig)

	app := fiber.New(fiberConfig)
	app.Use(cors.New())

	app.Use(logger.New())

	app.Post("/contact-us", ContactUsHandler)
	app.Post("/welcome", WelcomeHandler)

	return app
}
