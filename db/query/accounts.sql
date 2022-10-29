-- name: CreateAccounts :one
INSERT INTO accounts(
    owner,
    is_private
) VALUES(
    $1,$2
) RETURNING *;

-- name: GetAccounts :one
SELECT * FROM accounts
WHERE accounts_id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE owner = $1
ORDER BY accounts_id
LIMIT $2
OFFSET $3;

-- name: GetAccountsOwner :one
SELECT * FROM accounts
WHERE owner = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE accounts_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: AddAccountFollowing :one
UPDATE accounts
SET following = following + 1
WHERE accounts_id = @accounts_id
RETURNING *;

-- name: AddAccountFollower :one
UPDATE accounts
SET follower = follower + 1
WHERE accounts_id = @accounts_id
RETURNING *;