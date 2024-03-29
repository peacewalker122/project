package token

import (
	"errors"
	"golang.org/x/crypto/sha3"
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
	AccountID int64     `json:"account_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type PayloadRequest struct {
	Username  string        `json:"username"`
	AccountID int64         `json:"account_id"`
	Duration  time.Duration `json:"duration"`
}

func Newpayload(param *PayloadRequest) (*Payload, error) {
	uid, _ := uuid.NewRandom()
	token := uuid.NewHash(sha3.New256(), uid, nil, 4)

	payload := &Payload{
		ID:        token,
		Username:  param.Username,
		AccountID: param.AccountID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(param.Duration),
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpired
	}
	return nil
}
