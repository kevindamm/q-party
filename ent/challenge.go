// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kevindamm/q-party/ent/challenge"
)

// Challenge is the model entity for the Challenge schema.
type Challenge struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Media holds the value of the "media" field.
	Media string `json:"media,omitempty"`
	// May be a URL path if the media type is defined.
	Prompt string `json:"prompt,omitempty"`
	// Response holds the value of the "response" field.
	Response string `json:"response,omitempty"`
	// If positive, the board's given value; if negative, the player's wagered value.
	Value int `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChallengeQuery when eager-loading is set.
	Edges        ChallengeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ChallengeEdges holds the relations/edges for other nodes in the graph.
type ChallengeEdges struct {
	// Column holds the value of the column edge.
	Column []*ChallengeGroup `json:"column,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ColumnOrErr returns the Column value or an error if the edge
// was not loaded in eager-loading.
func (e ChallengeEdges) ColumnOrErr() ([]*ChallengeGroup, error) {
	if e.loadedTypes[0] {
		return e.Column, nil
	}
	return nil, &NotLoadedError{edge: "column"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Challenge) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case challenge.FieldID, challenge.FieldValue:
			values[i] = new(sql.NullInt64)
		case challenge.FieldMedia, challenge.FieldPrompt, challenge.FieldResponse:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Challenge fields.
func (c *Challenge) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case challenge.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case challenge.FieldMedia:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field media", values[i])
			} else if value.Valid {
				c.Media = value.String
			}
		case challenge.FieldPrompt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prompt", values[i])
			} else if value.Valid {
				c.Prompt = value.String
			}
		case challenge.FieldResponse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field response", values[i])
			} else if value.Valid {
				c.Response = value.String
			}
		case challenge.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				c.Value = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Challenge.
// This includes values selected through modifiers, order, etc.
func (c *Challenge) GetValue(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryColumn queries the "column" edge of the Challenge entity.
func (c *Challenge) QueryColumn() *ChallengeGroupQuery {
	return NewChallengeClient(c.config).QueryColumn(c)
}

// Update returns a builder for updating this Challenge.
// Note that you need to call Challenge.Unwrap() before calling this method if this Challenge
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Challenge) Update() *ChallengeUpdateOne {
	return NewChallengeClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Challenge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Challenge) Unwrap() *Challenge {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Challenge is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Challenge) String() string {
	var builder strings.Builder
	builder.WriteString("Challenge(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("media=")
	builder.WriteString(c.Media)
	builder.WriteString(", ")
	builder.WriteString("prompt=")
	builder.WriteString(c.Prompt)
	builder.WriteString(", ")
	builder.WriteString("response=")
	builder.WriteString(c.Response)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", c.Value))
	builder.WriteByte(')')
	return builder.String()
}

// Challenges is a parsable slice of Challenge.
type Challenges []*Challenge
