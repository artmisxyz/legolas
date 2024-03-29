// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/predicate"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolburn"
)

// UniswapV3PoolBurnQuery is the builder for querying UniswapV3PoolBurn entities.
type UniswapV3PoolBurnQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UniswapV3PoolBurn
	// eager-loading edges.
	withEvent *EventQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UniswapV3PoolBurnQuery builder.
func (uvbq *UniswapV3PoolBurnQuery) Where(ps ...predicate.UniswapV3PoolBurn) *UniswapV3PoolBurnQuery {
	uvbq.predicates = append(uvbq.predicates, ps...)
	return uvbq
}

// Limit adds a limit step to the query.
func (uvbq *UniswapV3PoolBurnQuery) Limit(limit int) *UniswapV3PoolBurnQuery {
	uvbq.limit = &limit
	return uvbq
}

// Offset adds an offset step to the query.
func (uvbq *UniswapV3PoolBurnQuery) Offset(offset int) *UniswapV3PoolBurnQuery {
	uvbq.offset = &offset
	return uvbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uvbq *UniswapV3PoolBurnQuery) Unique(unique bool) *UniswapV3PoolBurnQuery {
	uvbq.unique = &unique
	return uvbq
}

// Order adds an order step to the query.
func (uvbq *UniswapV3PoolBurnQuery) Order(o ...OrderFunc) *UniswapV3PoolBurnQuery {
	uvbq.order = append(uvbq.order, o...)
	return uvbq
}

// QueryEvent chains the current query on the "event" edge.
func (uvbq *UniswapV3PoolBurnQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: uvbq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uvbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uvbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(uniswapv3poolburn.Table, uniswapv3poolburn.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, uniswapv3poolburn.EventTable, uniswapv3poolburn.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(uvbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UniswapV3PoolBurn entity from the query.
// Returns a *NotFoundError when no UniswapV3PoolBurn was found.
func (uvbq *UniswapV3PoolBurnQuery) First(ctx context.Context) (*UniswapV3PoolBurn, error) {
	nodes, err := uvbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{uniswapv3poolburn.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) FirstX(ctx context.Context) *UniswapV3PoolBurn {
	node, err := uvbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UniswapV3PoolBurn ID from the query.
// Returns a *NotFoundError when no UniswapV3PoolBurn ID was found.
func (uvbq *UniswapV3PoolBurnQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{uniswapv3poolburn.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) FirstIDX(ctx context.Context) int {
	id, err := uvbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UniswapV3PoolBurn entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UniswapV3PoolBurn entity is not found.
// Returns a *NotFoundError when no UniswapV3PoolBurn entities are found.
func (uvbq *UniswapV3PoolBurnQuery) Only(ctx context.Context) (*UniswapV3PoolBurn, error) {
	nodes, err := uvbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{uniswapv3poolburn.Label}
	default:
		return nil, &NotSingularError{uniswapv3poolburn.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) OnlyX(ctx context.Context) *UniswapV3PoolBurn {
	node, err := uvbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UniswapV3PoolBurn ID in the query.
// Returns a *NotSingularError when exactly one UniswapV3PoolBurn ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uvbq *UniswapV3PoolBurnQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = &NotSingularError{uniswapv3poolburn.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) OnlyIDX(ctx context.Context) int {
	id, err := uvbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UniswapV3PoolBurns.
func (uvbq *UniswapV3PoolBurnQuery) All(ctx context.Context) ([]*UniswapV3PoolBurn, error) {
	if err := uvbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uvbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) AllX(ctx context.Context) []*UniswapV3PoolBurn {
	nodes, err := uvbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UniswapV3PoolBurn IDs.
func (uvbq *UniswapV3PoolBurnQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uvbq.Select(uniswapv3poolburn.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) IDsX(ctx context.Context) []int {
	ids, err := uvbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uvbq *UniswapV3PoolBurnQuery) Count(ctx context.Context) (int, error) {
	if err := uvbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uvbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) CountX(ctx context.Context) int {
	count, err := uvbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uvbq *UniswapV3PoolBurnQuery) Exist(ctx context.Context) (bool, error) {
	if err := uvbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uvbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uvbq *UniswapV3PoolBurnQuery) ExistX(ctx context.Context) bool {
	exist, err := uvbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UniswapV3PoolBurnQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uvbq *UniswapV3PoolBurnQuery) Clone() *UniswapV3PoolBurnQuery {
	if uvbq == nil {
		return nil
	}
	return &UniswapV3PoolBurnQuery{
		config:     uvbq.config,
		limit:      uvbq.limit,
		offset:     uvbq.offset,
		order:      append([]OrderFunc{}, uvbq.order...),
		predicates: append([]predicate.UniswapV3PoolBurn{}, uvbq.predicates...),
		withEvent:  uvbq.withEvent.Clone(),
		// clone intermediate query.
		sql:  uvbq.sql.Clone(),
		path: uvbq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (uvbq *UniswapV3PoolBurnQuery) WithEvent(opts ...func(*EventQuery)) *UniswapV3PoolBurnQuery {
	query := &EventQuery{config: uvbq.config}
	for _, opt := range opts {
		opt(query)
	}
	uvbq.withEvent = query
	return uvbq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Owner string `json:"owner,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UniswapV3PoolBurn.Query().
//		GroupBy(uniswapv3poolburn.FieldOwner).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uvbq *UniswapV3PoolBurnQuery) GroupBy(field string, fields ...string) *UniswapV3PoolBurnGroupBy {
	group := &UniswapV3PoolBurnGroupBy{config: uvbq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uvbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uvbq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Owner string `json:"owner,omitempty"`
//	}
//
//	client.UniswapV3PoolBurn.Query().
//		Select(uniswapv3poolburn.FieldOwner).
//		Scan(ctx, &v)
//
func (uvbq *UniswapV3PoolBurnQuery) Select(fields ...string) *UniswapV3PoolBurnSelect {
	uvbq.fields = append(uvbq.fields, fields...)
	return &UniswapV3PoolBurnSelect{UniswapV3PoolBurnQuery: uvbq}
}

func (uvbq *UniswapV3PoolBurnQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uvbq.fields {
		if !uniswapv3poolburn.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uvbq.path != nil {
		prev, err := uvbq.path(ctx)
		if err != nil {
			return err
		}
		uvbq.sql = prev
	}
	return nil
}

func (uvbq *UniswapV3PoolBurnQuery) sqlAll(ctx context.Context) ([]*UniswapV3PoolBurn, error) {
	var (
		nodes       = []*UniswapV3PoolBurn{}
		withFKs     = uvbq.withFKs
		_spec       = uvbq.querySpec()
		loadedTypes = [1]bool{
			uvbq.withEvent != nil,
		}
	)
	if uvbq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolburn.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UniswapV3PoolBurn{config: uvbq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, uvbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uvbq.withEvent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UniswapV3PoolBurn)
		for i := range nodes {
			if nodes[i].event_id == nil {
				continue
			}
			fk := *nodes[i].event_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(event.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "event_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Event = n
			}
		}
	}

	return nodes, nil
}

func (uvbq *UniswapV3PoolBurnQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uvbq.querySpec()
	return sqlgraph.CountNodes(ctx, uvbq.driver, _spec)
}

func (uvbq *UniswapV3PoolBurnQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uvbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uvbq *UniswapV3PoolBurnQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3poolburn.Table,
			Columns: uniswapv3poolburn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolburn.FieldID,
			},
		},
		From:   uvbq.sql,
		Unique: true,
	}
	if unique := uvbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uvbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolburn.FieldID)
		for i := range fields {
			if fields[i] != uniswapv3poolburn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uvbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uvbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uvbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uvbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uvbq *UniswapV3PoolBurnQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uvbq.driver.Dialect())
	t1 := builder.Table(uniswapv3poolburn.Table)
	columns := uvbq.fields
	if len(columns) == 0 {
		columns = uniswapv3poolburn.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uvbq.sql != nil {
		selector = uvbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range uvbq.predicates {
		p(selector)
	}
	for _, p := range uvbq.order {
		p(selector)
	}
	if offset := uvbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uvbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UniswapV3PoolBurnGroupBy is the group-by builder for UniswapV3PoolBurn entities.
type UniswapV3PoolBurnGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uvbgb *UniswapV3PoolBurnGroupBy) Aggregate(fns ...AggregateFunc) *UniswapV3PoolBurnGroupBy {
	uvbgb.fns = append(uvbgb.fns, fns...)
	return uvbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uvbgb *UniswapV3PoolBurnGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uvbgb.path(ctx)
	if err != nil {
		return err
	}
	uvbgb.sql = query
	return uvbgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uvbgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uvbgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uvbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) StringsX(ctx context.Context) []string {
	v, err := uvbgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvbgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) StringX(ctx context.Context) string {
	v, err := uvbgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uvbgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uvbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) IntsX(ctx context.Context) []int {
	v, err := uvbgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvbgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) IntX(ctx context.Context) int {
	v, err := uvbgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvbgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uvbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uvbgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvbgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uvbgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uvbgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uvbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uvbgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvbgb *UniswapV3PoolBurnGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvbgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvbgb *UniswapV3PoolBurnGroupBy) BoolX(ctx context.Context) bool {
	v, err := uvbgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvbgb *UniswapV3PoolBurnGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uvbgb.fields {
		if !uniswapv3poolburn.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uvbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uvbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uvbgb *UniswapV3PoolBurnGroupBy) sqlQuery() *sql.Selector {
	selector := uvbgb.sql.Select()
	aggregation := make([]string, 0, len(uvbgb.fns))
	for _, fn := range uvbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uvbgb.fields)+len(uvbgb.fns))
		for _, f := range uvbgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uvbgb.fields...)...)
}

// UniswapV3PoolBurnSelect is the builder for selecting fields of UniswapV3PoolBurn entities.
type UniswapV3PoolBurnSelect struct {
	*UniswapV3PoolBurnQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uvbs *UniswapV3PoolBurnSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uvbs.prepareQuery(ctx); err != nil {
		return err
	}
	uvbs.sql = uvbs.UniswapV3PoolBurnQuery.sqlQuery(ctx)
	return uvbs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uvbs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uvbs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uvbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) StringsX(ctx context.Context) []string {
	v, err := uvbs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvbs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) StringX(ctx context.Context) string {
	v, err := uvbs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uvbs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uvbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) IntsX(ctx context.Context) []int {
	v, err := uvbs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvbs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) IntX(ctx context.Context) int {
	v, err := uvbs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvbs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uvbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uvbs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvbs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) Float64X(ctx context.Context) float64 {
	v, err := uvbs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uvbs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolBurnSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uvbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) BoolsX(ctx context.Context) []bool {
	v, err := uvbs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uvbs *UniswapV3PoolBurnSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvbs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolburn.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolBurnSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvbs *UniswapV3PoolBurnSelect) BoolX(ctx context.Context) bool {
	v, err := uvbs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvbs *UniswapV3PoolBurnSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uvbs.sql.Query()
	if err := uvbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
