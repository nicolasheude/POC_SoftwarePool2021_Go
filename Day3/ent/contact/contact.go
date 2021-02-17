// Code generated by entc, DO NOT EDIT.

package contact

const (
	// Label holds the string label denoting the contact type in the database.
	Label = "contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldGithub holds the string denoting the github field in the database.
	FieldGithub = "github"
	// FieldLinkedin holds the string denoting the linkedin field in the database.
	FieldLinkedin = "linkedin"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the contact in the database.
	Table = "contacts"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "contacts"
	// OwnerInverseTable is the table name for the Developper entity.
	// It exists in this package in order to avoid circular dependency with the "developper" package.
	OwnerInverseTable = "developpers"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "developper_contact"
)

// Columns holds all SQL columns for contact fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPhone,
	FieldGithub,
	FieldLinkedin,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Contact type.
var ForeignKeys = []string{
	"developper_contact",
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)