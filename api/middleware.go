package api

import (
	"github.com/labstack/echo/v4/middleware"
)

func Logger() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format:           "method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}
}
