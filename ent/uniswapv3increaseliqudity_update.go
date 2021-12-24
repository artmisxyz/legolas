// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/predicate"
	"github.com/artmisxyz/legolas/ent/uniswapv3increaseliqudity"
)

// UniswapV3IncreaseLiqudityUpdate is the builder for updating UniswapV3IncreaseLiqudity entities.
type UniswapV3IncreaseLiqudityUpdate struct {
	config
	hooks    []Hook
	mutation *UniswapV3IncreaseLiqudityMutation
}

// Where appends a list predicates to the UniswapV3IncreaseLiqudityUpdate builder.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) Where(ps ...predicate.UniswapV3IncreaseLiqudity) *UniswapV3IncreaseLiqudityUpdate {
	uvlu.mutation.Where(ps...)
	return uvlu
}

// SetTokenID sets the "token_id" field.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) SetTokenID(s string) *UniswapV3IncreaseLiqudityUpdate {
	uvlu.mutation.SetTokenID(s)
	return uvlu
}

// SetLiquidity sets the "liquidity" field.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) SetLiquidity(s string) *UniswapV3IncreaseLiqudityUpdate {
	uvlu.mutation.SetLiquidity(s)
	return uvlu
}

// SetAmount0 sets the "amount0" field.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) SetAmount0(s string) *UniswapV3IncreaseLiqudityUpdate {
	uvlu.mutation.SetAmount0(s)
	return uvlu
}

// SetAmount1 sets the "amount1" field.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) SetAmount1(s string) *UniswapV3IncreaseLiqudityUpdate {
	uvlu.mutation.SetAmount1(s)
	return uvlu
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) SetEventID(id int) *UniswapV3IncreaseLiqudityUpdate {
	uvlu.mutation.SetEventID(id)
	return uvlu
}

// SetEvent sets the "event" edge to the Event entity.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) SetEvent(e *Event) *UniswapV3IncreaseLiqudityUpdate {
	return uvlu.SetEventID(e.ID)
}

// Mutation returns the UniswapV3IncreaseLiqudityMutation object of the builder.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) Mutation() *UniswapV3IncreaseLiqudityMutation {
	return uvlu.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) ClearEvent() *UniswapV3IncreaseLiqudityUpdate {
	uvlu.mutation.ClearEvent()
	return uvlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uvlu.hooks) == 0 {
		if err = uvlu.check(); err != nil {
			return 0, err
		}
		affected, err = uvlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3IncreaseLiqudityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvlu.check(); err != nil {
				return 0, err
			}
			uvlu.mutation = mutation
			affected, err = uvlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uvlu.hooks) - 1; i >= 0; i-- {
			if uvlu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) SaveX(ctx context.Context) int {
	affected, err := uvlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) Exec(ctx context.Context) error {
	_, err := uvlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) ExecX(ctx context.Context) {
	if err := uvlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvlu *UniswapV3IncreaseLiqudityUpdate) check() error {
	if _, ok := uvlu.mutation.EventID(); uvlu.mutation.EventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"event\"")
	}
	return nil
}

func (uvlu *UniswapV3IncreaseLiqudityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3increaseliqudity.Table,
			Columns: uniswapv3increaseliqudity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3increaseliqudity.FieldID,
			},
		},
	}
	if ps := uvlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvlu.mutation.TokenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldTokenID,
		})
	}
	if value, ok := uvlu.mutation.Liquidity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldLiquidity,
		})
	}
	if value, ok := uvlu.mutation.Amount0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldAmount0,
		})
	}
	if value, ok := uvlu.mutation.Amount1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldAmount1,
		})
	}
	if uvlu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3increaseliqudity.EventTable,
			Columns: []string{uniswapv3increaseliqudity.EventColumn},
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
	if nodes := uvlu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3increaseliqudity.EventTable,
			Columns: []string{uniswapv3increaseliqudity.EventColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uvlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uniswapv3increaseliqudity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UniswapV3IncreaseLiqudityUpdateOne is the builder for updating a single UniswapV3IncreaseLiqudity entity.
type UniswapV3IncreaseLiqudityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UniswapV3IncreaseLiqudityMutation
}

// SetTokenID sets the "token_id" field.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) SetTokenID(s string) *UniswapV3IncreaseLiqudityUpdateOne {
	uvluo.mutation.SetTokenID(s)
	return uvluo
}

// SetLiquidity sets the "liquidity" field.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) SetLiquidity(s string) *UniswapV3IncreaseLiqudityUpdateOne {
	uvluo.mutation.SetLiquidity(s)
	return uvluo
}

// SetAmount0 sets the "amount0" field.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) SetAmount0(s string) *UniswapV3IncreaseLiqudityUpdateOne {
	uvluo.mutation.SetAmount0(s)
	return uvluo
}

// SetAmount1 sets the "amount1" field.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) SetAmount1(s string) *UniswapV3IncreaseLiqudityUpdateOne {
	uvluo.mutation.SetAmount1(s)
	return uvluo
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) SetEventID(id int) *UniswapV3IncreaseLiqudityUpdateOne {
	uvluo.mutation.SetEventID(id)
	return uvluo
}

// SetEvent sets the "event" edge to the Event entity.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) SetEvent(e *Event) *UniswapV3IncreaseLiqudityUpdateOne {
	return uvluo.SetEventID(e.ID)
}

// Mutation returns the UniswapV3IncreaseLiqudityMutation object of the builder.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) Mutation() *UniswapV3IncreaseLiqudityMutation {
	return uvluo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) ClearEvent() *UniswapV3IncreaseLiqudityUpdateOne {
	uvluo.mutation.ClearEvent()
	return uvluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) Select(field string, fields ...string) *UniswapV3IncreaseLiqudityUpdateOne {
	uvluo.fields = append([]string{field}, fields...)
	return uvluo
}

// Save executes the query and returns the updated UniswapV3IncreaseLiqudity entity.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) Save(ctx context.Context) (*UniswapV3IncreaseLiqudity, error) {
	var (
		err  error
		node *UniswapV3IncreaseLiqudity
	)
	if len(uvluo.hooks) == 0 {
		if err = uvluo.check(); err != nil {
			return nil, err
		}
		node, err = uvluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3IncreaseLiqudityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvluo.check(); err != nil {
				return nil, err
			}
			uvluo.mutation = mutation
			node, err = uvluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uvluo.hooks) - 1; i >= 0; i-- {
			if uvluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) SaveX(ctx context.Context) *UniswapV3IncreaseLiqudity {
	node, err := uvluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) Exec(ctx context.Context) error {
	_, err := uvluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) ExecX(ctx context.Context) {
	if err := uvluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) check() error {
	if _, ok := uvluo.mutation.EventID(); uvluo.mutation.EventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"event\"")
	}
	return nil
}

func (uvluo *UniswapV3IncreaseLiqudityUpdateOne) sqlSave(ctx context.Context) (_node *UniswapV3IncreaseLiqudity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3increaseliqudity.Table,
			Columns: uniswapv3increaseliqudity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3increaseliqudity.FieldID,
			},
		},
	}
	id, ok := uvluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UniswapV3IncreaseLiqudity.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uvluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3increaseliqudity.FieldID)
		for _, f := range fields {
			if !uniswapv3increaseliqudity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != uniswapv3increaseliqudity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uvluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvluo.mutation.TokenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldTokenID,
		})
	}
	if value, ok := uvluo.mutation.Liquidity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldLiquidity,
		})
	}
	if value, ok := uvluo.mutation.Amount0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldAmount0,
		})
	}
	if value, ok := uvluo.mutation.Amount1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldAmount1,
		})
	}
	if uvluo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3increaseliqudity.EventTable,
			Columns: []string{uniswapv3increaseliqudity.EventColumn},
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
	if nodes := uvluo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3increaseliqudity.EventTable,
			Columns: []string{uniswapv3increaseliqudity.EventColumn},
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
	_node = &UniswapV3IncreaseLiqudity{config: uvluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uvluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uniswapv3increaseliqudity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
