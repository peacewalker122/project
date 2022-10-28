package api

import (
	"time"

	"github.com/google/uuid"
	db "github.com/peacewalker122/project/db/sqlc"
)

type (
	CreateUserResponse struct {
		Username  string                 `json:"username"`
		FullName  string                 `json:"full_name"`
		Email     string                 `json:"email"`
		Account   CreateAccountsResponse `json:"account"`
		CreatedAt int64                  `json:"created_at"`
	}
	CreateAccountsResponse struct {
		ID          int64  `json:"id"`
		Owner       string `json:"owner"`
		AccountType bool   `json:"is_private"`
		CreatedAt   int64  `json:"created_at"`
	}
	CreatePostResponse struct {
		ID                 int64               `json:"post_id"`
		PictureDescription string              `json:"picture_description"`
		PostFeature        postfeatureresponse `json:"post_feature"`
		IsRetweet          bool                `json:"is_retweet"`
		CreatedAt          int64               `json:"created_at"`
	}
	postfeatureresponse struct {
		ID              int64 `json:"post_id"`
		SumComment      int64 `json:"sum_comment"`
		SumLike         int64 `json:"sum_like"`
		SumRetweet      int64 `json:"sum_retweet"`
		SumQouteRetweet int64 `json:"sum_qoute_retweet"`
		CreatedAt       int64 `json:"created_at"`
	}
	GetPostResponses struct {
		ID                 int64          `json:"id"`
		PictureDescription string         `json:"picture_description"`
		PostFeature        db.PostFeature `json:"post_feature"`
		PostComment        []commentresp  `json:"post_comment"`
		CreatedAt          int64          `json:"created_at"`
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
		PostID  int64 `json:"id"`
		SumLike int64 `json:"like"`
		LikeAT  int64 `json:"like_at"`
	}
	CommentPostResp struct {
		PostID     int64  `json:"id"`
		Comment    string `json:"comment"`
		SumComment int64  `json:"sum_comment"`
		LikeAT     int64  `json:"like_at"`
	}
	commentresp struct {
		FromAccountID int64  `json:"from_account_id"`
		CommentID     int64  `json:"comment_id"`
		Comment       string `json:"comment"`
		SumLike       int64  `json:"sum_like"`
		CreatedAt     int64  `json:"created_at"`
	}
	RetweetPostResp struct {
		PostID     int64 `json:"id"`
		SumRetweet int64 `json:"sum_retweet"`
		RetweetAt  int64 `json:"retweet_at"`
	}
	QouteRetweetPostResp struct {
		Qoute       string             `json:"qoute"`
		PostFeature CreatePostResponse `json:"post_feature"`
		RetweetAt   int64              `json:"retweet_at"`
	}
)

func CreateUserResponses(input db.User, input2 CreateAccountsResponse) CreateUserResponse {
	return CreateUserResponse{
		Username:  input.Username,
		FullName:  input.FullName,
		Email:     input.Email,
		Account:   input2,
		CreatedAt: input.CreatedAt.Unix(),
	}
}
func UserResponse(input db.User) CreateUserResponse {
	return CreateUserResponse{
		Username:  input.Username,
		FullName:  input.FullName,
		Email:     input.Email,
		CreatedAt: input.CreatedAt.Unix(),
	}
}

func AccountResponse(input db.Account) CreateAccountsResponse {
	return CreateAccountsResponse{
		ID:          input.AccountsID,
		Owner:       input.Owner,
		AccountType: input.IsPrivate,
		CreatedAt:   input.CreatedAt.Unix(),
	}
}
func postfeatureresp(input db.PostFeature) postfeatureresponse {
	return postfeatureresponse{
		ID:              input.PostID,
		SumComment:      input.SumComment,
		SumLike:         input.SumLike,
		SumRetweet:      input.SumRetweet,
		SumQouteRetweet: input.SumQouteRetweet,
		CreatedAt:       input.CreatedAt.Unix(),
	}
}

func PostResponse(input db.Post, input2 db.PostFeature) CreatePostResponse {
	return CreatePostResponse{
		ID:                 input.PostID,
		PictureDescription: input.PictureDescription,
		PostFeature:        postfeatureresp(input2),
		IsRetweet:          input.IsRetweet,
		CreatedAt:          input.CreatedAt.Unix(),
	}
}
func PostResponsePointer(input *db.Post, input2 db.PostFeature) CreatePostResponse {
	return CreatePostResponse{
		ID:                 input.PostID,
		PictureDescription: input.PictureDescription,
		PostFeature:        postfeatureresp(input2),
		IsRetweet:          input.IsRetweet,
		CreatedAt:          input.CreatedAt.Unix(),
	}
}
func GetPostResponse(input db.Post, input2 db.PostFeature, comment []db.ListCommentRow) GetPostResponses {
	return GetPostResponses{
		ID:                 input.PostID,
		PictureDescription: input.PictureDescription,
		PostFeature:        input2,
		PostComment:        commentconverter(comment),
		CreatedAt:          input.CreatedAt.Unix(),
	}
}

func likeResponse(arg db.PostFeature) LikePostResp {
	return LikePostResp{
		PostID:  arg.PostID,
		SumLike: arg.SumLike,
		LikeAT:  arg.CreatedAt.UTC().Unix(),
	}
}

func commentResponse(comment string, arg db.PostFeature) CommentPostResp {
	return CommentPostResp{
		PostID:     arg.PostID,
		Comment:    comment,
		SumComment: arg.SumComment,
		LikeAT:     arg.CreatedAt.UTC().Unix(),
	}
}

func retweetResponse(arg db.PostFeature) RetweetPostResp {
	return RetweetPostResp{
		PostID:     arg.PostID,
		SumRetweet: arg.SumRetweet,
		RetweetAt:  arg.CreatedAt.Unix(),
	}
}

func qouteretweetResponse(post db.Post, postFeature db.PostFeature, qoute string) QouteRetweetPostResp {
	return QouteRetweetPostResp{
		Qoute:       qoute,
		PostFeature: PostResponse(post, postFeature),
		RetweetAt:   postFeature.CreatedAt.Unix(),
	}
}

// TO BE IMPLEMENTED (GENERIC RETURN)
// type anyFeature interface {
// 	LikePostResp | CommentPostResp
// }

// func FeatureResponse[v anyFeature](arg v) v {
// 	return v[string]
// }

func commentconverter(arg []db.ListCommentRow) []commentresp {
	result := make([]commentresp, len(arg))
	for i := range arg {
		result[i].Comment = arg[i].Comment
		result[i].FromAccountID = arg[i].FromAccountID
		result[i].CommentID = arg[i].CommentID
		result[i].SumLike = arg[i].SumLike
		result[i].SumLike = arg[i].SumLike
		result[i].CreatedAt = arg[i].CreatedAt.Unix()
	}
	return result
}
