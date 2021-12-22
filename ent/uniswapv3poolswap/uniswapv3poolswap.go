// Code generated by entc, DO NOT EDIT.

package uniswapv3poolswap

const (
	// Label holds the string label denoting the uniswapv3poolswap type in the database.
	Label = "uniswap_v3pool_swap"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSender holds the string denoting the sender field in the database.
	FieldSender = "sender"
	// FieldRecipient holds the string denoting the recipient field in the database.
	FieldRecipient = "recipient"
	// FieldAmount0 holds the string denoting the amount0 field in the database.
	FieldAmount0 = "amount0"
	// FieldAmount1 holds the string denoting the amount1 field in the database.
	FieldAmount1 = "amount1"
	// FieldSqrtPriceX96 holds the string denoting the sqrt_price_x96 field in the database.
	FieldSqrtPriceX96 = "sqrt_price_x96"
	// FieldLiquidity holds the string denoting the liquidity field in the database.
	FieldLiquidity = "liquidity"
	// FieldTick holds the string denoting the tick field in the database.
	FieldTick = "tick"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the uniswapv3poolswap in the database.
	Table = "uniswap_v3pool_swaps"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "uniswap_v3pool_swaps"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_id"
)

// Columns holds all SQL columns for uniswapv3poolswap fields.
var Columns = []string{
	FieldID,
	FieldSender,
	FieldRecipient,
	FieldAmount0,
	FieldAmount1,
	FieldSqrtPriceX96,
	FieldLiquidity,
	FieldTick,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "uniswap_v3pool_swaps"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_id",
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

var (
	// SenderValidator is a validator for the "sender" field. It is called by the builders before save.
	SenderValidator func(string) error
	// RecipientValidator is a validator for the "recipient" field. It is called by the builders before save.
	RecipientValidator func(string) error
)