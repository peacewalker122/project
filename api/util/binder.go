package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// using default binder
	db := &echo.DefaultBinder{}
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return errors.New("unsupported Type")
	}

	return
}

type customValidator struct {
	validate *validator.Validate
}

func newValidator(arg *validator.Validate) *customValidator {
	return &customValidator{
		validate: arg,
	}
}

func (v *customValidator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

func (u *Util) HTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errmsg := []string{}
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, e := range castedObject {
			errmsg = append(errmsg, fmt.Sprintf("error happen in %s, due %s", e.Field(), e.ActualTag()))
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, errmsg)
}
