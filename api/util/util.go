package api

import (
	"github.com/go-playground/validator/v10"
)

type (
	Util struct {
		Middleware *CustomBinder
		*customValidator
	}
)

func NewUtil(val *validator.Validate) *Util {
	return &Util{
		Middleware:      &CustomBinder{},
		customValidator: newValidator(val),
	}
}
