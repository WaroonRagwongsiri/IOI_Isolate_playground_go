package router

import (
	"ioitest/controller"

	"github.com/gofiber/fiber/v2"
)

func RunCRoutes(app *fiber.App) {
	app.Post("/run_C", func(c *fiber.Ctx) error {
		var req controller.RunC

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON body",
			})
		}

		return c.JSON(controller.RunCController(req))
	})
}

func JobFromIDRoutes(app *fiber.App) {
	app.Get("/job_id", func(c *fiber.Ctx) error {
		var req controller.JobFromId

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON body",
			})
		}

		return c.JSON(controller.JobFromIDController(req))
	})
}
