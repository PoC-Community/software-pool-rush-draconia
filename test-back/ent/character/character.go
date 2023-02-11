// Code generated by ent, DO NOT EDIT.

package character

const (
	// Label holds the string label denoting the character type in the database.
	Label = "character"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// Table holds the table name of the character in the database.
	Table = "characters"
)

// Columns holds all SQL columns for character fields.
var Columns = []string{
	FieldID,
	FieldPrice,
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