package auth

import "github.com/google/uuid"

type AuthParams struct {
	UUID     uuid.UUID
	Email    string
	FullName string
	ClientIp string
}

type ChangePassParams struct {
	UUID     string
	Password string
	Email    string
	ClientIp string
}