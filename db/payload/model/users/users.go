package users

import (
	"context"

	"github.com/peacewalker122/project/db/ent"
	"github.com/peacewalker122/project/db/ent/users"
)

type UsersQuery interface {
	SetUser(ctx context.Context, Params *UsersParam) (*ent.Users, error)
	//GetUser(ctx context.Context, email string) (*ent.Users, error)
	UpdateUser(ctx context.Context, Params *UsersParam) error
	IsUserExist(ctx context.Context, email string) (bool, error)
}

type UserQueries struct {
	client *ent.Client
}

func (s *UserQueries) SetUser(ctx context.Context, Params *UsersParam) (*ent.Users, error) {
	return s.client.Users.
		Create().
		SetEmail(Params.Email).
		SetFullName(Params.FullName).
		SetUsername(Params.Username).
		Save(ctx)
}

func (s *UserQueries) GetUser(ctx context.Context, email string) (*ent.Users, error) {
	return s.client.Users.
		Query().
		Where(users.Email(email)).
		Only(ctx)
}

func (s *UserQueries) UpdateUser(ctx context.Context, Params *UsersParam) error {
	_, err := s.client.Users.
		Update().
		Where(
			users.Or(
				users.Email(Params.Email),
				users.Username(Params.Username),
			),
		).
		SetFullName(Params.FullName).
		SetUsername(Params.Username).
		Save(ctx)
	return err
}

func (s *UserQueries) IsUserExist(ctx context.Context, email string) (bool, error) {
	return s.client.Users.
		Query().
		Where(users.Email(email)).
		Exist(ctx)
}

func NewUserQuery(Client *ent.Client) *UserQueries {
	return &UserQueries{
		client: Client,
	}
}
