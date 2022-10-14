package api

import (
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
	resp := UserResponse(user)
	return c.JSON(http.StatusOK, resp)
}

func ValidationCreateUserRequest(input *CreateUserParam) (errors []string) {
	if err := ValidateAlphanum(input.Username); err != nil {
		errors = append(errors, ValidateError("username", err.Error()))
	}
	if err := ValidateAlpha(input.FullName); err != nil {
		errors = append(errors, ValidateError("full_name", err.Error()))
	}
	if err := ValidateEmail(input.Email); err != nil {
		errors = append(errors, ValidateError("email", err.Error()))
	}
	if err := validatePassword(input.HashedPassword); err != nil {
		errors = append(errors, ValidateError("password", err.Error()))
	}
	return errors
}
