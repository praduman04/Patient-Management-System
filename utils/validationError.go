package utils

import "github.com/go-playground/validator/v10"

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if _, ok := err.(*validator.InvalidValidationError); ok {
		errors["error"] = err.Error()
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = getErrorMessage(err)
	}
	return errors
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "min":
		return err.Field() + " must be at least " + err.Param() + " characters long"
	default:
		return "Invalid value for " + err.Field()
	}
}
