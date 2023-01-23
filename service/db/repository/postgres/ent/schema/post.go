package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("owner").
			NotEmpty(),
		field.Bool("is_private").
			Default(false),
		field.Time("created_at").
			Default(time.Now()),
		field.Int64("follower").
			Default(0),
		field.Int64("following").
			Default(0),
		field.String("photo_dir").
			Optional(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
