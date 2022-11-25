package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
)

type CreateUserParam struct {
	Username       string `json:"username" validate:"required,min=4,max=100"`
	HashedPassword string `json:"password" validate:"required,min=6,max=100"`
	FullName       string `json:"full_name" validate:"required,min=3,max=100"`
	Email          string `json:"email" validate:"required,email"`
}

func (s *Server) createUser(c echo.Context) error {
	req := new(CreateUserParam)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if errors := ValidationCreateUserRequest(req); errors != nil {
		return c.JSONPretty(http.StatusBadRequest, errors, "    ")
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

func (s *Server) login(c echo.Context) error {
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

	err = util.CheckPassword(req.Password, username.HashedPassword)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ValidateError("password", err.Error()))
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
		User:                  UserResponse(username),
		AccesToken:            token,
		AccesTokenExpiresAt:   Accespayload.ExpiredAt.UTC().Local(),
	}
	//c.Response().Header().Add("refreshtoken",refreshToken)
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
