package auth

import (
	"context"

	"github.com/peacewalker122/project/usecase/auth"
	"github.com/peacewalker122/project/util"
)

type AuthContract interface {
	CreateRequest(ctx context.Context, params auth.AuthParams) *util.Error
	ChangePasswordAuth(ctx context.Context, req auth.ChangePassParams) *util.Error
}
