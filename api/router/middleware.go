package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).
		WithTime(time.Now()).
		Info("Request")
		return next(c)
	}
}

func Redirect() middleware.RedirectConfig {
	return middleware.RedirectConfig{
		Code: http.StatusMovedPermanently,
	}
}

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"At": time.Now().Format("15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"Time":   time.Now().Format("15:04:05"),
		"Method": c.Request().Method,
		"Status": c.Response().Status,
	})
}
