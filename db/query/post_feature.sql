-- name: CreateLike_feature :one
INSERT INTO like_feature(
    from_account_id,
    is_like,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING *;

-- name: CreateComment_feature :one
INSERT INTO comment_feature(
    from_account_id,
    comment,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING *;

-- name: CreateRetweet_feature :one
INSERT INTO retweet_feature(
    from_account_id,
    retweet,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING *;

-- name: CreateQouteRetweet_feature :one
INSERT INTO qoute_retweet_feature(
    from_account_id,
    qoute_retweet,
    qoute,
    post_id
) VALUES(
    $1,$2,$3,$4
) RETURNING *;

-- name: CreatePost_feature :one
INSERT INTO post_feature(
    post_id
) values(
    $1
) RETURNING *;

-- name: GetPost_feature :one
SELECT * FROM post_feature
WHERE post_id = $1 LIMIT 1;

-- name: GetPost_feature_Update :one
SELECT * FROM post_feature
WHERE post_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdatePost_feature :one
UPDATE post_feature
SET sum_comment = $2, sum_like = $3, sum_retweet = $4, sum_qoute_retweet =$5
WHERE post_id = $1
RETURNING *; 