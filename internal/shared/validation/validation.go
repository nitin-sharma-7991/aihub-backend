package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

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
			errors[field] = append(errors[field], fieldError.Field()+" is required")

		case "email":
			errors[field] = append(errors[field], "Invalid email address")

		case "min":
			errors[field] = append(errors[field], fieldError.Field()+" must be at least "+fieldError.Param()+" characters")

		case "max":
			errors[field] = append(errors[field], fieldError.Field()+" must be at most "+fieldError.Param()+" characters")

		default:
			errors[field] = append(errors[field], fieldError.Error())
		}
	}

	return errors
}
