package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3DecreaseLiqudity holds the schema definition for the UniswapV3DecreaseLiqudity entity.
type UniswapV3DecreaseLiqudity struct {
	ent.Schema
}

// Fields of the UniswapV3DecreaseLiqudity.
func (UniswapV3DecreaseLiqudity) Fields() []ent.Field {
	return []ent.Field{
		field.String("token_id").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("liquidity").GoType(new(BigInt)).
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

// Edges of the UniswapV3DecreaseLiqudity.
func (UniswapV3DecreaseLiqudity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("event", Event.Type),
	}
}
