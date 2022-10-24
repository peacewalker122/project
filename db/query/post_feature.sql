-- name: CreateLike_feature :one
INSERT INTO like_feature(
    from_account_id,
    is_like,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING is_like;

-- name: CreateComment_feature :one
INSERT INTO comment_feature(
    from_account_id,
    comment,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING comment;

-- name: CreateRetweet_feature :one
INSERT INTO retweet_feature(
    from_account_id,
    retweet,
    post_id
) VALUES(
    $1,$2,$3
) RETURNING retweet;

-- name: CreateQouteRetweet_feature :one
INSERT INTO qoute_retweet_feature(
    from_account_id,
    qoute_retweet,
    qoute,
    post_id
) VALUES(
    $1,$2,$3,$4
) RETURNING qoute_retweet;

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

-- name: GetPostJoin :one
SELECT post.post_id,post.account_id FROM post
INNER JOIN post_feature ON post_feature.post_id = post.post_id
WHERE post.post_id = $1;

-- name: GetLikejoin :one
SELECT like_feature.is_like from like_feature
INNER JOIN post ON post.post_id = like_feature.post_id
WHERE post.post_id = $1;

-- name: GetLikeInfo :one
SELECT * from like_feature
WHERE from_account_id = $1 and post_id = $2 LIMIT 1;

-- name: UpdateLike :one
UPDATE like_feature
set is_like = $1
WHERE post_id = $2 and from_account_id = $3
RETURNING is_like;

-- name: GetCommentInfo :one
SELECT * from comment_feature
WHERE from_account_id = $1 and post_id = $2 and comment = $3 LIMIT 1;