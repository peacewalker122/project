package post

type CreatePostParams struct {
	PictureDescription string `json:"picture_description" form:"picture_description" validate:"required"`
	AccountID          int64  `json:"account_id" form:"id" validate:"required"`
}
