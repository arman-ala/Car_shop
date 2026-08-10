package validations

import (
	"github.com/arman-ala/Car_shop/common"
	"github.com/go-playground/validator/v10"
)

// ^09[0-9]{9}|^98[0-9]{9}|^\+98[0-9]{9}
func IranianPhoneNumberValidator(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.IsPhoneNumberIranian(value)
}
