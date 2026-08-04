package validation

import "github.com/go-playground/validator/v10"

// Validator exposes a shared validator instance.
//
// Custom validation rules can be registered here in the future.
var Validator = validator.New()
