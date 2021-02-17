package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Developper holds the schema definition for the Developper entity.
type Developper struct {
	ent.Schema
}

// Fields of the Developper.
func (Developper) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.String("name"),
		field.Int("age").Positive(),
		field.String("school"),
		field.Int("experience"),
	}
}

// Edges of the Developper.
func (Developper) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contact", Contact.Type).
			Unique(),
			edge.To("competence", Competence.Type),
	}
}
