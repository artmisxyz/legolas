package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// UniswapV3PoolInitialize holds the schema definition for the UniswapV3PoolInitialize entity.
type UniswapV3PoolInitialize struct {
	ent.Schema
}

// Fields of the UniswapV3PoolInitialize.
func (UniswapV3PoolInitialize) Fields() []ent.Field {
	return []ent.Field{
		BigIntField("sqrt_price_x96"),
		BigIntField("tick"),
	}
}

// Edges of the UniswapV3PoolInitialize.
func (UniswapV3PoolInitialize) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_initialize").
			Unique().
			Required(),
	}
}
