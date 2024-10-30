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
		field.Int("episode").Unique(),
		field.Time("aired"),

		field.Enum("difficulty").
			Values(
				"UNKNOWN",
				"Kids",
				"College",
				"Standard",
				"Champions"),

		field.Bytes("content_hash").Nillable(),
		field.Time("fetched").Nillable().
			Comment("when this episode was fetched from j-archive, nil if not yet fetched"),
		field.Time("converted").Nillable().
			Comment("when this episode was parsed and encoded to JSON, nil if not yet converted"),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rounds", EpisodeRound.Type),
	}
}
