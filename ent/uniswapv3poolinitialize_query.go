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
	"github.com/artmisxyz/legolas/ent/uniswapv3poolinitialize"
)

// UniswapV3PoolInitializeQuery is the builder for querying UniswapV3PoolInitialize entities.
type UniswapV3PoolInitializeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UniswapV3PoolInitialize
	// eager-loading edges.
	withEvent *EventQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UniswapV3PoolInitializeQuery builder.
func (uviq *UniswapV3PoolInitializeQuery) Where(ps ...predicate.UniswapV3PoolInitialize) *UniswapV3PoolInitializeQuery {
	uviq.predicates = append(uviq.predicates, ps...)
	return uviq
}

// Limit adds a limit step to the query.
func (uviq *UniswapV3PoolInitializeQuery) Limit(limit int) *UniswapV3PoolInitializeQuery {
	uviq.limit = &limit
	return uviq
}

// Offset adds an offset step to the query.
func (uviq *UniswapV3PoolInitializeQuery) Offset(offset int) *UniswapV3PoolInitializeQuery {
	uviq.offset = &offset
	return uviq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uviq *UniswapV3PoolInitializeQuery) Unique(unique bool) *UniswapV3PoolInitializeQuery {
	uviq.unique = &unique
	return uviq
}

// Order adds an order step to the query.
func (uviq *UniswapV3PoolInitializeQuery) Order(o ...OrderFunc) *UniswapV3PoolInitializeQuery {
	uviq.order = append(uviq.order, o...)
	return uviq
}

// QueryEvent chains the current query on the "event" edge.
func (uviq *UniswapV3PoolInitializeQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: uviq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uviq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uviq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(uniswapv3poolinitialize.Table, uniswapv3poolinitialize.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, uniswapv3poolinitialize.EventTable, uniswapv3poolinitialize.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(uviq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UniswapV3PoolInitialize entity from the query.
// Returns a *NotFoundError when no UniswapV3PoolInitialize was found.
func (uviq *UniswapV3PoolInitializeQuery) First(ctx context.Context) (*UniswapV3PoolInitialize, error) {
	nodes, err := uviq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{uniswapv3poolinitialize.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) FirstX(ctx context.Context) *UniswapV3PoolInitialize {
	node, err := uviq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UniswapV3PoolInitialize ID from the query.
// Returns a *NotFoundError when no UniswapV3PoolInitialize ID was found.
func (uviq *UniswapV3PoolInitializeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uviq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{uniswapv3poolinitialize.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) FirstIDX(ctx context.Context) int {
	id, err := uviq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UniswapV3PoolInitialize entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UniswapV3PoolInitialize entity is not found.
// Returns a *NotFoundError when no UniswapV3PoolInitialize entities are found.
func (uviq *UniswapV3PoolInitializeQuery) Only(ctx context.Context) (*UniswapV3PoolInitialize, error) {
	nodes, err := uviq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		return nil, &NotSingularError{uniswapv3poolinitialize.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) OnlyX(ctx context.Context) *UniswapV3PoolInitialize {
	node, err := uviq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UniswapV3PoolInitialize ID in the query.
// Returns a *NotSingularError when exactly one UniswapV3PoolInitialize ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uviq *UniswapV3PoolInitializeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uviq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = &NotSingularError{uniswapv3poolinitialize.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) OnlyIDX(ctx context.Context) int {
	id, err := uviq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UniswapV3PoolInitializes.
func (uviq *UniswapV3PoolInitializeQuery) All(ctx context.Context) ([]*UniswapV3PoolInitialize, error) {
	if err := uviq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uviq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) AllX(ctx context.Context) []*UniswapV3PoolInitialize {
	nodes, err := uviq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UniswapV3PoolInitialize IDs.
func (uviq *UniswapV3PoolInitializeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uviq.Select(uniswapv3poolinitialize.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) IDsX(ctx context.Context) []int {
	ids, err := uviq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uviq *UniswapV3PoolInitializeQuery) Count(ctx context.Context) (int, error) {
	if err := uviq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uviq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) CountX(ctx context.Context) int {
	count, err := uviq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uviq *UniswapV3PoolInitializeQuery) Exist(ctx context.Context) (bool, error) {
	if err := uviq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uviq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uviq *UniswapV3PoolInitializeQuery) ExistX(ctx context.Context) bool {
	exist, err := uviq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UniswapV3PoolInitializeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uviq *UniswapV3PoolInitializeQuery) Clone() *UniswapV3PoolInitializeQuery {
	if uviq == nil {
		return nil
	}
	return &UniswapV3PoolInitializeQuery{
		config:     uviq.config,
		limit:      uviq.limit,
		offset:     uviq.offset,
		order:      append([]OrderFunc{}, uviq.order...),
		predicates: append([]predicate.UniswapV3PoolInitialize{}, uviq.predicates...),
		withEvent:  uviq.withEvent.Clone(),
		// clone intermediate query.
		sql:  uviq.sql.Clone(),
		path: uviq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (uviq *UniswapV3PoolInitializeQuery) WithEvent(opts ...func(*EventQuery)) *UniswapV3PoolInitializeQuery {
	query := &EventQuery{config: uviq.config}
	for _, opt := range opts {
		opt(query)
	}
	uviq.withEvent = query
	return uviq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SqrtPriceX96 *schema.BigInt `json:"sqrt_price_x96,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UniswapV3PoolInitialize.Query().
//		GroupBy(uniswapv3poolinitialize.FieldSqrtPriceX96).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uviq *UniswapV3PoolInitializeQuery) GroupBy(field string, fields ...string) *UniswapV3PoolInitializeGroupBy {
	group := &UniswapV3PoolInitializeGroupBy{config: uviq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uviq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uviq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SqrtPriceX96 *schema.BigInt `json:"sqrt_price_x96,omitempty"`
//	}
//
//	client.UniswapV3PoolInitialize.Query().
//		Select(uniswapv3poolinitialize.FieldSqrtPriceX96).
//		Scan(ctx, &v)
//
func (uviq *UniswapV3PoolInitializeQuery) Select(fields ...string) *UniswapV3PoolInitializeSelect {
	uviq.fields = append(uviq.fields, fields...)
	return &UniswapV3PoolInitializeSelect{UniswapV3PoolInitializeQuery: uviq}
}

func (uviq *UniswapV3PoolInitializeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uviq.fields {
		if !uniswapv3poolinitialize.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uviq.path != nil {
		prev, err := uviq.path(ctx)
		if err != nil {
			return err
		}
		uviq.sql = prev
	}
	return nil
}

func (uviq *UniswapV3PoolInitializeQuery) sqlAll(ctx context.Context) ([]*UniswapV3PoolInitialize, error) {
	var (
		nodes       = []*UniswapV3PoolInitialize{}
		withFKs     = uviq.withFKs
		_spec       = uviq.querySpec()
		loadedTypes = [1]bool{
			uviq.withEvent != nil,
		}
	)
	if uviq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolinitialize.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UniswapV3PoolInitialize{config: uviq.config}
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
	if err := sqlgraph.QueryNodes(ctx, uviq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uviq.withEvent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UniswapV3PoolInitialize)
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

func (uviq *UniswapV3PoolInitializeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uviq.querySpec()
	return sqlgraph.CountNodes(ctx, uviq.driver, _spec)
}

func (uviq *UniswapV3PoolInitializeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uviq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uviq *UniswapV3PoolInitializeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3poolinitialize.Table,
			Columns: uniswapv3poolinitialize.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolinitialize.FieldID,
			},
		},
		From:   uviq.sql,
		Unique: true,
	}
	if unique := uviq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uviq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolinitialize.FieldID)
		for i := range fields {
			if fields[i] != uniswapv3poolinitialize.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uviq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uviq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uviq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uviq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uviq *UniswapV3PoolInitializeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uviq.driver.Dialect())
	t1 := builder.Table(uniswapv3poolinitialize.Table)
	columns := uviq.fields
	if len(columns) == 0 {
		columns = uniswapv3poolinitialize.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uviq.sql != nil {
		selector = uviq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range uviq.predicates {
		p(selector)
	}
	for _, p := range uviq.order {
		p(selector)
	}
	if offset := uviq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uviq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UniswapV3PoolInitializeGroupBy is the group-by builder for UniswapV3PoolInitialize entities.
type UniswapV3PoolInitializeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uvigb *UniswapV3PoolInitializeGroupBy) Aggregate(fns ...AggregateFunc) *UniswapV3PoolInitializeGroupBy {
	uvigb.fns = append(uvigb.fns, fns...)
	return uvigb
}

// Scan applies the group-by query and scans the result into the given value.
func (uvigb *UniswapV3PoolInitializeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uvigb.path(ctx)
	if err != nil {
		return err
	}
	uvigb.sql = query
	return uvigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uvigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uvigb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uvigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) StringsX(ctx context.Context) []string {
	v, err := uvigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) StringX(ctx context.Context) string {
	v, err := uvigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uvigb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uvigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) IntsX(ctx context.Context) []int {
	v, err := uvigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) IntX(ctx context.Context) int {
	v, err := uvigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvigb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uvigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uvigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uvigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uvigb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uvigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uvigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvigb *UniswapV3PoolInitializeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvigb *UniswapV3PoolInitializeGroupBy) BoolX(ctx context.Context) bool {
	v, err := uvigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvigb *UniswapV3PoolInitializeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uvigb.fields {
		if !uniswapv3poolinitialize.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uvigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uvigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uvigb *UniswapV3PoolInitializeGroupBy) sqlQuery() *sql.Selector {
	selector := uvigb.sql.Select()
	aggregation := make([]string, 0, len(uvigb.fns))
	for _, fn := range uvigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uvigb.fields)+len(uvigb.fns))
		for _, f := range uvigb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uvigb.fields...)...)
}

// UniswapV3PoolInitializeSelect is the builder for selecting fields of UniswapV3PoolInitialize entities.
type UniswapV3PoolInitializeSelect struct {
	*UniswapV3PoolInitializeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uvis *UniswapV3PoolInitializeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uvis.prepareQuery(ctx); err != nil {
		return err
	}
	uvis.sql = uvis.UniswapV3PoolInitializeQuery.sqlQuery(ctx)
	return uvis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uvis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uvis.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uvis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) StringsX(ctx context.Context) []string {
	v, err := uvis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) StringX(ctx context.Context) string {
	v, err := uvis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uvis.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uvis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) IntsX(ctx context.Context) []int {
	v, err := uvis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) IntX(ctx context.Context) int {
	v, err := uvis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvis.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uvis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uvis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) Float64X(ctx context.Context) float64 {
	v, err := uvis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uvis.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolInitializeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uvis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) BoolsX(ctx context.Context) []bool {
	v, err := uvis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uvis *UniswapV3PoolInitializeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolInitializeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvis *UniswapV3PoolInitializeSelect) BoolX(ctx context.Context) bool {
	v, err := uvis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvis *UniswapV3PoolInitializeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uvis.sql.Query()
	if err := uvis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
