package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.String("sender").NotEmpty(),
		field.String("recipient").NotEmpty(),
		field.String("amount0").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("amount1").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("sqrt_price_x96").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("liquidity").GoType(new(BigInt)).
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

// Edges of the UniswapV3PoolSwap.
func (UniswapV3PoolSwap) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_swap").
			Unique().
			Required(),
	}
}
