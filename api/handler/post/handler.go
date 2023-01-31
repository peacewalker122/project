package post

import (
	"github.com/labstack/echo/v4"
	handler "github.com/peacewalker122/project/api/handler"
	"github.com/peacewalker122/project/contract/post"
)

type PostHandler struct {
	post   post.PostContract
	helper handler.Helper
}

func NewPostHandler(post post.PostContract, helper handler.Helper) *PostHandler {
	return &PostHandler{
		post:   post,
		helper: helper,
	}
}

func (p *PostHandler) PostRouter(e *echo.Group) {
	e.POST("/post", p.CreatePost)
	e.POST("/post/retweet", p.CreateRetweet)
	e.POST("/post", p.CreatePost)
	e.POST("/post/like", p.LikePost)
	e.POST("/post/comment", p.CreateComment)
	e.POST("/post/qoute/retweet", p.CreateQouteRetweet)
	e.PUT("/post/like", p.UnlikePost)
}
