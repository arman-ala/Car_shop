package validations

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ^09[0-9]{9}|^98[0-9]{9}|^\+98[0-9]{9}
func IranianPhoneNumberValidator(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if !ok {
		return false
	}

	result, err := regexp.MatchString("^09[0-9]{9}$|^98[0-9]{9}$|^\\+98[0-9]{9}$", value)
	if err != nil {
		log.Printf("error happened while validating phone number: \n%s\n", err)
		return false
	}
	
	return result
}
