-- name: CreatePost :one
INSERT INTO post(
    account_id,
    picture_description
) VALUES(
    $1,$2
) RETURNING *;

-- name: GetPost :one
SELECT * FROM post
WHERE post_id = $1 LIMIT 1;

-- name: ListPost :many
SELECT * FROM post
ORDER BY post_id
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE post
SET picture_description = $2
WHERE post_id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM post
WHERE post_id = $1;