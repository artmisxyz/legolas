package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time"),
		field.String("name").NotEmpty(),
		field.String("signature").NotEmpty(),
		field.String("address").NotEmpty(),
		field.Uint64("block_number"),
		field.String("tx_hash").NotEmpty(),
		field.Uint("tx_index"),
		field.String("block_hash").NotEmpty(),
		field.Uint("index"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("increase_liquidity", UniswapV3IncreaseLiqudity.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("decrease_liquidity", UniswapV3DecreaseLiqudity.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("collect", UniswapV3Collect.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("transfer", UniswapV3Transfer.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("pool_created", UniswapV3PoolCreated.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("pool_initialize", UniswapV3PoolInitialize.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("pool_swap", UniswapV3PoolSwap.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("pool_mint", UniswapV3PoolMint.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("pool_burn", UniswapV3PoolBurn.Type).Unique().StorageKey(edge.Column("event_id")),
		edge.To("pool_flash", UniswapV3PoolFlash.Type).Unique().StorageKey(edge.Column("event_id")),
	}
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address"),
		index.Fields("index", "tx_hash").Unique(),
	}
}

func BigIntField(name string) ent.Field {
	return field.String(name).SchemaType(map[string]string{
		dialect.SQLite:   "integer",
		dialect.Postgres: "numeric(50, 0)",
	})
}
