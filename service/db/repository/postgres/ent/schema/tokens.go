package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Tokens holds the schema definition for the Tokens entity.
type Tokens struct {
	ent.Schema
}

// Fields of the Tokens.
func (Tokens) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id",uuid.UUID{}).
			Default(uuid.New),

		field.String("email").
			Unique().
			MaxLen(255).
			NotEmpty(),

		field.String("access_token").
			NotEmpty(),

		field.String("refresh_token"),

		field.String("token_type").
			MaxLen(255),

		field.Time("expiry"),

		field.JSON("raw", map[string]interface{}{}),
	}
}

// Edges of the Tokens.
func (Tokens) Edges() []ent.Edge {
	return nil
}
