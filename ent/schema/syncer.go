package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Syncer holds the schema definition for the Syncer entity.
type Syncer struct {
	ent.Schema
}

// Fields of the Syncer.
func (Syncer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Uint64("start"),
		field.Uint64("finish"),
		field.Uint64("current"),
	}
}

// Edges of the Syncer.
func (Syncer) Edges() []ent.Edge {
	return nil
}
