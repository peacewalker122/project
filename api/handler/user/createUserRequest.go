package user

import (
	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
	"net/http"
)

func (u *UserHandler) CreateUserRequest(c echo.Context) error {
	req := new(CreateUserParam)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	uid, reqErr := u.user.CreateNewUserRequest(c.Request().Context(), db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	})
	if reqErr != nil {
		return c.JSON(http.StatusInternalServerError, reqErr.Error())
	}

	c.Response().Header().Add("uuid", uid.String())
	return c.JSON(http.StatusOK, "success")
}
