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
	CreateComment_feature(ctx context.Context, arg CreateComment_featureParams) (CommentFeature, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error)
	CreateLike_feature(ctx context.Context, arg CreateLike_featureParams) (LikeFeature, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreatePost_feature(ctx context.Context, arg CreatePost_featureParams) (PostFeature, error)
	CreateQouteRetweet_feature(ctx context.Context, arg CreateQouteRetweet_featureParams) (QouteRetweetFeature, error)
	CreateRetweet_feature(ctx context.Context, arg CreateRetweet_featureParams) (RetweetFeature, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeletePost(ctx context.Context, postID int64) error
	GetAccounts(ctx context.Context, accountsID int64) (Account, error)
	GetAccountsOwner(ctx context.Context, owner string) (Account, error)
	GetEntries(ctx context.Context, entriesID int64) (Entry, error)
	GetPost(ctx context.Context, postID int64) (Post, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListPost(ctx context.Context, arg ListPostParams) ([]Post, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
}

var _ Querier = (*Queries)(nil)
