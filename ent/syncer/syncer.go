// Code generated by entc, DO NOT EDIT.

package syncer

const (
	// Label holds the string label denoting the syncer type in the database.
	Label = "syncer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldFinish holds the string denoting the finish field in the database.
	FieldFinish = "finish"
	// FieldCurrent holds the string denoting the current field in the database.
	FieldCurrent = "current"
	// Table holds the table name of the syncer in the database.
	Table = "syncers"
)

// Columns holds all SQL columns for syncer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStart,
	FieldFinish,
	FieldCurrent,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
