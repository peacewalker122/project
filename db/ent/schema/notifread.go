package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NotifRead holds the schema definition for the NotifRead entity.
type NotifRead struct {
	ent.Schema
}

// Fields of the NotifRead.
func (NotifRead) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("notif_id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.Int64("account_id"),
		field.Time("read_at").
			Nillable(),
	}
}

// Edges of the NotifRead.
func (NotifRead) Edges() []ent.Edge {
	return nil
}
