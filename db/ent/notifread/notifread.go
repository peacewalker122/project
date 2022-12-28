// Code generated by ent, DO NOT EDIT.

package notifread

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the notifread type in the database.
	Label = "notif_read"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNotifID holds the string denoting the notif_id field in the database.
	FieldNotifID = "notif_id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldReadAt holds the string denoting the read_at field in the database.
	FieldReadAt = "read_at"
	// Table holds the table name of the notifread in the database.
	Table = "notif_reads"
)

// Columns holds all SQL columns for notifread fields.
var Columns = []string{
	FieldID,
	FieldNotifID,
	FieldAccountID,
	FieldReadAt,
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
	// DefaultNotifID holds the default value on creation for the "notif_id" field.
	DefaultNotifID func() uuid.UUID
)
