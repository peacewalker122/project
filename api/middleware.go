package api

import "github.com/labstack/echo/v4/middleware"

func Logger() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, host:${host}, user_agent:${user_agent}\n",
	}
}