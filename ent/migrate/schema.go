// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "owner", Type: field.TypeBytes},
		{Name: "pool", Type: field.TypeBytes},
		{Name: "token0", Type: field.TypeBytes},
		{Name: "token1", Type: field.TypeBytes},
		{Name: "tick_lower", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "tick_upper", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "liquidity", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "deposited_token0", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "deposited_token1", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "withdrawn_token0", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "withdrawn_token1", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "collected_token0", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "collected_token1", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "collected_fees_token0", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "collected_fees_token1", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "fee_growth_inside0last_x128", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
		{Name: "fee_growth_inside1last_x128", Type: field.TypeString, SchemaType: map[string]string{"postgres": "numeric(18, 0)"}},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:       "positions",
		Columns:    PositionsColumns,
		PrimaryKey: []*schema.Column{PositionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PositionsTable,
	}
)

func init() {
}
