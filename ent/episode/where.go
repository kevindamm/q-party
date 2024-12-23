// Code generated by ent, DO NOT EDIT.

package episode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kevindamm/q-party/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldID, id))
}

// Aired applies equality check predicate on the "aired" field. It's identical to AiredEQ.
func Aired(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAired, v))
}

// AiredEQ applies the EQ predicate on the "aired" field.
func AiredEQ(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAired, v))
}

// AiredNEQ applies the NEQ predicate on the "aired" field.
func AiredNEQ(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldAired, v))
}

// AiredIn applies the In predicate on the "aired" field.
func AiredIn(vs ...time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldAired, vs...))
}

// AiredNotIn applies the NotIn predicate on the "aired" field.
func AiredNotIn(vs ...time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldAired, vs...))
}

// AiredGT applies the GT predicate on the "aired" field.
func AiredGT(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldAired, v))
}

// AiredGTE applies the GTE predicate on the "aired" field.
func AiredGTE(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldAired, v))
}

// AiredLT applies the LT predicate on the "aired" field.
func AiredLT(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldAired, v))
}

// AiredLTE applies the LTE predicate on the "aired" field.
func AiredLTE(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldAired, v))
}

// DifficultyEQ applies the EQ predicate on the "difficulty" field.
func DifficultyEQ(v Difficulty) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldDifficulty, v))
}

// DifficultyNEQ applies the NEQ predicate on the "difficulty" field.
func DifficultyNEQ(v Difficulty) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldDifficulty, v))
}

// DifficultyIn applies the In predicate on the "difficulty" field.
func DifficultyIn(vs ...Difficulty) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldDifficulty, vs...))
}

// DifficultyNotIn applies the NotIn predicate on the "difficulty" field.
func DifficultyNotIn(vs ...Difficulty) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldDifficulty, vs...))
}

// HasRounds applies the HasEdge predicate on the "rounds" edge.
func HasRounds() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RoundsTable, RoundsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoundsWith applies the HasEdge predicate on the "rounds" edge with a given conditions (other predicates).
func HasRoundsWith(preds ...predicate.EpisodeRound) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := newRoundsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.NotPredicates(p))
}
