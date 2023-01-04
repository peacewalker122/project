package tokens

import (
	"context"

	"github.com/peacewalker122/project/db/ent"
	"github.com/peacewalker122/project/db/ent/tokens"
)

type TokensQuery interface {
	SetToken(ctx context.Context, Params *TokensParams) (*ent.Tokens, error)
	GetToken(ctx context.Context, email string) (*ent.Tokens, error)
	UpdateToken(ctx context.Context, Params *TokensParams) error
	IsTokenExist(ctx context.Context, email string) (bool, error)
}

type TokenQueries struct {
	client *ent.Client
}

func (s *TokenQueries) IsTokenExist(ctx context.Context, email string) (bool, error) {
	res, err := s.client.Tokens.
		Query().
		Where(tokens.Email(email)).
		Exist(ctx)
	if err != nil {
		return false, err
	}
	return res, nil
}

// GetToken implements TokensQuery
func (s *TokenQueries) GetToken(ctx context.Context, email string) (*ent.Tokens, error) {
	res, err := s.client.Tokens.
		Query().
		Where(tokens.Email(email)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SetToken implements TokensQuery
func (s *TokenQueries) SetToken(ctx context.Context, Params *TokensParams) (*ent.Tokens, error) {
	res, err := s.client.Tokens.
		Create().
		SetEmail(Params.Email).
		SetAccessToken(Params.AccessToken).
		SetRefreshToken(Params.RefreshToken).
		SetExpiry(Params.ExpiresIn).
		SetTokenType(Params.TokenType).
		SetRaw(Params.Raw).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TokenQueries) DeleteToken(ctx context.Context, email string) error {
	_, err := s.client.Tokens.
		Delete().
		Where(tokens.Email(email)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *TokenQueries) UpdateToken(ctx context.Context, Params *TokensParams) error {
	err := s.client.Tokens.
		Update().
		Where(tokens.AccessToken(Params.AccessToken)).
		SetAccessToken(Params.AccessToken).
		SetRefreshToken(Params.RefreshToken).
		SetExpiry(Params.ExpiresIn).
		SetTokenType(Params.TokenType).
		SetRaw(Params.Raw).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewTokenQuery(Client *ent.Client) *TokenQueries {
	return &TokenQueries{
		client: Client,
	}
}
