package validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func(cv *CustomValidator) Validate(i interface{}) error {
	return validator.New().Struct(i)
}
