package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		field.String("password_changed_at").
			NotEmpty().
			Default("0001-01-01 00:00:00Z"),
		field.String("created_at").
			NotEmpty().
			Default("now()"),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
