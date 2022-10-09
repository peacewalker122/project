-- name: CreatePost :one
INSERT INTO post(
    account_id,
    post_word,
    post_picture
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
SET post_word = $2, 
    post_picture = $3
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM post
WHERE id = $1;