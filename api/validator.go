package api

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// A for Allow
var (
	UsernameCheck = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	FullNameCheck = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
	InvalidChar   = fmt.Errorf("invalid type must be string or number")
)

type customValidator struct {
	validate *validator.Validate
}

func (v *customValidator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

func HTTPErrorHandler(err error, c echo.Context) {
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

func validateString(target string, minChar, maxChar int) error {
	if len(target) < minChar || len(target) > maxChar {
		return fmt.Errorf("invalid length of string, must contain %d-%d character", minChar, maxChar)
	}
	return nil
}

func ValidateUsername(username string) error {
	if err := validateString(username, 4, 100); err != nil {
		return err
	}

	if !UsernameCheck(username) {
		return fmt.Errorf("illegal type must be string and number")
	}
	return nil
}

func ValidateFullname(fullname string) error {
	if err := validateString(fullname, 3, 100); err != nil {
		return err
	}
	if !FullNameCheck(fullname) {
		return fmt.Errorf("illegal type must be string")
	}
	return nil
}

func ValidateEmail(email string) error {
	if err := validateString(email, 4, 100); err != nil {
		return err
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("not valid email: %v", err)
	}
	return nil
}

func ValidateError(errorTag, errorString string) string {
	return fmt.Sprintf("error happen in %v due: %v", errorTag, errorString)
}

func validatePassword(pass string) error {
	if err := validateString(pass, 5, 100); err != nil {
		return err
	}
	return nil
}