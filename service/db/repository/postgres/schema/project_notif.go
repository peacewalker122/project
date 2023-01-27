package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Notif holds the schema definition for the Notif entity.
type AccountNotifs struct {
	ent.Schema
}

// Fields of the Notif.
func (AccountNotifs) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),

		field.Int64("account_id").
			Optional().
			Nillable().
			Comment("for deploy purpose"),

		field.String("notif_type").
			MaxLen(255),

		field.String("notif_title").
			MaxLen(50).
			Optional(),

		field.String("notif_content").
			Optional().
			MaxLen(255),

		field.Time("notif_time").
			Optional().
			Nillable().
			Default(time.Now).
			Comment("for deploy purpose"),

		field.String("username").
			MaxLen(255).
			Optional().
			Nillable().
			Comment("for deploy purpose"),

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the AccountNotif.
func (AccountNotifs) Edges() []ent.Edge {
	return nil
}

// Indexes of the AccountNotif.
func (AccountNotifs) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
