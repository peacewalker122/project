package token

import (
	"context"
	"github.com/peacewalker122/project/usecase/token"
)

type TokenContract interface {
	RefreshToken(ctx context.Context, token string) (*token.AccesTokenResp, error)
}
