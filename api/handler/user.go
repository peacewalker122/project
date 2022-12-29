package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	api "github.com/peacewalker122/project/api/util"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
)

var (
	Errors = map[string]string{}
)

type userService interface {
	CreateRequestUser(c echo.Context) error
	CreateUser(c echo.Context) error
	Login(c echo.Context) error
}

type CreateUserParam struct {
	Username       string `json:"username" form:"username" validate:"required,min=4,max=100"`
	HashedPassword string `json:"password" form:"password" validate:"required,min=6,max=100"`
	FullName       string `json:"full_name" form:"full_name" validate:"required,min=3,max=100"`
	Email          string `json:"email" form:"email" validate:"required,email"`
}

type CreatingUser struct {
	Token int `json:"token" form:"token" validate:"required"`
}

func (s *Handler) CreateRequestUser(c echo.Context) error {
	req := new(CreateUserParam)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if errors := ValidationCreateUserRequest(req); errors != nil {
		return c.JSONPretty(http.StatusBadRequest, errors, "    ")
	}

	test, err := s.store.GetEmail(c.Request().Context(), db.GetEmailParams{Email: req.Email})
	if err == nil {
		Errors["email"] = errors.New("email already exist").Error()
	}

	log.Println(test)

	_, err = s.store.GetEmail(c.Request().Context(), db.GetEmailParams{Username: req.Username})
	if err == nil {
		Errors["username"] = errors.New("username already exist").Error()
	}

	if len(Errors) > 0 {
		return c.JSONPretty(http.StatusBadRequest, Errors, "    ")
	}

	var wg sync.WaitGroup
	uuidchan := make(chan uuid.UUID, 1)
	errchan := make(chan error, 2)

	wg.Add(1)
	go func(errchan chan error, uuidchan chan uuid.UUID) {
		defer wg.Done()
		uid, err := s.util.CreateEmailAuth(c.Request().Context(), req.Email)
		errchan <- err
		uuidchan <- uid
		// here we set the key to redis
		// to get the key we use the uuid
		err = s.redis.Set(c.Request().Context(), uid.String()+"h", req, 3*time.Minute)
		errchan <- err
	}(errchan, uuidchan)

	uid := <-uuidchan

	for v := range errchan {
		if v != nil {
			return c.JSON(http.StatusBadRequest, v.Error())
		}
	}
	wg.Wait()
	// here we send the email
	c.Response().Header().Add("uuid", uid.String())
	return c.JSON(http.StatusOK, "success")
}

func (s *Handler) CreateUser(c echo.Context) error {
	reqs := new(CreatingUser)
	if err := c.Bind(reqs); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(reqs); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	requid := c.Param("uuid")

	var req CreateUserParam
	_, err = s.util.VerifyEmailAuth(c.Request().Context(), requid, reqs.Token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	// here we get the key from redis
	result, err := s.redis.Get(c.Request().Context(), requid+"h")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := json.Unmarshal([]byte(result), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	hashpass, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashpass,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := s.store.CreateUser(c.Request().Context(), arg)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code.Name() {
			case "unique_violation":
				return c.JSON(http.StatusForbidden, err.Error())
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	argaccount := db.CreateAccountsParams{
		Owner: req.Username,
	}

	res, err := s.store.CreateAccounts(c.Request().Context(), argaccount)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return c.JSON(http.StatusForbidden, err.Error())
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	output := AccountResponse(res)
	resp := CreateUserResponses(user, output)
	return c.JSON(http.StatusOK, resp)
}

type LoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *Handler) Login(c echo.Context) error {
	req := new(LoginParams)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ValidationLoginRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	username, err := s.store.GetUser(c.Request().Context(), req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	account, err := s.store.GetAccountsOwner(c.Request().Context(), req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = util.CheckPassword(req.Password, username.HashedPassword)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ValidateError("password", err.Error()))
	}

	err = s.util.SendEmailWithNotif(c.Request().Context(), api.SendEmail{
		AccountID: []int64{account.AccountsID},
		Params:    []string{username.Email, c.RealIP()},
		Type:      "login",
		TimeSend:  time.Now().UTC().Local(),
	})
	if err != nil {
		log.Panic(err.Error())
	}

	token, Accespayload, err := s.token.CreateToken(req.Username, s.config.TokenDuration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	refreshToken, refreshPayload, err := s.token.CreateToken(req.Username, s.config.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     req.Username,
		RefreshToken: refreshToken,
		UserAgent:    c.Request().UserAgent(),
		ClientIp:     c.RealIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := s.store.CreateSession(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp := loginResp{
		SessionID:             session.ID,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt.UTC().Local(),
		User:                  UserResponse(username, account),
		AccesToken:            token,
		AccesTokenExpiresAt:   Accespayload.ExpiredAt.UTC().Local(),
	}
	return c.JSON(http.StatusOK, resp)
}

func ValidationCreateUserRequest(input *CreateUserParam) (errors []string) {
	if err := ValidateAlphanum(input.Username, 4, 20); err != nil {
		errors = append(errors, ValidateError("username", err.Error()))
	}
	if err := ValidateAlpha(input.FullName, 3, 50); err != nil {
		errors = append(errors, ValidateError("full_name", err.Error()))
	}
	if err := ValidateEmail(input.Email, 5, 50); err != nil {
		errors = append(errors, ValidateError("email", err.Error()))
	}
	if err := validatePassword(input.HashedPassword, 5, 30); err != nil {
		errors = append(errors, ValidateError("password", err.Error()))
	}
	return errors
}

func ValidationLoginRequest(input *LoginParams) (errors []string) {
	if err := ValidateAlphanum(input.Username, 4, 20); err != nil {
		errors = append(errors, ValidateError("username", err.Error()))
	}
	if err := validatePassword(input.Password, 5, 30); err != nil {
		errors = append(errors, ValidateError("password", err.Error()))
	}
	return errors
}
