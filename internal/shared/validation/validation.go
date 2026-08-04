package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatErrors converts validator errors into a consistent API response.
func FormatErrors(err error) map[string][]string {

	errors := make(map[string][]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["error"] = []string{err.Error()}
		return errors
	}

	for _, fieldError := range validationErrors {

		field := strings.ToLower(fieldError.Field())

		switch fieldError.Tag() {

		case "required":
			errors[field] = append(errors[field],
				field+" is required")

		case "email":
			errors[field] = append(errors[field],
				"invalid email address")

		case "min":
			errors[field] = append(errors[field],
				field+" must be at least "+fieldError.Param()+" characters")

		case "max":
			errors[field] = append(errors[field],
				field+" must be at most "+fieldError.Param()+" characters")

		case "oneof":
			errors[field] = append(errors[field],
				field+" must be one of: "+fieldError.Param())

		default:
			errors[field] = append(errors[field],
				fieldError.Error())
		}
	}

	return errors
}
