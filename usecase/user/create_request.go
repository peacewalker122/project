package user

import (
	"context"
	"errors"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	"sync"
	"time"

	"github.com/google/uuid"
)

func (s *UserUsecase) CreateNewUserRequest(ctx context.Context, req db.CreateUserParams) (uuid.UUID, error) {

	// multierror := &util.Error{}
	errs := map[string]string{}

	_, err := s.postgre.GetEmail(ctx, db.GetEmailParams{Email: req.Email})
	if err == nil {
		errs["email"] = "email already exist"
	}
	_, err = s.postgre.GetEmail(ctx, db.GetEmailParams{Username: req.Username})
	if err == nil {
		errs["username"] = "username already exist"
	}

	if len(errs) > 0 {
		err := errors.New("failed to create new user request: " + errs["email"] + errs["username"])
		return uuid.Nil, err
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
			return uuid.Nil, err
		}
	case uid = <-uuidchan:
	}
	// here we send the email
	wg.Wait()

	return uid, nil
}
