package token

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

const minSecretkeySize = 32

type Jwt struct {
	secretkey string
}

func NewJwt(secretkey string) (Maker, error) {
	if len(secretkey) < minSecretkeySize {
		return nil, fmt.Errorf("invalid Key Size must be %v length", minSecretkeySize)
	}
	return &Jwt{secretkey}, nil
}

func (j *Jwt) CreateToken(param *PayloadRequest) (string, *Payload, error) {
	payload, err := Newpayload(&PayloadRequest{
		Username:  param.Username,
		AccountID: param.AccountID,
		Duration:  param.Duration,
	})
	if err != nil {
		return "", payload, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(j.secretkey))
	return token, payload, err
}

func (j *Jwt) VerifyToken(token string) (*Payload, error) {
	Keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrToken
		}
		return []byte(j.secretkey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, Keyfunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpired) {
			return nil, ErrExpired
		}
		return nil, ErrToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrToken
	}
	return payload, nil
}
