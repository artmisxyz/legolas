package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3PoolFlash holds the schema definition for the UniswapV3PoolFlash entity.
type UniswapV3PoolFlash struct {
	ent.Schema
}

// Fields of the UniswapV3PoolFlash.
func (UniswapV3PoolFlash) Fields() []ent.Field {
	return []ent.Field{
		field.String("sender"),
		field.String("recipient"),
		BigIntField("amount0"),
		BigIntField("amount1"),
		BigIntField("paid0"),
		BigIntField("paid1"),
	}
}

// Edges of the UniswapV3PoolFlash.
func (UniswapV3PoolFlash) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_flash").
			Unique().
			Required(),
	}
}
