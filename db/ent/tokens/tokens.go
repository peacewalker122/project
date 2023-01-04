// Code generated by ent, DO NOT EDIT.

package tokens

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tokens type in the database.
	Label = "tokens"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldRefreshToken holds the string denoting the refresh_token field in the database.
	FieldRefreshToken = "refresh_token"
	// FieldTokenType holds the string denoting the token_type field in the database.
	FieldTokenType = "token_type"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// Table holds the table name of the tokens in the database.
	Table = "tokens"
)

// Columns holds all SQL columns for tokens fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldAccessToken,
	FieldRefreshToken,
	FieldTokenType,
	FieldExpiry,
	FieldRaw,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// AccessTokenValidator is a validator for the "access_token" field. It is called by the builders before save.
	AccessTokenValidator func(string) error
	// TokenTypeValidator is a validator for the "token_type" field. It is called by the builders before save.
	TokenTypeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
