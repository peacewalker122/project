package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Notif holds the schema definition for the Notif entity.
type Notif struct {
	ent.Schema
}

// Fields of the Notif.
func (Notif) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("notif_id", uuid.UUID{}).
			StorageKey("oid").
			Default(uuid.New),

		field.Int64("account_id"),

		field.String("notif_type").
			MaxLen(255),

		field.String("notif_title").
			MaxLen(50),

		field.String("notif_content").
			MaxLen(255),

		field.Time("notif_time").
			Optional().
			Nillable().
			Default(time.Now).
			Comment("for deploy purpose"),

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Notif.
func (Notif) Edges() []ent.Edge {
	return nil
}

// Indexes of the Notif.
func (Notif) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
