-- name: CreateEntries :one
INSERT INTO entries (
    from_account_id,
    post_id,
    type_entries
 ) VALUES (
    $1,$2,$3
  ) RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE entries_id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE post_id = $1
ORDER BY entries_id
LIMIT $2
OFFSET $3;

-- name: GetEntriesFull :exec
SELECT * FROM entries
WHERE post_id = $1 and from_account_id = $2 and type_entries = $3 LIMIT 1;