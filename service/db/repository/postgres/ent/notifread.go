// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/notifread"
)

// NotifRead is the model entity for the NotifRead schema.
type NotifRead struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NotifID holds the value of the "notif_id" field.
	NotifID uuid.UUID `json:"notif_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID int64 `json:"account_id,omitempty"`
	// ReadAt holds the value of the "read_at" field.
	ReadAt *time.Time `json:"read_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotifRead) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notifread.FieldID, notifread.FieldAccountID:
			values[i] = new(sql.NullInt64)
		case notifread.FieldReadAt:
			values[i] = new(sql.NullTime)
		case notifread.FieldNotifID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NotifRead", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotifRead fields.
func (nr *NotifRead) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notifread.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nr.ID = int(value.Int64)
		case notifread.FieldNotifID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field notif_id", values[i])
			} else if value != nil {
				nr.NotifID = *value
			}
		case notifread.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				nr.AccountID = value.Int64
			}
		case notifread.FieldReadAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field read_at", values[i])
			} else if value.Valid {
				nr.ReadAt = new(time.Time)
				*nr.ReadAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this NotifRead.
// Note that you need to call NotifRead.Unwrap() before calling this method if this NotifRead
// was returned from a transaction, and the transaction was committed or rolled back.
func (nr *NotifRead) Update() *NotifReadUpdateOne {
	return (&NotifReadClient{config: nr.config}).UpdateOne(nr)
}

// Unwrap unwraps the NotifRead entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nr *NotifRead) Unwrap() *NotifRead {
	_tx, ok := nr.config.driver.(*txDriver)
	if !ok {
		panic("ent: NotifRead is not a transactional entity")
	}
	nr.config.driver = _tx.drv
	return nr
}

// String implements the fmt.Stringer.
func (nr *NotifRead) String() string {
	var builder strings.Builder
	builder.WriteString("NotifRead(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nr.ID))
	builder.WriteString("notif_id=")
	builder.WriteString(fmt.Sprintf("%v", nr.NotifID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", nr.AccountID))
	builder.WriteString(", ")
	if v := nr.ReadAt; v != nil {
		builder.WriteString("read_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NotifReads is a parsable slice of NotifRead.
type NotifReads []*NotifRead

func (nr NotifReads) config(cfg config) {
	for _i := range nr {
		nr[_i].config = cfg
	}
}
