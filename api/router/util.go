package api

import (
	"github.com/go-playground/validator/v10"
)

type (
	Util struct {
		Middleware *CustomBinder
		Validator *customValidator
	}
)

func NewUtil(val *validator.Validate) *Util {
	return &Util{
		Middleware:      &CustomBinder{},
		Validator: NewValidator(val),
	}
}
