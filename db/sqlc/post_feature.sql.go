// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: post_feature.sql

package db

import (
	"context"
	"time"
)

const createComment_feature = `-- name: CreateComment_feature :one
INSERT INTO comment_feature(
    from_account_id,
    comment,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING comment
`

type CreateComment_featureParams struct {
	FromAccountID int64  `json:"from_account_id"`
	Comment       string `json:"comment"`
	PostID        int64  `json:"post_id"`
}

func (q *Queries) CreateComment_feature(ctx context.Context, arg CreateComment_featureParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createComment_feature, arg.FromAccountID, arg.Comment, arg.PostID)
	var comment string
	err := row.Scan(&comment)
	return comment, err
}

const createLike_feature = `-- name: CreateLike_feature :exec
INSERT INTO like_feature(
    from_account_id,
    is_like,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING is_like
`

type CreateLike_featureParams struct {
	FromAccountID int64 `json:"from_account_id"`
	IsLike        bool  `json:"is_like"`
	PostID        int64 `json:"post_id"`
}

func (q *Queries) CreateLike_feature(ctx context.Context, arg CreateLike_featureParams) error {
	_, err := q.db.ExecContext(ctx, createLike_feature, arg.FromAccountID, arg.IsLike, arg.PostID)
	return err
}

const createPost_feature = `-- name: CreatePost_feature :one
INSERT INTO post_feature(
    post_id
) values(
    $1
) RETURNING post_id, sum_comment, sum_like, sum_retweet, sum_qoute_retweet, created_at
`

func (q *Queries) CreatePost_feature(ctx context.Context, postID int64) (PostFeature, error) {
	row := q.db.QueryRowContext(ctx, createPost_feature, postID)
	var i PostFeature
	err := row.Scan(
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
) RETURNING qoute_retweet
`

type CreateQouteRetweet_featureParams struct {
	FromAccountID int64  `json:"from_account_id"`
	QouteRetweet  bool   `json:"qoute_retweet"`
	Qoute         string `json:"qoute"`
	PostID        int64  `json:"post_id"`
}

func (q *Queries) CreateQouteRetweet_feature(ctx context.Context, arg CreateQouteRetweet_featureParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, createQouteRetweet_feature,
		arg.FromAccountID,
		arg.QouteRetweet,
		arg.Qoute,
		arg.PostID,
	)
	var qoute_retweet bool
	err := row.Scan(&qoute_retweet)
	return qoute_retweet, err
}

const createRetweet_feature = `-- name: CreateRetweet_feature :exec
INSERT INTO retweet_feature(
    from_account_id,
    retweet,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING retweet
`

type CreateRetweet_featureParams struct {
	FromAccountID int64 `json:"from_account_id"`
	Retweet       bool  `json:"retweet"`
	PostID        int64 `json:"post_id"`
}

func (q *Queries) CreateRetweet_feature(ctx context.Context, arg CreateRetweet_featureParams) error {
	_, err := q.db.ExecContext(ctx, createRetweet_feature, arg.FromAccountID, arg.Retweet, arg.PostID)
	return err
}

const getLikeInfo = `-- name: GetLikeInfo :one
SELECT from_account_id, is_like, post_id, created_at from like_feature
WHERE from_account_id = $1 and post_id = $2 LIMIT 1
`

type GetLikeInfoParams struct {
	FromAccountID int64 `json:"from_account_id"`
	PostID        int64 `json:"post_id"`
}

func (q *Queries) GetLikeInfo(ctx context.Context, arg GetLikeInfoParams) (LikeFeature, error) {
	row := q.db.QueryRowContext(ctx, getLikeInfo, arg.FromAccountID, arg.PostID)
	var i LikeFeature
	err := row.Scan(
		&i.FromAccountID,
		&i.IsLike,
		&i.PostID,
		&i.CreatedAt,
	)
	return i, err
}

const getLikejoin = `-- name: GetLikejoin :one
SELECT like_feature.is_like from like_feature
INNER JOIN post ON post.post_id = like_feature.post_id
WHERE post.post_id = $1
`

func (q *Queries) GetLikejoin(ctx context.Context, postID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, getLikejoin, postID)
	var is_like bool
	err := row.Scan(&is_like)
	return is_like, err
}

const getPostJoin = `-- name: GetPostJoin :one
SELECT post.post_id,post.account_id FROM post
INNER JOIN post_feature ON post_feature.post_id = post.post_id
WHERE post.post_id = $1
`

type GetPostJoinRow struct {
	PostID    int64 `json:"post_id"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) GetPostJoin(ctx context.Context, postID int64) (GetPostJoinRow, error) {
	row := q.db.QueryRowContext(ctx, getPostJoin, postID)
	var i GetPostJoinRow
	err := row.Scan(&i.PostID, &i.AccountID)
	return i, err
}

const getPost_feature = `-- name: GetPost_feature :one
SELECT post_id, sum_comment, sum_like, sum_retweet, sum_qoute_retweet, created_at FROM post_feature
WHERE post_id = $1 LIMIT 1
`

func (q *Queries) GetPost_feature(ctx context.Context, postID int64) (PostFeature, error) {
	row := q.db.QueryRowContext(ctx, getPost_feature, postID)
	var i PostFeature
	err := row.Scan(
		&i.PostID,
		&i.SumComment,
		&i.SumLike,
		&i.SumRetweet,
		&i.SumQouteRetweet,
		&i.CreatedAt,
	)
	return i, err
}

const getPost_feature_Update = `-- name: GetPost_feature_Update :one
SELECT post_id, sum_comment, sum_like, sum_retweet, sum_qoute_retweet, created_at FROM post_feature
WHERE post_id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetPost_feature_Update(ctx context.Context, postID int64) (PostFeature, error) {
	row := q.db.QueryRowContext(ctx, getPost_feature_Update, postID)
	var i PostFeature
	err := row.Scan(
		&i.PostID,
		&i.SumComment,
		&i.SumLike,
		&i.SumRetweet,
		&i.SumQouteRetweet,
		&i.CreatedAt,
	)
	return i, err
}

const getRetweet = `-- name: GetRetweet :one
SELECT from_account_id, retweet, post_id, created_at from retweet_feature
WHERE from_account_id = $1 and post_id = $2 LIMIT 1
`

type GetRetweetParams struct {
	FromAccountID int64 `json:"from_account_id"`
	PostID        int64 `json:"post_id"`
}

func (q *Queries) GetRetweet(ctx context.Context, arg GetRetweetParams) (RetweetFeature, error) {
	row := q.db.QueryRowContext(ctx, getRetweet, arg.FromAccountID, arg.PostID)
	var i RetweetFeature
	err := row.Scan(
		&i.FromAccountID,
		&i.Retweet,
		&i.PostID,
		&i.CreatedAt,
	)
	return i, err
}

const getRetweetJoin = `-- name: GetRetweetJoin :one
SELECT retweet_feature.retweet from retweet_feature
INNER JOIN post ON post.post_id = retweet_feature.post_id
WHERE post.post_id = $1
`

func (q *Queries) GetRetweetJoin(ctx context.Context, postID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, getRetweetJoin, postID)
	var retweet bool
	err := row.Scan(&retweet)
	return retweet, err
}

const listComment = `-- name: ListComment :many
SELECT from_account_id,comment,created_at from comment_feature
WHERE post_id = $1
ORDER by from_account_id
LIMIT $2
OFFSET $3
`

type ListCommentParams struct {
	PostID int64 `json:"post_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListCommentRow struct {
	FromAccountID int64     `json:"from_account_id"`
	Comment       string    `json:"comment"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) ListComment(ctx context.Context, arg ListCommentParams) ([]ListCommentRow, error) {
	rows, err := q.db.QueryContext(ctx, listComment, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCommentRow{}
	for rows.Next() {
		var i ListCommentRow
		if err := rows.Scan(&i.FromAccountID, &i.Comment, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLike = `-- name: UpdateLike :exec
UPDATE like_feature
set is_like = $1
WHERE post_id = $2 and from_account_id = $3
RETURNING is_like
`

type UpdateLikeParams struct {
	IsLike        bool  `json:"is_like"`
	PostID        int64 `json:"post_id"`
	FromAccountID int64 `json:"from_account_id"`
}

func (q *Queries) UpdateLike(ctx context.Context, arg UpdateLikeParams) error {
	_, err := q.db.ExecContext(ctx, updateLike, arg.IsLike, arg.PostID, arg.FromAccountID)
	return err
}

const updatePost_feature = `-- name: UpdatePost_feature :one
UPDATE post_feature
SET sum_comment = $2, sum_like = $3, sum_retweet = $4, sum_qoute_retweet =$5
WHERE post_id = $1
RETURNING post_id, sum_comment, sum_like, sum_retweet, sum_qoute_retweet, created_at
`

type UpdatePost_featureParams struct {
	PostID          int64 `json:"post_id"`
	SumComment      int64 `json:"sum_comment"`
	SumLike         int64 `json:"sum_like"`
	SumRetweet      int64 `json:"sum_retweet"`
	SumQouteRetweet int64 `json:"sum_qoute_retweet"`
}

func (q *Queries) UpdatePost_feature(ctx context.Context, arg UpdatePost_featureParams) (PostFeature, error) {
	row := q.db.QueryRowContext(ctx, updatePost_feature,
		arg.PostID,
		arg.SumComment,
		arg.SumLike,
		arg.SumRetweet,
		arg.SumQouteRetweet,
	)
	var i PostFeature
	err := row.Scan(
		&i.PostID,
		&i.SumComment,
		&i.SumLike,
		&i.SumRetweet,
		&i.SumQouteRetweet,
		&i.CreatedAt,
	)
	return i, err
}

const updateRetweet = `-- name: UpdateRetweet :exec
UPDATE retweet_feature
set retweet = $1
WHERE post_id = $2 and from_account_id = $3
RETURNING retweet
`

type UpdateRetweetParams struct {
	Retweet       bool  `json:"retweet"`
	PostID        int64 `json:"post_id"`
	FromAccountID int64 `json:"from_account_id"`
}

func (q *Queries) UpdateRetweet(ctx context.Context, arg UpdateRetweetParams) error {
	_, err := q.db.ExecContext(ctx, updateRetweet, arg.Retweet, arg.PostID, arg.FromAccountID)
	return err
}
