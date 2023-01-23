package api

import (
	"time"

	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"

	"github.com/google/uuid"
)

type BasicResponse map[string]interface{}

type (
	QueueResponse struct {
		Owner     string `json:"owner"`
		AccountID int64  `json:"accountid"`
	}
	OwnerGetAccountResponse struct {
		Account      db2.Account     `json:"account"`
		QueueAccount []QueueResponse `json:"queue"`
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
		PhototDir   string `json:"photo_dir"`
	}
	CreatePostResponse struct {
		ID                 uuid.UUID           `json:"post_id"`
		PictureDescription string              `json:"picture_description"`
		PostFeature        postfeatureresponse `json:"post_feature"`
		IsRetweet          bool                `json:"is_retweet"`
		CreatedAt          int64               `json:"created_at"`
	}
	postfeatureresponse struct {
		ID              uuid.UUID `json:"post_id"`
		SumComment      int64     `json:"sum_comment"`
		SumLike         int64     `json:"sum_like"`
		SumRetweet      int64     `json:"sum_retweet"`
		SumQouteRetweet int64     `json:"sum_qoute_retweet"`
		CreatedAt       int64     `json:"created_at"`
	}
	GetPostResponses struct {
		ID                 uuid.UUID       `json:"id"`
		PictureDescription string          `json:"picture_description"`
		PostFeature        db2.PostFeature `json:"post_feature"`
		PostComment        []commentresp   `json:"post_comment"`
		CreatedAt          int64           `json:"created_at"`
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
		PostID  uuid.UUID `json:"id"`
		SumLike int64     `json:"like"`
		LikeAT  int64     `json:"like_at"`
	}
	CommentPostResp struct {
		PostID     uuid.UUID `json:"id"`
		Comment    string    `json:"comment"`
		SumComment int64     `json:"sum_comment"`
		LikeAT     int64     `json:"like_at"`
	}
	commentresp struct {
		FromAccountID int64     `json:"from_account_id"`
		CommentID     uuid.UUID `json:"comment_id"`
		Comment       string    `json:"comment"`
		SumLike       int64     `json:"sum_like"`
		CreatedAt     int64     `json:"created_at"`
	}
	RetweetPostResp struct {
		PostID      uuid.UUID          `json:"id"`
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
	RetweetResponse struct {
		Post    db2.Post
		Feature db2.PostFeature
	}
	PublicAccountResponse struct {
		ID          int64  `json:"id"`
		Owner       string `json:"owner"`
		AccountType bool   `json:"is_private"`
		Follower    int64  `json:"follower"`
		Following   int64  `json:"following"`
		CreatedAt   int64  `json:"created_at"`
		PhotoDir    string `json:"photo_dir"`
	}
	PublicAccountResp struct {
		Accounts []PublicAccountResponse `json:"accounts"`
		PageInfo map[string]interface{}  `json:"page_info"`
	}
)

func createaccountfollowresp(arg db2.AccountsFollow) accountfollowresp {
	return accountfollowresp{
		FromAccountID: arg.FromAccountID,
		ToAccountID:   arg.ToAccountID,
		FollowAt:      arg.FollowAt.Unix(),
	}
}

func CreateUserResponses(input db2.User, input2 CreateAccountsResponse) CreateUserResponse {
	return CreateUserResponse{
		Username:  input.Username,
		FullName:  input.FullName,
		Email:     input.Email,
		Account:   input2,
		CreatedAt: input.CreatedAt.Unix(),
	}
}
func UserResponse(input db2.User, account db2.Account) CreateUserResponse {
	return CreateUserResponse{
		Username:  input.Username,
		FullName:  input.FullName,
		Email:     input.Email,
		Account:   AccountResponse(account),
		CreatedAt: input.CreatedAt.Unix(),
	}
}

func OwnerAccountResponse(Account db2.Account, Queue ...db2.ListQueueRow) OwnerGetAccountResponse {
	return OwnerGetAccountResponse{
		Account:      Account,
		QueueAccount: queueconverter(Queue),
	}
}

func AccountResponse(input db2.Account) CreateAccountsResponse {
	var ProfilePhoto string
	if input.PhotoDir.Valid {
		ProfilePhoto = input.PhotoDir.String
	}

	return CreateAccountsResponse{
		ID:          input.ID,
		Owner:       input.Owner,
		AccountType: input.IsPrivate,
		Follower:    input.Follower,
		Following:   input.Following,
		PhototDir:   ProfilePhoto,
		CreatedAt:   input.CreatedAt.Unix(),
	}
}
func postfeatureresp(input db2.PostFeature) postfeatureresponse {
	return postfeatureresponse{
		ID:              input.PostID,
		SumComment:      input.SumComment,
		SumLike:         input.SumLike,
		SumRetweet:      input.SumRetweet,
		SumQouteRetweet: input.SumQouteRetweet,
		CreatedAt:       input.CreatedAt.Unix(),
	}
}

func PostResponse(input db2.Post, input2 db2.PostFeature) CreatePostResponse {
	return CreatePostResponse{
		ID:                 input.ID,
		PictureDescription: input.PictureDescription,
		PostFeature:        postfeatureresp(input2),
		IsRetweet:          input.IsRetweet,
		CreatedAt:          input.CreatedAt.Unix(),
	}
}
func PostResponsePointer(input *db2.Post, input2 db2.PostFeature) CreatePostResponse {
	return CreatePostResponse{
		ID:                 input.ID,
		PictureDescription: input.PictureDescription,
		PostFeature:        postfeatureresp(input2),
		IsRetweet:          input.IsRetweet,
		CreatedAt:          input.CreatedAt.Unix(),
	}
}
func GetPostResponse(input db2.Post, input2 db2.PostFeature, comment []db2.ListCommentRow) GetPostResponses {
	return GetPostResponses{
		ID:                 input.ID,
		PictureDescription: input.PictureDescription,
		PostFeature:        input2,
		PostComment:        commentconverter(comment),
		CreatedAt:          input.CreatedAt.Unix(),
	}
}

func likeResponse(arg db2.PostFeature) LikePostResp {
	return LikePostResp{
		PostID:  arg.PostID,
		SumLike: arg.SumLike,
		LikeAT:  arg.CreatedAt.Unix(),
	}
}

func commentResponse(comment string, arg db2.PostFeature) CommentPostResp {
	return CommentPostResp{
		PostID:     arg.PostID,
		Comment:    comment,
		SumComment: arg.SumComment,
		LikeAT:     arg.CreatedAt.Unix(),
	}
}

func retweetResponse(postFeature db2.PostFeature, post db2.Post) RetweetPostResp {
	return RetweetPostResp{
		PostID:      post.ID,
		Postfeature: PostResponse(post, postFeature),
		RetweetAt:   post.CreatedAt.Unix(),
	}
}

func qouteretweetResponse(post db2.Post, postFeature db2.PostFeature, qoute string) QouteRetweetPostResp {
	return QouteRetweetPostResp{
		Qoute:       qoute,
		PostFeature: PostResponse(post, postFeature),
		RetweetAt:   postFeature.CreatedAt.Unix(),
	}
}

// func followResponse(follow db.FollowTXResult) FollowResponse {
// 	return FollowResponse{
// 		Follow:      createaccountfollowresp(follow.Follow),
// 		FromAccount: AccountResponse(follow.FromAcc),
// 		ToAccount:   AccountResponse(follow.ToAcc),
// 	}
// }

// TO BE IMPLEMENTED (GENERIC RETURN)
// type anyFeature interface {
// 	LikePostResp | CommentPostResp
// }

// func FeatureResponse[v anyFeature](arg v) v {
// 	return v[string]
// }

func queueconverter(arg []db2.ListQueueRow) []QueueResponse {
	result := make([]QueueResponse, len(arg))

	for i := range arg {
		result[i].Owner = arg[i].Owner
		result[i].AccountID = arg[i].FromAccountID.Int64
	}
	return result
}

func commentconverter(arg []db2.ListCommentRow) []commentresp {
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

func FuncPublicAccountResponse(input []db2.Account) []PublicAccountResponse {
	result := make([]PublicAccountResponse, len(input))
	for i := range input {
		result[i].ID = input[i].ID
		result[i].Owner = input[i].Owner
		result[i].AccountType = input[i].IsPrivate
		result[i].Follower = input[i].Follower
		result[i].Following = input[i].Following
		result[i].CreatedAt = input[i].CreatedAt.Unix()
		result[i].PhotoDir = input[i].PhotoDir.String
	}
	return result
}

func FuncPublicAccountsResp(input []db2.Account, pageInfo map[string]interface{}) PublicAccountResp {
	accounts := FuncPublicAccountResponse(input)
	return PublicAccountResp{
		Accounts: accounts,
		PageInfo: pageInfo,
	}
}
