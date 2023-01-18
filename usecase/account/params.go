package account

type GetAccountParams struct {
	Offset      int32
	ToAccountID int64
	Limit       int32
	Username    string
}
