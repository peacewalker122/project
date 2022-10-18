-- name: CreateAccounts :one
INSERT INTO accounts(
    owner,
    account_type
) VALUES(
    $1,$2
) RETURNING *;

-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetAccountsOwner :one
SELECT * FROM accounts
WHERE owner = $1 LIMIT 1;