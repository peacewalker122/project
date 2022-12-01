package api

import (
	"errors"

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
