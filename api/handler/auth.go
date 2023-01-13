package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/api/util"
	"github.com/peacewalker122/project/db/repository/postgres/ent"
	"golang.org/x/crypto/bcrypt"
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

	errchan := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		err := s.util.ChangePasswordAuth(c.Request().Context(), util.SendEmail{
			Params:   []string{req.Email, c.RealIP(), id.String()},
			Type:     "change_password",
			TimeSend: time.Now(),
		})
		if err != nil {
			errchan <- err
		}

		tempVal, err := s.store.GetAllWithEmail(c.Request().Context(), req.Email)
		if err != nil {
			errchan <- err
		}

		err = s.redis.Set(c.Request().Context(), id.String(), tempVal, 5*time.Minute)
		if err != nil {
			errchan <- err
		}

		done <- struct{}{}
	}()
	select {
	case <-done:
	case err := <-errchan:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

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

	if err = c.Validate(req); err != nil {
		return err
	}

	var payload *ent.Users

	err = s.util.GetRedisPayload(ctx, uid, &payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	accountID, _ := s.store.GetAccountsOwner(ctx, payload.Username)

	pass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	errchan := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		err = s.util.SendEmailWithNotif(ctx, util.SendEmail{
			AccountID: []int64{accountID.ID},
			Type:      "password-changing",
			Params:    []string{payload.Email, payload.Username, c.RealIP()},
		})
		if err != nil {
			errchan <- err
		}
		done <- struct{}{}
	}()
	select {
	case <-done:
	case err := <-errchan:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = s.store.SetPassword(ctx, payload.Username, string(pass))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}
