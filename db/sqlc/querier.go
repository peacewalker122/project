// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error)
	CreateComment_feature(ctx context.Context, arg CreateComment_featureParams) (string, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error)
	CreateLike_feature(ctx context.Context, arg CreateLike_featureParams) error
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreatePost_feature(ctx context.Context, postID int64) (PostFeature, error)
	CreateQouteRetweet_feature(ctx context.Context, arg CreateQouteRetweet_featureParams) (string, error)
	CreateRetweet_feature(ctx context.Context, arg CreateRetweet_featureParams) error
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeletePost(ctx context.Context, postID int64) error
	DeletePostFeature(ctx context.Context, postID int64) error
	DeleteQouteRetweet(ctx context.Context, arg DeleteQouteRetweetParams) error
	GetAccounts(ctx context.Context, accountsID int64) (Account, error)
	GetAccountsOwner(ctx context.Context, owner string) (Account, error)
	GetEntries(ctx context.Context, entriesID int64) (Entry, error)
	GetEntriesFull(ctx context.Context, arg GetEntriesFullParams) error
	GetLikeInfo(ctx context.Context, arg GetLikeInfoParams) (LikeFeature, error)
	GetLikejoin(ctx context.Context, postID int64) (bool, error)
	GetPost(ctx context.Context, postID int64) (Post, error)
	GetPostInfoJoin(ctx context.Context, arg GetPostInfoJoinParams) (int64, error)
	GetPostJoin(ctx context.Context, postID int64) (GetPostJoinRow, error)
	GetPostJoin_QouteRetweet(ctx context.Context, arg GetPostJoin_QouteRetweetParams) (bool, error)
	GetPost_feature(ctx context.Context, postID int64) (PostFeature, error)
	GetPost_feature_Update(ctx context.Context, postID int64) (PostFeature, error)
	GetQouteRetweet(ctx context.Context, arg GetQouteRetweetParams) (QouteRetweetFeature, error)
	GetQouteRetweetJoin(ctx context.Context, postID int64) (bool, error)
	GetQouteRetweetRows(ctx context.Context, arg GetQouteRetweetRowsParams) (int64, error)
	GetRetweet(ctx context.Context, arg GetRetweetParams) (RetweetFeature, error)
	GetRetweetJoin(ctx context.Context, postID int64) (bool, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSessionuser(ctx context.Context, username string) (GetSessionuserRow, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListComment(ctx context.Context, arg ListCommentParams) ([]ListCommentRow, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListPost(ctx context.Context, arg ListPostParams) ([]Post, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	UpdateLike(ctx context.Context, arg UpdateLikeParams) error
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdatePost_feature(ctx context.Context, arg UpdatePost_featureParams) (PostFeature, error)
	UpdateQouteRetweet(ctx context.Context, arg UpdateQouteRetweetParams) error
	UpdateRetweet(ctx context.Context, arg UpdateRetweetParams) error
}

var _ Querier = (*Queries)(nil)
