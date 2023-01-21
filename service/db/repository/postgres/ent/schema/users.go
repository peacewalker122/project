package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),

		field.String("username").
			Unique().
			NotEmpty(),
		field.String("hashed_password").
			NotEmpty().
			Optional(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("full_name").
			NotEmpty(),
		field.Time("password_changed_at").
			Optional(),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
