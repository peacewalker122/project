package post

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/post"
)

func (p *PostHandler) CreateQouteRetweet(c echo.Context) error {
	req := new(QouteRetweetParams)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	postID, err := uuid.Parse(req.PostID)
	if err != nil {
		return c.JSON(400, err)
	}

	errNum, payload, err := p.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err)
	}

	err = p.post.CreateQouteRetweet(c.Request().Context(), &post.QouteRetweetRequest{
		AccountID: payload.AccountID,
		PostID:    postID,
	})
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(201, "success")
}
