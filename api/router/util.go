package api

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	Util struct {
		Middleware *CustomBinder
		Validator  *customValidator
	}
)

func NewUtil(val *validator.Validate) *Util {
	return &Util{
		Middleware: &CustomBinder{},
		Validator:  NewValidator(val),
	}
}

func EndpointLog(arg []*echo.Route, logger echo.Logger) []interface{} {
	dummy := make([]interface{}, 0)
	for _, v := range arg {
		logS := fmt.Sprintf("%s \t %s \n ", v.Method, v.Path)
		dummy = append(dummy, logS)
	}
	return dummy
}
