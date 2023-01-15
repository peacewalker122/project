package user

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/util"
)

func (s *UserUsecase) CreateNewUserRequest(ctx context.Context, req db.CreateUserParams) (uuid.UUID, *util.Error) {

	multierror := &util.Error{}

	_, err := s.postgre.GetEmail(ctx, db.GetEmailParams{Email: req.Email})
	if err == nil {
		multierror.Important(errors.New("email already exist").Error(), "email")
	}
	_, err = s.postgre.GetEmail(ctx, db.GetEmailParams{Username: req.Username})
	if err == nil {
		multierror.Important(errors.New("username already exist").Error(), "username")
	}

	if multierror.HasError() {
		return uuid.Nil, multierror
	}

	var wg sync.WaitGroup
	var uid uuid.UUID

	uuidchan := make(chan uuid.UUID, 1)
	errchan := make(chan error, 2)

	wg.Add(1)
	go func(errchan chan error, uuidchan chan uuid.UUID) {
		defer wg.Done()
		uid, err := s.email.CreateEmailAuth(ctx, req.Email)
		if err != nil {
			errchan <- errors.New("failed to create email auth: " + err.Error())
		}
		// here we set the key to redis
		// to get the key we use the uuid
		err = s.redis.Set(ctx, uid.String()+"h", req, 3*time.Minute)
		if err != nil {
			errchan <- err
		}
		uuidchan <- uid
	}(errchan, uuidchan)
	select {
	case err := <-errchan:
		if err != nil {
			multierror.Error(err.Error())
			return uuid.Nil, multierror
		}
	case uid = <-uuidchan:
	}
	// here we send the email
	wg.Wait()

	return uid, nil
}
