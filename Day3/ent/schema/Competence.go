package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Competence holds the schema definition for the Competence entity.
type Competence struct {
	ent.Schema
}

// Fields of the Competence.
func (Competence) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.String("name"),
		field.Int("level").Min(0).Max(10),
	}
}

// Edges of the Competence.
func (Competence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Developper.Type).
			Ref("competence").
			Unique().
			Required(),
	}
}
