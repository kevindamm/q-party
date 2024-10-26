package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// ChallengeGroup holds the schema definition for the ChallengeGroup entity.
type ChallengeGroup struct {
	ent.Schema
}

// Fields of the ChallengeGroup.
func (ChallengeGroup) Fields() []ent.Field {
	return nil
}

// Edges of the ChallengeGroup.
func (ChallengeGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("challenges"),
		edge.To("challenges", Challenge.Type),

		edge.From("episode_round", EpisodeRound.Type).
			Ref("columns"),
	}
}
