package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tokens holds the schema definition for the Tokens entity.
type Tokens struct {
	ent.Schema
}

// Fields of the Tokens.
func (Tokens) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			MaxLen(255).
			NotEmpty(),

		field.String("access_token").
			MaxLen(255).
			NotEmpty(),

		field.String("refresh_token").
			MaxLen(255),

		field.String("token_type").
			MaxLen(255),

		field.Time("expires_in"),

		field.JSON("raw", map[string]interface{}{}),
	}
}

// Edges of the Tokens.
func (Tokens) Edges() []ent.Edge {
	return nil
}
