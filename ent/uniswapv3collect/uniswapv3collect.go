// Code generated by entc, DO NOT EDIT.

package uniswapv3collect

const (
	// Label holds the string label denoting the uniswapv3collect type in the database.
	Label = "uniswap_v3collect"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTokenID holds the string denoting the token_id field in the database.
	FieldTokenID = "token_id"
	// FieldRecipient holds the string denoting the recipient field in the database.
	FieldRecipient = "recipient"
	// FieldAmount0 holds the string denoting the amount0 field in the database.
	FieldAmount0 = "amount0"
	// FieldAmount1 holds the string denoting the amount1 field in the database.
	FieldAmount1 = "amount1"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the uniswapv3collect in the database.
	Table = "uniswap_v3collects"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "uniswap_v3collects"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_collect"
)

// Columns holds all SQL columns for uniswapv3collect fields.
var Columns = []string{
	FieldID,
	FieldTokenID,
	FieldRecipient,
	FieldAmount0,
	FieldAmount1,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "uniswap_v3collects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_collect",
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
	// RecipientValidator is a validator for the "recipient" field. It is called by the builders before save.
	RecipientValidator func(string) error
)
