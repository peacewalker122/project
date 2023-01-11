-- name: CreateAccounts :one
INSERT INTO accounts(
    owner,
    is_private,
    photo_dir
) VALUES(
    $1,$2,$3
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

-- name: DeleteAccountQueue :exec
Delete from accounts_queue
WHERE from_account_id = @FromAccountID and to_account_id = @ToAccountID;

-- name: PrivateAccount :exec
UPDATE accounts
SET is_private = $1
WHERE owner = @username
RETURNING is_private;

-- name: GetQueueRows :one
SELECT COUNT(*) from accounts_queue
WHERE from_account_id = @FromAccountID and to_account_id = @ToAccountID;

-- name: UpdateAccountQueue :exec
UPDATE accounts_queue
SET queue = $1
WHERE from_account_id = @FromAccountID and to_account_id = @ToAccountID;

-- name: ListQueue :many
select a."owner" ,aq.from_account_id  from accounts a
left join accounts_queue aq ON a.accounts_id = aq.from_account_id 
where aq.to_account_id  = @AccountID
order by a.accounts_id
limit $1
offset $2;

-- name: UpdatePhoto :exec
UPDATE accounts
SET photo_dir = @FileDirectory
WHERE owner = @username or accounts_id = @AccountID;

-- name: GetAccountByEmail :one
SELECT accounts_id,owner,is_private,follower,following,photo_dir from accounts a
left join users u on a.owner = u.username
where u.email = @Email LIMIT 1;