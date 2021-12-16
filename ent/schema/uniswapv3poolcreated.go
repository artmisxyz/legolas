package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.String("token0").NotEmpty(),
		field.String("token1").NotEmpty(),
		field.String("fee").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("tick_spacing").GoType(new(BigInt)).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18, 0)",
				dialect.SQLite:   "numeric(18, 0)",
			}),
		field.String("pool").NotEmpty().Unique(),
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
