package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrToken   = errors.New("token invalid")
	ErrExpired = errors.New("token expired")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func Newpayload(username string, duration time.Duration) (*Payload, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        token,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpired
	}
	return nil
}
