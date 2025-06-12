package utils

import (
	"fmt"
	"net/http"

	"github.com/anonystick/go-drunk-backend-api-by-ddd-java/pkg/response"
	"github.com/go-playground/validator"
)

// ValidateStruct validates a struct and returns a formatted APIError if validation fails
func ValidateStruct(data interface{}, validate *validator.Validate) *response.APIError {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return response.NewAPIError(http.StatusBadRequest, "Validation failed", err.Error())
	}

	errorMessages := make(map[string]string)
	for _, fieldErr := range validationErrs {
		errorMessages[fieldErr.Field()] = fmt.Sprintf(
			"Field validation for '%s' failed on the '%s' tag (value: '%v', param: '%s')",
			fieldErr.Field(),
			fieldErr.Tag(),
			fieldErr.Value(),
			fieldErr.Param(),
		)
	}

	return response.NewAPIError(http.StatusBadRequest, "Validation failed", errorMessages)
}
