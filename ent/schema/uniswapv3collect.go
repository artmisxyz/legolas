package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3Collect holds the schema definition for the UniswapV3Collect entity.
type UniswapV3Collect struct {
	ent.Schema
}

// Fields of the UniswapV3Collect.
func (UniswapV3Collect) Fields() []ent.Field {
	return []ent.Field{
		BigIntField("token_id"),
		field.String("recipient"),
		BigIntField("amount0"),
		BigIntField("amount1"),
	}
}

// Edges of the UniswapV3Collect.
func (UniswapV3Collect) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("collect").
			Unique().
			Required(),
	}
}
