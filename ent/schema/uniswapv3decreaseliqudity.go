package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// UniswapV3DecreaseLiqudity holds the schema definition for the UniswapV3DecreaseLiqudity entity.
type UniswapV3DecreaseLiqudity struct {
	ent.Schema
}

// Fields of the UniswapV3DecreaseLiqudity.
func (UniswapV3DecreaseLiqudity) Fields() []ent.Field {
	return []ent.Field{
		BigIntField("token_id"),
		BigIntField("liquidity"),
		BigIntField("amount0"),
		BigIntField("amount1"),
	}
}

// Edges of the UniswapV3DecreaseLiqudity.
func (UniswapV3DecreaseLiqudity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("decrease_liquidity").
			Unique().
			Required(),
	}
}
