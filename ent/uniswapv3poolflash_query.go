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
	"github.com/artmisxyz/legolas/ent/uniswapv3poolflash"
)

// UniswapV3PoolFlashQuery is the builder for querying UniswapV3PoolFlash entities.
type UniswapV3PoolFlashQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UniswapV3PoolFlash
	// eager-loading edges.
	withEvent *EventQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UniswapV3PoolFlashQuery builder.
func (uvfq *UniswapV3PoolFlashQuery) Where(ps ...predicate.UniswapV3PoolFlash) *UniswapV3PoolFlashQuery {
	uvfq.predicates = append(uvfq.predicates, ps...)
	return uvfq
}

// Limit adds a limit step to the query.
func (uvfq *UniswapV3PoolFlashQuery) Limit(limit int) *UniswapV3PoolFlashQuery {
	uvfq.limit = &limit
	return uvfq
}

// Offset adds an offset step to the query.
func (uvfq *UniswapV3PoolFlashQuery) Offset(offset int) *UniswapV3PoolFlashQuery {
	uvfq.offset = &offset
	return uvfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uvfq *UniswapV3PoolFlashQuery) Unique(unique bool) *UniswapV3PoolFlashQuery {
	uvfq.unique = &unique
	return uvfq
}

// Order adds an order step to the query.
func (uvfq *UniswapV3PoolFlashQuery) Order(o ...OrderFunc) *UniswapV3PoolFlashQuery {
	uvfq.order = append(uvfq.order, o...)
	return uvfq
}

// QueryEvent chains the current query on the "event" edge.
func (uvfq *UniswapV3PoolFlashQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: uvfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uvfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uvfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(uniswapv3poolflash.Table, uniswapv3poolflash.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, uniswapv3poolflash.EventTable, uniswapv3poolflash.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(uvfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UniswapV3PoolFlash entity from the query.
// Returns a *NotFoundError when no UniswapV3PoolFlash was found.
func (uvfq *UniswapV3PoolFlashQuery) First(ctx context.Context) (*UniswapV3PoolFlash, error) {
	nodes, err := uvfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{uniswapv3poolflash.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) FirstX(ctx context.Context) *UniswapV3PoolFlash {
	node, err := uvfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UniswapV3PoolFlash ID from the query.
// Returns a *NotFoundError when no UniswapV3PoolFlash ID was found.
func (uvfq *UniswapV3PoolFlashQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{uniswapv3poolflash.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) FirstIDX(ctx context.Context) int {
	id, err := uvfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UniswapV3PoolFlash entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UniswapV3PoolFlash entity is not found.
// Returns a *NotFoundError when no UniswapV3PoolFlash entities are found.
func (uvfq *UniswapV3PoolFlashQuery) Only(ctx context.Context) (*UniswapV3PoolFlash, error) {
	nodes, err := uvfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{uniswapv3poolflash.Label}
	default:
		return nil, &NotSingularError{uniswapv3poolflash.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) OnlyX(ctx context.Context) *UniswapV3PoolFlash {
	node, err := uvfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UniswapV3PoolFlash ID in the query.
// Returns a *NotSingularError when exactly one UniswapV3PoolFlash ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uvfq *UniswapV3PoolFlashQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uvfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = &NotSingularError{uniswapv3poolflash.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) OnlyIDX(ctx context.Context) int {
	id, err := uvfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UniswapV3PoolFlashes.
func (uvfq *UniswapV3PoolFlashQuery) All(ctx context.Context) ([]*UniswapV3PoolFlash, error) {
	if err := uvfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uvfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) AllX(ctx context.Context) []*UniswapV3PoolFlash {
	nodes, err := uvfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UniswapV3PoolFlash IDs.
func (uvfq *UniswapV3PoolFlashQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uvfq.Select(uniswapv3poolflash.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) IDsX(ctx context.Context) []int {
	ids, err := uvfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uvfq *UniswapV3PoolFlashQuery) Count(ctx context.Context) (int, error) {
	if err := uvfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uvfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) CountX(ctx context.Context) int {
	count, err := uvfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uvfq *UniswapV3PoolFlashQuery) Exist(ctx context.Context) (bool, error) {
	if err := uvfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uvfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uvfq *UniswapV3PoolFlashQuery) ExistX(ctx context.Context) bool {
	exist, err := uvfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UniswapV3PoolFlashQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uvfq *UniswapV3PoolFlashQuery) Clone() *UniswapV3PoolFlashQuery {
	if uvfq == nil {
		return nil
	}
	return &UniswapV3PoolFlashQuery{
		config:     uvfq.config,
		limit:      uvfq.limit,
		offset:     uvfq.offset,
		order:      append([]OrderFunc{}, uvfq.order...),
		predicates: append([]predicate.UniswapV3PoolFlash{}, uvfq.predicates...),
		withEvent:  uvfq.withEvent.Clone(),
		// clone intermediate query.
		sql:  uvfq.sql.Clone(),
		path: uvfq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (uvfq *UniswapV3PoolFlashQuery) WithEvent(opts ...func(*EventQuery)) *UniswapV3PoolFlashQuery {
	query := &EventQuery{config: uvfq.config}
	for _, opt := range opts {
		opt(query)
	}
	uvfq.withEvent = query
	return uvfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Sender string `json:"sender,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UniswapV3PoolFlash.Query().
//		GroupBy(uniswapv3poolflash.FieldSender).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uvfq *UniswapV3PoolFlashQuery) GroupBy(field string, fields ...string) *UniswapV3PoolFlashGroupBy {
	group := &UniswapV3PoolFlashGroupBy{config: uvfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uvfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uvfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Sender string `json:"sender,omitempty"`
//	}
//
//	client.UniswapV3PoolFlash.Query().
//		Select(uniswapv3poolflash.FieldSender).
//		Scan(ctx, &v)
//
func (uvfq *UniswapV3PoolFlashQuery) Select(fields ...string) *UniswapV3PoolFlashSelect {
	uvfq.fields = append(uvfq.fields, fields...)
	return &UniswapV3PoolFlashSelect{UniswapV3PoolFlashQuery: uvfq}
}

func (uvfq *UniswapV3PoolFlashQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uvfq.fields {
		if !uniswapv3poolflash.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uvfq.path != nil {
		prev, err := uvfq.path(ctx)
		if err != nil {
			return err
		}
		uvfq.sql = prev
	}
	return nil
}

func (uvfq *UniswapV3PoolFlashQuery) sqlAll(ctx context.Context) ([]*UniswapV3PoolFlash, error) {
	var (
		nodes       = []*UniswapV3PoolFlash{}
		withFKs     = uvfq.withFKs
		_spec       = uvfq.querySpec()
		loadedTypes = [1]bool{
			uvfq.withEvent != nil,
		}
	)
	if uvfq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolflash.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UniswapV3PoolFlash{config: uvfq.config}
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
	if err := sqlgraph.QueryNodes(ctx, uvfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uvfq.withEvent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UniswapV3PoolFlash)
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

func (uvfq *UniswapV3PoolFlashQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uvfq.querySpec()
	return sqlgraph.CountNodes(ctx, uvfq.driver, _spec)
}

func (uvfq *UniswapV3PoolFlashQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uvfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uvfq *UniswapV3PoolFlashQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3poolflash.Table,
			Columns: uniswapv3poolflash.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolflash.FieldID,
			},
		},
		From:   uvfq.sql,
		Unique: true,
	}
	if unique := uvfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uvfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolflash.FieldID)
		for i := range fields {
			if fields[i] != uniswapv3poolflash.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uvfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uvfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uvfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uvfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uvfq *UniswapV3PoolFlashQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uvfq.driver.Dialect())
	t1 := builder.Table(uniswapv3poolflash.Table)
	columns := uvfq.fields
	if len(columns) == 0 {
		columns = uniswapv3poolflash.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uvfq.sql != nil {
		selector = uvfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range uvfq.predicates {
		p(selector)
	}
	for _, p := range uvfq.order {
		p(selector)
	}
	if offset := uvfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uvfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UniswapV3PoolFlashGroupBy is the group-by builder for UniswapV3PoolFlash entities.
type UniswapV3PoolFlashGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uvfgb *UniswapV3PoolFlashGroupBy) Aggregate(fns ...AggregateFunc) *UniswapV3PoolFlashGroupBy {
	uvfgb.fns = append(uvfgb.fns, fns...)
	return uvfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uvfgb *UniswapV3PoolFlashGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uvfgb.path(ctx)
	if err != nil {
		return err
	}
	uvfgb.sql = query
	return uvfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uvfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uvfgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uvfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) StringsX(ctx context.Context) []string {
	v, err := uvfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) StringX(ctx context.Context) string {
	v, err := uvfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uvfgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uvfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) IntsX(ctx context.Context) []int {
	v, err := uvfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) IntX(ctx context.Context) int {
	v, err := uvfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvfgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uvfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uvfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uvfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uvfgb.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uvfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uvfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uvfgb *UniswapV3PoolFlashGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvfgb *UniswapV3PoolFlashGroupBy) BoolX(ctx context.Context) bool {
	v, err := uvfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvfgb *UniswapV3PoolFlashGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uvfgb.fields {
		if !uniswapv3poolflash.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uvfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uvfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uvfgb *UniswapV3PoolFlashGroupBy) sqlQuery() *sql.Selector {
	selector := uvfgb.sql.Select()
	aggregation := make([]string, 0, len(uvfgb.fns))
	for _, fn := range uvfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uvfgb.fields)+len(uvfgb.fns))
		for _, f := range uvfgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uvfgb.fields...)...)
}

// UniswapV3PoolFlashSelect is the builder for selecting fields of UniswapV3PoolFlash entities.
type UniswapV3PoolFlashSelect struct {
	*UniswapV3PoolFlashQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uvfs *UniswapV3PoolFlashSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uvfs.prepareQuery(ctx); err != nil {
		return err
	}
	uvfs.sql = uvfs.UniswapV3PoolFlashQuery.sqlQuery(ctx)
	return uvfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uvfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uvfs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uvfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) StringsX(ctx context.Context) []string {
	v, err := uvfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uvfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) StringX(ctx context.Context) string {
	v, err := uvfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uvfs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uvfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) IntsX(ctx context.Context) []int {
	v, err := uvfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uvfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) IntX(ctx context.Context) int {
	v, err := uvfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uvfs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uvfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uvfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uvfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) Float64X(ctx context.Context) float64 {
	v, err := uvfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uvfs.fields) > 1 {
		return nil, errors.New("ent: UniswapV3PoolFlashSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uvfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) BoolsX(ctx context.Context) []bool {
	v, err := uvfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uvfs *UniswapV3PoolFlashSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uvfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{uniswapv3poolflash.Label}
	default:
		err = fmt.Errorf("ent: UniswapV3PoolFlashSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uvfs *UniswapV3PoolFlashSelect) BoolX(ctx context.Context) bool {
	v, err := uvfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uvfs *UniswapV3PoolFlashSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uvfs.sql.Query()
	if err := uvfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
