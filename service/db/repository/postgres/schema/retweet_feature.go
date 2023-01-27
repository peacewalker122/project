package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Retweet_feature holds the schema definition for the Retweet_feature entity.
type Retweet_feature struct {
	ent.Schema
}

// Fields of the Retweet_feature.
func (Retweet_feature) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("from_account_id"),
		field.Bool("retweet").
			Default(false),
		field.UUID("post_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the Retweet_feature.
func (Retweet_feature) Edges() []ent.Edge {
	return nil
}
