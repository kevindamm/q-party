// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kevindamm/q-party/ent/category"
	"github.com/kevindamm/q-party/ent/challenge"
	"github.com/kevindamm/q-party/ent/challengegroup"
	"github.com/kevindamm/q-party/ent/episoderound"
)

// ChallengeGroupCreate is the builder for creating a ChallengeGroup entity.
type ChallengeGroupCreate struct {
	config
	mutation *ChallengeGroupMutation
	hooks    []Hook
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (cgc *ChallengeGroupCreate) AddCategoryIDs(ids ...int) *ChallengeGroupCreate {
	cgc.mutation.AddCategoryIDs(ids...)
	return cgc
}

// AddCategory adds the "category" edges to the Category entity.
func (cgc *ChallengeGroupCreate) AddCategory(c ...*Category) *ChallengeGroupCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cgc.AddCategoryIDs(ids...)
}

// AddChallengeIDs adds the "challenges" edge to the Challenge entity by IDs.
func (cgc *ChallengeGroupCreate) AddChallengeIDs(ids ...int) *ChallengeGroupCreate {
	cgc.mutation.AddChallengeIDs(ids...)
	return cgc
}

// AddChallenges adds the "challenges" edges to the Challenge entity.
func (cgc *ChallengeGroupCreate) AddChallenges(c ...*Challenge) *ChallengeGroupCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cgc.AddChallengeIDs(ids...)
}

// AddEpisodeRoundIDs adds the "episode_round" edge to the EpisodeRound entity by IDs.
func (cgc *ChallengeGroupCreate) AddEpisodeRoundIDs(ids ...int) *ChallengeGroupCreate {
	cgc.mutation.AddEpisodeRoundIDs(ids...)
	return cgc
}

// AddEpisodeRound adds the "episode_round" edges to the EpisodeRound entity.
func (cgc *ChallengeGroupCreate) AddEpisodeRound(e ...*EpisodeRound) *ChallengeGroupCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cgc.AddEpisodeRoundIDs(ids...)
}

// Mutation returns the ChallengeGroupMutation object of the builder.
func (cgc *ChallengeGroupCreate) Mutation() *ChallengeGroupMutation {
	return cgc.mutation
}

// Save creates the ChallengeGroup in the database.
func (cgc *ChallengeGroupCreate) Save(ctx context.Context) (*ChallengeGroup, error) {
	return withHooks(ctx, cgc.sqlSave, cgc.mutation, cgc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cgc *ChallengeGroupCreate) SaveX(ctx context.Context) *ChallengeGroup {
	v, err := cgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cgc *ChallengeGroupCreate) Exec(ctx context.Context) error {
	_, err := cgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgc *ChallengeGroupCreate) ExecX(ctx context.Context) {
	if err := cgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cgc *ChallengeGroupCreate) check() error {
	return nil
}

func (cgc *ChallengeGroupCreate) sqlSave(ctx context.Context) (*ChallengeGroup, error) {
	if err := cgc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cgc.mutation.id = &_node.ID
	cgc.mutation.done = true
	return _node, nil
}

func (cgc *ChallengeGroupCreate) createSpec() (*ChallengeGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &ChallengeGroup{config: cgc.config}
		_spec = sqlgraph.NewCreateSpec(challengegroup.Table, sqlgraph.NewFieldSpec(challengegroup.FieldID, field.TypeInt))
	)
	if nodes := cgc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   challengegroup.CategoryTable,
			Columns: challengegroup.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cgc.mutation.ChallengesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   challengegroup.ChallengesTable,
			Columns: challengegroup.ChallengesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(challenge.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cgc.mutation.EpisodeRoundIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   challengegroup.EpisodeRoundTable,
			Columns: challengegroup.EpisodeRoundPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episoderound.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChallengeGroupCreateBulk is the builder for creating many ChallengeGroup entities in bulk.
type ChallengeGroupCreateBulk struct {
	config
	err      error
	builders []*ChallengeGroupCreate
}

// Save creates the ChallengeGroup entities in the database.
func (cgcb *ChallengeGroupCreateBulk) Save(ctx context.Context) ([]*ChallengeGroup, error) {
	if cgcb.err != nil {
		return nil, cgcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cgcb.builders))
	nodes := make([]*ChallengeGroup, len(cgcb.builders))
	mutators := make([]Mutator, len(cgcb.builders))
	for i := range cgcb.builders {
		func(i int, root context.Context) {
			builder := cgcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChallengeGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, cgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cgcb *ChallengeGroupCreateBulk) SaveX(ctx context.Context) []*ChallengeGroup {
	v, err := cgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cgcb *ChallengeGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := cgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgcb *ChallengeGroupCreateBulk) ExecX(ctx context.Context) {
	if err := cgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
