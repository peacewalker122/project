package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Qoute_retweet_feature holds the schema definition for the Qoute_retweet_feature entity.
type Qoute_retweet_feature struct {
	ent.Schema
}

// Fields of the Qoute_retweet_feature.
func (Qoute_retweet_feature) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("from_account_id"),
		field.Bool("qoute_retweet").
			Default(false),
		field.String("qoute").
			NotEmpty(),
		field.UUID("post_id", uuid.UUID{}),
		field.Time("created_at"),
	}
}

// Edges of the Qoute_retweet_feature.
func (Qoute_retweet_feature) Edges() []ent.Edge {
	return nil
}
