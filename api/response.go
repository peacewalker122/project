package api

import (
	"time"

	"github.com/google/uuid"
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

type CreatePostResponse struct {
	ID                 int64     `json:"id"`
	PictureDescription string    `json:"picture_description"`
	PictureID          int64     `json:"picture_id"`
	CreatedAt          time.Time `json:"created_at"`
}

func PostResponse(input db.Post) CreatePostResponse {
	return CreatePostResponse{
		ID:                 input.ID,
		PictureDescription: input.PictureDescription.String,
		PictureID:          input.PictureID,
		CreatedAt:          input.CreatedAt,
	}
}

type loginResp struct {
	SessionID             uuid.UUID          `json:"session_id"`
	RefreshToken          string             `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time          `json:"refresh_token_expires_at"`
	User                  CreateUserResponse `json:"user"`
	AccesToken            string             `json:"acc_token"`
	AccesTokenExpiresAt   time.Time          `json:"acces_token_expire_sat"`
}
