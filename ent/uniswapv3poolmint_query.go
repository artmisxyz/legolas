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
	"github.com/artmisxyz/blockinspector/ent/event"
	"github.com/artmisxyz/blockinspector/ent/predicate"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3poolmint"
)

// UniswapV3PoolMintQuery is the builder for querying UniswapV3PoolMint entities.
type UniswapV3PoolMintQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UniswapV3PoolMint
	// eager-loading edges.
	withEvent *EventQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UniswapV3PoolMintQuery builder.
func (uvmq *UniswapV3PoolMintQuery) Where(ps ...predicate.UniswapV3PoolMint) *UniswapV3PoolMintQuery {
	uvmq.predicates = append(uvmq.predicates, ps...)
	return uvmq
}

// Limit adds a limit step to the query.
func (uvmq *UniswapV3PoolMintQuery) Limit(limit int) *UniswapV3PoolMintQuery {
	uvmq.limit = &limit
	return uvmq
}

// Offset adds an offset step to the query.
func (uvmq *UniswapV3PoolMintQuery) Offset(offset int) *UniswapV3PoolMintQuery {
	uvmq.offset = &offset
	return uvmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uvmq *UniswapV3PoolMintQuery) Unique(unique bool) *UniswapV3PoolMintQuery {
	uvmq.unique = &unique
	return uvmq
}

// Order adds an order step to the query.
func (uvmq *UniswapV3PoolMintQuery) Order(o ...OrderFunc) *UniswapV3PoolMintQuery {
	uvmq.order = append(uvmq.order, o...)
	return uvmq
}

// QueryEvent chains the current query on the "event" edge.
func (uvmq *UniswapV3PoolMintQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: uvmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uvmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uvmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(uniswapv3poolmint.Table, uniswapv3poolmint.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, uniswapv3poolmint.EventTable, uniswapv3poolmint.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(uvmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UniswapV3PoolMint entity from the query.
// Returns a *NotFoundError when no UniswapV3PoolMint was found.
func (uvmq *UniswapV3PoolMintQuery) First(ctx context.Context) (*UniswapV3PoolMint, error) {
	nodes, err := uvmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{uniswapv3poolmint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) FirstX(ctx context.Context) *UniswapV3PoolMint {
	node, err := uvmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UniswapV3PoolMint ID from the query.
// Returns a *NotFoundError when no UniswapV3PoolMint ID was found.
func (uvmq *UniswapV3PoolMintQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{uniswapv3poolmint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) FirstIDX(ctx context.Context) int {
	id, err := uvmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UniswapV3PoolMint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UniswapV3PoolMint entity is not found.
// Returns a *NotFoundError when no UniswapV3PoolMint entities are found.
func (uvmq *UniswapV3PoolMintQuery) Only(ctx context.Context) (*UniswapV3PoolMint, error) {
	nodes, err := uvmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{uniswapv3poolmint.Label}
	default:
		return nil, &NotSingularError{uniswapv3poolmint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) OnlyX(ctx context.Context) *UniswapV3PoolMint {
	node, err := uvmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UniswapV3PoolMint ID in the query.
// Returns a *NotSingularError when exactly one UniswapV3PoolMint ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uvmq *UniswapV3PoolMintQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = &NotSingularError{uniswapv3poolmint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) OnlyIDX(ctx context.Context) int {
	id, err := uvmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UniswapV3PoolMints.
func (uvmq *UniswapV3PoolMintQuery) All(ctx context.Context) ([]*UniswapV3PoolMint, error) {
	if err := uvmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uvmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) AllX(ctx context.Context) []*UniswapV3PoolMint {
	nodes, err := uvmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UniswapV3PoolMint IDs.
func (uvmq *UniswapV3PoolMintQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uvmq.Select(uniswapv3poolmint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) IDsX(ctx context.Context) []int {
	ids, err := uvmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uvmq *UniswapV3PoolMintQuery) Count(ctx context.Context) (int, error) {
	if err := uvmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uvmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) CountX(ctx context.Context) int {
	count, err := uvmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uvmq *UniswapV3PoolMintQuery) Exist(ctx context.Context) (bool, error) {
	if err := uvmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uvmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uvmq *UniswapV3PoolMintQuery) ExistX(ctx context.Context) bool {
	exist, err := uvmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UniswapV3PoolMintQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uvmq *UniswapV3PoolMintQuery) Clone() *UniswapV3PoolMintQuery {
	if uvmq == nil {
		return nil
	}
	return &UniswapV3PoolMintQuery{
		config:     uvmq.config,
		limit:      uvmq.limit,
		offset:     uvmq.offset,
		order:      append([]OrderFunc{}, uvmq.order...),
		predicates: append([]predicate.UniswapV3PoolMint{}, uvmq.predicates...),
		withEvent:  uvmq.withEvent.Clone(),
		// clone intermediate query.
		sql:  uvmq.sql.Clone(),
		path: uvmq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (uvmq *UniswapV3PoolMintQuery) WithEvent(opts ...func(*EventQuery)) *UniswapV3PoolMintQuery {
	query := &EventQuery{config: uvmq.config}
	for _, opt := range opts {
		opt(query)
	}
	uvmq.withEvent = query
	return uvmq
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
//	client.UniswapV3PoolMint.Query().
//		GroupBy(uniswapv3poolmint.FieldOwner).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uvmq *UniswapV3PoolMintQuery) GroupBy(field string, fields ...string) *UniswapV3PoolMintGroupBy {
	group := &UniswapV3PoolMintGroupBy{config: uvmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uvmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uvmq.sqlQuery(ctx), nil
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
//	client.UniswapV3PoolMint.Query().
//		Select(uniswapv3poolmint.FieldOwner).
//		Scan(ctx, &v)
//
func (uvmq *UniswapV3PoolMintQuery) Select(fields ...string) *UniswapV3PoolMintSelect {
	uvmq.fields = append(uvmq.fields, fields...)
	return &UniswapV3PoolMintSelect{UniswapV3PoolMintQuery: uvmq}
}

func (uvmq *UniswapV3PoolMintQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uvmq.fields {
		if !uniswapv3poolmint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uvmq.path != nil {
		prev, err := uvmq.path(ctx)
		if err != nil {
			return err
		}
		uvmq.sql = prev
	}
	return nil
}

func (uvmq *UniswapV3PoolMintQuery) sqlAll(ctx context.Context) ([]*UniswapV3PoolMint, error) {
	var (
		nodes       = []*UniswapV3PoolMint{}
		withFKs     = uvmq.withFKs
		_spec       = uvmq.querySpec()
		loadedTypes = [1]bool{
			uvmq.withEvent != nil,
		}
	)
	if uvmq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolmint.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UniswapV3PoolMint{config: uvmq.config}
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
	if err := sqlgraph.QueryNodes(ctx, uvmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uvmq.withEvent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UniswapV3PoolMint)
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

func (uvmq *UniswapV3PoolMintQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uvmq.querySpec()
	return sqlgraph.CountNodes(ctx, uvmq.driver, _spec)
}

func (uvmq *UniswapV3PoolMintQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uvmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uvmq *UniswapV3PoolMintQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3poolmint.Table,
			Columns: uniswapv3poolmint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolmint.FieldID,
			},
		},
		From:   uvmq.sql,
		Unique: true,
	}
	if unique := uvmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uvmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolmint.FieldID)
		for i := range fields {
			if fields[i] != uniswapv3poolmint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uvmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uvmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uvmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uvmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uvmq *UniswapV3PoolMintQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uvmq.driver.Dialect())
	t1 := builder.Table(uniswapv3poolmint.Table)
	columns := uvmq.fields
	if len(columns) == 0 {
		columns = uniswapv3poolmint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uvmq.sql != nil {
		selector = uvmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range uvmq.predicates {
		p(selector)
	}
	for _, p := range uvmq.order {
		p(selector)
	}
	if offset := uvmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uvmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UniswapV3PoolMintGroupBy is the group-by builder for UniswapV3PoolMint entities.
type UniswapV3PoolMintGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uvmgb *UniswapV3PoolMintGroupBy) Aggregate(fns ...AggregateFunc) *UniswapV3PoolMintGroupBy {
	uvmgb.fns = append(uvmgb.fns, fns...)
	return uvmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uvmgb *UniswapV3PoolMintGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uvmgb.path(ctx)
	if err != nil {
		return err
	}
	uvmgb.sql = query
	return uvmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uvmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uvmgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uvmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) StringsX(ctx context.Context) []string {
	v, err := uvmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) StringX(ctx context.Context) string {
	v, err := uvmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uvmgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uvmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) IntsX(ctx context.Context) []int {
	v, err := uvmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) IntX(ctx context.Context) int {
	v, err := uvmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvmgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uvmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uvmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uvmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uvmgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uvmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uvmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvmgb *UniswapV3PoolMintGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvmgb *UniswapV3PoolMintGroupBy) BoolX(ctx context.Context) bool {
	v, err := uvmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvmgb *UniswapV3PoolMintGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uvmgb.fields {
		if !uniswapv3poolmint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uvmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uvmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uvmgb *UniswapV3PoolMintGroupBy) sqlQuery() *sql.Selector {
	selector := uvmgb.sql.Select()
	aggregation := make([]string, 0, len(uvmgb.fns))
	for _, fn := range uvmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uvmgb.fields)+len(uvmgb.fns))
		for _, f := range uvmgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uvmgb.fields...)...)
}

// UniswapV3PoolMintSelect is the builder for selecting fields of UniswapV3PoolMint entities.
type UniswapV3PoolMintSelect struct {
	*UniswapV3PoolMintQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uvms *UniswapV3PoolMintSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uvms.prepareQuery(ctx); err != nil {
		return err
	}
	uvms.sql = uvms.UniswapV3PoolMintQuery.sqlQuery(ctx)
	return uvms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uvms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uvms.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uvms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) StringsX(ctx context.Context) []string {
	v, err := uvms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) StringX(ctx context.Context) string {
	v, err := uvms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uvms.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uvms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) IntsX(ctx context.Context) []int {
	v, err := uvms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) IntX(ctx context.Context) int {
	v, err := uvms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvms.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uvms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uvms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) Float64X(ctx context.Context) float64 {
	v, err := uvms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uvms.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolMintSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uvms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) BoolsX(ctx context.Context) []bool {
	v, err := uvms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uvms *UniswapV3PoolMintSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolmint.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolMintSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvms *UniswapV3PoolMintSelect) BoolX(ctx context.Context) bool {
	v, err := uvms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvms *UniswapV3PoolMintSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uvms.sql.Query()
	if err := uvms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
