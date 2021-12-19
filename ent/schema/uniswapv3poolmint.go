package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UniswapV3PoolMint holds the schema definition for the UniswapV3PoolMint entity.
type UniswapV3PoolMint struct {
	ent.Schema
}

//
//Owner     common.Address
//TickLower *big.Int
//TickUpper *big.Int
//Amount    *big.Int
//Amount0   *big.Int
//Amount1   *big.Int

// Fields of the UniswapV3PoolMint.
func (UniswapV3PoolMint) Fields() []ent.Field {
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

// Edges of the UniswapV3PoolMint.
func (UniswapV3PoolMint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("pool_mint").
			Unique().
			Required(),
	}
}
