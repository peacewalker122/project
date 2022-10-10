package api

import (
	"time"

	db "github.com/peacewalker122/project/db/sqlc"
)

type CreateUserResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func UserResponse(input db.User) CreateUserResponse {
	return CreateUserResponse{
		Username:  input.Username,
		FullName:  input.FullName,
		Email:     input.Email,
		CreatedAt: input.CreatedAt,
	}
}

type CreateAccountsResponse struct {
	Owner       string    `json:"owner"`
	AccountType bool      `json:"account_type"`
	CreatedAt   time.Time `json:"created_at"`
}

func AccountResponse(input db.Account) CreateAccountsResponse {
	return CreateAccountsResponse{
		Owner:       input.Owner,
		AccountType: input.AccountType,
		CreatedAt:   input.CreatedAt,
	}
}
