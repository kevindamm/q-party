package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Challenge holds the schema definition for the Challenge entity.
type Challenge struct {
	ent.Schema
}

// Fields of the Challenge.
func (Challenge) Fields() []ent.Field {
	return []ent.Field{
		field.String("media"),
		field.String("prompt").
			Comment("May be a URL path if the media type is defined."),

		field.String("response"),

		field.Int("value").
			Comment("If positive, the board's given value; if negative, the player's wagered value."),
	}
}

// Edges of the Challenge.
func (Challenge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("challenge_group", ChallengeGroup.Type).
			Ref("challenges"),
	}
}
