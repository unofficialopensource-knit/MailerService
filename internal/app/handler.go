package app

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func ContactUsHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var payload ContactUsInput

	err := c.BodyParser(&payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusUnprocessableEntity)
	}
	err = service.SendContactUsMail(payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func WelcomeHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var payload WelcomeInput

	err := c.BodyParser(&payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusUnprocessableEntity)
	}
	err = service.SendWelcomeMail(payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func PasswordResetHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var payload PasswordResetInput

	err := c.BodyParser(&payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusUnprocessableEntity)
	}
	err = service.SendPasswordResetMail(payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func OrderReceiptHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var payload OrderReceiptInput

	err := c.BodyParser(&payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusUnprocessableEntity)
	}
	err = service.SendOrderReceiptMail(payload)
	if err != nil {
		slog.Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusAccepted)
}
