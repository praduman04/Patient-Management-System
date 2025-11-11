package routes

import (
	"pms/handlers"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo, handler *handlers.PatientHandler) {
	patientGroup := e.Group("/patients")

	patientGroup.POST("/create", handler.CreatePatient)
	patientGroup.GET("/list", handler.GetAll)
	patientGroup.GET("/:id", handler.GetById)
	patientGroup.DELETE("/:id", handler.DeletePatient)
}
