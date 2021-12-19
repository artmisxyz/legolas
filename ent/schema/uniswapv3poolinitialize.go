package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3PoolInitialize holds the schema definition for the UniswapV3PoolInitialize entity.
type UniswapV3PoolInitialize struct {
	ent.Schema
}

// Fields of the UniswapV3PoolInitialize.
func (UniswapV3PoolInitialize) Fields() []ent.Field {
	return []ent.Field{
		field.String("sqrt_price_x96").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("tick").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
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
