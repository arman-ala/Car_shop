package validations

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PasswordValidator(field validator.FieldLevel) (result bool) {
	value, ok := field.Field().Interface().(string)
	// result = false by default because of zero-value of bool
	if !ok {
		log.Printf("error happened while validating password: %v", value)
		return
	}

	// Here we check if the password satisfies the regex
	result, err := regexp.MatchString(`^((?=\S*?[A-Z])(?=\S*?[a-z])(?=\S*?[0-9]).{6,})\S$`, value)
	if err != nil {
		log.Printf("error happened while validating password: \n%s\n", err)
		return
	}

	result = true
	return
}
