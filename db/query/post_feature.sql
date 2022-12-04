-- name: CreateLike_feature :exec
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

-- name: CreateRetweet_feature :exec
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
) RETURNING qoute;

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
SELECT sum_comment,sum_like,sum_retweet,sum_qoute_retweet FROM post_feature as pf
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

-- name: UpdateLike :exec
UPDATE like_feature
set is_like = $1
WHERE post_id = $2 and from_account_id = $3
RETURNING is_like;

-- name: ListComment :many
SELECT comment_id,from_account_id,comment,sum_like,created_at from comment_feature
WHERE post_id = $1
ORDER by from_account_id
LIMIT $2
OFFSET $3;

-- name: GetRetweet :one
SELECT * from retweet_feature
WHERE from_account_id = $1 and post_id = $2 LIMIT 1;

-- name: UpdateRetweet :exec
UPDATE retweet_feature
set retweet = $1
WHERE post_id = $2 and from_account_id = $3
RETURNING retweet;

-- name: GetRetweetJoin :one
SELECT rf.retweet from retweet_feature rf
INNER JOIN post ON post.post_id = rf.post_id
WHERE post.post_id = @PostID and rf.from_account_id  = @FromAccountID;

-- name: GetQouteRetweet :one
SELECT * from qoute_retweet_feature
WHERE from_account_id=$1 and post_id = $2 LIMIT 1;

-- name: GetQouteRetweetRows :execrows
SELECT * from qoute_retweet_feature
WHERE from_account_id=$1 and post_id = $2 LIMIT 1;

-- name: GetQouteRetweetJoin :one
SELECT qoute_retweet_feature.qoute_retweet from qoute_retweet_feature
INNER JOIN post on post.post_id = qoute_retweet_feature.post_id
WHERE qoute_retweet_feature.post_id = $1;

-- name: UpdateQouteRetweet :exec
UPDATE qoute_retweet_feature
set qoute_retweet = $1
WHERE post_id = $2 and from_account_id = $3;

-- name: DeleteQouteRetweet :exec
delete from qoute_retweet_feature
WHERE post_id = $1 and from_account_id = $2;

-- name: DeletePostFeature :exec
delete from post_feature p
WHERE p.post_id = $1;

-- name: DeleteRetweet :exec
delete from retweet_feature
WHERE post_id=$1 and from_account_id=$2;