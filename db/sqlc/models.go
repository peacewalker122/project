// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type Account struct {
	AccountsID int64          `json:"accounts_id"`
	Owner      string         `json:"owner"`
	IsPrivate  bool           `json:"is_private"`
	CreatedAt  time.Time      `json:"created_at"`
	Follower   int64          `json:"follower"`
	Following  int64          `json:"following"`
	PhotoDir   sql.NullString `json:"photo_dir"`
}

type AccountsFollow struct {
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Follow        bool      `json:"follow"`
	FollowAt      time.Time `json:"follow_at"`
}

type AccountsQueue struct {
	FromAccountID int64     `json:"from_account_id"`
	Queue         bool      `json:"queue"`
	ToAccountID   int64     `json:"to_account_id"`
	QueueAt       time.Time `json:"queue_at"`
}

type CommentFeature struct {
	CommentID     int64     `json:"comment_id"`
	FromAccountID int64     `json:"from_account_id"`
	Comment       string    `json:"comment"`
	SumLike       int64     `json:"sum_like"`
	PostID        int64     `json:"post_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type Entry struct {
	EntriesID     int64     `json:"entries_id"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	PostID        int64     `json:"post_id"`
	TypeEntries   string    `json:"type_entries"`
	CreatedAt     time.Time `json:"created_at"`
}

type LikeFeature struct {
	FromAccountID int64     `json:"from_account_id"`
	IsLike        bool      `json:"is_like"`
	PostID        int64     `json:"post_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type Post struct {
	PostID             int64          `json:"post_id"`
	AccountID          int64          `json:"account_id"`
	PictureDescription string         `json:"picture_description"`
	PhotoDir           sql.NullString `json:"photo_dir"`
	IsRetweet          bool           `json:"is_retweet"`
	CreatedAt          time.Time      `json:"created_at"`
}

type PostFeature struct {
	PostID          int64     `json:"post_id"`
	SumComment      int64     `json:"sum_comment"`
	SumLike         int64     `json:"sum_like"`
	SumRetweet      int64     `json:"sum_retweet"`
	SumQouteRetweet int64     `json:"sum_qoute_retweet"`
	CreatedAt       time.Time `json:"created_at"`
}

type QouteRetweetFeature struct {
	FromAccountID int64     `json:"from_account_id"`
	QouteRetweet  bool      `json:"qoute_retweet"`
	Qoute         string    `json:"qoute"`
	PostID        int64     `json:"post_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type RetweetFeature struct {
	FromAccountID int64     `json:"from_account_id"`
	Retweet       bool      `json:"retweet"`
	PostID        int64     `json:"post_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Token struct {
	Email        string                `json:"email"`
	AccessToken  string                `json:"access_token"`
	RefreshToken sql.NullString        `json:"refresh_token"`
	TokenType    sql.NullString        `json:"token_type"`
	Expiry       sql.NullTime          `json:"expiry"`
	Raw          pqtype.NullRawMessage `json:"raw"`
	ID           uuid.UUID             `json:"id"`
}

type User struct {
	Username          string        `json:"username"`
	HashedPassword    string        `json:"hashed_password"`
	FullName          string        `json:"full_name"`
	Email             string        `json:"email"`
	PasswordChangedAt time.Time     `json:"password_changed_at"`
	CreatedAt         time.Time     `json:"created_at"`
	ID                uuid.NullUUID `json:"id"`
}
