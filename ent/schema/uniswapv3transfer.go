package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3Transfer holds the schema definition for the UniswapV3Transfer entity.
type UniswapV3Transfer struct {
	ent.Schema
}

// Fields of the UniswapV3Transfer.
func (UniswapV3Transfer) Fields() []ent.Field {
	return []ent.Field{
		BigIntField("token_id"),
		field.String("from"),
		field.String("to"),
	}
}

// Edges of the UniswapV3Transfer.
func (UniswapV3Transfer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("transfer").
			Unique().
			Required(),
	}
}
