// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kevindamm/q-party/ent/category"
	"github.com/kevindamm/q-party/ent/challenge"
	"github.com/kevindamm/q-party/ent/challengegroup"
	"github.com/kevindamm/q-party/ent/episoderound"
	"github.com/kevindamm/q-party/ent/predicate"
)

// ChallengeGroupQuery is the builder for querying ChallengeGroup entities.
type ChallengeGroupQuery struct {
	config
	ctx              *QueryContext
	order            []challengegroup.OrderOption
	inters           []Interceptor
	predicates       []predicate.ChallengeGroup
	withChallenges   *ChallengeQuery
	withCategory     *CategoryQuery
	withEpisodeRound *EpisodeRoundQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChallengeGroupQuery builder.
func (cgq *ChallengeGroupQuery) Where(ps ...predicate.ChallengeGroup) *ChallengeGroupQuery {
	cgq.predicates = append(cgq.predicates, ps...)
	return cgq
}

// Limit the number of records to be returned by this query.
func (cgq *ChallengeGroupQuery) Limit(limit int) *ChallengeGroupQuery {
	cgq.ctx.Limit = &limit
	return cgq
}

// Offset to start from.
func (cgq *ChallengeGroupQuery) Offset(offset int) *ChallengeGroupQuery {
	cgq.ctx.Offset = &offset
	return cgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cgq *ChallengeGroupQuery) Unique(unique bool) *ChallengeGroupQuery {
	cgq.ctx.Unique = &unique
	return cgq
}

// Order specifies how the records should be ordered.
func (cgq *ChallengeGroupQuery) Order(o ...challengegroup.OrderOption) *ChallengeGroupQuery {
	cgq.order = append(cgq.order, o...)
	return cgq
}

// QueryChallenges chains the current query on the "challenges" edge.
func (cgq *ChallengeGroupQuery) QueryChallenges() *ChallengeQuery {
	query := (&ChallengeClient{config: cgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(challengegroup.Table, challengegroup.FieldID, selector),
			sqlgraph.To(challenge.Table, challenge.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, challengegroup.ChallengesTable, challengegroup.ChallengesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategory chains the current query on the "category" edge.
func (cgq *ChallengeGroupQuery) QueryCategory() *CategoryQuery {
	query := (&CategoryClient{config: cgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(challengegroup.Table, challengegroup.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, challengegroup.CategoryTable, challengegroup.CategoryPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEpisodeRound chains the current query on the "episode_round" edge.
func (cgq *ChallengeGroupQuery) QueryEpisodeRound() *EpisodeRoundQuery {
	query := (&EpisodeRoundClient{config: cgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(challengegroup.Table, challengegroup.FieldID, selector),
			sqlgraph.To(episoderound.Table, episoderound.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, challengegroup.EpisodeRoundTable, challengegroup.EpisodeRoundPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ChallengeGroup entity from the query.
// Returns a *NotFoundError when no ChallengeGroup was found.
func (cgq *ChallengeGroupQuery) First(ctx context.Context) (*ChallengeGroup, error) {
	nodes, err := cgq.Limit(1).All(setContextOp(ctx, cgq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{challengegroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) FirstX(ctx context.Context) *ChallengeGroup {
	node, err := cgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChallengeGroup ID from the query.
// Returns a *NotFoundError when no ChallengeGroup ID was found.
func (cgq *ChallengeGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cgq.Limit(1).IDs(setContextOp(ctx, cgq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{challengegroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := cgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChallengeGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChallengeGroup entity is found.
// Returns a *NotFoundError when no ChallengeGroup entities are found.
func (cgq *ChallengeGroupQuery) Only(ctx context.Context) (*ChallengeGroup, error) {
	nodes, err := cgq.Limit(2).All(setContextOp(ctx, cgq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{challengegroup.Label}
	default:
		return nil, &NotSingularError{challengegroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) OnlyX(ctx context.Context) *ChallengeGroup {
	node, err := cgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChallengeGroup ID in the query.
// Returns a *NotSingularError when more than one ChallengeGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (cgq *ChallengeGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cgq.Limit(2).IDs(setContextOp(ctx, cgq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{challengegroup.Label}
	default:
		err = &NotSingularError{challengegroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := cgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChallengeGroups.
func (cgq *ChallengeGroupQuery) All(ctx context.Context) ([]*ChallengeGroup, error) {
	ctx = setContextOp(ctx, cgq.ctx, ent.OpQueryAll)
	if err := cgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChallengeGroup, *ChallengeGroupQuery]()
	return withInterceptors[[]*ChallengeGroup](ctx, cgq, qr, cgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) AllX(ctx context.Context) []*ChallengeGroup {
	nodes, err := cgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChallengeGroup IDs.
func (cgq *ChallengeGroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cgq.ctx.Unique == nil && cgq.path != nil {
		cgq.Unique(true)
	}
	ctx = setContextOp(ctx, cgq.ctx, ent.OpQueryIDs)
	if err = cgq.Select(challengegroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := cgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cgq *ChallengeGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cgq.ctx, ent.OpQueryCount)
	if err := cgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cgq, querierCount[*ChallengeGroupQuery](), cgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) CountX(ctx context.Context) int {
	count, err := cgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cgq *ChallengeGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cgq.ctx, ent.OpQueryExist)
	switch _, err := cgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cgq *ChallengeGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := cgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChallengeGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cgq *ChallengeGroupQuery) Clone() *ChallengeGroupQuery {
	if cgq == nil {
		return nil
	}
	return &ChallengeGroupQuery{
		config:           cgq.config,
		ctx:              cgq.ctx.Clone(),
		order:            append([]challengegroup.OrderOption{}, cgq.order...),
		inters:           append([]Interceptor{}, cgq.inters...),
		predicates:       append([]predicate.ChallengeGroup{}, cgq.predicates...),
		withChallenges:   cgq.withChallenges.Clone(),
		withCategory:     cgq.withCategory.Clone(),
		withEpisodeRound: cgq.withEpisodeRound.Clone(),
		// clone intermediate query.
		sql:  cgq.sql.Clone(),
		path: cgq.path,
	}
}

// WithChallenges tells the query-builder to eager-load the nodes that are connected to
// the "challenges" edge. The optional arguments are used to configure the query builder of the edge.
func (cgq *ChallengeGroupQuery) WithChallenges(opts ...func(*ChallengeQuery)) *ChallengeGroupQuery {
	query := (&ChallengeClient{config: cgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cgq.withChallenges = query
	return cgq
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (cgq *ChallengeGroupQuery) WithCategory(opts ...func(*CategoryQuery)) *ChallengeGroupQuery {
	query := (&CategoryClient{config: cgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cgq.withCategory = query
	return cgq
}

// WithEpisodeRound tells the query-builder to eager-load the nodes that are connected to
// the "episode_round" edge. The optional arguments are used to configure the query builder of the edge.
func (cgq *ChallengeGroupQuery) WithEpisodeRound(opts ...func(*EpisodeRoundQuery)) *ChallengeGroupQuery {
	query := (&EpisodeRoundClient{config: cgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cgq.withEpisodeRound = query
	return cgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (cgq *ChallengeGroupQuery) GroupBy(field string, fields ...string) *ChallengeGroupGroupBy {
	cgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChallengeGroupGroupBy{build: cgq}
	grbuild.flds = &cgq.ctx.Fields
	grbuild.label = challengegroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (cgq *ChallengeGroupQuery) Select(fields ...string) *ChallengeGroupSelect {
	cgq.ctx.Fields = append(cgq.ctx.Fields, fields...)
	sbuild := &ChallengeGroupSelect{ChallengeGroupQuery: cgq}
	sbuild.label = challengegroup.Label
	sbuild.flds, sbuild.scan = &cgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChallengeGroupSelect configured with the given aggregations.
func (cgq *ChallengeGroupQuery) Aggregate(fns ...AggregateFunc) *ChallengeGroupSelect {
	return cgq.Select().Aggregate(fns...)
}

func (cgq *ChallengeGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cgq); err != nil {
				return err
			}
		}
	}
	for _, f := range cgq.ctx.Fields {
		if !challengegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cgq.path != nil {
		prev, err := cgq.path(ctx)
		if err != nil {
			return err
		}
		cgq.sql = prev
	}
	return nil
}

func (cgq *ChallengeGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChallengeGroup, error) {
	var (
		nodes       = []*ChallengeGroup{}
		_spec       = cgq.querySpec()
		loadedTypes = [3]bool{
			cgq.withChallenges != nil,
			cgq.withCategory != nil,
			cgq.withEpisodeRound != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChallengeGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChallengeGroup{config: cgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cgq.withChallenges; query != nil {
		if err := cgq.loadChallenges(ctx, query, nodes,
			func(n *ChallengeGroup) { n.Edges.Challenges = []*Challenge{} },
			func(n *ChallengeGroup, e *Challenge) { n.Edges.Challenges = append(n.Edges.Challenges, e) }); err != nil {
			return nil, err
		}
	}
	if query := cgq.withCategory; query != nil {
		if err := cgq.loadCategory(ctx, query, nodes,
			func(n *ChallengeGroup) { n.Edges.Category = []*Category{} },
			func(n *ChallengeGroup, e *Category) { n.Edges.Category = append(n.Edges.Category, e) }); err != nil {
			return nil, err
		}
	}
	if query := cgq.withEpisodeRound; query != nil {
		if err := cgq.loadEpisodeRound(ctx, query, nodes,
			func(n *ChallengeGroup) { n.Edges.EpisodeRound = []*EpisodeRound{} },
			func(n *ChallengeGroup, e *EpisodeRound) { n.Edges.EpisodeRound = append(n.Edges.EpisodeRound, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cgq *ChallengeGroupQuery) loadChallenges(ctx context.Context, query *ChallengeQuery, nodes []*ChallengeGroup, init func(*ChallengeGroup), assign func(*ChallengeGroup, *Challenge)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ChallengeGroup)
	nids := make(map[int]map[*ChallengeGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(challengegroup.ChallengesTable)
		s.Join(joinT).On(s.C(challenge.FieldID), joinT.C(challengegroup.ChallengesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(challengegroup.ChallengesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(challengegroup.ChallengesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ChallengeGroup]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Challenge](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "challenges" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cgq *ChallengeGroupQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*ChallengeGroup, init func(*ChallengeGroup), assign func(*ChallengeGroup, *Category)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ChallengeGroup)
	nids := make(map[int]map[*ChallengeGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(challengegroup.CategoryTable)
		s.Join(joinT).On(s.C(category.FieldID), joinT.C(challengegroup.CategoryPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(challengegroup.CategoryPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(challengegroup.CategoryPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ChallengeGroup]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Category](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "category" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cgq *ChallengeGroupQuery) loadEpisodeRound(ctx context.Context, query *EpisodeRoundQuery, nodes []*ChallengeGroup, init func(*ChallengeGroup), assign func(*ChallengeGroup, *EpisodeRound)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ChallengeGroup)
	nids := make(map[int]map[*ChallengeGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(challengegroup.EpisodeRoundTable)
		s.Join(joinT).On(s.C(episoderound.FieldID), joinT.C(challengegroup.EpisodeRoundPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(challengegroup.EpisodeRoundPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(challengegroup.EpisodeRoundPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ChallengeGroup]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*EpisodeRound](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "episode_round" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cgq *ChallengeGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cgq.querySpec()
	_spec.Node.Columns = cgq.ctx.Fields
	if len(cgq.ctx.Fields) > 0 {
		_spec.Unique = cgq.ctx.Unique != nil && *cgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cgq.driver, _spec)
}

func (cgq *ChallengeGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(challengegroup.Table, challengegroup.Columns, sqlgraph.NewFieldSpec(challengegroup.FieldID, field.TypeInt))
	_spec.From = cgq.sql
	if unique := cgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cgq.path != nil {
		_spec.Unique = true
	}
	if fields := cgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, challengegroup.FieldID)
		for i := range fields {
			if fields[i] != challengegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cgq *ChallengeGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cgq.driver.Dialect())
	t1 := builder.Table(challengegroup.Table)
	columns := cgq.ctx.Fields
	if len(columns) == 0 {
		columns = challengegroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cgq.sql != nil {
		selector = cgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cgq.ctx.Unique != nil && *cgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cgq.predicates {
		p(selector)
	}
	for _, p := range cgq.order {
		p(selector)
	}
	if offset := cgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChallengeGroupGroupBy is the group-by builder for ChallengeGroup entities.
type ChallengeGroupGroupBy struct {
	selector
	build *ChallengeGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cggb *ChallengeGroupGroupBy) Aggregate(fns ...AggregateFunc) *ChallengeGroupGroupBy {
	cggb.fns = append(cggb.fns, fns...)
	return cggb
}

// Scan applies the selector query and scans the result into the given value.
func (cggb *ChallengeGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cggb.build.ctx, ent.OpQueryGroupBy)
	if err := cggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChallengeGroupQuery, *ChallengeGroupGroupBy](ctx, cggb.build, cggb, cggb.build.inters, v)
}

func (cggb *ChallengeGroupGroupBy) sqlScan(ctx context.Context, root *ChallengeGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cggb.fns))
	for _, fn := range cggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cggb.flds)+len(cggb.fns))
		for _, f := range *cggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChallengeGroupSelect is the builder for selecting fields of ChallengeGroup entities.
type ChallengeGroupSelect struct {
	*ChallengeGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cgs *ChallengeGroupSelect) Aggregate(fns ...AggregateFunc) *ChallengeGroupSelect {
	cgs.fns = append(cgs.fns, fns...)
	return cgs
}

// Scan applies the selector query and scans the result into the given value.
func (cgs *ChallengeGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgs.ctx, ent.OpQuerySelect)
	if err := cgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChallengeGroupQuery, *ChallengeGroupSelect](ctx, cgs.ChallengeGroupQuery, cgs, cgs.inters, v)
}

func (cgs *ChallengeGroupSelect) sqlScan(ctx context.Context, root *ChallengeGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cgs.fns))
	for _, fn := range cgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
