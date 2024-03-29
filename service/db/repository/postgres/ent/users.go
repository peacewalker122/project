// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/users"
)

// Users is the model entity for the Users schema.
type Users struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// HashedPassword holds the value of the "hashed_password" field.
	HashedPassword string `json:"hashed_password,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// PasswordChangedAt holds the value of the "password_changed_at" field.
	PasswordChangedAt time.Time `json:"password_changed_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Users) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case users.FieldUsername, users.FieldHashedPassword, users.FieldEmail, users.FieldFullName:
			values[i] = new(sql.NullString)
		case users.FieldPasswordChangedAt, users.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case users.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Users", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Users fields.
func (u *Users) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case users.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case users.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case users.FieldHashedPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hashed_password", values[i])
			} else if value.Valid {
				u.HashedPassword = value.String
			}
		case users.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case users.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				u.FullName = value.String
			}
		case users.FieldPasswordChangedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field password_changed_at", values[i])
			} else if value.Valid {
				u.PasswordChangedAt = value.Time
			}
		case users.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Users.
// Note that you need to call Users.Unwrap() before calling this method if this Users
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Users) Update() *UsersUpdateOne {
	return (&UsersClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the Users entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Users) Unwrap() *Users {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Users is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Users) String() string {
	var builder strings.Builder
	builder.WriteString("Users(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("hashed_password=")
	builder.WriteString(u.HashedPassword)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(u.FullName)
	builder.WriteString(", ")
	builder.WriteString("password_changed_at=")
	builder.WriteString(u.PasswordChangedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UsersSlice is a parsable slice of Users.
type UsersSlice []*Users

func (u UsersSlice) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
