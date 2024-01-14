package config

import "github.com/go-playground/validator/v10"

/*
	NewValidatorRegister func create a new validator instance
	Adds a custom validations with given tags in this config package
	And returning pointer validator.Validate
*/
func NewValidatorRegister() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("alphabet_only", AlphabetOnly)
	return v
}
