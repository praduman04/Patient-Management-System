package handlers

import (
	"net/http"
	"pms/models"
	"pms/services"

	"pms/utils"

	"github.com/labstack/echo/v4"
)

type PatientHandler struct {
	service *services.PatientService
}

func NewPatientHandler(service *services.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}
func (h *PatientHandler) CreatePatient(c echo.Context) error {
	var patient models.Patient
	if err := c.Bind(&patient); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if err := utils.Validator.Struct(patient); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": utils.FormatValidationError(err)})
	}

	if err := h.service.CreatePatient(c.Request().Context(), patient); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Patient created Successfully"})

}
func (h *PatientHandler) GetAll(c echo.Context) error {

	patients, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, patients)
}
func (h *PatientHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	patient, err := h.service.GetById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Patient Not Found"})
	}
	return c.JSON(http.StatusOK, patient)
}
func (h *PatientHandler) DeletePatient(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Patient Deleted Successfully"})
}

// func (h *PatientHandler)update(c echo.Context)error{
// 	id:=c.Param("id")
// 	var patient models.Patient
// 	if err:=c.Bind(patient);err!=nil{
// 		return c.JSON(http.StatusBadRequest,echo.Map{"error":err.Error()})
// 	}

// }
