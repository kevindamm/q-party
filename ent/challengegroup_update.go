// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kevindamm/q-party/ent/category"
	"github.com/kevindamm/q-party/ent/challenge"
	"github.com/kevindamm/q-party/ent/challengegroup"
	"github.com/kevindamm/q-party/ent/episoderound"
	"github.com/kevindamm/q-party/ent/predicate"
)

// ChallengeGroupUpdate is the builder for updating ChallengeGroup entities.
type ChallengeGroupUpdate struct {
	config
	hooks    []Hook
	mutation *ChallengeGroupMutation
}

// Where appends a list predicates to the ChallengeGroupUpdate builder.
func (cgu *ChallengeGroupUpdate) Where(ps ...predicate.ChallengeGroup) *ChallengeGroupUpdate {
	cgu.mutation.Where(ps...)
	return cgu
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (cgu *ChallengeGroupUpdate) AddCategoryIDs(ids ...int) *ChallengeGroupUpdate {
	cgu.mutation.AddCategoryIDs(ids...)
	return cgu
}

// AddCategory adds the "category" edges to the Category entity.
func (cgu *ChallengeGroupUpdate) AddCategory(c ...*Category) *ChallengeGroupUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cgu.AddCategoryIDs(ids...)
}

// AddChallengeIDs adds the "challenges" edge to the Challenge entity by IDs.
func (cgu *ChallengeGroupUpdate) AddChallengeIDs(ids ...int) *ChallengeGroupUpdate {
	cgu.mutation.AddChallengeIDs(ids...)
	return cgu
}

// AddChallenges adds the "challenges" edges to the Challenge entity.
func (cgu *ChallengeGroupUpdate) AddChallenges(c ...*Challenge) *ChallengeGroupUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cgu.AddChallengeIDs(ids...)
}

// AddEpisodeRoundIDs adds the "episode_round" edge to the EpisodeRound entity by IDs.
func (cgu *ChallengeGroupUpdate) AddEpisodeRoundIDs(ids ...int) *ChallengeGroupUpdate {
	cgu.mutation.AddEpisodeRoundIDs(ids...)
	return cgu
}

// AddEpisodeRound adds the "episode_round" edges to the EpisodeRound entity.
func (cgu *ChallengeGroupUpdate) AddEpisodeRound(e ...*EpisodeRound) *ChallengeGroupUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cgu.AddEpisodeRoundIDs(ids...)
}

// Mutation returns the ChallengeGroupMutation object of the builder.
func (cgu *ChallengeGroupUpdate) Mutation() *ChallengeGroupMutation {
	return cgu.mutation
}

// ClearCategory clears all "category" edges to the Category entity.
func (cgu *ChallengeGroupUpdate) ClearCategory() *ChallengeGroupUpdate {
	cgu.mutation.ClearCategory()
	return cgu
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (cgu *ChallengeGroupUpdate) RemoveCategoryIDs(ids ...int) *ChallengeGroupUpdate {
	cgu.mutation.RemoveCategoryIDs(ids...)
	return cgu
}

// RemoveCategory removes "category" edges to Category entities.
func (cgu *ChallengeGroupUpdate) RemoveCategory(c ...*Category) *ChallengeGroupUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cgu.RemoveCategoryIDs(ids...)
}

// ClearChallenges clears all "challenges" edges to the Challenge entity.
func (cgu *ChallengeGroupUpdate) ClearChallenges() *ChallengeGroupUpdate {
	cgu.mutation.ClearChallenges()
	return cgu
}

// RemoveChallengeIDs removes the "challenges" edge to Challenge entities by IDs.
func (cgu *ChallengeGroupUpdate) RemoveChallengeIDs(ids ...int) *ChallengeGroupUpdate {
	cgu.mutation.RemoveChallengeIDs(ids...)
	return cgu
}

// RemoveChallenges removes "challenges" edges to Challenge entities.
func (cgu *ChallengeGroupUpdate) RemoveChallenges(c ...*Challenge) *ChallengeGroupUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cgu.RemoveChallengeIDs(ids...)
}

// ClearEpisodeRound clears all "episode_round" edges to the EpisodeRound entity.
func (cgu *ChallengeGroupUpdate) ClearEpisodeRound() *ChallengeGroupUpdate {
	cgu.mutation.ClearEpisodeRound()
	return cgu
}

// RemoveEpisodeRoundIDs removes the "episode_round" edge to EpisodeRound entities by IDs.
func (cgu *ChallengeGroupUpdate) RemoveEpisodeRoundIDs(ids ...int) *ChallengeGroupUpdate {
	cgu.mutation.RemoveEpisodeRoundIDs(ids...)
	return cgu
}

// RemoveEpisodeRound removes "episode_round" edges to EpisodeRound entities.
func (cgu *ChallengeGroupUpdate) RemoveEpisodeRound(e ...*EpisodeRound) *ChallengeGroupUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cgu.RemoveEpisodeRoundIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cgu *ChallengeGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cgu.sqlSave, cgu.mutation, cgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cgu *ChallengeGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := cgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cgu *ChallengeGroupUpdate) Exec(ctx context.Context) error {
	_, err := cgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgu *ChallengeGroupUpdate) ExecX(ctx context.Context) {
	if err := cgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cgu *ChallengeGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(challengegroup.Table, challengegroup.Columns, sqlgraph.NewFieldSpec(challengegroup.FieldID, field.TypeInt))
	if ps := cgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cgu.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !cgu.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.CategoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cgu.mutation.ChallengesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.RemovedChallengesIDs(); len(nodes) > 0 && !cgu.mutation.ChallengesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.ChallengesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cgu.mutation.EpisodeRoundCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.RemovedEpisodeRoundIDs(); len(nodes) > 0 && !cgu.mutation.EpisodeRoundCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.EpisodeRoundIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challengegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cgu.mutation.done = true
	return n, nil
}

// ChallengeGroupUpdateOne is the builder for updating a single ChallengeGroup entity.
type ChallengeGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChallengeGroupMutation
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (cguo *ChallengeGroupUpdateOne) AddCategoryIDs(ids ...int) *ChallengeGroupUpdateOne {
	cguo.mutation.AddCategoryIDs(ids...)
	return cguo
}

// AddCategory adds the "category" edges to the Category entity.
func (cguo *ChallengeGroupUpdateOne) AddCategory(c ...*Category) *ChallengeGroupUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cguo.AddCategoryIDs(ids...)
}

// AddChallengeIDs adds the "challenges" edge to the Challenge entity by IDs.
func (cguo *ChallengeGroupUpdateOne) AddChallengeIDs(ids ...int) *ChallengeGroupUpdateOne {
	cguo.mutation.AddChallengeIDs(ids...)
	return cguo
}

// AddChallenges adds the "challenges" edges to the Challenge entity.
func (cguo *ChallengeGroupUpdateOne) AddChallenges(c ...*Challenge) *ChallengeGroupUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cguo.AddChallengeIDs(ids...)
}

// AddEpisodeRoundIDs adds the "episode_round" edge to the EpisodeRound entity by IDs.
func (cguo *ChallengeGroupUpdateOne) AddEpisodeRoundIDs(ids ...int) *ChallengeGroupUpdateOne {
	cguo.mutation.AddEpisodeRoundIDs(ids...)
	return cguo
}

// AddEpisodeRound adds the "episode_round" edges to the EpisodeRound entity.
func (cguo *ChallengeGroupUpdateOne) AddEpisodeRound(e ...*EpisodeRound) *ChallengeGroupUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cguo.AddEpisodeRoundIDs(ids...)
}

// Mutation returns the ChallengeGroupMutation object of the builder.
func (cguo *ChallengeGroupUpdateOne) Mutation() *ChallengeGroupMutation {
	return cguo.mutation
}

// ClearCategory clears all "category" edges to the Category entity.
func (cguo *ChallengeGroupUpdateOne) ClearCategory() *ChallengeGroupUpdateOne {
	cguo.mutation.ClearCategory()
	return cguo
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (cguo *ChallengeGroupUpdateOne) RemoveCategoryIDs(ids ...int) *ChallengeGroupUpdateOne {
	cguo.mutation.RemoveCategoryIDs(ids...)
	return cguo
}

// RemoveCategory removes "category" edges to Category entities.
func (cguo *ChallengeGroupUpdateOne) RemoveCategory(c ...*Category) *ChallengeGroupUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cguo.RemoveCategoryIDs(ids...)
}

// ClearChallenges clears all "challenges" edges to the Challenge entity.
func (cguo *ChallengeGroupUpdateOne) ClearChallenges() *ChallengeGroupUpdateOne {
	cguo.mutation.ClearChallenges()
	return cguo
}

// RemoveChallengeIDs removes the "challenges" edge to Challenge entities by IDs.
func (cguo *ChallengeGroupUpdateOne) RemoveChallengeIDs(ids ...int) *ChallengeGroupUpdateOne {
	cguo.mutation.RemoveChallengeIDs(ids...)
	return cguo
}

// RemoveChallenges removes "challenges" edges to Challenge entities.
func (cguo *ChallengeGroupUpdateOne) RemoveChallenges(c ...*Challenge) *ChallengeGroupUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cguo.RemoveChallengeIDs(ids...)
}

// ClearEpisodeRound clears all "episode_round" edges to the EpisodeRound entity.
func (cguo *ChallengeGroupUpdateOne) ClearEpisodeRound() *ChallengeGroupUpdateOne {
	cguo.mutation.ClearEpisodeRound()
	return cguo
}

// RemoveEpisodeRoundIDs removes the "episode_round" edge to EpisodeRound entities by IDs.
func (cguo *ChallengeGroupUpdateOne) RemoveEpisodeRoundIDs(ids ...int) *ChallengeGroupUpdateOne {
	cguo.mutation.RemoveEpisodeRoundIDs(ids...)
	return cguo
}

// RemoveEpisodeRound removes "episode_round" edges to EpisodeRound entities.
func (cguo *ChallengeGroupUpdateOne) RemoveEpisodeRound(e ...*EpisodeRound) *ChallengeGroupUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cguo.RemoveEpisodeRoundIDs(ids...)
}

// Where appends a list predicates to the ChallengeGroupUpdate builder.
func (cguo *ChallengeGroupUpdateOne) Where(ps ...predicate.ChallengeGroup) *ChallengeGroupUpdateOne {
	cguo.mutation.Where(ps...)
	return cguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cguo *ChallengeGroupUpdateOne) Select(field string, fields ...string) *ChallengeGroupUpdateOne {
	cguo.fields = append([]string{field}, fields...)
	return cguo
}

// Save executes the query and returns the updated ChallengeGroup entity.
func (cguo *ChallengeGroupUpdateOne) Save(ctx context.Context) (*ChallengeGroup, error) {
	return withHooks(ctx, cguo.sqlSave, cguo.mutation, cguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cguo *ChallengeGroupUpdateOne) SaveX(ctx context.Context) *ChallengeGroup {
	node, err := cguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cguo *ChallengeGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := cguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cguo *ChallengeGroupUpdateOne) ExecX(ctx context.Context) {
	if err := cguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cguo *ChallengeGroupUpdateOne) sqlSave(ctx context.Context) (_node *ChallengeGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(challengegroup.Table, challengegroup.Columns, sqlgraph.NewFieldSpec(challengegroup.FieldID, field.TypeInt))
	id, ok := cguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChallengeGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, challengegroup.FieldID)
		for _, f := range fields {
			if !challengegroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != challengegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cguo.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !cguo.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.CategoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cguo.mutation.ChallengesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.RemovedChallengesIDs(); len(nodes) > 0 && !cguo.mutation.ChallengesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.ChallengesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cguo.mutation.EpisodeRoundCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.RemovedEpisodeRoundIDs(); len(nodes) > 0 && !cguo.mutation.EpisodeRoundCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.EpisodeRoundIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ChallengeGroup{config: cguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challengegroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cguo.mutation.done = true
	return _node, nil
}
