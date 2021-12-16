// Code generated by entc, DO NOT EDIT.

package uniswapv3increaseliqudity

const (
	// Label holds the string label denoting the uniswapv3increaseliqudity type in the database.
	Label = "uniswap_v3increase_liqudity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTokenID holds the string denoting the token_id field in the database.
	FieldTokenID = "token_id"
	// FieldLiquidity holds the string denoting the liquidity field in the database.
	FieldLiquidity = "liquidity"
	// FieldAmount0 holds the string denoting the amount0 field in the database.
	FieldAmount0 = "amount0"
	// FieldAmount1 holds the string denoting the amount1 field in the database.
	FieldAmount1 = "amount1"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the uniswapv3increaseliqudity in the database.
	Table = "uniswap_v3increase_liqudities"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "uniswap_v3increase_liqudities"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_increase_liquidity"
)

// Columns holds all SQL columns for uniswapv3increaseliqudity fields.
var Columns = []string{
	FieldID,
	FieldTokenID,
	FieldLiquidity,
	FieldAmount0,
	FieldAmount1,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "uniswap_v3increase_liqudities"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_increase_liquidity",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
