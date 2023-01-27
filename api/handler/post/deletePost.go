package post

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (p *PostHandler) DeletePost(c echo.Context) error {
	req := new(DeletePostParams)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	postid, err := uuid.Parse(req.PostID)
	if err != nil {
		return err
	}

	errNum, payload, err := p.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err)
	}

	err = p.post.DeletePost(c.Request().Context(), payload.AccountID, postid)
	if err != nil {
		return err
	}

	return nil
}
