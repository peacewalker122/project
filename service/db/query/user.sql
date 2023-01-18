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