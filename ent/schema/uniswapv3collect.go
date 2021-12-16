package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3Collect holds the schema definition for the UniswapV3Collect entity.
type UniswapV3Collect struct {
	ent.Schema
}

// Fields of the UniswapV3Collect.
func (UniswapV3Collect) Fields() []ent.Field {
	return []ent.Field{
		field.String("token_id").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
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
	}
}

// Edges of the UniswapV3Collect.
func (UniswapV3Collect) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("collect").
			Unique().
			Required(),
	}
}
