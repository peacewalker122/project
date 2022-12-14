// Code generated by ent, DO NOT EDIT.

package accountnotifs

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the accountnotifs type in the database.
	Label = "account_notifs"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldNotifType holds the string denoting the notif_type field in the database.
	FieldNotifType = "notif_type"
	// FieldNotifTitle holds the string denoting the notif_title field in the database.
	FieldNotifTitle = "notif_title"
	// FieldNotifContent holds the string denoting the notif_content field in the database.
	FieldNotifContent = "notif_content"
	// FieldNotifTime holds the string denoting the notif_time field in the database.
	FieldNotifTime = "notif_time"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the accountnotifs in the database.
	Table = "account_notifs"
)

// Columns holds all SQL columns for accountnotifs fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldNotifType,
	FieldNotifTitle,
	FieldNotifContent,
	FieldNotifTime,
	FieldUsername,
	FieldCreatedAt,
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
	// NotifTypeValidator is a validator for the "notif_type" field. It is called by the builders before save.
	NotifTypeValidator func(string) error
	// NotifTitleValidator is a validator for the "notif_title" field. It is called by the builders before save.
	NotifTitleValidator func(string) error
	// NotifContentValidator is a validator for the "notif_content" field. It is called by the builders before save.
	NotifContentValidator func(string) error
	// DefaultNotifTime holds the default value on creation for the "notif_time" field.
	DefaultNotifTime func() time.Time
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
