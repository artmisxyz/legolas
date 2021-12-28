package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3PoolMint holds the schema definition for the UniswapV3PoolMint entity.
type UniswapV3PoolMint struct {
	ent.Schema
}

// Fields of the UniswapV3PoolMint.
func (UniswapV3PoolMint) Fields() []ent.Field {
	return []ent.Field{
		field.String("sender").Optional(),
		field.String("owner"),
		BigIntField("tick_lower"),
		BigIntField("tick_upper"),
		BigIntField("amount"),
		BigIntField("amount0"),
		BigIntField("amount1"),
	}
}

// Edges of the UniswapV3PoolMint.
func (UniswapV3PoolMint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_mint").
			Unique().
			Required(),
	}
}
