package result

import (
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

type PostTXResult struct {
	Post        db.Post        `json:"post"`
	PostFeature db.PostFeature `json:"post_feature"`
	FileURL     string         `json:"file_url"`
	Err         error
}

type LikeTXResult struct {
	PostFeature db.PostFeature `json:"post_feature"`
	ErrCode     int            `json:"err_code"`
}

type RetweetTXResult struct {
	Err         error
	ErrCode     int
	Post        db.Post
	PostFeature db.PostFeature
	Retweet     db.CreateRetweetResult
}
