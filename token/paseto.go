package token

import (
	"fmt"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPaseto(secretkey string) (Maker, error) {
	if len(secretkey) < minSecretkeySize {
		return nil, fmt.Errorf("invalid Key Size must be %v length", minSecretkeySize)
	}

	pasetoMaker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(secretkey),
	}

	return pasetoMaker, nil
}

func (p *PasetoMaker) CreateToken(param *PayloadRequest) (string, *Payload, error) {
	payload, err := Newpayload(param)
	if err != nil {
		return "", nil, err
	}

	token, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)

	return token, payload, err
}

func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
