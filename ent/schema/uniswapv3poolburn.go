package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3PoolBurn holds the schema definition for the UniswapV3PoolBurn entity.
type UniswapV3PoolBurn struct {
	ent.Schema
}

// Fields of the UniswapV3PoolBurn.
func (UniswapV3PoolBurn) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner"),
		BigIntField("tick_lower"),
		BigIntField("tick_upper"),
		BigIntField("amount"),
		BigIntField("amount0"),
		BigIntField("amount1"),
	}
}

// Edges of the UniswapV3PoolBurn.
func (UniswapV3PoolBurn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_burn").
			Unique().
			Required(),
	}
}
