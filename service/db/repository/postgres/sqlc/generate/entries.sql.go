// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: entries.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createEntries = `-- name: CreateEntries :one
INSERT INTO entries (
    from_account_id,
    to_account_id,
    post_id,
    type_entries
 ) VALUES (
    $1,$2,$3,$4
  ) RETURNING entries_id, from_account_id, to_account_id, post_id, type_entries, created_at
`

type CreateEntriesParams struct {
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	PostID        uuid.UUID `json:"post_id"`
	TypeEntries   string    `json:"type_entries"`
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntries,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.PostID,
		arg.TypeEntries,
	)
	var i Entry
	err := row.Scan(
		&i.EntriesID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.PostID,
		&i.TypeEntries,
		&i.CreatedAt,
	)
	return i, err
}

const getEntries = `-- name: GetEntries :one
SELECT entries_id, from_account_id, to_account_id, post_id, type_entries, created_at FROM entries
WHERE entries_id = $1 LIMIT 1
`

func (q *Queries) GetEntries(ctx context.Context, entriesID int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntries, entriesID)
	var i Entry
	err := row.Scan(
		&i.EntriesID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.PostID,
		&i.TypeEntries,
		&i.CreatedAt,
	)
	return i, err
}

const getEntriesFull = `-- name: GetEntriesFull :exec
SELECT entries_id, from_account_id, to_account_id, post_id, type_entries, created_at FROM entries
WHERE post_id = $1 and from_account_id = $2 and type_entries = $3 LIMIT 1
`

type GetEntriesFullParams struct {
	PostID        uuid.UUID `json:"post_id"`
	FromAccountID int64     `json:"from_account_id"`
	TypeEntries   string    `json:"type_entries"`
}

func (q *Queries) GetEntriesFull(ctx context.Context, arg GetEntriesFullParams) error {
	_, err := q.db.ExecContext(ctx, getEntriesFull, arg.PostID, arg.FromAccountID, arg.TypeEntries)
	return err
}

const listEntries = `-- name: ListEntries :many
SELECT entries_id, from_account_id, to_account_id, post_id, type_entries, created_at FROM entries
WHERE post_id = $1
ORDER BY entries_id
LIMIT $2
OFFSET $3
`

type ListEntriesParams struct {
	PostID uuid.UUID `json:"post_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.EntriesID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.PostID,
			&i.TypeEntries,
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
