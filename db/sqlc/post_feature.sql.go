// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: post_feature.sql

package db

import (
	"context"
)

const createComment_feature = `-- name: CreateComment_feature :one
INSERT INTO comment_feature(
    from_account_id,
    comment,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING from_account_id, comment, post_id, created_at
`

type CreateComment_featureParams struct {
	FromAccountID int64  `json:"from_account_id"`
	Comment       string `json:"comment"`
	PostID        int64  `json:"post_id"`
}

func (q *Queries) CreateComment_feature(ctx context.Context, arg CreateComment_featureParams) (CommentFeature, error) {
	row := q.db.QueryRowContext(ctx, createComment_feature, arg.FromAccountID, arg.Comment, arg.PostID)
	var i CommentFeature
	err := row.Scan(
		&i.FromAccountID,
		&i.Comment,
		&i.PostID,
		&i.CreatedAt,
	)
	return i, err
}

const createLike_feature = `-- name: CreateLike_feature :one
INSERT INTO like_feature(
    from_account_id,
    is_like,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING from_account_id, is_like, post_id, created_at
`

type CreateLike_featureParams struct {
	FromAccountID int64 `json:"from_account_id"`
	IsLike        bool  `json:"is_like"`
	PostID        int64 `json:"post_id"`
}

func (q *Queries) CreateLike_feature(ctx context.Context, arg CreateLike_featureParams) (LikeFeature, error) {
	row := q.db.QueryRowContext(ctx, createLike_feature, arg.FromAccountID, arg.IsLike, arg.PostID)
	var i LikeFeature
	err := row.Scan(
		&i.FromAccountID,
		&i.IsLike,
		&i.PostID,
		&i.CreatedAt,
	)
	return i, err
}

const createPost_feature = `-- name: CreatePost_feature :one
INSERT INTO post_feature(
    from_account_id,
    post_id,
    sum_comment,
    sum_like,
    sum_retweet,
    sum_qoute_retweet
) values(
    $1,$2,$3,$4,$5,$6
) RETURNING from_account_id, post_id, sum_comment, sum_like, sum_retweet, sum_qoute_retweet, created_at
`

type CreatePost_featureParams struct {
	FromAccountID   int64  `json:"from_account_id"`
	PostID          int64  `json:"post_id"`
	SumComment      string `json:"sum_comment"`
	SumLike         int64  `json:"sum_like"`
	SumRetweet      int64  `json:"sum_retweet"`
	SumQouteRetweet int64  `json:"sum_qoute_retweet"`
}

func (q *Queries) CreatePost_feature(ctx context.Context, arg CreatePost_featureParams) (PostFeature, error) {
	row := q.db.QueryRowContext(ctx, createPost_feature,
		arg.FromAccountID,
		arg.PostID,
		arg.SumComment,
		arg.SumLike,
		arg.SumRetweet,
		arg.SumQouteRetweet,
	)
	var i PostFeature
	err := row.Scan(
		&i.FromAccountID,
		&i.PostID,
		&i.SumComment,
		&i.SumLike,
		&i.SumRetweet,
		&i.SumQouteRetweet,
		&i.CreatedAt,
	)
	return i, err
}

const createQouteRetweet_feature = `-- name: CreateQouteRetweet_feature :one
INSERT INTO qoute_retweet_feature(
    from_account_id,
    qoute_retweet,
    qoute,
    post_id
) VALUES(
    $1,$2,$3,$4
) RETURNING from_account_id, qoute_retweet, qoute, post_id, created_at
`

type CreateQouteRetweet_featureParams struct {
	FromAccountID int64  `json:"from_account_id"`
	QouteRetweet  bool   `json:"qoute_retweet"`
	Qoute         string `json:"qoute"`
	PostID        int64  `json:"post_id"`
}

func (q *Queries) CreateQouteRetweet_feature(ctx context.Context, arg CreateQouteRetweet_featureParams) (QouteRetweetFeature, error) {
	row := q.db.QueryRowContext(ctx, createQouteRetweet_feature,
		arg.FromAccountID,
		arg.QouteRetweet,
		arg.Qoute,
		arg.PostID,
	)
	var i QouteRetweetFeature
	err := row.Scan(
		&i.FromAccountID,
		&i.QouteRetweet,
		&i.Qoute,
		&i.PostID,
		&i.CreatedAt,
	)
	return i, err
}

const createRetweet_feature = `-- name: CreateRetweet_feature :one
INSERT INTO retweet_feature(
    from_account_id,
    retweet,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING from_account_id, retweet, post_id, created_at
`

type CreateRetweet_featureParams struct {
	FromAccountID int64 `json:"from_account_id"`
	Retweet       bool  `json:"retweet"`
	PostID        int64 `json:"post_id"`
}

func (q *Queries) CreateRetweet_feature(ctx context.Context, arg CreateRetweet_featureParams) (RetweetFeature, error) {
	row := q.db.QueryRowContext(ctx, createRetweet_feature, arg.FromAccountID, arg.Retweet, arg.PostID)
	var i RetweetFeature
	err := row.Scan(
		&i.FromAccountID,
		&i.Retweet,
		&i.PostID,
		&i.CreatedAt,
	)
	return i, err
}
