// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: accounts.sql

package db

import (
	"context"
)

const createAccounts = `-- name: CreateAccounts :one
INSERT INTO accounts(
    owner,
    is_private
) VALUES(
    $1,$2
) RETURNING accounts_id, owner, is_private, created_at, follower, following
`

type CreateAccountsParams struct {
	Owner     string `json:"owner"`
	IsPrivate bool   `json:"is_private"`
}

func (q *Queries) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccounts, arg.Owner, arg.IsPrivate)
	var i Account
	err := row.Scan(
		&i.AccountsID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
	)
	return i, err
}

const createPrivateQueue = `-- name: CreatePrivateQueue :one
INSERT INTO accounts_queue(
    from_account_id,
    to_account_id,
    queue
) VALUES(
    $1, $2, true
) RETURNING from_account_id, queue, to_account_id, queue_at
`

type CreatePrivateQueueParams struct {
	Fromaccountid int64 `json:"fromaccountid"`
	Toaccountid   int64 `json:"toaccountid"`
}

func (q *Queries) CreatePrivateQueue(ctx context.Context, arg CreatePrivateQueueParams) (AccountsQueue, error) {
	row := q.db.QueryRowContext(ctx, createPrivateQueue, arg.Fromaccountid, arg.Toaccountid)
	var i AccountsQueue
	err := row.Scan(
		&i.FromAccountID,
		&i.Queue,
		&i.ToAccountID,
		&i.QueueAt,
	)
	return i, err
}

const deleteAccountQueue = `-- name: DeleteAccountQueue :exec
Delete from accounts_queue
WHERE from_account_id = $1 and to_account_id = $2
`

type DeleteAccountQueueParams struct {
	Fromaccountid int64 `json:"fromaccountid"`
	Toaccountid   int64 `json:"toaccountid"`
}

func (q *Queries) DeleteAccountQueue(ctx context.Context, arg DeleteAccountQueueParams) error {
	_, err := q.db.ExecContext(ctx, deleteAccountQueue, arg.Fromaccountid, arg.Toaccountid)
	return err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT accounts_id, owner, is_private, created_at, follower, following FROM accounts
WHERE accounts_id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, accountsID int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, accountsID)
	var i Account
	err := row.Scan(
		&i.AccountsID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :one
SELECT accounts_id, owner, is_private, created_at, follower, following FROM accounts
WHERE accounts_id = $1 LIMIT 1
`

func (q *Queries) GetAccounts(ctx context.Context, accountsID int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccounts, accountsID)
	var i Account
	err := row.Scan(
		&i.AccountsID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
	)
	return i, err
}

const getAccountsInfo = `-- name: GetAccountsInfo :one
SELECT is_private,accounts_id FROM accounts
WHERE accounts_id = $1 LIMIT 1
`

type GetAccountsInfoRow struct {
	IsPrivate  bool  `json:"is_private"`
	AccountsID int64 `json:"accounts_id"`
}

func (q *Queries) GetAccountsInfo(ctx context.Context, accountsID int64) (GetAccountsInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountsInfo, accountsID)
	var i GetAccountsInfoRow
	err := row.Scan(&i.IsPrivate, &i.AccountsID)
	return i, err
}

const getAccountsOwner = `-- name: GetAccountsOwner :one
SELECT accounts_id, owner, is_private, created_at, follower, following FROM accounts
WHERE owner = $1 LIMIT 1
`

func (q *Queries) GetAccountsOwner(ctx context.Context, owner string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountsOwner, owner)
	var i Account
	err := row.Scan(
		&i.AccountsID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT accounts_id, owner, is_private, created_at, follower, following FROM accounts
WHERE owner = $1
ORDER BY accounts_id
LIMIT $2
OFFSET $3
`

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.AccountsID,
			&i.Owner,
			&i.IsPrivate,
			&i.CreatedAt,
			&i.Follower,
			&i.Following,
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

const updateAccountFollower = `-- name: UpdateAccountFollower :one
UPDATE accounts
SET follower = follower + $1
WHERE accounts_id = $2
RETURNING accounts_id, owner, is_private, created_at, follower, following
`

type UpdateAccountFollowerParams struct {
	Num        int64 `json:"num"`
	AccountsID int64 `json:"accounts_id"`
}

func (q *Queries) UpdateAccountFollower(ctx context.Context, arg UpdateAccountFollowerParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountFollower, arg.Num, arg.AccountsID)
	var i Account
	err := row.Scan(
		&i.AccountsID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
	)
	return i, err
}

const updateAccountFollowing = `-- name: UpdateAccountFollowing :one
UPDATE accounts
SET following = following + $1
WHERE accounts_id = $2
RETURNING accounts_id, owner, is_private, created_at, follower, following
`

type UpdateAccountFollowingParams struct {
	Num        int64 `json:"num"`
	AccountsID int64 `json:"accounts_id"`
}

func (q *Queries) UpdateAccountFollowing(ctx context.Context, arg UpdateAccountFollowingParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountFollowing, arg.Num, arg.AccountsID)
	var i Account
	err := row.Scan(
		&i.AccountsID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
	)
	return i, err
}
