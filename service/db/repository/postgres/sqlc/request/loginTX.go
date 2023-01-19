package request

import (
	"time"
)

type LoginTXRequest struct {
	Username              string        `json:"username"`
	RefreshTokenDuration  time.Duration `json:"refresh_token_duration"`
	AccessTokenDuration   time.Duration `json:"access_token_duration"`
	UserAgent             string        `json:"user_agent"`
	ClientIp              string        `json:"client_ip"`
	IsBlocked             bool          `json:"is_blocked"`
	AccessTokenExpiresAt  time.Time     `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time     `json:"refresh_token_expires_at"`
}
