package main

import (
	"pms/config"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	config.Connect("mongodb://localhost:27017")
	app.Start(":3000")
}
