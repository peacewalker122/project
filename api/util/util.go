package util

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/repository/postgres"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/db/repository/redis"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"golang.org/x/oauth2"
)

const (
	AuthRefresh    = "RefreshToken"
	AuthHeaderkey  = "authorization"
	AuthTypeBearer = "bearer"
	AuthPayload    = "authorization_payload"
)

type utilTools struct {
	store db.PostgresStore
	redis redis.Store
	cfg   util.Config
}

type UtilTools interface {
	accountFeature
	GetRedisPayload(ctx context.Context, uid string, payload interface{}) error
	TokenHelper(ctx context.Context, token oauth2.TokenSource) (*oauth2.Token, error)
	AuthPayload(c echo.Context) (*token.Payload, error)
}

func NewApiUtil(store db.PostgresStore, redis redis.Store, cfg util.Config) UtilTools {
	return &utilTools{
		store: store,
		redis: redis,
		cfg:   cfg,
	}
}

func (s *utilTools) TokenHelper(ctx context.Context, token oauth2.TokenSource) (*oauth2.Token, error) {
	t, err := token.Token()
	if err != nil {
		return nil, err
	}
	if t.Valid() {
		return t, nil
	}
	t, err = s.RefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *utilTools) RefreshToken(ctx context.Context, token oauth2.TokenSource) (*oauth2.Token, error) {
	var res oauth2.Token

	newToken, err := oauth2.ReuseTokenSource(&res, token).Token()
	if err != nil {
		return nil, err
	}

	err = s.store.UpdateToken(ctx, &tokens.TokensParams{
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
		ExpiresIn:    newToken.Expiry,
		TokenType:    newToken.TokenType,
	})
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func (s *utilTools) AuthPayload(c echo.Context) (*token.Payload, error) {
	authParam, ok := c.Get(AuthPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return nil, err
	}

	return authParam, nil
}

func (s *utilTools) GetRedisPayload(ctx context.Context, uid string, payload interface{}) error {
	tempVal, err := s.redis.Get(ctx, uid)
	if err != nil {
		return err
	}

	err = s.redis.Del(ctx, uid)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(tempVal), &payload)
	if err != nil {
		return err
	}

	return nil
}
