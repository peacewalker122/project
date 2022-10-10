package api

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
)

type CreatePostParams struct {
	AccountID   int64          `json:"account_id"`
	PostWord    sql.NullString `json:"post_word"`
	PostPicture []byte         `form:"post_picture"`
}

func (s *Server) CreatePost(c echo.Context) error {
	req := new(CreatePostParams)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	arg := db.CreatePostParams{
		AccountID:   req.AccountID,
		PostWord:    req.PostWord,
		PostPicture: req.PostPicture,
	}
	post, err := s.store.CreatePost(c.Request().Context(), arg)
	if err != nil {
		return nil, err
	}
}
