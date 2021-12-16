package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("signature").NotEmpty(),
		field.String("address").NotEmpty(),
		field.Uint64("block_number"),
		field.String("tx_hash").NotEmpty(),
		field.Uint("tx_index"),
		field.String("block_hash").NotEmpty(),
		field.Uint("index"),
		field.String("hash").Unique().NotEmpty(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
