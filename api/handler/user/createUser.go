package user

import (
	"github.com/labstack/echo/v4"
	api "github.com/peacewalker122/project/api/handler"
	"net/http"
)

func (u *UserHandler) CreateUser(c echo.Context) error {
	reqs := new(CreatingUser)
	if err := c.Bind(reqs); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(reqs); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	requid := c.Param("uuid")

	res, err := u.user.CreateUser(c.Request().Context(), requid, reqs.Token)
	if err != nil {
		c.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	output := api.AccountResponse(res.Account)
	resp := api.CreateUserResponses(res.User, output)
	return c.JSON(http.StatusCreated, resp)
}
