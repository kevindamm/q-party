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
	"github.com/kevindamm/q-party/ent/episode"
	"github.com/kevindamm/q-party/ent/episoderound"
	"github.com/kevindamm/q-party/ent/predicate"
)

// EpisodeQuery is the builder for querying Episode entities.
type EpisodeQuery struct {
	config
	ctx        *QueryContext
	order      []episode.OrderOption
	inters     []Interceptor
	predicates []predicate.Episode
	withRounds *EpisodeRoundQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EpisodeQuery builder.
func (eq *EpisodeQuery) Where(ps ...predicate.Episode) *EpisodeQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EpisodeQuery) Limit(limit int) *EpisodeQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EpisodeQuery) Offset(offset int) *EpisodeQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EpisodeQuery) Unique(unique bool) *EpisodeQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EpisodeQuery) Order(o ...episode.OrderOption) *EpisodeQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryRounds chains the current query on the "rounds" edge.
func (eq *EpisodeQuery) QueryRounds() *EpisodeRoundQuery {
	query := (&EpisodeRoundClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, selector),
			sqlgraph.To(episoderound.Table, episoderound.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, episode.RoundsTable, episode.RoundsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Episode entity from the query.
// Returns a *NotFoundError when no Episode was found.
func (eq *EpisodeQuery) First(ctx context.Context) (*Episode, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{episode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EpisodeQuery) FirstX(ctx context.Context) *Episode {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Episode ID from the query.
// Returns a *NotFoundError when no Episode ID was found.
func (eq *EpisodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{episode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EpisodeQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Episode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Episode entity is found.
// Returns a *NotFoundError when no Episode entities are found.
func (eq *EpisodeQuery) Only(ctx context.Context) (*Episode, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{episode.Label}
	default:
		return nil, &NotSingularError{episode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EpisodeQuery) OnlyX(ctx context.Context) *Episode {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Episode ID in the query.
// Returns a *NotSingularError when more than one Episode ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EpisodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{episode.Label}
	default:
		err = &NotSingularError{episode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EpisodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Episodes.
func (eq *EpisodeQuery) All(ctx context.Context) ([]*Episode, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryAll)
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Episode, *EpisodeQuery]()
	return withInterceptors[[]*Episode](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EpisodeQuery) AllX(ctx context.Context) []*Episode {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Episode IDs.
func (eq *EpisodeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryIDs)
	if err = eq.Select(episode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EpisodeQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EpisodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryCount)
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EpisodeQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EpisodeQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EpisodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, ent.OpQueryExist)
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EpisodeQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EpisodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EpisodeQuery) Clone() *EpisodeQuery {
	if eq == nil {
		return nil
	}
	return &EpisodeQuery{
		config:     eq.config,
		ctx:        eq.ctx.Clone(),
		order:      append([]episode.OrderOption{}, eq.order...),
		inters:     append([]Interceptor{}, eq.inters...),
		predicates: append([]predicate.Episode{}, eq.predicates...),
		withRounds: eq.withRounds.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithRounds tells the query-builder to eager-load the nodes that are connected to
// the "rounds" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EpisodeQuery) WithRounds(opts ...func(*EpisodeRoundQuery)) *EpisodeQuery {
	query := (&EpisodeRoundClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withRounds = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Aired time.Time `json:"aired,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Episode.Query().
//		GroupBy(episode.FieldAired).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EpisodeQuery) GroupBy(field string, fields ...string) *EpisodeGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EpisodeGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = episode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Aired time.Time `json:"aired,omitempty"`
//	}
//
//	client.Episode.Query().
//		Select(episode.FieldAired).
//		Scan(ctx, &v)
func (eq *EpisodeQuery) Select(fields ...string) *EpisodeSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EpisodeSelect{EpisodeQuery: eq}
	sbuild.label = episode.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EpisodeSelect configured with the given aggregations.
func (eq *EpisodeQuery) Aggregate(fns ...AggregateFunc) *EpisodeSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EpisodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !episode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EpisodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Episode, error) {
	var (
		nodes       = []*Episode{}
		_spec       = eq.querySpec()
		loadedTypes = [1]bool{
			eq.withRounds != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Episode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Episode{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withRounds; query != nil {
		if err := eq.loadRounds(ctx, query, nodes,
			func(n *Episode) { n.Edges.Rounds = []*EpisodeRound{} },
			func(n *Episode, e *EpisodeRound) { n.Edges.Rounds = append(n.Edges.Rounds, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EpisodeQuery) loadRounds(ctx context.Context, query *EpisodeRoundQuery, nodes []*Episode, init func(*Episode), assign func(*Episode, *EpisodeRound)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Episode)
	nids := make(map[int]map[*Episode]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(episode.RoundsTable)
		s.Join(joinT).On(s.C(episoderound.FieldID), joinT.C(episode.RoundsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(episode.RoundsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(episode.RoundsPrimaryKey[0]))
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
					nids[inValue] = map[*Episode]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "rounds" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (eq *EpisodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EpisodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(episode.Table, episode.Columns, sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, episode.FieldID)
		for i := range fields {
			if fields[i] != episode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EpisodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(episode.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = episode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EpisodeGroupBy is the group-by builder for Episode entities.
type EpisodeGroupBy struct {
	selector
	build *EpisodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EpisodeGroupBy) Aggregate(fns ...AggregateFunc) *EpisodeGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EpisodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, ent.OpQueryGroupBy)
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EpisodeQuery, *EpisodeGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EpisodeGroupBy) sqlScan(ctx context.Context, root *EpisodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EpisodeSelect is the builder for selecting fields of Episode entities.
type EpisodeSelect struct {
	*EpisodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EpisodeSelect) Aggregate(fns ...AggregateFunc) *EpisodeSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EpisodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, ent.OpQuerySelect)
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EpisodeQuery, *EpisodeSelect](ctx, es.EpisodeQuery, es, es.inters, v)
}

func (es *EpisodeSelect) sqlScan(ctx context.Context, root *EpisodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
