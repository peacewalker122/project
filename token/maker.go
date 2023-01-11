package token

type Maker interface {
	CreateToken(param *PayloadRequest) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
