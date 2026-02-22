package main

import (
	"log"

	"ioitest/controller"
	"ioitest/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	controller.StartWorkers()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected")
	})

	router.RunCRoutes(app)
	router.JobFromIDRoutes(app)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}