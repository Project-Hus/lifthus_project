// Code generated by ent, DO NOT EDIT.

package connectedsession

const (
	// Label holds the string label denoting the connectedsession type in the database.
	Label = "connected_session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHsid holds the string denoting the hsid field in the database.
	FieldHsid = "hsid"
	// FieldService holds the string denoting the service field in the database.
	FieldService = "service"
	// FieldCsid holds the string denoting the csid field in the database.
	FieldCsid = "csid"
	// EdgeHusSession holds the string denoting the hus_session edge name in mutations.
	EdgeHusSession = "hus_session"
	// Table holds the table name of the connectedsession in the database.
	Table = "connected_sessions"
	// HusSessionTable is the table that holds the hus_session relation/edge.
	HusSessionTable = "connected_sessions"
	// HusSessionInverseTable is the table name for the HusSession entity.
	// It exists in this package in order to avoid circular dependency with the "hussession" package.
	HusSessionInverseTable = "hus_sessions"
	// HusSessionColumn is the table column denoting the hus_session relation/edge.
	HusSessionColumn = "hsid"
)

// Columns holds all SQL columns for connectedsession fields.
var Columns = []string{
	FieldID,
	FieldHsid,
	FieldService,
	FieldCsid,
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
