package post

import (
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/post"
	"net/http"
)

func (p *PostHandler) CreatePost(c echo.Context) error {
	req := new(CreatePostParams)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	errNum, payload, err := p.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err)
	}

	postFile, postFileHeader, err := c.Request().FormFile("file")
	if err != nil {
		if err != http.ErrMissingFile {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	postRequest := post.CreatePostRequest{
		File:               postFile,
		FileHeader:         postFileHeader,
		AccountID:          payload.AccountID,
		PictureDescription: req.PictureDescription,
	}

	postData, err := p.post.CreatePost(c.Request().Context(), &postRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, postData)
}
