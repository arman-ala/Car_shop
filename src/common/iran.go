package common

import (
	"log"
	"regexp"
)

const IRANIAN_PHONE_NUMBER_REGEX = "^09[0-9]{9}$|^98[0-9]{9}$|^\\+98[0-9]{9}$"

func IsPhoneNumberIranian(value string) bool {
	result, err := regexp.MatchString(IRANIAN_PHONE_NUMBER_REGEX, value)
	if err != nil {
		log.Printf("error happened while validating phone number: \n%s\n", err)
		return false
	}

	return result
}
