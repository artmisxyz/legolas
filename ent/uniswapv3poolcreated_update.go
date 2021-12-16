// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/blockinspector/ent/event"
	"github.com/artmisxyz/blockinspector/ent/predicate"
	"github.com/artmisxyz/blockinspector/ent/schema"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3poolcreated"
)

// UniswapV3PoolCreatedUpdate is the builder for updating UniswapV3PoolCreated entities.
type UniswapV3PoolCreatedUpdate struct {
	config
	hooks    []Hook
	mutation *UniswapV3PoolCreatedMutation
}

// Where appends a list predicates to the UniswapV3PoolCreatedUpdate builder.
func (uvcu *UniswapV3PoolCreatedUpdate) Where(ps ...predicate.UniswapV3PoolCreated) *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.Where(ps...)
	return uvcu
}

// SetToken0 sets the "token0" field.
func (uvcu *UniswapV3PoolCreatedUpdate) SetToken0(s string) *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.SetToken0(s)
	return uvcu
}

// SetToken1 sets the "token1" field.
func (uvcu *UniswapV3PoolCreatedUpdate) SetToken1(s string) *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.SetToken1(s)
	return uvcu
}

// SetFee sets the "fee" field.
func (uvcu *UniswapV3PoolCreatedUpdate) SetFee(si *schema.BigInt) *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.SetFee(si)
	return uvcu
}

// SetTickSpacing sets the "tick_spacing" field.
func (uvcu *UniswapV3PoolCreatedUpdate) SetTickSpacing(si *schema.BigInt) *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.SetTickSpacing(si)
	return uvcu
}

// SetPool sets the "pool" field.
func (uvcu *UniswapV3PoolCreatedUpdate) SetPool(s string) *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.SetPool(s)
	return uvcu
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvcu *UniswapV3PoolCreatedUpdate) SetEventID(id int) *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.SetEventID(id)
	return uvcu
}

// SetEvent sets the "event" edge to the Event entity.
func (uvcu *UniswapV3PoolCreatedUpdate) SetEvent(e *Event) *UniswapV3PoolCreatedUpdate {
	return uvcu.SetEventID(e.ID)
}

// Mutation returns the UniswapV3PoolCreatedMutation object of the builder.
func (uvcu *UniswapV3PoolCreatedUpdate) Mutation() *UniswapV3PoolCreatedMutation {
	return uvcu.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (uvcu *UniswapV3PoolCreatedUpdate) ClearEvent() *UniswapV3PoolCreatedUpdate {
	uvcu.mutation.ClearEvent()
	return uvcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uvcu *UniswapV3PoolCreatedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uvcu.hooks) == 0 {
		if err = uvcu.check(); err != nil {
			return 0, err
		}
		affected, err = uvcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3PoolCreatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvcu.check(); err != nil {
				return 0, err
			}
			uvcu.mutation = mutation
			affected, err = uvcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uvcu.hooks) - 1; i >= 0; i-- {
			if uvcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uvcu *UniswapV3PoolCreatedUpdate) SaveX(ctx context.Context) int {
	affected, err := uvcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uvcu *UniswapV3PoolCreatedUpdate) Exec(ctx context.Context) error {
	_, err := uvcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvcu *UniswapV3PoolCreatedUpdate) ExecX(ctx context.Context) {
	if err := uvcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvcu *UniswapV3PoolCreatedUpdate) check() error {
	if v, ok := uvcu.mutation.Token0(); ok {
		if err := uniswapv3poolcreated.Token0Validator(v); err != nil {
			return &ValidationError{Name: "token0", err: fmt.Errorf("ent: validator failed for field \"token0\": %w", err)}
		}
	}
	if v, ok := uvcu.mutation.Token1(); ok {
		if err := uniswapv3poolcreated.Token1Validator(v); err != nil {
			return &ValidationError{Name: "token1", err: fmt.Errorf("ent: validator failed for field \"token1\": %w", err)}
		}
	}
	if v, ok := uvcu.mutation.Pool(); ok {
		if err := uniswapv3poolcreated.PoolValidator(v); err != nil {
			return &ValidationError{Name: "pool", err: fmt.Errorf("ent: validator failed for field \"pool\": %w", err)}
		}
	}
	if _, ok := uvcu.mutation.EventID(); uvcu.mutation.EventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"event\"")
	}
	return nil
}

func (uvcu *UniswapV3PoolCreatedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3poolcreated.Table,
			Columns: uniswapv3poolcreated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolcreated.FieldID,
			},
		},
	}
	if ps := uvcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvcu.mutation.Token0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldToken0,
		})
	}
	if value, ok := uvcu.mutation.Token1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldToken1,
		})
	}
	if value, ok := uvcu.mutation.Fee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldFee,
		})
	}
	if value, ok := uvcu.mutation.TickSpacing(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldTickSpacing,
		})
	}
	if value, ok := uvcu.mutation.Pool(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldPool,
		})
	}
	if uvcu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3poolcreated.EventTable,
			Columns: []string{uniswapv3poolcreated.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uvcu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3poolcreated.EventTable,
			Columns: []string{uniswapv3poolcreated.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uvcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uniswapv3poolcreated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UniswapV3PoolCreatedUpdateOne is the builder for updating a single UniswapV3PoolCreated entity.
type UniswapV3PoolCreatedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UniswapV3PoolCreatedMutation
}

// SetToken0 sets the "token0" field.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SetToken0(s string) *UniswapV3PoolCreatedUpdateOne {
	uvcuo.mutation.SetToken0(s)
	return uvcuo
}

// SetToken1 sets the "token1" field.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SetToken1(s string) *UniswapV3PoolCreatedUpdateOne {
	uvcuo.mutation.SetToken1(s)
	return uvcuo
}

// SetFee sets the "fee" field.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SetFee(si *schema.BigInt) *UniswapV3PoolCreatedUpdateOne {
	uvcuo.mutation.SetFee(si)
	return uvcuo
}

// SetTickSpacing sets the "tick_spacing" field.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SetTickSpacing(si *schema.BigInt) *UniswapV3PoolCreatedUpdateOne {
	uvcuo.mutation.SetTickSpacing(si)
	return uvcuo
}

// SetPool sets the "pool" field.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SetPool(s string) *UniswapV3PoolCreatedUpdateOne {
	uvcuo.mutation.SetPool(s)
	return uvcuo
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SetEventID(id int) *UniswapV3PoolCreatedUpdateOne {
	uvcuo.mutation.SetEventID(id)
	return uvcuo
}

// SetEvent sets the "event" edge to the Event entity.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SetEvent(e *Event) *UniswapV3PoolCreatedUpdateOne {
	return uvcuo.SetEventID(e.ID)
}

// Mutation returns the UniswapV3PoolCreatedMutation object of the builder.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) Mutation() *UniswapV3PoolCreatedMutation {
	return uvcuo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) ClearEvent() *UniswapV3PoolCreatedUpdateOne {
	uvcuo.mutation.ClearEvent()
	return uvcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) Select(field string, fields ...string) *UniswapV3PoolCreatedUpdateOne {
	uvcuo.fields = append([]string{field}, fields...)
	return uvcuo
}

// Save executes the query and returns the updated UniswapV3PoolCreated entity.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) Save(ctx context.Context) (*UniswapV3PoolCreated, error) {
	var (
		err  error
		node *UniswapV3PoolCreated
	)
	if len(uvcuo.hooks) == 0 {
		if err = uvcuo.check(); err != nil {
			return nil, err
		}
		node, err = uvcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3PoolCreatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvcuo.check(); err != nil {
				return nil, err
			}
			uvcuo.mutation = mutation
			node, err = uvcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uvcuo.hooks) - 1; i >= 0; i-- {
			if uvcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) SaveX(ctx context.Context) *UniswapV3PoolCreated {
	node, err := uvcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) Exec(ctx context.Context) error {
	_, err := uvcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) ExecX(ctx context.Context) {
	if err := uvcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvcuo *UniswapV3PoolCreatedUpdateOne) check() error {
	if v, ok := uvcuo.mutation.Token0(); ok {
		if err := uniswapv3poolcreated.Token0Validator(v); err != nil {
			return &ValidationError{Name: "token0", err: fmt.Errorf("ent: validator failed for field \"token0\": %w", err)}
		}
	}
	if v, ok := uvcuo.mutation.Token1(); ok {
		if err := uniswapv3poolcreated.Token1Validator(v); err != nil {
			return &ValidationError{Name: "token1", err: fmt.Errorf("ent: validator failed for field \"token1\": %w", err)}
		}
	}
	if v, ok := uvcuo.mutation.Pool(); ok {
		if err := uniswapv3poolcreated.PoolValidator(v); err != nil {
			return &ValidationError{Name: "pool", err: fmt.Errorf("ent: validator failed for field \"pool\": %w", err)}
		}
	}
	if _, ok := uvcuo.mutation.EventID(); uvcuo.mutation.EventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"event\"")
	}
	return nil
}

func (uvcuo *UniswapV3PoolCreatedUpdateOne) sqlSave(ctx context.Context) (_node *UniswapV3PoolCreated, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3poolcreated.Table,
			Columns: uniswapv3poolcreated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolcreated.FieldID,
			},
		},
	}
	id, ok := uvcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UniswapV3PoolCreated.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uvcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3poolcreated.FieldID)
		for _, f := range fields {
			if !uniswapv3poolcreated.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != uniswapv3poolcreated.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uvcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvcuo.mutation.Token0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldToken0,
		})
	}
	if value, ok := uvcuo.mutation.Token1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldToken1,
		})
	}
	if value, ok := uvcuo.mutation.Fee(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldFee,
		})
	}
	if value, ok := uvcuo.mutation.TickSpacing(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldTickSpacing,
		})
	}
	if value, ok := uvcuo.mutation.Pool(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolcreated.FieldPool,
		})
	}
	if uvcuo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3poolcreated.EventTable,
			Columns: []string{uniswapv3poolcreated.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uvcuo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3poolcreated.EventTable,
			Columns: []string{uniswapv3poolcreated.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UniswapV3PoolCreated{config: uvcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uvcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uniswapv3poolcreated.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
