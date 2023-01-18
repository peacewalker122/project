package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
	"mime/multipart"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/peacewalker122/project/token"
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
	Like       = "like"
	Unlike     = "unlike"
	Retweet    = "retweet"
	Unretweet  = "unretweet"
	Comment    = "comment"
	Qretweet   = "qoute-retweet"
	Unqretweet = "unqoute-retweet(Delete)"
	Posttag    = "post"
	Accountag  = "account"
)

type Helper interface {
	AuthAccount(c echo.Context) (int, *token.Payload, error)
}

func (s *Handler) AuthAccount(c echo.Context) (int, *token.Payload, error) {
	var sessionS db.Session

	authParam, ok := c.Get(AuthPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return http.StatusInternalServerError, nil, err
	}

	key := "session:AccountsID:" + strconv.Itoa(int(authParam.AccountID))

	// Here we get from our redis cache value
	session, err := s.redis.Get(c.Request().Context(), key)
	if err != nil {
		// if it doesn't exist then query it from thes postgres
		sessionS, err := s.store.GetSession(c.Request().Context(), authParam.ID)
		if num, err := GetErrorValidator(c, err, Accountag); err != nil {
			return num, nil, err
		}

		if sessionS.ExpiresAt.Unix() < time.Now().Unix() {
			return http.StatusUnauthorized, nil, errors.New("session expired")
		}

		if sessionS.IsBlocked {
			return http.StatusUnauthorized, nil, errors.New("already LOGOUT")
		}

		// set into redis cache
		err = s.redis.Set(c.Request().Context(), key, sessionS, 15*time.Minute)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		return 0, authParam, nil
	}

	err = json.Unmarshal([]byte(session), &sessionS)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if sessionS.ExpiresAt.Unix() < time.Now().Unix() {
		return http.StatusUnauthorized, nil, errors.New("session expired")
	}

	if sessionS.IsBlocked {
		return http.StatusUnauthorized, nil, errors.New("already LOGOUT")
	}

	return 0, authParam, nil
}

func validateString(target string, minChar, maxChar int) error {
	if target == "" {
		return fmt.Errorf("empty! must contain %d-%d character", minChar, maxChar)
	}

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

func ValidateURI[T int64 | int](context echo.Context, URIparam string) (T, error) {
	n := ConverterParam(context, URIparam)
	if err := ValidateID(n); err != nil {
		return 0, err
	}
	id := T(n)
	return id, nil
}

func (s *GetPostParam) ValidateURIPost(context echo.Context, URIparam string) error {
	n := ConverterParam(context, URIparam)
	if err := ValidateID(n); err != nil {
		return err
	}
	s.PostID = int64(n)
	return nil
}
func (s *GetImageParam) ValidateURIPost(context echo.Context, URIparam string) error {
	n := ConverterParam(context, URIparam)
	if err := ValidateID(n); err != nil {
		return err
	}
	s.PostID = int64(n)
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

func GetErrorValidator(c echo.Context, err error, tag string) (int, error) {
	if err != nil {
		if err == sql.ErrNoRows {
			err := fmt.Errorf("%v not found", tag)
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}
	return 0, err
}

func CreateErrorValidator(c echo.Context, err error) error {
	if err != nil {
		c.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func ValidateCreateListAccount(req *listAccountRequest) (errors []string) {
	if err := ValidateNum(int(req.PageID)); err != nil {
		errors = append(errors, ValidateError("page_id", err.Error()))
	}
	if err := ValidateNum(int(req.PageSize)); err != nil {
		errors = append(errors, ValidateError("page_size", err.Error()))
	}
	return errors
}

func ValidateFileType(input multipart.File) error {
	byte := make([]byte, 512)
	if _, err := input.Read(byte); err != nil {
		return err
	}
	file := http.DetectContentType(byte)
	log.Error(file)
	//
	if file == "image/jpg" || file == "image/jpeg" || file == "image/gif" || file == "image/png" || file == "image/webp" || file == "video/mp4" {
		return nil
	}
	return errors.New("invalid type! must be jpeg/jpg/gif/png/mp4")
}
