package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// LikeFeature holds the schema definition for the LikeFeature entity.
type LikeFeature struct {
	ent.Schema
}

// Fields of the LikeFeature.
func (LikeFeature) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("from_account_id"),
		field.Bool("is_like").
			Default(false),
		field.UUID("post_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the LikeFeature.
func (LikeFeature) Edges() []ent.Edge {
	return nil
}
