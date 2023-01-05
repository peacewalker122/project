package tokens

import "time"

type TokensParams struct {
	Email        string
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    time.Time
	Raw          map[string]interface{}
}
