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
	"github.com/artmisxyz/blockinspector/ent/uniswapv3increaseliqudity"
)

// UniswapV3IncreaseLiqudityQuery is the builder for querying UniswapV3IncreaseLiqudity entities.
type UniswapV3IncreaseLiqudityQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UniswapV3IncreaseLiqudity
	// eager-loading edges.
	withEvent *EventQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UniswapV3IncreaseLiqudityQuery builder.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Where(ps ...predicate.UniswapV3IncreaseLiqudity) *UniswapV3IncreaseLiqudityQuery {
	uvlq.predicates = append(uvlq.predicates, ps...)
	return uvlq
}

// Limit adds a limit step to the query.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Limit(limit int) *UniswapV3IncreaseLiqudityQuery {
	uvlq.limit = &limit
	return uvlq
}

// Offset adds an offset step to the query.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Offset(offset int) *UniswapV3IncreaseLiqudityQuery {
	uvlq.offset = &offset
	return uvlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Unique(unique bool) *UniswapV3IncreaseLiqudityQuery {
	uvlq.unique = &unique
	return uvlq
}

// Order adds an order step to the query.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Order(o ...OrderFunc) *UniswapV3IncreaseLiqudityQuery {
	uvlq.order = append(uvlq.order, o...)
	return uvlq
}

// QueryEvent chains the current query on the "event" edge.
func (uvlq *UniswapV3IncreaseLiqudityQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: uvlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uvlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uvlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(uniswapv3increaseliqudity.Table, uniswapv3increaseliqudity.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, uniswapv3increaseliqudity.EventTable, uniswapv3increaseliqudity.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(uvlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UniswapV3IncreaseLiqudity entity from the query.
// Returns a *NotFoundError when no UniswapV3IncreaseLiqudity was found.
func (uvlq *UniswapV3IncreaseLiqudityQuery) First(ctx context.Context) (*UniswapV3IncreaseLiqudity, error) {
	nodes, err := uvlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{uniswapv3increaseliqudity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) FirstX(ctx context.Context) *UniswapV3IncreaseLiqudity {
	node, err := uvlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UniswapV3IncreaseLiqudity ID from the query.
// Returns a *NotFoundError when no UniswapV3IncreaseLiqudity ID was found.
func (uvlq *UniswapV3IncreaseLiqudityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) FirstIDX(ctx context.Context) int {
	id, err := uvlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UniswapV3IncreaseLiqudity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UniswapV3IncreaseLiqudity entity is not found.
// Returns a *NotFoundError when no UniswapV3IncreaseLiqudity entities are found.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Only(ctx context.Context) (*UniswapV3IncreaseLiqudity, error) {
	nodes, err := uvlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		return nil, &NotSingularError{uniswapv3increaseliqudity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) OnlyX(ctx context.Context) *UniswapV3IncreaseLiqudity {
	node, err := uvlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UniswapV3IncreaseLiqudity ID in the query.
// Returns a *NotSingularError when exactly one UniswapV3IncreaseLiqudity ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uvlq *UniswapV3IncreaseLiqudityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = &NotSingularError{uniswapv3increaseliqudity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) OnlyIDX(ctx context.Context) int {
	id, err := uvlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UniswapV3IncreaseLiqudities.
func (uvlq *UniswapV3IncreaseLiqudityQuery) All(ctx context.Context) ([]*UniswapV3IncreaseLiqudity, error) {
	if err := uvlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uvlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) AllX(ctx context.Context) []*UniswapV3IncreaseLiqudity {
	nodes, err := uvlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UniswapV3IncreaseLiqudity IDs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uvlq.Select(uniswapv3increaseliqudity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) IDsX(ctx context.Context) []int {
	ids, err := uvlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Count(ctx context.Context) (int, error) {
	if err := uvlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uvlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) CountX(ctx context.Context) int {
	count, err := uvlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Exist(ctx context.Context) (bool, error) {
	if err := uvlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uvlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uvlq *UniswapV3IncreaseLiqudityQuery) ExistX(ctx context.Context) bool {
	exist, err := uvlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UniswapV3IncreaseLiqudityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uvlq *UniswapV3IncreaseLiqudityQuery) Clone() *UniswapV3IncreaseLiqudityQuery {
	if uvlq == nil {
		return nil
	}
	return &UniswapV3IncreaseLiqudityQuery{
		config:     uvlq.config,
		limit:      uvlq.limit,
		offset:     uvlq.offset,
		order:      append([]OrderFunc{}, uvlq.order...),
		predicates: append([]predicate.UniswapV3IncreaseLiqudity{}, uvlq.predicates...),
		withEvent:  uvlq.withEvent.Clone(),
		// clone intermediate query.
		sql:  uvlq.sql.Clone(),
		path: uvlq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (uvlq *UniswapV3IncreaseLiqudityQuery) WithEvent(opts ...func(*EventQuery)) *UniswapV3IncreaseLiqudityQuery {
	query := &EventQuery{config: uvlq.config}
	for _, opt := range opts {
		opt(query)
	}
	uvlq.withEvent = query
	return uvlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TokenID *schema.BigInt `json:"token_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UniswapV3IncreaseLiqudity.Query().
//		GroupBy(uniswapv3increaseliqudity.FieldTokenID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uvlq *UniswapV3IncreaseLiqudityQuery) GroupBy(field string, fields ...string) *UniswapV3IncreaseLiqudityGroupBy {
	group := &UniswapV3IncreaseLiqudityGroupBy{config: uvlq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uvlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uvlq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TokenID *schema.BigInt `json:"token_id,omitempty"`
//	}
//
//	client.UniswapV3IncreaseLiqudity.Query().
//		Select(uniswapv3increaseliqudity.FieldTokenID).
//		Scan(ctx, &v)
//
func (uvlq *UniswapV3IncreaseLiqudityQuery) Select(fields ...string) *UniswapV3IncreaseLiquditySelect {
	uvlq.fields = append(uvlq.fields, fields...)
	return &UniswapV3IncreaseLiquditySelect{UniswapV3IncreaseLiqudityQuery: uvlq}
}

func (uvlq *UniswapV3IncreaseLiqudityQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uvlq.fields {
		if !uniswapv3increaseliqudity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uvlq.path != nil {
		prev, err := uvlq.path(ctx)
		if err != nil {
			return err
		}
		uvlq.sql = prev
	}
	return nil
}

func (uvlq *UniswapV3IncreaseLiqudityQuery) sqlAll(ctx context.Context) ([]*UniswapV3IncreaseLiqudity, error) {
	var (
		nodes       = []*UniswapV3IncreaseLiqudity{}
		withFKs     = uvlq.withFKs
		_spec       = uvlq.querySpec()
		loadedTypes = [1]bool{
			uvlq.withEvent != nil,
		}
	)
	if uvlq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3increaseliqudity.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UniswapV3IncreaseLiqudity{config: uvlq.config}
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
	if err := sqlgraph.QueryNodes(ctx, uvlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uvlq.withEvent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UniswapV3IncreaseLiqudity)
		for i := range nodes {
			if nodes[i].event_increase_liquidity == nil {
				continue
			}
			fk := *nodes[i].event_increase_liquidity
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
				return nil, fmt.Errorf(`unexpected foreign-key "event_increase_liquidity" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Event = n
			}
		}
	}

	return nodes, nil
}

func (uvlq *UniswapV3IncreaseLiqudityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uvlq.querySpec()
	return sqlgraph.CountNodes(ctx, uvlq.driver, _spec)
}

func (uvlq *UniswapV3IncreaseLiqudityQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uvlq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uvlq *UniswapV3IncreaseLiqudityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3increaseliqudity.Table,
			Columns: uniswapv3increaseliqudity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3increaseliqudity.FieldID,
			},
		},
		From:   uvlq.sql,
		Unique: true,
	}
	if unique := uvlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uvlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3increaseliqudity.FieldID)
		for i := range fields {
			if fields[i] != uniswapv3increaseliqudity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uvlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uvlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uvlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uvlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uvlq *UniswapV3IncreaseLiqudityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uvlq.driver.Dialect())
	t1 := builder.Table(uniswapv3increaseliqudity.Table)
	columns := uvlq.fields
	if len(columns) == 0 {
		columns = uniswapv3increaseliqudity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uvlq.sql != nil {
		selector = uvlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range uvlq.predicates {
		p(selector)
	}
	for _, p := range uvlq.order {
		p(selector)
	}
	if offset := uvlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uvlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UniswapV3IncreaseLiqudityGroupBy is the group-by builder for UniswapV3IncreaseLiqudity entities.
type UniswapV3IncreaseLiqudityGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Aggregate(fns ...AggregateFunc) *UniswapV3IncreaseLiqudityGroupBy {
	uvlgb.fns = append(uvlgb.fns, fns...)
	return uvlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uvlgb.path(ctx)
	if err != nil {
		return err
	}
	uvlgb.sql = query
	return uvlgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uvlgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uvlgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiqudityGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uvlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) StringsX(ctx context.Context) []string {
	v, err := uvlgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvlgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiqudityGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) StringX(ctx context.Context) string {
	v, err := uvlgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uvlgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiqudityGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uvlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) IntsX(ctx context.Context) []int {
	v, err := uvlgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvlgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiqudityGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) IntX(ctx context.Context) int {
	v, err := uvlgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvlgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiqudityGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uvlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uvlgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvlgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiqudityGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uvlgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uvlgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiqudityGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uvlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uvlgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvlgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiqudityGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) BoolX(ctx context.Context) bool {
	v, err := uvlgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uvlgb.fields {
		if !uniswapv3increaseliqudity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uvlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uvlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uvlgb *UniswapV3IncreaseLiqudityGroupBy) sqlQuery() *sql.Selector {
	selector := uvlgb.sql.Select()
	aggregation := make([]string, 0, len(uvlgb.fns))
	for _, fn := range uvlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uvlgb.fields)+len(uvlgb.fns))
		for _, f := range uvlgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uvlgb.fields...)...)
}

// UniswapV3IncreaseLiquditySelect is the builder for selecting fields of UniswapV3IncreaseLiqudity entities.
type UniswapV3IncreaseLiquditySelect struct {
	*UniswapV3IncreaseLiqudityQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uvls *UniswapV3IncreaseLiquditySelect) Scan(ctx context.Context, v interface{}) error {
	if err := uvls.prepareQuery(ctx); err != nil {
		return err
	}
	uvls.sql = uvls.UniswapV3IncreaseLiqudityQuery.sqlQuery(ctx)
	return uvls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) ScanX(ctx context.Context, v interface{}) {
	if err := uvls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) Strings(ctx context.Context) ([]string, error) {
	if len(uvls.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiquditySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uvls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) StringsX(ctx context.Context) []string {
	v, err := uvls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiquditySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) StringX(ctx context.Context) string {
	v, err := uvls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) Ints(ctx context.Context) ([]int, error) {
	if len(uvls.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiquditySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uvls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) IntsX(ctx context.Context) []int {
	v, err := uvls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiquditySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) IntX(ctx context.Context) int {
	v, err := uvls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvls.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiquditySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uvls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) Float64sX(ctx context.Context) []float64 {
	v, err := uvls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiquditySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) Float64X(ctx context.Context) float64 {
	v, err := uvls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uvls.fields) > 1 {
		return nil, errors.New("ent: UniswapV3IncreaseLiquditySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uvls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) BoolsX(ctx context.Context) []bool {
	v, err := uvls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uvls *UniswapV3IncreaseLiquditySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3increaseliqudity.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3IncreaseLiquditySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvls *UniswapV3IncreaseLiquditySelect) BoolX(ctx context.Context) bool {
	v, err := uvls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvls *UniswapV3IncreaseLiquditySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uvls.sql.Query()
	if err := uvls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
