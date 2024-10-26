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
	"github.com/kevindamm/q-party/ent/challenge"
	"github.com/kevindamm/q-party/ent/challengegroup"
	"github.com/kevindamm/q-party/ent/predicate"
)

// ChallengeQuery is the builder for querying Challenge entities.
type ChallengeQuery struct {
	config
	ctx        *QueryContext
	order      []challenge.OrderOption
	inters     []Interceptor
	predicates []predicate.Challenge
	withColumn *ChallengeGroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChallengeQuery builder.
func (cq *ChallengeQuery) Where(ps ...predicate.Challenge) *ChallengeQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ChallengeQuery) Limit(limit int) *ChallengeQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ChallengeQuery) Offset(offset int) *ChallengeQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ChallengeQuery) Unique(unique bool) *ChallengeQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ChallengeQuery) Order(o ...challenge.OrderOption) *ChallengeQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryColumn chains the current query on the "column" edge.
func (cq *ChallengeQuery) QueryColumn() *ChallengeGroupQuery {
	query := (&ChallengeGroupClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(challenge.Table, challenge.FieldID, selector),
			sqlgraph.To(challengegroup.Table, challengegroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, challenge.ColumnTable, challenge.ColumnPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Challenge entity from the query.
// Returns a *NotFoundError when no Challenge was found.
func (cq *ChallengeQuery) First(ctx context.Context) (*Challenge, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{challenge.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChallengeQuery) FirstX(ctx context.Context) *Challenge {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Challenge ID from the query.
// Returns a *NotFoundError when no Challenge ID was found.
func (cq *ChallengeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{challenge.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ChallengeQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Challenge entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Challenge entity is found.
// Returns a *NotFoundError when no Challenge entities are found.
func (cq *ChallengeQuery) Only(ctx context.Context) (*Challenge, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{challenge.Label}
	default:
		return nil, &NotSingularError{challenge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChallengeQuery) OnlyX(ctx context.Context) *Challenge {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Challenge ID in the query.
// Returns a *NotSingularError when more than one Challenge ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ChallengeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{challenge.Label}
	default:
		err = &NotSingularError{challenge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChallengeQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Challenges.
func (cq *ChallengeQuery) All(ctx context.Context) ([]*Challenge, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Challenge, *ChallengeQuery]()
	return withInterceptors[[]*Challenge](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChallengeQuery) AllX(ctx context.Context) []*Challenge {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Challenge IDs.
func (cq *ChallengeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(challenge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChallengeQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChallengeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ChallengeQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChallengeQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChallengeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChallengeQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChallengeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChallengeQuery) Clone() *ChallengeQuery {
	if cq == nil {
		return nil
	}
	return &ChallengeQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]challenge.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Challenge{}, cq.predicates...),
		withColumn: cq.withColumn.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithColumn tells the query-builder to eager-load the nodes that are connected to
// the "column" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChallengeQuery) WithColumn(opts ...func(*ChallengeGroupQuery)) *ChallengeQuery {
	query := (&ChallengeGroupClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withColumn = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Media string `json:"media,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Challenge.Query().
//		GroupBy(challenge.FieldMedia).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ChallengeQuery) GroupBy(field string, fields ...string) *ChallengeGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChallengeGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = challenge.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Media string `json:"media,omitempty"`
//	}
//
//	client.Challenge.Query().
//		Select(challenge.FieldMedia).
//		Scan(ctx, &v)
func (cq *ChallengeQuery) Select(fields ...string) *ChallengeSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ChallengeSelect{ChallengeQuery: cq}
	sbuild.label = challenge.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChallengeSelect configured with the given aggregations.
func (cq *ChallengeQuery) Aggregate(fns ...AggregateFunc) *ChallengeSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ChallengeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !challenge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ChallengeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Challenge, error) {
	var (
		nodes       = []*Challenge{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withColumn != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Challenge).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Challenge{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withColumn; query != nil {
		if err := cq.loadColumn(ctx, query, nodes,
			func(n *Challenge) { n.Edges.Column = []*ChallengeGroup{} },
			func(n *Challenge, e *ChallengeGroup) { n.Edges.Column = append(n.Edges.Column, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ChallengeQuery) loadColumn(ctx context.Context, query *ChallengeGroupQuery, nodes []*Challenge, init func(*Challenge), assign func(*Challenge, *ChallengeGroup)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Challenge)
	nids := make(map[int]map[*Challenge]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(challenge.ColumnTable)
		s.Join(joinT).On(s.C(challengegroup.FieldID), joinT.C(challenge.ColumnPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(challenge.ColumnPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(challenge.ColumnPrimaryKey[1]))
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
					nids[inValue] = map[*Challenge]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ChallengeGroup](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "column" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *ChallengeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChallengeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(challenge.Table, challenge.Columns, sqlgraph.NewFieldSpec(challenge.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, challenge.FieldID)
		for i := range fields {
			if fields[i] != challenge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ChallengeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(challenge.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = challenge.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChallengeGroupBy is the group-by builder for Challenge entities.
type ChallengeGroupBy struct {
	selector
	build *ChallengeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChallengeGroupBy) Aggregate(fns ...AggregateFunc) *ChallengeGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ChallengeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChallengeQuery, *ChallengeGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ChallengeGroupBy) sqlScan(ctx context.Context, root *ChallengeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChallengeSelect is the builder for selecting fields of Challenge entities.
type ChallengeSelect struct {
	*ChallengeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ChallengeSelect) Aggregate(fns ...AggregateFunc) *ChallengeSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ChallengeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChallengeQuery, *ChallengeSelect](ctx, cs.ChallengeQuery, cs, cs.inters, v)
}

func (cs *ChallengeSelect) sqlScan(ctx context.Context, root *ChallengeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
