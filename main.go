package main

import (
	"fmt"
	"log"

	"ioitest/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected")
	})

	router.RunCRoutes(app)
	router.JobFromIDRoutes(app)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}