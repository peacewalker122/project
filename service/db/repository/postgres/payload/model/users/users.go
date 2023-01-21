package users

import (
	"context"
	ent2 "github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/users"
	"time"
)

type UsersQuery interface {
	SetUser(ctx context.Context, Params *UsersParam) (*ent2.Users, error)

	UpdateUser(ctx context.Context, Params *UsersParam) error
	SetPassword(ctx context.Context, username string, password string) error
	GetAllWithEmail(ctx context.Context, email string) (*ent2.Users, error)
	IsUserExist(ctx context.Context, email string) (bool, error)
}

type UserQueries struct {
	client *ent2.Client
}

func (s *UserQueries) SetUser(ctx context.Context, Params *UsersParam) (*ent2.Users, error) {
	return s.client.Users.
		Create().
		SetEmail(Params.Email).
		SetFullName(Params.FullName).
		SetUsername(Params.Username).
		Save(ctx)
}

//func (s *UserQueries) GetUser(ctx context.Context, email string) (*ent2.Users, error) {
//	return s.client.Users.
//		Query().
//		Where(users.Email(email)).
//		Only(ctx)
//}

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

func (s *UserQueries) SetPassword(ctx context.Context, username string, password string) error {
	_, err := s.client.Users.
		Update().
		Where(users.Username(username)).
		SetPasswordChangedAt(time.Now()).
		SetHashedPassword(password).
		Save(ctx)
	return err
}

func (s *UserQueries) GetAllWithEmail(ctx context.Context, email string) (*ent2.Users, error) {
	res, err := s.client.Users.
		Query().
		Where(users.Email(email)).
		Only(ctx)
	return res, err
}

func (s *UserQueries) IsUserExist(ctx context.Context, email string) (bool, error) {
	return s.client.Users.
		Query().
		Where(users.Email(email)).
		Exist(ctx)
}

func NewUserQuery(Client *ent2.Client) *UserQueries {
	return &UserQueries{
		client: Client,
	}
}
