package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3IncreaseLiqudity holds the schema definition for the UniswapV3IncreaseLiqudity entity.
type UniswapV3IncreaseLiqudity struct {
	ent.Schema
}

// Fields of the UniswapV3IncreaseLiqudity.
func (UniswapV3IncreaseLiqudity) Fields() []ent.Field {
	return []ent.Field{
		field.String("token_id").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("liquidity").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("amount0").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
		field.String("amount1").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
			}),
	}
}

// Edges of the UniswapV3IncreaseLiqudity.
func (UniswapV3IncreaseLiqudity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("event", Event.Type),
	}
}
