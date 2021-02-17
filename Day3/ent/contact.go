// Code generated by entc, DO NOT EDIT.

package ent

import (
	"SofwareGoDay3/ent/contact"
	"SofwareGoDay3/ent/developper"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Github holds the value of the "github" field.
	Github string `json:"github,omitempty"`
	// Linkedin holds the value of the "linkedin" field.
	Linkedin string `json:"linkedin,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContactQuery when eager-loading is set.
	Edges              ContactEdges `json:"edges"`
	developper_contact *int
}

// ContactEdges holds the relations/edges for other nodes in the graph.
type ContactEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Developper `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) OwnerOrErr() (*Developper, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: developper.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			values[i] = &sql.NullInt64{}
		case contact.FieldEmail, contact.FieldPhone, contact.FieldGithub, contact.FieldLinkedin:
			values[i] = &sql.NullString{}
		case contact.ForeignKeys[0]: // developper_contact
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case contact.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case contact.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case contact.FieldGithub:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field github", values[i])
			} else if value.Valid {
				c.Github = value.String
			}
		case contact.FieldLinkedin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field linkedin", values[i])
			} else if value.Valid {
				c.Linkedin = value.String
			}
		case contact.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field developper_contact", value)
			} else if value.Valid {
				c.developper_contact = new(int)
				*c.developper_contact = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Contact entity.
func (c *Contact) QueryOwner() *DevelopperQuery {
	return (&ContactClient{config: c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return (&ContactClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contact is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", email=")
	builder.WriteString(c.Email)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", github=")
	builder.WriteString(c.Github)
	builder.WriteString(", linkedin=")
	builder.WriteString(c.Linkedin)
	builder.WriteByte(')')
	return builder.String()
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact

func (c Contacts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
