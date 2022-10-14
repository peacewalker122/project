-- name: CreatePost :one
INSERT INTO post(
    account_id,
    picture_description,
    picture_id
) VALUES(
    $1,$2,$3
) RETURNING *;

-- name: GetPost :one
SELECT * FROM post
WHERE id = $1 LIMIT 1;

-- name: ListPost :many
SELECT * FROM post
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE post
SET picture_description = $2, 
    picture_id = $3
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM post
WHERE id = $1;