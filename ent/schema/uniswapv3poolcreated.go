package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3PoolCreated holds the schema definition for the UniswapV3PoolCreated entity.
type UniswapV3PoolCreated struct {
	ent.Schema
}

// Fields of the UniswapV3PoolCreated.
func (UniswapV3PoolCreated) Fields() []ent.Field {
	return []ent.Field{
		field.String("token0"),
		field.String("token1"),
		BigIntField("fee"),
		BigIntField("tick_spacing"),
		field.String("pool").Unique(),
	}
}

// Edges of the UniswapV3PoolCreated.
func (UniswapV3PoolCreated) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_created").
			Unique().
			Required(),
	}
}
