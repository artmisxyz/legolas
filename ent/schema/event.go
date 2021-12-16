package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("signature").NotEmpty(),
		field.String("address").NotEmpty(),
		field.Uint64("block_number"),
		field.String("tx_hash").NotEmpty(),
		field.Uint("tx_index"),
		field.String("block_hash").NotEmpty(),
		field.Uint("index"),
		field.String("hash").Unique().NotEmpty(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("increase_liquidity", UniswapV3IncreaseLiqudity.Type),
		edge.To("decrease_liquidity", UniswapV3DecreaseLiqudity.Type),
		edge.To("collect", UniswapV3Collect.Type),
		edge.To("transfer", UniswapV3Transfer.Type),
	}
}
