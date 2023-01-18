package account

type GetAccountsParams struct {
	Limit       int32 `json:"limit" form:"limit" query:"limit" validate:"required,min=1,max=50"`
	Page        int32 `json:"page" form:"page" query:"page" validate:"required,min=0"`
	ToAccountID int64 `uri:"id" query:"to_account_id" validate:"required,min=1"`
}

type GetAccountParams struct {
	ToAccountID int64 `json:"to_account_id" form:"to_account_id" query:"to_account_id" validate:"required,min=1"`
}
