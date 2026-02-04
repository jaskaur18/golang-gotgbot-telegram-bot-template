package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	if validate == nil {
		Init()
	}
	return validate.Struct(s)
}

func Get() *validator.Validate {
	if validate == nil {
		Init()
	}
	return validate
}
