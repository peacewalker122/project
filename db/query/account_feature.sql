-- name: CreateAccountsFollow :one
INSERT INTO accounts_follow(
    from_account_id,
    to_account_id,
    follow
) VALUES(
    $1,$2,$3
) RETURNING *;

-- name: CreateAccountQueue :one
INSERT INTO accounts_queue(
    from_account_id,
    to_account_id
) VALUES(
    $1,$2
) RETURNING *;

-- name: GetAccountsFollowRows :execrows
SELECT follow FROM accounts_follow
WHERE from_account_id = @fromID and to_account_id = @toID LIMIT 1;

-- name: GetAccountsFollow :one
SELECT follow FROM accounts_follow
WHERE from_account_id = @fromID and to_account_id = @toID LIMIT 1;

-- name: DeleteAccountsFollow :exec
DElete from accounts_follow
WHERE from_account_id = @fromID and to_account_id = @toID;