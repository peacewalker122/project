package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	db "github.com/peacewalker122/project/db/sqlc"
)

type CreateUserParam struct {
	Username       string `json:"username" validate:"required,alphanum"`
	HashedPassword string `json:"hashed_password" validate:"required"`
	FullName       string `json:"full_name" validate:"required,alpha"`
	Email          string `json:"email" validate:"required,email"`
}

func (s *Server) createUser(c echo.Context) error {
	req := new(CreateUserParam)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
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
