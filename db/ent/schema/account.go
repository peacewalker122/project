package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner").
			NotEmpty().
			Unique(),

		field.Bool("is_private").
			Default(true),

		field.Time("created_at").
			Default(time.Now()),

		field.Int64("follower").
			Default(0),

		field.Int64("following").
			Default(0),

		field.String("photo_dir").
			Optional().
			MaxLen(70).
			Nillable(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
