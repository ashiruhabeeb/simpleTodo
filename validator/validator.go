package validator

import "github.com/go-playground/validator/v10"

func Validate(i interface{}) error {
	return validator.New().Struct(i)
}