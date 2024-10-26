package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Episode holds the schema definition for the Episode entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
	return []ent.Field{
		field.Time("aired"),

		field.Enum("difficulty").
			Values(
				"UNKNOWN",
				"Kids",
				"College",
				"Standard",
				"Champions"),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rounds", EpisodeRound.Type),
	}
}
