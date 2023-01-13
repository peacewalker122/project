package account

import (
	"context"

	"github.com/peacewalker122/project/db/repository/postgres/ent"
	"github.com/peacewalker122/project/db/repository/postgres/ent/account"
)

type AccountQuery interface {
	SetAccount(ctx context.Context, Params *AccountParam) (*ent.Account, error)
	GetAccount(ctx context.Context, owner string) (*ent.Account, error)
	UpdateAccount(ctx context.Context, Params *AccountParam) error
}

type AccountQueries struct {
	client *ent.Client
}

func (s *AccountQueries) SetAccount(ctx context.Context, Params *AccountParam) (*ent.Account, error) {
	return s.client.Account.
		Create().
		SetOwner(Params.Owner).
		SetIsPrivate(Params.IsPrivate).
		SetPhotoDir(Params.PhotoDir.String).
		Save(ctx)
}

func (s *AccountQueries) GetAccount(ctx context.Context, owner string) (*ent.Account, error) {
	return s.client.Account.
		Query().
		Where(account.Owner(owner)).
		Only(ctx)
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

func NewAccountQuery(client *ent.Client) *AccountQueries {
	return &AccountQueries{
		client: client,
	}
}
