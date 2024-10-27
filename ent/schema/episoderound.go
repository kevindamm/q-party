package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EpisodeRound holds the schema definition for the EpisodeRound entity.
type EpisodeRound struct {
	ent.Schema
}

// Fields of the EpisodeRound.
func (EpisodeRound) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("round").
			Values(
				"UNKNOWN",
				"Single",
				"Double",
				"Final",
				"Tiebreaker"),
	}
}

// Edges of the EpisodeRound.
func (EpisodeRound) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categories", ChallengeGroup.Type),

		edge.From("episode", Episode.Type).
			Ref("rounds"),
	}
}
