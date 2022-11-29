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

-- name: UpdateAccountFollowing :one
UPDATE accounts
SET following = following + @num
WHERE accounts_id = @accounts_id
RETURNING *;

-- name: UpdateAccountFollower :one
UPDATE accounts
SET follower = follower + @num
WHERE accounts_id = @accounts_id
RETURNING *;

-- name: GetAccountsInfo :one
SELECT is_private,accounts_id FROM accounts
WHERE accounts_id = @accounts_id LIMIT 1;

-- name: CreatePrivateQueue :one
INSERT INTO accounts_queue(
    from_account_id,
    to_account_id,
    queue
) VALUES(
    @FromAccountID, @ToAccountID, true
) RETURNING *;

-- name: UpdateAccountQueue :exec
UPDATE accounts_queue
set queue = @Queue
WHERE  from_account_id = @FromAccountID and to_account_id = @ToAccountID;