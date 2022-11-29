package api

import (
	"net/http"
	"os"
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

func (s *Server) TimeoutPost() middleware.TimeoutConfig {
	return middleware.TimeoutConfig{
		ErrorMessage: "timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			// we delete the file if its already timeout
			if _, err := os.Stat(s.fileString); err == nil {
				os.Remove(s.fileString)
			}
			c.Error(err)
			c.SetHandler(s.timeout)
		},
		Timeout: 8 * time.Second,
	}
}
