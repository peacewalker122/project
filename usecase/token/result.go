package token

import "time"

type AccesTokenResp struct {
	AccesToken          string    `json:"access_token"`
	AccesTokenExpiresAt time.Time `json:"access_token_expires_at"`
}
