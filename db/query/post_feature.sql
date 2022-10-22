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
    from_account_id,
    post_id,
    sum_comment,
    sum_like,
    sum_retweet,
    sum_qoute_retweet
) values(
    $1,$2,$3,$4,$5,$6
) RETURNING *;

