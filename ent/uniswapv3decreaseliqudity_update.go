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
	"github.com/artmisxyz/legolas/ent/uniswapv3decreaseliqudity"
)

// UniswapV3DecreaseLiqudityUpdate is the builder for updating UniswapV3DecreaseLiqudity entities.
type UniswapV3DecreaseLiqudityUpdate struct {
	config
	hooks    []Hook
	mutation *UniswapV3DecreaseLiqudityMutation
}

// Where appends a list predicates to the UniswapV3DecreaseLiqudityUpdate builder.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) Where(ps ...predicate.UniswapV3DecreaseLiqudity) *UniswapV3DecreaseLiqudityUpdate {
	uvlu.mutation.Where(ps...)
	return uvlu
}

// SetTokenID sets the "token_id" field.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) SetTokenID(s string) *UniswapV3DecreaseLiqudityUpdate {
	uvlu.mutation.SetTokenID(s)
	return uvlu
}

// SetLiquidity sets the "liquidity" field.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) SetLiquidity(s string) *UniswapV3DecreaseLiqudityUpdate {
	uvlu.mutation.SetLiquidity(s)
	return uvlu
}

// SetAmount0 sets the "amount0" field.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) SetAmount0(s string) *UniswapV3DecreaseLiqudityUpdate {
	uvlu.mutation.SetAmount0(s)
	return uvlu
}

// SetAmount1 sets the "amount1" field.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) SetAmount1(s string) *UniswapV3DecreaseLiqudityUpdate {
	uvlu.mutation.SetAmount1(s)
	return uvlu
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) SetEventID(id int) *UniswapV3DecreaseLiqudityUpdate {
	uvlu.mutation.SetEventID(id)
	return uvlu
}

// SetEvent sets the "event" edge to the Event entity.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) SetEvent(e *Event) *UniswapV3DecreaseLiqudityUpdate {
	return uvlu.SetEventID(e.ID)
}

// Mutation returns the UniswapV3DecreaseLiqudityMutation object of the builder.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) Mutation() *UniswapV3DecreaseLiqudityMutation {
	return uvlu.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) ClearEvent() *UniswapV3DecreaseLiqudityUpdate {
	uvlu.mutation.ClearEvent()
	return uvlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*UniswapV3DecreaseLiqudityMutation)
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
func (uvlu *UniswapV3DecreaseLiqudityUpdate) SaveX(ctx context.Context) int {
	affected, err := uvlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) Exec(ctx context.Context) error {
	_, err := uvlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) ExecX(ctx context.Context) {
	if err := uvlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvlu *UniswapV3DecreaseLiqudityUpdate) check() error {
	if _, ok := uvlu.mutation.EventID(); uvlu.mutation.EventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"event\"")
	}
	return nil
}

func (uvlu *UniswapV3DecreaseLiqudityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3decreaseliqudity.Table,
			Columns: uniswapv3decreaseliqudity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3decreaseliqudity.FieldID,
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
			Column: uniswapv3decreaseliqudity.FieldTokenID,
		})
	}
	if value, ok := uvlu.mutation.Liquidity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3decreaseliqudity.FieldLiquidity,
		})
	}
	if value, ok := uvlu.mutation.Amount0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3decreaseliqudity.FieldAmount0,
		})
	}
	if value, ok := uvlu.mutation.Amount1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3decreaseliqudity.FieldAmount1,
		})
	}
	if uvlu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3decreaseliqudity.EventTable,
			Columns: []string{uniswapv3decreaseliqudity.EventColumn},
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
			Table:   uniswapv3decreaseliqudity.EventTable,
			Columns: []string{uniswapv3decreaseliqudity.EventColumn},
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
			err = &NotFoundError{uniswapv3decreaseliqudity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UniswapV3DecreaseLiqudityUpdateOne is the builder for updating a single UniswapV3DecreaseLiqudity entity.
type UniswapV3DecreaseLiqudityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UniswapV3DecreaseLiqudityMutation
}

// SetTokenID sets the "token_id" field.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) SetTokenID(s string) *UniswapV3DecreaseLiqudityUpdateOne {
	uvluo.mutation.SetTokenID(s)
	return uvluo
}

// SetLiquidity sets the "liquidity" field.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) SetLiquidity(s string) *UniswapV3DecreaseLiqudityUpdateOne {
	uvluo.mutation.SetLiquidity(s)
	return uvluo
}

// SetAmount0 sets the "amount0" field.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) SetAmount0(s string) *UniswapV3DecreaseLiqudityUpdateOne {
	uvluo.mutation.SetAmount0(s)
	return uvluo
}

// SetAmount1 sets the "amount1" field.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) SetAmount1(s string) *UniswapV3DecreaseLiqudityUpdateOne {
	uvluo.mutation.SetAmount1(s)
	return uvluo
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) SetEventID(id int) *UniswapV3DecreaseLiqudityUpdateOne {
	uvluo.mutation.SetEventID(id)
	return uvluo
}

// SetEvent sets the "event" edge to the Event entity.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) SetEvent(e *Event) *UniswapV3DecreaseLiqudityUpdateOne {
	return uvluo.SetEventID(e.ID)
}

// Mutation returns the UniswapV3DecreaseLiqudityMutation object of the builder.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) Mutation() *UniswapV3DecreaseLiqudityMutation {
	return uvluo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) ClearEvent() *UniswapV3DecreaseLiqudityUpdateOne {
	uvluo.mutation.ClearEvent()
	return uvluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) Select(field string, fields ...string) *UniswapV3DecreaseLiqudityUpdateOne {
	uvluo.fields = append([]string{field}, fields...)
	return uvluo
}

// Save executes the query and returns the updated UniswapV3DecreaseLiqudity entity.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) Save(ctx context.Context) (*UniswapV3DecreaseLiqudity, error) {
	var (
		err  error
		node *UniswapV3DecreaseLiqudity
	)
	if len(uvluo.hooks) == 0 {
		if err = uvluo.check(); err != nil {
			return nil, err
		}
		node, err = uvluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3DecreaseLiqudityMutation)
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
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) SaveX(ctx context.Context) *UniswapV3DecreaseLiqudity {
	node, err := uvluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) Exec(ctx context.Context) error {
	_, err := uvluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) ExecX(ctx context.Context) {
	if err := uvluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) check() error {
	if _, ok := uvluo.mutation.EventID(); uvluo.mutation.EventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"event\"")
	}
	return nil
}

func (uvluo *UniswapV3DecreaseLiqudityUpdateOne) sqlSave(ctx context.Context) (_node *UniswapV3DecreaseLiqudity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   uniswapv3decreaseliqudity.Table,
			Columns: uniswapv3decreaseliqudity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3decreaseliqudity.FieldID,
			},
		},
	}
	id, ok := uvluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UniswapV3DecreaseLiqudity.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uvluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uniswapv3decreaseliqudity.FieldID)
		for _, f := range fields {
			if !uniswapv3decreaseliqudity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != uniswapv3decreaseliqudity.FieldID {
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
			Column: uniswapv3decreaseliqudity.FieldTokenID,
		})
	}
	if value, ok := uvluo.mutation.Liquidity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3decreaseliqudity.FieldLiquidity,
		})
	}
	if value, ok := uvluo.mutation.Amount0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3decreaseliqudity.FieldAmount0,
		})
	}
	if value, ok := uvluo.mutation.Amount1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3decreaseliqudity.FieldAmount1,
		})
	}
	if uvluo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3decreaseliqudity.EventTable,
			Columns: []string{uniswapv3decreaseliqudity.EventColumn},
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
			Table:   uniswapv3decreaseliqudity.EventTable,
			Columns: []string{uniswapv3decreaseliqudity.EventColumn},
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
	_node = &UniswapV3DecreaseLiqudity{config: uvluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uvluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uniswapv3decreaseliqudity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
