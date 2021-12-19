package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.String("owner").NotEmpty(),
		field.String("tick_lower").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("tick_upper").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("amount").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
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
