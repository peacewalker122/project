package api

import (
	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
)

type Server struct {
	store *db.Store  
	router *echo.Echo
}

func Newserver(store *db.Store) *Server{
	server := &Server{store: store}
	router := echo.New()

	router.POST("/user",server.CreateUser)

	server.router = router
	return server
}