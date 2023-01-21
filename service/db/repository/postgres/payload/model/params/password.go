package params

import "context"

type ChangePasswordParam struct {
	UUID     string
	Password string
	Username string
	Email    string
	ClientIp string
	RedisDel func(ctx context.Context, key string) error
}
