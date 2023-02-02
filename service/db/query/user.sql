-- name: CreateUser :one
INSERT INTO users(
    username,
    hashed_password,
    full_name,
    email
) VALUES(
    $1,$2,$3,$4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetEmail :one
SELECT username,email FROM users
WHERE email = @Email or username = @Username LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateUserData :exec
UPDATE users
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email),
    username = COALESCE(sqlc.narg(username), username)
WHERE username = sqlc.arg(username)
RETURNING *;