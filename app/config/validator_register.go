package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func AlphabetOnly(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	isAlpha := regexp.MustCompile(`^[a-zA-Z]+(\s[a-zA-Z]+)?$`).MatchString

	if len(field) == 0 {
		return true
	}
	return isAlpha(field)
}
