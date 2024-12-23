// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kevindamm/q-party/ent/challenge"
	"github.com/kevindamm/q-party/ent/challengegroup"
)

// ChallengeCreate is the builder for creating a Challenge entity.
type ChallengeCreate struct {
	config
	mutation *ChallengeMutation
	hooks    []Hook
}

// SetMedia sets the "media" field.
func (cc *ChallengeCreate) SetMedia(s string) *ChallengeCreate {
	cc.mutation.SetMedia(s)
	return cc
}

// SetPrompt sets the "prompt" field.
func (cc *ChallengeCreate) SetPrompt(s string) *ChallengeCreate {
	cc.mutation.SetPrompt(s)
	return cc
}

// SetResponse sets the "response" field.
func (cc *ChallengeCreate) SetResponse(s string) *ChallengeCreate {
	cc.mutation.SetResponse(s)
	return cc
}

// SetValue sets the "value" field.
func (cc *ChallengeCreate) SetValue(i int) *ChallengeCreate {
	cc.mutation.SetValue(i)
	return cc
}

// AddChallengeGroupIDs adds the "challenge_group" edge to the ChallengeGroup entity by IDs.
func (cc *ChallengeCreate) AddChallengeGroupIDs(ids ...int) *ChallengeCreate {
	cc.mutation.AddChallengeGroupIDs(ids...)
	return cc
}

// AddChallengeGroup adds the "challenge_group" edges to the ChallengeGroup entity.
func (cc *ChallengeCreate) AddChallengeGroup(c ...*ChallengeGroup) *ChallengeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddChallengeGroupIDs(ids...)
}

// Mutation returns the ChallengeMutation object of the builder.
func (cc *ChallengeCreate) Mutation() *ChallengeMutation {
	return cc.mutation
}

// Save creates the Challenge in the database.
func (cc *ChallengeCreate) Save(ctx context.Context) (*Challenge, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChallengeCreate) SaveX(ctx context.Context) *Challenge {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChallengeCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChallengeCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChallengeCreate) check() error {
	if _, ok := cc.mutation.Media(); !ok {
		return &ValidationError{Name: "media", err: errors.New(`ent: missing required field "Challenge.media"`)}
	}
	if _, ok := cc.mutation.Prompt(); !ok {
		return &ValidationError{Name: "prompt", err: errors.New(`ent: missing required field "Challenge.prompt"`)}
	}
	if _, ok := cc.mutation.Response(); !ok {
		return &ValidationError{Name: "response", err: errors.New(`ent: missing required field "Challenge.response"`)}
	}
	if _, ok := cc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Challenge.value"`)}
	}
	return nil
}

func (cc *ChallengeCreate) sqlSave(ctx context.Context) (*Challenge, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChallengeCreate) createSpec() (*Challenge, *sqlgraph.CreateSpec) {
	var (
		_node = &Challenge{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(challenge.Table, sqlgraph.NewFieldSpec(challenge.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Media(); ok {
		_spec.SetField(challenge.FieldMedia, field.TypeString, value)
		_node.Media = value
	}
	if value, ok := cc.mutation.Prompt(); ok {
		_spec.SetField(challenge.FieldPrompt, field.TypeString, value)
		_node.Prompt = value
	}
	if value, ok := cc.mutation.Response(); ok {
		_spec.SetField(challenge.FieldResponse, field.TypeString, value)
		_node.Response = value
	}
	if value, ok := cc.mutation.Value(); ok {
		_spec.SetField(challenge.FieldValue, field.TypeInt, value)
		_node.Value = value
	}
	if nodes := cc.mutation.ChallengeGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   challenge.ChallengeGroupTable,
			Columns: challenge.ChallengeGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(challengegroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChallengeCreateBulk is the builder for creating many Challenge entities in bulk.
type ChallengeCreateBulk struct {
	config
	err      error
	builders []*ChallengeCreate
}

// Save creates the Challenge entities in the database.
func (ccb *ChallengeCreateBulk) Save(ctx context.Context) ([]*Challenge, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Challenge, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChallengeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChallengeCreateBulk) SaveX(ctx context.Context) []*Challenge {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChallengeCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChallengeCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
