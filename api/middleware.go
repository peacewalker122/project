package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format:           "method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}
}

func Redirect() middleware.RedirectConfig {
	return middleware.RedirectConfig{
		Code: http.StatusMovedPermanently,
	}
}

func (s *Server) Timeout() middleware.TimeoutConfig {
	return middleware.TimeoutConfig{
		//ErrorMessage: "timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			c.Error(err)
			c.SetHandler(s.timeout)
		},
		Timeout: 10 * time.Second,
	}
}
