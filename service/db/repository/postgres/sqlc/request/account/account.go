package account

type FollowParams struct {
	FromAccountID int64 `json:"account_id"`
	ToAccountID   int64 `json:"follow_id"`
}