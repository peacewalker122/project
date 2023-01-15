package auth

import (
	"context"

	"github.com/peacewalker122/project/usecase/auth"
)

type AuthContract interface {
	CreateRequest(ctx context.Context, params auth.AuthParams) error
	ChangePasswordAuth(ctx context.Context, req auth.ChangePassParams) error
}
