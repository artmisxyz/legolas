// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolswap"
)

// UniswapV3PoolSwapCreate is the builder for creating a UniswapV3PoolSwap entity.
type UniswapV3PoolSwapCreate struct {
	config
	mutation *UniswapV3PoolSwapMutation
	hooks    []Hook
}

// SetSender sets the "sender" field.
func (uvsc *UniswapV3PoolSwapCreate) SetSender(s string) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetSender(s)
	return uvsc
}

// SetRecipient sets the "recipient" field.
func (uvsc *UniswapV3PoolSwapCreate) SetRecipient(s string) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetRecipient(s)
	return uvsc
}

// SetAmount0 sets the "amount0" field.
func (uvsc *UniswapV3PoolSwapCreate) SetAmount0(s string) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetAmount0(s)
	return uvsc
}

// SetAmount1 sets the "amount1" field.
func (uvsc *UniswapV3PoolSwapCreate) SetAmount1(s string) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetAmount1(s)
	return uvsc
}

// SetSqrtPriceX96 sets the "sqrt_price_x96" field.
func (uvsc *UniswapV3PoolSwapCreate) SetSqrtPriceX96(s string) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetSqrtPriceX96(s)
	return uvsc
}

// SetLiquidity sets the "liquidity" field.
func (uvsc *UniswapV3PoolSwapCreate) SetLiquidity(s string) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetLiquidity(s)
	return uvsc
}

// SetTick sets the "tick" field.
func (uvsc *UniswapV3PoolSwapCreate) SetTick(s string) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetTick(s)
	return uvsc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvsc *UniswapV3PoolSwapCreate) SetEventID(id int) *UniswapV3PoolSwapCreate {
	uvsc.mutation.SetEventID(id)
	return uvsc
}

// SetEvent sets the "event" edge to the Event entity.
func (uvsc *UniswapV3PoolSwapCreate) SetEvent(e *Event) *UniswapV3PoolSwapCreate {
	return uvsc.SetEventID(e.ID)
}

// Mutation returns the UniswapV3PoolSwapMutation object of the builder.
func (uvsc *UniswapV3PoolSwapCreate) Mutation() *UniswapV3PoolSwapMutation {
	return uvsc.mutation
}

// Save creates the UniswapV3PoolSwap in the database.
func (uvsc *UniswapV3PoolSwapCreate) Save(ctx context.Context) (*UniswapV3PoolSwap, error) {
	var (
		err  error
		node *UniswapV3PoolSwap
	)
	if len(uvsc.hooks) == 0 {
		if err = uvsc.check(); err != nil {
			return nil, err
		}
		node, err = uvsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3PoolSwapMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvsc.check(); err != nil {
				return nil, err
			}
			uvsc.mutation = mutation
			if node, err = uvsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uvsc.hooks) - 1; i >= 0; i-- {
			if uvsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uvsc *UniswapV3PoolSwapCreate) SaveX(ctx context.Context) *UniswapV3PoolSwap {
	v, err := uvsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvsc *UniswapV3PoolSwapCreate) Exec(ctx context.Context) error {
	_, err := uvsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvsc *UniswapV3PoolSwapCreate) ExecX(ctx context.Context) {
	if err := uvsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvsc *UniswapV3PoolSwapCreate) check() error {
	if _, ok := uvsc.mutation.Sender(); !ok {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required field "sender"`)}
	}
	if _, ok := uvsc.mutation.Recipient(); !ok {
		return &ValidationError{Name: "recipient", err: errors.New(`ent: missing required field "recipient"`)}
	}
	if _, ok := uvsc.mutation.Amount0(); !ok {
		return &ValidationError{Name: "amount0", err: errors.New(`ent: missing required field "amount0"`)}
	}
	if _, ok := uvsc.mutation.Amount1(); !ok {
		return &ValidationError{Name: "amount1", err: errors.New(`ent: missing required field "amount1"`)}
	}
	if _, ok := uvsc.mutation.SqrtPriceX96(); !ok {
		return &ValidationError{Name: "sqrt_price_x96", err: errors.New(`ent: missing required field "sqrt_price_x96"`)}
	}
	if _, ok := uvsc.mutation.Liquidity(); !ok {
		return &ValidationError{Name: "liquidity", err: errors.New(`ent: missing required field "liquidity"`)}
	}
	if _, ok := uvsc.mutation.Tick(); !ok {
		return &ValidationError{Name: "tick", err: errors.New(`ent: missing required field "tick"`)}
	}
	if _, ok := uvsc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New("ent: missing required edge \"event\"")}
	}
	return nil
}

func (uvsc *UniswapV3PoolSwapCreate) sqlSave(ctx context.Context) (*UniswapV3PoolSwap, error) {
	_node, _spec := uvsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uvsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uvsc *UniswapV3PoolSwapCreate) createSpec() (*UniswapV3PoolSwap, *sqlgraph.CreateSpec) {
	var (
		_node = &UniswapV3PoolSwap{config: uvsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: uniswapv3poolswap.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolswap.FieldID,
			},
		}
	)
	if value, ok := uvsc.mutation.Sender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolswap.FieldSender,
		})
		_node.Sender = value
	}
	if value, ok := uvsc.mutation.Recipient(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolswap.FieldRecipient,
		})
		_node.Recipient = value
	}
	if value, ok := uvsc.mutation.Amount0(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolswap.FieldAmount0,
		})
		_node.Amount0 = value
	}
	if value, ok := uvsc.mutation.Amount1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolswap.FieldAmount1,
		})
		_node.Amount1 = value
	}
	if value, ok := uvsc.mutation.SqrtPriceX96(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolswap.FieldSqrtPriceX96,
		})
		_node.SqrtPriceX96 = value
	}
	if value, ok := uvsc.mutation.Liquidity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolswap.FieldLiquidity,
		})
		_node.Liquidity = value
	}
	if value, ok := uvsc.mutation.Tick(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolswap.FieldTick,
		})
		_node.Tick = value
	}
	if nodes := uvsc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3poolswap.EventTable,
			Columns: []string{uniswapv3poolswap.EventColumn},
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
		_node.event_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UniswapV3PoolSwapCreateBulk is the builder for creating many UniswapV3PoolSwap entities in bulk.
type UniswapV3PoolSwapCreateBulk struct {
	config
	builders []*UniswapV3PoolSwapCreate
}

// Save creates the UniswapV3PoolSwap entities in the database.
func (uvscb *UniswapV3PoolSwapCreateBulk) Save(ctx context.Context) ([]*UniswapV3PoolSwap, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uvscb.builders))
	nodes := make([]*UniswapV3PoolSwap, len(uvscb.builders))
	mutators := make([]Mutator, len(uvscb.builders))
	for i := range uvscb.builders {
		func(i int, root context.Context) {
			builder := uvscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UniswapV3PoolSwapMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uvscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uvscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uvscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uvscb *UniswapV3PoolSwapCreateBulk) SaveX(ctx context.Context) []*UniswapV3PoolSwap {
	v, err := uvscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvscb *UniswapV3PoolSwapCreateBulk) Exec(ctx context.Context) error {
	_, err := uvscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvscb *UniswapV3PoolSwapCreateBulk) ExecX(ctx context.Context) {
	if err := uvscb.Exec(ctx); err != nil {
		panic(err)
	}
}
