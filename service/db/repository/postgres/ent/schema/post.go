package schema

import "entgo.io/ent"

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
