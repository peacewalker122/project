package api

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type H = map[string]interface{}

// A for Allow
var (
	AlphaNumCheck = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	AlphaCheck    = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
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

func ValidateAlphanum(username string) error {
	if err := validateString(username, 4, 100); err != nil {
		return err
	}

	if !AlphaNumCheck(username) {
		return fmt.Errorf("illegal type must be string and number")
	}
	return nil
}

func ValidateAlpha(fullname string) error {
	if err := validateString(fullname, 3, 100); err != nil {
		return err
	}
	if !AlphaCheck(fullname) {
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

func ConverterParam(param string) int {
	var c echo.Context
	id := c.Param(param)

	num, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return num
}

func ValidateID(num int) error {
	var c echo.Context
	if num == 0 {
		return c.JSON(http.StatusBadRequest, ValidateError("id", fmt.Errorf("nvalid number").Error()))
	}
	return nil
}

func ValidateURI(params *GetAccountsParam, URIparam string) error {
	n := ConverterParam(URIparam)
	if err := ValidateID(n); err != nil {
		return err
	}
	params.ID = n
	return nil
}
