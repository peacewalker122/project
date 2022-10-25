package api

import (
	"time"

	"github.com/google/uuid"
	db "github.com/peacewalker122/project/db/sqlc"
)

type (
	CreateUserResponse struct {
		Username  string    `json:"username"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}
	CreateAccountsResponse struct {
		ID          int64     `json:"id"`
		Owner       string    `json:"owner"`
		AccountType bool      `json:"account_type"`
		CreatedAt   time.Time `json:"created_at"`
	}
	CreatePostResponse struct {
		ID                 int64          `json:"id"`
		PictureDescription string         `json:"picture_description"`
		PostFeature        db.PostFeature `json:"post_feature"`
		CreatedAt          time.Time      `json:"created_at"`
	}
	GetPostResponses struct {
		ID                 int64               `json:"id"`
		PictureDescription string              `json:"picture_description"`
		PostFeature        db.PostFeature      `json:"post_feature"`
		PostComment        []db.ListCommentRow `json:"post_comment"`
		CreatedAt          time.Time           `json:"created_at"`
	}
	loginResp struct {
		SessionID             uuid.UUID          `json:"session_id"`
		RefreshToken          string             `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time          `json:"refresh_token_expires_at"`
		User                  CreateUserResponse `json:"user"`
		AccesToken            string             `json:"acc_token"`
		AccesTokenExpiresAt   time.Time          `json:"acces_token_expire_sat"`
	}
	AccesTokenResp struct {
		AccesToken          string    `json:"access_token"`
		AccesTokenExpiresAt time.Time `json:"access_token_expires_at"`
	}
	LikePostResp struct {
		PostID  int64     `json:"id"`
		SumLike int64     `json:"like"`
		LikeAT  time.Time `json:"like_at"`
	}
	CommentPostResp struct {
		PostID     int64     `json:"id"`
		Comment    string    `json:"comment"`
		SumComment int64     `json:"sum_comment"`
		LikeAT     time.Time `json:"like_at"`
	}
	// commentresp struct {
	// 	FromAccountID int64  `json:"from_account_id"`
	// 	Comment       string `json:"comment"`
	// 	CreatedAt     int    `json:"created_at"`
	// }
	RetweetPostResp struct {
		PostID     int64     `json:"id"`
		SumRetweet int64     `json:"sum_retweet"`
		RetweetAt  time.Time `json:"retweet_at"`
	}
)

func UserResponse(input db.User) CreateUserResponse {
	return CreateUserResponse{
		Username:  input.Username,
		FullName:  input.FullName,
		Email:     input.Email,
		CreatedAt: input.CreatedAt,
	}
}

func AccountResponse(input db.Account) CreateAccountsResponse {
	return CreateAccountsResponse{
		ID:          input.AccountsID,
		Owner:       input.Owner,
		AccountType: input.IsPrivate,
		CreatedAt:   input.CreatedAt,
	}
}

func PostResponse(input db.Post, input2 db.PostFeature) CreatePostResponse {
	return CreatePostResponse{
		ID:                 input.PostID,
		PictureDescription: input.PictureDescription,
		PostFeature:        input2,
		CreatedAt:          input.CreatedAt,
	}
}
func GetPostResponse(input db.Post, input2 db.PostFeature, comment []db.ListCommentRow) GetPostResponses {
	return GetPostResponses{
		ID:                 input.PostID,
		PictureDescription: input.PictureDescription,
		PostFeature:        input2,
		PostComment:        comment,
		CreatedAt:          input.CreatedAt,
	}
}

func likeResponse(arg db.PostFeature) LikePostResp {
	return LikePostResp{
		PostID:  arg.PostID,
		SumLike: arg.SumLike,
		LikeAT:  arg.CreatedAt.UTC(),
	}
}

func commentResponse(comment string, arg db.PostFeature) CommentPostResp {
	return CommentPostResp{
		PostID:     arg.PostID,
		Comment:    comment,
		SumComment: arg.SumComment,
		LikeAT:     arg.CreatedAt.UTC(),
	}
}

func retweetResponse(arg db.PostFeature) RetweetPostResp {
	return RetweetPostResp{
		PostID:     arg.PostID,
		SumRetweet: arg.SumRetweet,
		RetweetAt:  arg.CreatedAt,
	}
}

// TO BE IMPLEMENTED (GENERIC RETURN)
// type anyFeature interface {
// 	LikePostResp | CommentPostResp
// }

// func FeatureResponse[v anyFeature](arg v) v {
// 	return v[string]
// }

// func commentconverter(arg []db.ListCommentRow) <-chan []commentresp {
// 	lenarg := len(arg)
// 	res := make(chan []commentresp, lenarg)
// 	for i := range arg {
// 		select {
// 		case <-res:

// 		}
// 	}
// 	return res
// }
