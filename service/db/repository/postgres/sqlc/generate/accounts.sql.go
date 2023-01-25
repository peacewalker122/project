// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: accounts.sql

package db

import (
	"context"
	"database/sql"
)

const createAccounts = `-- name: CreateAccounts :one
INSERT INTO accounts(
    owner,
    is_private,
    photo_dir
) VALUES(
    $1,$2,$3
) RETURNING id, owner, is_private, created_at, follower, following, photo_dir
`

type CreateAccountsParams struct {
	Owner     string         `json:"owner"`
	IsPrivate bool           `json:"is_private"`
	PhotoDir  sql.NullString `json:"photo_dir"`
}

func (q *Queries) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccounts, arg.Owner, arg.IsPrivate, arg.PhotoDir)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
		&i.PhotoDir,
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

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT a.id,owner,is_private,follower,following,photo_dir from accounts a
left join users u on a.owner = u.username
where u.email = $1 LIMIT 1
`

type GetAccountByEmailRow struct {
	ID        int64          `json:"id"`
	Owner     string         `json:"owner"`
	IsPrivate bool           `json:"is_private"`
	Follower  int64          `json:"follower"`
	Following int64          `json:"following"`
	PhotoDir  sql.NullString `json:"photo_dir"`
}

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountByEmail, email)
	var i GetAccountByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.IsPrivate,
		&i.Follower,
		&i.Following,
		&i.PhotoDir,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, owner, is_private, created_at, follower, following, photo_dir FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
		&i.PhotoDir,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :one
SELECT id, owner, is_private, created_at, follower, following, photo_dir FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccounts(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccounts, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
		&i.PhotoDir,
	)
	return i, err
}

const getAccountsInfo = `-- name: GetAccountsInfo :one
SELECT is_private,id FROM accounts
WHERE id = $1 LIMIT 1
`

type GetAccountsInfoRow struct {
	IsPrivate bool  `json:"is_private"`
	ID        int64 `json:"id"`
}

func (q *Queries) GetAccountsInfo(ctx context.Context, id int64) (GetAccountsInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountsInfo, id)
	var i GetAccountsInfoRow
	err := row.Scan(&i.IsPrivate, &i.ID)
	return i, err
}

const getAccountsOwner = `-- name: GetAccountsOwner :one
SELECT id, owner, is_private, created_at, follower, following, photo_dir FROM accounts
WHERE owner = $1 LIMIT 1
`

func (q *Queries) GetAccountsOwner(ctx context.Context, owner string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountsOwner, owner)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
		&i.PhotoDir,
	)
	return i, err
}

const getQueueRows = `-- name: GetQueueRows :one
SELECT COUNT(*) from accounts_queue
WHERE from_account_id = $1 and to_account_id = $2
`

type GetQueueRowsParams struct {
	Fromaccountid int64 `json:"fromaccountid"`
	Toaccountid   int64 `json:"toaccountid"`
}

func (q *Queries) GetQueueRows(ctx context.Context, arg GetQueueRowsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getQueueRows, arg.Fromaccountid, arg.Toaccountid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, is_private, created_at, follower, following, photo_dir FROM accounts
WHERE owner = $1
ORDER BY id
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
			&i.ID,
			&i.Owner,
			&i.IsPrivate,
			&i.CreatedAt,
			&i.Follower,
			&i.Following,
			&i.PhotoDir,
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

const listQueue = `-- name: ListQueue :many
select a."owner" ,aq.from_account_id  from accounts a
left join accounts_queue aq ON a.id = aq.from_account_id 
where aq.to_account_id  = $3
order by a.id
limit $1
offset $2
`

type ListQueueParams struct {
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
	Accountid int64 `json:"accountid"`
}

type ListQueueRow struct {
	Owner         string        `json:"owner"`
	FromAccountID sql.NullInt64 `json:"from_account_id"`
}

func (q *Queries) ListQueue(ctx context.Context, arg ListQueueParams) ([]ListQueueRow, error) {
	rows, err := q.db.QueryContext(ctx, listQueue, arg.Limit, arg.Offset, arg.Accountid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListQueueRow{}
	for rows.Next() {
		var i ListQueueRow
		if err := rows.Scan(&i.Owner, &i.FromAccountID); err != nil {
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

const privateAccount = `-- name: PrivateAccount :exec
UPDATE accounts
SET is_private = $1
WHERE owner = $2
RETURNING is_private
`

type PrivateAccountParams struct {
	IsPrivate bool   `json:"is_private"`
	Username  string `json:"username"`
}

func (q *Queries) PrivateAccount(ctx context.Context, arg PrivateAccountParams) error {
	_, err := q.db.ExecContext(ctx, privateAccount, arg.IsPrivate, arg.Username)
	return err
}

const updateAccountFollower = `-- name: UpdateAccountFollower :one
UPDATE accounts
SET follower = follower + $1
WHERE id = $2
RETURNING id, owner, is_private, created_at, follower, following, photo_dir
`

type UpdateAccountFollowerParams struct {
	Num int64 `json:"num"`
	ID  int64 `json:"id"`
}

func (q *Queries) UpdateAccountFollower(ctx context.Context, arg UpdateAccountFollowerParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountFollower, arg.Num, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
		&i.PhotoDir,
	)
	return i, err
}

const updateAccountFollowing = `-- name: UpdateAccountFollowing :one
UPDATE accounts
SET following = following + $1
WHERE id = $2
RETURNING id, owner, is_private, created_at, follower, following, photo_dir
`

type UpdateAccountFollowingParams struct {
	Num int64 `json:"num"`
	ID  int64 `json:"id"`
}

func (q *Queries) UpdateAccountFollowing(ctx context.Context, arg UpdateAccountFollowingParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountFollowing, arg.Num, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.Follower,
		&i.Following,
		&i.PhotoDir,
	)
	return i, err
}

const updateAccountQueue = `-- name: UpdateAccountQueue :exec
UPDATE accounts_queue
SET queue = $1
WHERE from_account_id = $2 and to_account_id = $3
`

type UpdateAccountQueueParams struct {
	Queue         bool  `json:"queue"`
	Fromaccountid int64 `json:"fromaccountid"`
	Toaccountid   int64 `json:"toaccountid"`
}

func (q *Queries) UpdateAccountQueue(ctx context.Context, arg UpdateAccountQueueParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountQueue, arg.Queue, arg.Fromaccountid, arg.Toaccountid)
	return err
}

const updatePhoto = `-- name: UpdatePhoto :exec
UPDATE accounts
SET photo_dir = $1
WHERE owner = $2 or id = $3
`

type UpdatePhotoParams struct {
	Filedirectory sql.NullString `json:"filedirectory"`
	Username      string         `json:"username"`
	Accountid     int64          `json:"accountid"`
}

func (q *Queries) UpdatePhoto(ctx context.Context, arg UpdatePhotoParams) error {
	_, err := q.db.ExecContext(ctx, updatePhoto, arg.Filedirectory, arg.Username, arg.Accountid)
	return err
}