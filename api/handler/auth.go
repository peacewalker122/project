package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/auth"
)

type AuthHandler interface {
	AuthUser(c echo.Context) error
	ChangePassword(c echo.Context) error
}

type ValidateChangePassRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

func (s *Handler) AuthUser(c echo.Context) error {
	req := new(ValidateChangePassRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := uuid.New()

	s.contract.CreateRequest(c.Request().Context(), auth.AuthParams{
		Email:    req.Email,
		UUID:     id,
		ClientIp: c.RealIP(),
	})

	return c.JSON(http.StatusCreated, id.String())
}

type ChangePasswordRequest struct {
	NewPassword string `json:"password" form:"password" validate:"required"`
}

// u need to login first
func (s *Handler) ChangePassword(c echo.Context) error {
	req := new(ChangePasswordRequest)
	ctx := c.Request().Context()

	if err = c.Bind(req); err != nil {
		return err
	}

	uid := c.Param("uid")

	s.contract.ChangePasswordAuth(ctx, auth.ChangePassParams{
		Password: req.NewPassword,
		UUID:     uid,
		ClientIp: c.RealIP(),
	})

	if err = c.Validate(req); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}
