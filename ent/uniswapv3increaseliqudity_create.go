// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/blockinspector/ent/event"
	"github.com/artmisxyz/blockinspector/ent/schema"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3increaseliqudity"
)

// UniswapV3IncreaseLiqudityCreate is the builder for creating a UniswapV3IncreaseLiqudity entity.
type UniswapV3IncreaseLiqudityCreate struct {
	config
	mutation *UniswapV3IncreaseLiqudityMutation
	hooks    []Hook
}

// SetTokenID sets the "token_id" field.
func (uvlc *UniswapV3IncreaseLiqudityCreate) SetTokenID(si *schema.BigInt) *UniswapV3IncreaseLiqudityCreate {
	uvlc.mutation.SetTokenID(si)
	return uvlc
}

// SetLiquidity sets the "liquidity" field.
func (uvlc *UniswapV3IncreaseLiqudityCreate) SetLiquidity(si *schema.BigInt) *UniswapV3IncreaseLiqudityCreate {
	uvlc.mutation.SetLiquidity(si)
	return uvlc
}

// SetAmount0 sets the "amount0" field.
func (uvlc *UniswapV3IncreaseLiqudityCreate) SetAmount0(si *schema.BigInt) *UniswapV3IncreaseLiqudityCreate {
	uvlc.mutation.SetAmount0(si)
	return uvlc
}

// SetAmount1 sets the "amount1" field.
func (uvlc *UniswapV3IncreaseLiqudityCreate) SetAmount1(si *schema.BigInt) *UniswapV3IncreaseLiqudityCreate {
	uvlc.mutation.SetAmount1(si)
	return uvlc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvlc *UniswapV3IncreaseLiqudityCreate) SetEventID(id int) *UniswapV3IncreaseLiqudityCreate {
	uvlc.mutation.SetEventID(id)
	return uvlc
}

// SetEvent sets the "event" edge to the Event entity.
func (uvlc *UniswapV3IncreaseLiqudityCreate) SetEvent(e *Event) *UniswapV3IncreaseLiqudityCreate {
	return uvlc.SetEventID(e.ID)
}

// Mutation returns the UniswapV3IncreaseLiqudityMutation object of the builder.
func (uvlc *UniswapV3IncreaseLiqudityCreate) Mutation() *UniswapV3IncreaseLiqudityMutation {
	return uvlc.mutation
}

// Save creates the UniswapV3IncreaseLiqudity in the database.
func (uvlc *UniswapV3IncreaseLiqudityCreate) Save(ctx context.Context) (*UniswapV3IncreaseLiqudity, error) {
	var (
		err  error
		node *UniswapV3IncreaseLiqudity
	)
	if len(uvlc.hooks) == 0 {
		if err = uvlc.check(); err != nil {
			return nil, err
		}
		node, err = uvlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3IncreaseLiqudityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvlc.check(); err != nil {
				return nil, err
			}
			uvlc.mutation = mutation
			if node, err = uvlc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uvlc.hooks) - 1; i >= 0; i-- {
			if uvlc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvlc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvlc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uvlc *UniswapV3IncreaseLiqudityCreate) SaveX(ctx context.Context) *UniswapV3IncreaseLiqudity {
	v, err := uvlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvlc *UniswapV3IncreaseLiqudityCreate) Exec(ctx context.Context) error {
	_, err := uvlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvlc *UniswapV3IncreaseLiqudityCreate) ExecX(ctx context.Context) {
	if err := uvlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvlc *UniswapV3IncreaseLiqudityCreate) check() error {
	if _, ok := uvlc.mutation.TokenID(); !ok {
		return &ValidationError{Name: "token_id", err: errors.New(`ent: missing required field "token_id"`)}
	}
	if _, ok := uvlc.mutation.Liquidity(); !ok {
		return &ValidationError{Name: "liquidity", err: errors.New(`ent: missing required field "liquidity"`)}
	}
	if _, ok := uvlc.mutation.Amount0(); !ok {
		return &ValidationError{Name: "amount0", err: errors.New(`ent: missing required field "amount0"`)}
	}
	if _, ok := uvlc.mutation.Amount1(); !ok {
		return &ValidationError{Name: "amount1", err: errors.New(`ent: missing required field "amount1"`)}
	}
	if _, ok := uvlc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New("ent: missing required edge \"event\"")}
	}
	return nil
}

func (uvlc *UniswapV3IncreaseLiqudityCreate) sqlSave(ctx context.Context) (*UniswapV3IncreaseLiqudity, error) {
	_node, _spec := uvlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uvlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uvlc *UniswapV3IncreaseLiqudityCreate) createSpec() (*UniswapV3IncreaseLiqudity, *sqlgraph.CreateSpec) {
	var (
		_node = &UniswapV3IncreaseLiqudity{config: uvlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: uniswapv3increaseliqudity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3increaseliqudity.FieldID,
			},
		}
	)
	if value, ok := uvlc.mutation.TokenID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldTokenID,
		})
		_node.TokenID = value
	}
	if value, ok := uvlc.mutation.Liquidity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldLiquidity,
		})
		_node.Liquidity = value
	}
	if value, ok := uvlc.mutation.Amount0(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldAmount0,
		})
		_node.Amount0 = value
	}
	if value, ok := uvlc.mutation.Amount1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3increaseliqudity.FieldAmount1,
		})
		_node.Amount1 = value
	}
	if nodes := uvlc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_node.event_increase_liquidity = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UniswapV3IncreaseLiqudityCreateBulk is the builder for creating many UniswapV3IncreaseLiqudity entities in bulk.
type UniswapV3IncreaseLiqudityCreateBulk struct {
	config
	builders []*UniswapV3IncreaseLiqudityCreate
}

// Save creates the UniswapV3IncreaseLiqudity entities in the database.
func (uvlcb *UniswapV3IncreaseLiqudityCreateBulk) Save(ctx context.Context) ([]*UniswapV3IncreaseLiqudity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uvlcb.builders))
	nodes := make([]*UniswapV3IncreaseLiqudity, len(uvlcb.builders))
	mutators := make([]Mutator, len(uvlcb.builders))
	for i := range uvlcb.builders {
		func(i int, root context.Context) {
			builder := uvlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UniswapV3IncreaseLiqudityMutation)
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
					_, err = mutators[i+1].Mutate(root, uvlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uvlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uvlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uvlcb *UniswapV3IncreaseLiqudityCreateBulk) SaveX(ctx context.Context) []*UniswapV3IncreaseLiqudity {
	v, err := uvlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvlcb *UniswapV3IncreaseLiqudityCreateBulk) Exec(ctx context.Context) error {
	_, err := uvlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvlcb *UniswapV3IncreaseLiqudityCreateBulk) ExecX(ctx context.Context) {
	if err := uvlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
