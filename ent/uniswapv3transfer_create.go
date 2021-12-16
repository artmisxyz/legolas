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
	"github.com/artmisxyz/blockinspector/ent/uniswapv3transfer"
)

// UniswapV3TransferCreate is the builder for creating a UniswapV3Transfer entity.
type UniswapV3TransferCreate struct {
	config
	mutation *UniswapV3TransferMutation
	hooks    []Hook
}

// SetTokenID sets the "token_id" field.
func (uvc *UniswapV3TransferCreate) SetTokenID(si *schema.BigInt) *UniswapV3TransferCreate {
	uvc.mutation.SetTokenID(si)
	return uvc
}

// SetFrom sets the "from" field.
func (uvc *UniswapV3TransferCreate) SetFrom(s string) *UniswapV3TransferCreate {
	uvc.mutation.SetFrom(s)
	return uvc
}

// SetTo sets the "to" field.
func (uvc *UniswapV3TransferCreate) SetTo(s string) *UniswapV3TransferCreate {
	uvc.mutation.SetTo(s)
	return uvc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvc *UniswapV3TransferCreate) SetEventID(id int) *UniswapV3TransferCreate {
	uvc.mutation.SetEventID(id)
	return uvc
}

// SetEvent sets the "event" edge to the Event entity.
func (uvc *UniswapV3TransferCreate) SetEvent(e *Event) *UniswapV3TransferCreate {
	return uvc.SetEventID(e.ID)
}

// Mutation returns the UniswapV3TransferMutation object of the builder.
func (uvc *UniswapV3TransferCreate) Mutation() *UniswapV3TransferMutation {
	return uvc.mutation
}

// Save creates the UniswapV3Transfer in the database.
func (uvc *UniswapV3TransferCreate) Save(ctx context.Context) (*UniswapV3Transfer, error) {
	var (
		err  error
		node *UniswapV3Transfer
	)
	if len(uvc.hooks) == 0 {
		if err = uvc.check(); err != nil {
			return nil, err
		}
		node, err = uvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3TransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvc.check(); err != nil {
				return nil, err
			}
			uvc.mutation = mutation
			if node, err = uvc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uvc.hooks) - 1; i >= 0; i-- {
			if uvc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uvc *UniswapV3TransferCreate) SaveX(ctx context.Context) *UniswapV3Transfer {
	v, err := uvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvc *UniswapV3TransferCreate) Exec(ctx context.Context) error {
	_, err := uvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvc *UniswapV3TransferCreate) ExecX(ctx context.Context) {
	if err := uvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvc *UniswapV3TransferCreate) check() error {
	if _, ok := uvc.mutation.TokenID(); !ok {
		return &ValidationError{Name: "token_id", err: errors.New(`ent: missing required field "token_id"`)}
	}
	if _, ok := uvc.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "from"`)}
	}
	if v, ok := uvc.mutation.From(); ok {
		if err := uniswapv3transfer.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "from": %w`, err)}
		}
	}
	if _, ok := uvc.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "to"`)}
	}
	if v, ok := uvc.mutation.To(); ok {
		if err := uniswapv3transfer.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "to": %w`, err)}
		}
	}
	if _, ok := uvc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New("ent: missing required edge \"event\"")}
	}
	return nil
}

func (uvc *UniswapV3TransferCreate) sqlSave(ctx context.Context) (*UniswapV3Transfer, error) {
	_node, _spec := uvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uvc *UniswapV3TransferCreate) createSpec() (*UniswapV3Transfer, *sqlgraph.CreateSpec) {
	var (
		_node = &UniswapV3Transfer{config: uvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: uniswapv3transfer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3transfer.FieldID,
			},
		}
	)
	if value, ok := uvc.mutation.TokenID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3transfer.FieldTokenID,
		})
		_node.TokenID = value
	}
	if value, ok := uvc.mutation.From(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3transfer.FieldFrom,
		})
		_node.From = value
	}
	if value, ok := uvc.mutation.To(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3transfer.FieldTo,
		})
		_node.To = value
	}
	if nodes := uvc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3transfer.EventTable,
			Columns: []string{uniswapv3transfer.EventColumn},
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
		_node.event_transfer = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UniswapV3TransferCreateBulk is the builder for creating many UniswapV3Transfer entities in bulk.
type UniswapV3TransferCreateBulk struct {
	config
	builders []*UniswapV3TransferCreate
}

// Save creates the UniswapV3Transfer entities in the database.
func (uvcb *UniswapV3TransferCreateBulk) Save(ctx context.Context) ([]*UniswapV3Transfer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uvcb.builders))
	nodes := make([]*UniswapV3Transfer, len(uvcb.builders))
	mutators := make([]Mutator, len(uvcb.builders))
	for i := range uvcb.builders {
		func(i int, root context.Context) {
			builder := uvcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UniswapV3TransferMutation)
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
					_, err = mutators[i+1].Mutate(root, uvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uvcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uvcb *UniswapV3TransferCreateBulk) SaveX(ctx context.Context) []*UniswapV3Transfer {
	v, err := uvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvcb *UniswapV3TransferCreateBulk) Exec(ctx context.Context) error {
	_, err := uvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvcb *UniswapV3TransferCreateBulk) ExecX(ctx context.Context) {
	if err := uvcb.Exec(ctx); err != nil {
		panic(err)
	}
}
