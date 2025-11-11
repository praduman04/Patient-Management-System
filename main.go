package main

import (
	"pms/config"
	"pms/handlers"
	"pms/repo"
	"pms/routes"
	"pms/services"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	config.Connect("mongodb://localhost:27017")
	patientRepo := repo.NewPatientRepo(config.Client.Database("pms"))
	patientService := services.NewPatientService(patientRepo)
	patientHandler := handlers.NewPatientHandler(patientService)
	routes.PatientRoutes(app, patientHandler)

	app.Start(":3000")
}
