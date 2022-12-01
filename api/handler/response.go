package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	auth "github.com/peacewalker122/project/api/auth"
	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

func NewHandler(store db.Store, redis redis.Store, util util.Config, token token.Maker) *Handler {
	return &Handler{
		store: store,
		redis: redis,
		util:  util,
		token: token,
	}
}

func (s *Handler) AuthAccount(c echo.Context, accountID int64) (int, error) {
	var acc db.Account

	authParam, ok := c.Get(auth.AuthPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return http.StatusInternalServerError, err
	}

	key := "session:AccountsID:" + strconv.Itoa(int(accountID))

	// Here we get from our redis cache value
	session, err := s.redis.Get(c.Request().Context(), key)
	if err != nil {
		// if it doesn't exist then query it from thes postgres
		acc, err = s.store.GetAccounts(c.Request().Context(), accountID)
		if num, err := GetErrorValidator(c, err, accountag); err != nil {
			return num, err
		}

		// set into redis cache
		err = s.redis.Set(c.Request().Context(), key, acc)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		// checking
		if acc.Owner != authParam.Username {
			err := errors.New("unauthorized Username for this account")
			return http.StatusUnauthorized, err
		}

		return 0, nil
	}
	// unmarshaling because we marshal into json for any value into set.
	err = json.Unmarshal([]byte(session), &acc)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// checking
	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized Username for this account")
		return http.StatusUnauthorized, err
	}
	return 0, nil
}

type (
	Handler struct {
		store db.Store
		redis redis.Store
		util  util.Config
		token token.Maker
	}

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
		Follower    int64  `json:"follower"`
		Following   int64  `json:"following"`
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
		PostID      int64              `json:"id"`
		Postfeature CreatePostResponse `json:"post_feature"`
		RetweetAt   int64              `json:"retweet_at"`
	}
	QouteRetweetPostResp struct {
		Qoute       string             `json:"qoute"`
		PostFeature CreatePostResponse `json:"post_feature"`
		RetweetAt   int64              `json:"retweet_at"`
	}
	accountfollowresp struct {
		FromAccountID int64 `json:"from_account_id"`
		ToAccountID   int64 `json:"to_account_id"`
		FollowAt      int64 `json:"follow_at"`
	}
	FollowResponse struct {
		Follow      accountfollowresp      `json:"follow_info"`
		FromAccount CreateAccountsResponse `json:"from_account"`
		ToAccount   CreateAccountsResponse `json:"to_account"`
	}
)

func createaccountfollowresp(arg db.AccountsFollow) accountfollowresp {
	return accountfollowresp{
		FromAccountID: arg.FromAccountID,
		ToAccountID:   arg.ToAccountID,
		FollowAt:      arg.FollowAt.Unix(),
	}
}

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
		Follower:    input.Follower,
		Following:   input.Following,
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
		LikeAT:  arg.CreatedAt.Unix(),
	}
}

func commentResponse(comment string, arg db.PostFeature) CommentPostResp {
	return CommentPostResp{
		PostID:     arg.PostID,
		Comment:    comment,
		SumComment: arg.SumComment,
		LikeAT:     arg.CreatedAt.Unix(),
	}
}

func retweetResponse(postFeature db.PostFeature, post db.Post) RetweetPostResp {
	return RetweetPostResp{
		PostID:      post.PostID,
		Postfeature: PostResponse(post, postFeature),
		RetweetAt:   post.CreatedAt.Unix(),
	}
}

func qouteretweetResponse(post db.Post, postFeature db.PostFeature, qoute string) QouteRetweetPostResp {
	return QouteRetweetPostResp{
		Qoute:       qoute,
		PostFeature: PostResponse(post, postFeature),
		RetweetAt:   postFeature.CreatedAt.Unix(),
	}
}

func followResponse(follow db.FollowTXResult) FollowResponse {
	return FollowResponse{
		Follow:      createaccountfollowresp(follow.Follow),
		FromAccount: AccountResponse(follow.FromAcc),
		ToAccount:   AccountResponse(follow.ToAcc),
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
