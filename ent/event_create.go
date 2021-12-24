// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/legolas/ent/event"
	"github.com/artmisxyz/legolas/ent/uniswapv3collect"
	"github.com/artmisxyz/legolas/ent/uniswapv3decreaseliqudity"
	"github.com/artmisxyz/legolas/ent/uniswapv3increaseliqudity"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolburn"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolcreated"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolflash"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolinitialize"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolmint"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolswap"
	"github.com/artmisxyz/legolas/ent/uniswapv3transfer"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *EventCreate) SetName(s string) *EventCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetSignature sets the "signature" field.
func (ec *EventCreate) SetSignature(s string) *EventCreate {
	ec.mutation.SetSignature(s)
	return ec
}

// SetAddress sets the "address" field.
func (ec *EventCreate) SetAddress(s string) *EventCreate {
	ec.mutation.SetAddress(s)
	return ec
}

// SetBlockNumber sets the "block_number" field.
func (ec *EventCreate) SetBlockNumber(u uint64) *EventCreate {
	ec.mutation.SetBlockNumber(u)
	return ec
}

// SetTxHash sets the "tx_hash" field.
func (ec *EventCreate) SetTxHash(s string) *EventCreate {
	ec.mutation.SetTxHash(s)
	return ec
}

// SetTxIndex sets the "tx_index" field.
func (ec *EventCreate) SetTxIndex(u uint) *EventCreate {
	ec.mutation.SetTxIndex(u)
	return ec
}

// SetBlockHash sets the "block_hash" field.
func (ec *EventCreate) SetBlockHash(s string) *EventCreate {
	ec.mutation.SetBlockHash(s)
	return ec
}

// SetIndex sets the "index" field.
func (ec *EventCreate) SetIndex(u uint) *EventCreate {
	ec.mutation.SetIndex(u)
	return ec
}

// SetIncreaseLiquidityID sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity by ID.
func (ec *EventCreate) SetIncreaseLiquidityID(id int) *EventCreate {
	ec.mutation.SetIncreaseLiquidityID(id)
	return ec
}

// SetNillableIncreaseLiquidityID sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableIncreaseLiquidityID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetIncreaseLiquidityID(*id)
	}
	return ec
}

// SetIncreaseLiquidity sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity.
func (ec *EventCreate) SetIncreaseLiquidity(u *UniswapV3IncreaseLiqudity) *EventCreate {
	return ec.SetIncreaseLiquidityID(u.ID)
}

// SetDecreaseLiquidityID sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity by ID.
func (ec *EventCreate) SetDecreaseLiquidityID(id int) *EventCreate {
	ec.mutation.SetDecreaseLiquidityID(id)
	return ec
}

// SetNillableDecreaseLiquidityID sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableDecreaseLiquidityID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetDecreaseLiquidityID(*id)
	}
	return ec
}

// SetDecreaseLiquidity sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity.
func (ec *EventCreate) SetDecreaseLiquidity(u *UniswapV3DecreaseLiqudity) *EventCreate {
	return ec.SetDecreaseLiquidityID(u.ID)
}

// SetCollectID sets the "collect" edge to the UniswapV3Collect entity by ID.
func (ec *EventCreate) SetCollectID(id int) *EventCreate {
	ec.mutation.SetCollectID(id)
	return ec
}

// SetNillableCollectID sets the "collect" edge to the UniswapV3Collect entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableCollectID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetCollectID(*id)
	}
	return ec
}

// SetCollect sets the "collect" edge to the UniswapV3Collect entity.
func (ec *EventCreate) SetCollect(u *UniswapV3Collect) *EventCreate {
	return ec.SetCollectID(u.ID)
}

// SetTransferID sets the "transfer" edge to the UniswapV3Transfer entity by ID.
func (ec *EventCreate) SetTransferID(id int) *EventCreate {
	ec.mutation.SetTransferID(id)
	return ec
}

// SetNillableTransferID sets the "transfer" edge to the UniswapV3Transfer entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableTransferID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetTransferID(*id)
	}
	return ec
}

// SetTransfer sets the "transfer" edge to the UniswapV3Transfer entity.
func (ec *EventCreate) SetTransfer(u *UniswapV3Transfer) *EventCreate {
	return ec.SetTransferID(u.ID)
}

// SetPoolCreatedID sets the "pool_created" edge to the UniswapV3PoolCreated entity by ID.
func (ec *EventCreate) SetPoolCreatedID(id int) *EventCreate {
	ec.mutation.SetPoolCreatedID(id)
	return ec
}

// SetNillablePoolCreatedID sets the "pool_created" edge to the UniswapV3PoolCreated entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillablePoolCreatedID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetPoolCreatedID(*id)
	}
	return ec
}

// SetPoolCreated sets the "pool_created" edge to the UniswapV3PoolCreated entity.
func (ec *EventCreate) SetPoolCreated(u *UniswapV3PoolCreated) *EventCreate {
	return ec.SetPoolCreatedID(u.ID)
}

// SetPoolInitializeID sets the "pool_initialize" edge to the UniswapV3PoolInitialize entity by ID.
func (ec *EventCreate) SetPoolInitializeID(id int) *EventCreate {
	ec.mutation.SetPoolInitializeID(id)
	return ec
}

// SetNillablePoolInitializeID sets the "pool_initialize" edge to the UniswapV3PoolInitialize entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillablePoolInitializeID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetPoolInitializeID(*id)
	}
	return ec
}

// SetPoolInitialize sets the "pool_initialize" edge to the UniswapV3PoolInitialize entity.
func (ec *EventCreate) SetPoolInitialize(u *UniswapV3PoolInitialize) *EventCreate {
	return ec.SetPoolInitializeID(u.ID)
}

// SetPoolSwapID sets the "pool_swap" edge to the UniswapV3PoolSwap entity by ID.
func (ec *EventCreate) SetPoolSwapID(id int) *EventCreate {
	ec.mutation.SetPoolSwapID(id)
	return ec
}

// SetNillablePoolSwapID sets the "pool_swap" edge to the UniswapV3PoolSwap entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillablePoolSwapID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetPoolSwapID(*id)
	}
	return ec
}

// SetPoolSwap sets the "pool_swap" edge to the UniswapV3PoolSwap entity.
func (ec *EventCreate) SetPoolSwap(u *UniswapV3PoolSwap) *EventCreate {
	return ec.SetPoolSwapID(u.ID)
}

// SetPoolMintID sets the "pool_mint" edge to the UniswapV3PoolMint entity by ID.
func (ec *EventCreate) SetPoolMintID(id int) *EventCreate {
	ec.mutation.SetPoolMintID(id)
	return ec
}

// SetNillablePoolMintID sets the "pool_mint" edge to the UniswapV3PoolMint entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillablePoolMintID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetPoolMintID(*id)
	}
	return ec
}

// SetPoolMint sets the "pool_mint" edge to the UniswapV3PoolMint entity.
func (ec *EventCreate) SetPoolMint(u *UniswapV3PoolMint) *EventCreate {
	return ec.SetPoolMintID(u.ID)
}

// SetPoolBurnID sets the "pool_burn" edge to the UniswapV3PoolBurn entity by ID.
func (ec *EventCreate) SetPoolBurnID(id int) *EventCreate {
	ec.mutation.SetPoolBurnID(id)
	return ec
}

// SetNillablePoolBurnID sets the "pool_burn" edge to the UniswapV3PoolBurn entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillablePoolBurnID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetPoolBurnID(*id)
	}
	return ec
}

// SetPoolBurn sets the "pool_burn" edge to the UniswapV3PoolBurn entity.
func (ec *EventCreate) SetPoolBurn(u *UniswapV3PoolBurn) *EventCreate {
	return ec.SetPoolBurnID(u.ID)
}

// SetPoolFlashID sets the "pool_flash" edge to the UniswapV3PoolFlash entity by ID.
func (ec *EventCreate) SetPoolFlashID(id int) *EventCreate {
	ec.mutation.SetPoolFlashID(id)
	return ec
}

// SetNillablePoolFlashID sets the "pool_flash" edge to the UniswapV3PoolFlash entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillablePoolFlashID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetPoolFlashID(*id)
	}
	return ec
}

// SetPoolFlash sets the "pool_flash" edge to the UniswapV3PoolFlash entity.
func (ec *EventCreate) SetPoolFlash(u *UniswapV3PoolFlash) *EventCreate {
	return ec.SetPoolFlashID(u.ID)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := ec.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Signature(); !ok {
		return &ValidationError{Name: "signature", err: errors.New(`ent: missing required field "signature"`)}
	}
	if v, ok := ec.mutation.Signature(); ok {
		if err := event.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf(`ent: validator failed for field "signature": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "address"`)}
	}
	if v, ok := ec.mutation.Address(); ok {
		if err := event.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "address": %w`, err)}
		}
	}
	if _, ok := ec.mutation.BlockNumber(); !ok {
		return &ValidationError{Name: "block_number", err: errors.New(`ent: missing required field "block_number"`)}
	}
	if _, ok := ec.mutation.TxHash(); !ok {
		return &ValidationError{Name: "tx_hash", err: errors.New(`ent: missing required field "tx_hash"`)}
	}
	if v, ok := ec.mutation.TxHash(); ok {
		if err := event.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf(`ent: validator failed for field "tx_hash": %w`, err)}
		}
	}
	if _, ok := ec.mutation.TxIndex(); !ok {
		return &ValidationError{Name: "tx_index", err: errors.New(`ent: missing required field "tx_index"`)}
	}
	if _, ok := ec.mutation.BlockHash(); !ok {
		return &ValidationError{Name: "block_hash", err: errors.New(`ent: missing required field "block_hash"`)}
	}
	if v, ok := ec.mutation.BlockHash(); ok {
		if err := event.BlockHashValidator(v); err != nil {
			return &ValidationError{Name: "block_hash", err: fmt.Errorf(`ent: validator failed for field "block_hash": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New(`ent: missing required field "index"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: event.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ec.mutation.Signature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSignature,
		})
		_node.Signature = value
	}
	if value, ok := ec.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := ec.mutation.BlockNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldBlockNumber,
		})
		_node.BlockNumber = value
	}
	if value, ok := ec.mutation.TxHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldTxHash,
		})
		_node.TxHash = value
	}
	if value, ok := ec.mutation.TxIndex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldTxIndex,
		})
		_node.TxIndex = value
	}
	if value, ok := ec.mutation.BlockHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldBlockHash,
		})
		_node.BlockHash = value
	}
	if value, ok := ec.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldIndex,
		})
		_node.Index = value
	}
	if nodes := ec.mutation.IncreaseLiquidityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.IncreaseLiquidityTable,
			Columns: []string{event.IncreaseLiquidityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3increaseliqudity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.DecreaseLiquidityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.DecreaseLiquidityTable,
			Columns: []string{event.DecreaseLiquidityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3decreaseliqudity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.CollectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.CollectTable,
			Columns: []string{event.CollectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3collect.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.TransferIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.TransferTable,
			Columns: []string{event.TransferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PoolCreatedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.PoolCreatedTable,
			Columns: []string{event.PoolCreatedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3poolcreated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PoolInitializeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.PoolInitializeTable,
			Columns: []string{event.PoolInitializeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3poolinitialize.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PoolSwapIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.PoolSwapTable,
			Columns: []string{event.PoolSwapColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3poolswap.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PoolMintIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.PoolMintTable,
			Columns: []string{event.PoolMintColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3poolmint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PoolBurnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.PoolBurnTable,
			Columns: []string{event.PoolBurnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3poolburn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PoolFlashIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.PoolFlashTable,
			Columns: []string{event.PoolFlashColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uniswapv3poolflash.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
