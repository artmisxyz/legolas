package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// UniswapV3IncreaseLiqudity holds the schema definition for the UniswapV3IncreaseLiqudity entity.
type UniswapV3IncreaseLiqudity struct {
	ent.Schema
}

// Fields of the UniswapV3IncreaseLiqudity.
func (UniswapV3IncreaseLiqudity) Fields() []ent.Field {
	return []ent.Field{
		BigIntField("token_id"),
		BigIntField("liquidity"),
		BigIntField("amount0"),
		BigIntField("amount1"),
	}
}

// Edges of the UniswapV3IncreaseLiqudity.
func (UniswapV3IncreaseLiqudity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("increase_liquidity").
			Unique().
			Required(),
	}
}
