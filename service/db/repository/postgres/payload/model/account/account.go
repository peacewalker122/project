package account

import (
	"context"
	ent2 "github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/account"
)

type AccountQuery interface {
	SetAccount(ctx context.Context, Params *AccountParam) (*ent2.Account, error)
	GetAccount(ctx context.Context, owner string) (*ent2.Account, error)
	UpdateAccount(ctx context.Context, Params *AccountParam) error
	GetAccountID(ctx context.Context, owner string) (int64, error)
}

type AccountQueries struct {
	client *ent2.Client
}

func (s *AccountQueries) SetAccount(ctx context.Context, Params *AccountParam) (*ent2.Account, error) {
	return s.client.Account.
		Create().
		SetOwner(Params.Owner).
		SetIsPrivate(Params.IsPrivate).
		SetPhotoDir(Params.PhotoDir.String).
		Save(ctx)
}

func (s *AccountQueries) GetAccount(ctx context.Context, owner string) (*ent2.Account, error) {
	return s.client.Account.
		Query().
		Where(account.Owner(owner)).
		Only(ctx)
}

func (s *AccountQueries) GetAccountID(ctx context.Context, owner string) (int64, error) {
	return s.client.Account.
		Query().
		Where(account.Owner(owner)).
		OnlyID(ctx)
}

func (s *AccountQueries) UpdateAccount(ctx context.Context, Params *AccountParam) error {
	_, err := s.client.Account.
		Update().
		Where(account.Owner(Params.Owner)).
		SetIsPrivate(Params.IsPrivate).
		SetPhotoDir(Params.PhotoDir.String).
		Save(ctx)
	return err
}

func NewAccountQuery(client *ent2.Client) *AccountQueries {
	return &AccountQueries{
		client: client,
	}
}
