package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// GetValidationErrors converts raw validator errors into a standardized list of structured error objects
// This function acts like a translator that takes technical error data and reformats it into an easy-to-read list
func GetValidationErrors(err error) *[]ValidationError {
	// Initialize an empty list to store our formatted validation errors (starts as an empty container)
	var ValidationErrors []ValidationError
	// Create a variable to hold the specialized validator error type (a container to hold technical error details)
	var ve validator.ValidationErrors
	// Check if the input error can be converted to the validator's error type (verify if the error is a validation failure)
	if errors.As(err, &ve) {
		// Loop through every individual validation error found in the technical error list
		for _, e := range ve {
			// Add a new formatted error to our list, extracting key details from the technical error
			// Pull out the name of the invalid field, the rule that failed, and the rule's required parameter
			ValidationErrors = append(ValidationErrors, ValidationError{
				Field: e.Field(),
				Tag:   e.Tag(),
				Value: e.Param(),
			})
		}
		// Return a pointer to our completed list of formatted errors (share the full list of validation failures)
		return &ValidationErrors
	}
	return nil
}
