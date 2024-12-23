// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kevindamm/q-party/ent/challengegroup"
)

// ChallengeGroup is the model entity for the ChallengeGroup schema.
type ChallengeGroup struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChallengeGroupQuery when eager-loading is set.
	Edges        ChallengeGroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ChallengeGroupEdges holds the relations/edges for other nodes in the graph.
type ChallengeGroupEdges struct {
	// Challenges holds the value of the challenges edge.
	Challenges []*Challenge `json:"challenges,omitempty"`
	// Category holds the value of the category edge.
	Category []*Category `json:"category,omitempty"`
	// EpisodeRound holds the value of the episode_round edge.
	EpisodeRound []*EpisodeRound `json:"episode_round,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ChallengesOrErr returns the Challenges value or an error if the edge
// was not loaded in eager-loading.
func (e ChallengeGroupEdges) ChallengesOrErr() ([]*Challenge, error) {
	if e.loadedTypes[0] {
		return e.Challenges, nil
	}
	return nil, &NotLoadedError{edge: "challenges"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading.
func (e ChallengeGroupEdges) CategoryOrErr() ([]*Category, error) {
	if e.loadedTypes[1] {
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// EpisodeRoundOrErr returns the EpisodeRound value or an error if the edge
// was not loaded in eager-loading.
func (e ChallengeGroupEdges) EpisodeRoundOrErr() ([]*EpisodeRound, error) {
	if e.loadedTypes[2] {
		return e.EpisodeRound, nil
	}
	return nil, &NotLoadedError{edge: "episode_round"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChallengeGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case challengegroup.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChallengeGroup fields.
func (cg *ChallengeGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case challengegroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cg.ID = int(value.Int64)
		default:
			cg.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ChallengeGroup.
// This includes values selected through modifiers, order, etc.
func (cg *ChallengeGroup) Value(name string) (ent.Value, error) {
	return cg.selectValues.Get(name)
}

// QueryChallenges queries the "challenges" edge of the ChallengeGroup entity.
func (cg *ChallengeGroup) QueryChallenges() *ChallengeQuery {
	return NewChallengeGroupClient(cg.config).QueryChallenges(cg)
}

// QueryCategory queries the "category" edge of the ChallengeGroup entity.
func (cg *ChallengeGroup) QueryCategory() *CategoryQuery {
	return NewChallengeGroupClient(cg.config).QueryCategory(cg)
}

// QueryEpisodeRound queries the "episode_round" edge of the ChallengeGroup entity.
func (cg *ChallengeGroup) QueryEpisodeRound() *EpisodeRoundQuery {
	return NewChallengeGroupClient(cg.config).QueryEpisodeRound(cg)
}

// Update returns a builder for updating this ChallengeGroup.
// Note that you need to call ChallengeGroup.Unwrap() before calling this method if this ChallengeGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (cg *ChallengeGroup) Update() *ChallengeGroupUpdateOne {
	return NewChallengeGroupClient(cg.config).UpdateOne(cg)
}

// Unwrap unwraps the ChallengeGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cg *ChallengeGroup) Unwrap() *ChallengeGroup {
	_tx, ok := cg.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChallengeGroup is not a transactional entity")
	}
	cg.config.driver = _tx.drv
	return cg
}

// String implements the fmt.Stringer.
func (cg *ChallengeGroup) String() string {
	var builder strings.Builder
	builder.WriteString("ChallengeGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", cg.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ChallengeGroups is a parsable slice of ChallengeGroup.
type ChallengeGroups []*ChallengeGroup
