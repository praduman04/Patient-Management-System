package main

import (
	"pms/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Connect("mongodb://localhost:27017")
	app.Listen(":3000")
}
