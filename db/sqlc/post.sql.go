// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: post.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO post(
    account_id,
    picture_description
) VALUES(
    $1,$2
) RETURNING id, account_id, picture_description, created_at
`

type CreatePostParams struct {
	AccountID          int64          `json:"account_id"`
	PictureDescription sql.NullString `json:"picture_description"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.AccountID, arg.PictureDescription)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PictureDescription,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM post
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, account_id, picture_description, created_at FROM post
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PictureDescription,
		&i.CreatedAt,
	)
	return i, err
}

const listPost = `-- name: ListPost :many
SELECT id, account_id, picture_description, created_at FROM post
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPostParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPost(ctx context.Context, arg ListPostParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPost, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.PictureDescription,
			&i.CreatedAt,
		); err != nil {
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

const updatePost = `-- name: UpdatePost :one
UPDATE post
SET picture_description = $2
WHERE id = $1
RETURNING id, account_id, picture_description, created_at
`

type UpdatePostParams struct {
	ID                 int64          `json:"id"`
	PictureDescription sql.NullString `json:"picture_description"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.ID, arg.PictureDescription)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PictureDescription,
		&i.CreatedAt,
	)
	return i, err
}
