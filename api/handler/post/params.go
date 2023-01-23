package post

type CreatePostParams struct {
	PictureDescription string `json:"picture_description" form:"picture_description" validate:"required"`
	AccountID          int64  `json:"account_id" form:"id" validate:"required"`
}

type LikeParams struct {
	AccountID int64  `json:"account_id" form:"id" validate:"required"`
	PostID    string `json:"post_id" form:"post_id" query:"post_id" validate:"required"`
}

type CommentParams struct {
	AccountID int64  `json:"account_id" form:"id" validate:"required"`
	PostID    string `json:"post_id" form:"post_id" query:"post_id" validate:"required"`
	Comment   string `json:"comment" form:"comment" query:"comment" validate:"required"`
}

type RetweetParams struct {
	AccountID int64  `json:"account_id" form:"id" validate:"required"`
	PostID    string `json:"post_id" form:"post_id" query:"post_id" validate:"required"`
}

type QouteRetweetParams struct {
	AccountID int64  `json:"account_id" form:"id" validate:"required"`
	PostID    string `json:"post_id" form:"post_id" query:"post_id" validate:"required"`
	Qoute     string `json:"qoute" form:"qoute" query:"qoute" validate:"required"`
}
