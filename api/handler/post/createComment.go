package post

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/post"
)

func (p *PostHandler) CreateComment(c echo.Context) error {
	req := new(CommentParams)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	postID, err := uuid.Parse(req.PostID)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	errNum, payload, err := p.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err)
	}

	errComment := p.post.CreateComment(c.Request().Context(), &post.CommentRequest{
		AccountID: payload.AccountID,
		PostID:    postID,
		Comment:   req.Comment,
	})
	if errComment.HasError() {
		return c.JSON(400, errComment.Error())
	}

	return c.JSON(201, "success")
}
