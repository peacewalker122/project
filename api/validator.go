package api

import (
	"database/sql"
	"errors"
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
	AlphaNumCheck = regexp.MustCompile(`^[a-zA-Z0-9_\s]+$`).MatchString
	AlphaCheck    = regexp.MustCompile(`^[a-zA-Z_\s]+$`).MatchString
	NumCheckByte  = regexp.MustCompile(`^[0-9]+$`).Match
	NumCheck      = regexp.MustCompile(`^[0-9]+$`).MatchString
	StringsCheck  = regexp.MustCompile(`^[a-zA-Z0-9_\s'"?!,.&%$@-]+$`).MatchString
)

const (
	like       = "like"
	unlike     = "unlike"
	retweet    = "retweet"
	unretweet  = "unretweet"
	comment    = "comment"
	qretweet   = "qoute-retweet"
	unqretweet = "unqoute-retweet(Delete)"
	posttag    = "post"
	accountag  = "account"
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

func ValidateAlphanum(username string, min, max int) error {
	if err := validateString(username, min, max); err != nil {
		return err
	}

	if !AlphaNumCheck(username) {
		return fmt.Errorf("illegal type must be string and number")
	}
	return nil
}

func ValidateAlpha(fullname string, min, max int) error {
	if err := validateString(fullname, min, max); err != nil {
		return err
	}
	if !AlphaCheck(fullname) {
		return fmt.Errorf("illegal type must be string")
	}
	return nil
}

func ValidateEmail(email string, min, max int) error {
	if err := validateString(email, min, max); err != nil {
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

func validatePassword(pass string, min, max int) error {
	if err := validateString(pass, min, max); err != nil {
		return err
	}
	return nil
}

func ValidateString(strings string, min, max int) error {
	if err := validateString(strings, min, max); err != nil {
		return err
	}

	if !StringsCheck(strings) {
		return errors.New("invalid character must be alphabet,num and symbol")
	}
	return nil
}

func ConverterParam(context echo.Context, param string) int {
	if !NumCheck(param) {
		id := context.Param(param)

		num, err := strconv.Atoi(id)
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}

func ValidateID(num int) error {
	if num < 1 {
		return fmt.Errorf("invalid number must be greater than 1")
	}
	return nil
}

func ValidateURIAccount(param *GetAccountsParams, context echo.Context, URIparam string) (*GetAccountsParams, error) {
	n := ConverterParam(context, URIparam)
	if err := ValidateID(n); err != nil {
		return nil, err
	}
	param.ID = n
	return param, nil
}

func ValidateURIPost(param *GetPostParam, context echo.Context, URIparam string) error {
	n := ConverterParam(context, URIparam)
	if err := ValidateID(n); err != nil {
		return err
	}
	param.ID = int64(n)
	return nil
}

func ValidateNum(num int) error {
	s := []byte("num")
	if NumCheckByte(s) {
		return errors.New("must be integer")
	}
	if num < 1 {
		return errors.New("must be integer that greater than 0")
	}
	return nil
}

func GetErrorValidator(c echo.Context, err error, tag string) error {
	if err != nil {
		if err == sql.ErrNoRows {
			err := fmt.Errorf("%v not found", tag)
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return err
}

func CreateErrorValidator(c echo.Context, err error) error {
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return err
}
