package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3PoolSwap holds the schema definition for the UniswapV3PoolSwap entity.
type UniswapV3PoolSwap struct {
	ent.Schema
}

// Fields of the UniswapV3PoolSwap.
func (UniswapV3PoolSwap) Fields() []ent.Field {
	return []ent.Field{
		field.String("sender"),
		field.String("recipient"),
		BigIntField("amount0"),
		BigIntField("amount1"),
		BigIntField("sqrt_price_x96"),
		BigIntField("liquidity"),
		BigIntField("tick"),
	}
}

// Edges of the UniswapV3PoolSwap.
func (UniswapV3PoolSwap) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_swap").
			Unique().
			Required(),
	}
}
