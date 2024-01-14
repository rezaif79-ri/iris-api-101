package config

import "github.com/go-playground/validator/v10"

func NewValidatorRegister() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("alphabet_only", AlphabetOnly)
	return v
}
