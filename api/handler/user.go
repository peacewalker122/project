package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/usecase/user"
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

	uid, reqErr := s.contract.CreateNewUserRequest(c.Request().Context(), db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	})
	if reqErr != nil {
		return c.JSON(http.StatusBadRequest, reqErr)
	}

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

	res, err := s.contract.CreateUser(c.Request().Context(), requid, reqs.Token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	output := AccountResponse(res.Account)
	resp := CreateUserResponses(res.User, output)
	return c.JSON(http.StatusCreated, resp)
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
			return c.JSON(http.StatusNotFound, "user not found")
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

	loginparam := user.SessionParams{
		ID:        &account.ID,
		Username:  username.Username,
		Email:     username.Email,
		ClientIp:  c.RealIP(),
		UserAgent: c.Request().UserAgent(),
		IsBlocked: false,
	}

	result, loginErr := s.contract.Login(c.Request().Context(), loginparam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, loginErr)
	}

	resp := loginResp{
		SessionID:             result.Session.ID,
		RefreshToken:          result.RefreshToken,
		RefreshTokenExpiresAt: result.RefreshPayload.ExpiredAt.UTC().Local(),
		User:                  UserResponse(username, account),
		AccesToken:            result.AccessToken,
		AccesTokenExpiresAt:   result.AccessPayload.ExpiredAt.UTC().Local(),
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
