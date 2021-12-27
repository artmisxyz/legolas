// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
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
	conflict []sql.ConflictOption
}

// SetTime sets the "time" field.
func (ec *EventCreate) SetTime(t time.Time) *EventCreate {
	ec.mutation.SetTime(t)
	return ec
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
	if _, ok := ec.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "time"`)}
	}
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
	_spec.OnConflict = ec.conflict
	if value, ok := ec.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldTime,
		})
		_node.Time = value
	}
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.Create().
//		SetTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetTime(v+v).
//		}).
//		Exec(ctx)
//
func (ec *EventCreate) OnConflict(opts ...sql.ConflictOption) *EventUpsertOne {
	ec.conflict = opts
	return &EventUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ec *EventCreate) OnConflictColumns(columns ...string) *EventUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertOne{
		create: ec,
	}
}

type (
	// EventUpsertOne is the builder for "upsert"-ing
	//  one Event node.
	EventUpsertOne struct {
		create *EventCreate
	}

	// EventUpsert is the "OnConflict" setter.
	EventUpsert struct {
		*sql.UpdateSet
	}
)

// SetTime sets the "time" field.
func (u *EventUpsert) SetTime(v time.Time) *EventUpsert {
	u.Set(event.FieldTime, v)
	return u
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *EventUpsert) UpdateTime() *EventUpsert {
	u.SetExcluded(event.FieldTime)
	return u
}

// SetName sets the "name" field.
func (u *EventUpsert) SetName(v string) *EventUpsert {
	u.Set(event.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *EventUpsert) UpdateName() *EventUpsert {
	u.SetExcluded(event.FieldName)
	return u
}

// SetSignature sets the "signature" field.
func (u *EventUpsert) SetSignature(v string) *EventUpsert {
	u.Set(event.FieldSignature, v)
	return u
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventUpsert) UpdateSignature() *EventUpsert {
	u.SetExcluded(event.FieldSignature)
	return u
}

// SetAddress sets the "address" field.
func (u *EventUpsert) SetAddress(v string) *EventUpsert {
	u.Set(event.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EventUpsert) UpdateAddress() *EventUpsert {
	u.SetExcluded(event.FieldAddress)
	return u
}

// SetBlockNumber sets the "block_number" field.
func (u *EventUpsert) SetBlockNumber(v uint64) *EventUpsert {
	u.Set(event.FieldBlockNumber, v)
	return u
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *EventUpsert) UpdateBlockNumber() *EventUpsert {
	u.SetExcluded(event.FieldBlockNumber)
	return u
}

// SetTxHash sets the "tx_hash" field.
func (u *EventUpsert) SetTxHash(v string) *EventUpsert {
	u.Set(event.FieldTxHash, v)
	return u
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *EventUpsert) UpdateTxHash() *EventUpsert {
	u.SetExcluded(event.FieldTxHash)
	return u
}

// SetTxIndex sets the "tx_index" field.
func (u *EventUpsert) SetTxIndex(v uint) *EventUpsert {
	u.Set(event.FieldTxIndex, v)
	return u
}

// UpdateTxIndex sets the "tx_index" field to the value that was provided on create.
func (u *EventUpsert) UpdateTxIndex() *EventUpsert {
	u.SetExcluded(event.FieldTxIndex)
	return u
}

// SetBlockHash sets the "block_hash" field.
func (u *EventUpsert) SetBlockHash(v string) *EventUpsert {
	u.Set(event.FieldBlockHash, v)
	return u
}

// UpdateBlockHash sets the "block_hash" field to the value that was provided on create.
func (u *EventUpsert) UpdateBlockHash() *EventUpsert {
	u.SetExcluded(event.FieldBlockHash)
	return u
}

// SetIndex sets the "index" field.
func (u *EventUpsert) SetIndex(v uint) *EventUpsert {
	u.Set(event.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *EventUpsert) UpdateIndex() *EventUpsert {
	u.SetExcluded(event.FieldIndex)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *EventUpsertOne) UpdateNewValues() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Event.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *EventUpsertOne) Ignore() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertOne) DoNothing() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreate.OnConflict
// documentation for more info.
func (u *EventUpsertOne) Update(set func(*EventUpsert)) *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetTime sets the "time" field.
func (u *EventUpsertOne) SetTime(v time.Time) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateTime() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTime()
	})
}

// SetName sets the "name" field.
func (u *EventUpsertOne) SetName(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateName() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateName()
	})
}

// SetSignature sets the "signature" field.
func (u *EventUpsertOne) SetSignature(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateSignature() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateSignature()
	})
}

// SetAddress sets the "address" field.
func (u *EventUpsertOne) SetAddress(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateAddress() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateAddress()
	})
}

// SetBlockNumber sets the "block_number" field.
func (u *EventUpsertOne) SetBlockNumber(v uint64) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetBlockNumber(v)
	})
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateBlockNumber() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateBlockNumber()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *EventUpsertOne) SetTxHash(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateTxHash() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTxHash()
	})
}

// SetTxIndex sets the "tx_index" field.
func (u *EventUpsertOne) SetTxIndex(v uint) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetTxIndex(v)
	})
}

// UpdateTxIndex sets the "tx_index" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateTxIndex() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTxIndex()
	})
}

// SetBlockHash sets the "block_hash" field.
func (u *EventUpsertOne) SetBlockHash(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetBlockHash(v)
	})
}

// UpdateBlockHash sets the "block_hash" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateBlockHash() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateBlockHash()
	})
}

// SetIndex sets the "index" field.
func (u *EventUpsertOne) SetIndex(v uint) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateIndex() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateIndex()
	})
}

// Exec executes the query.
func (u *EventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EventUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EventUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = ecb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetTime(v+v).
//		}).
//		Exec(ctx)
//
func (ecb *EventCreateBulk) OnConflict(opts ...sql.ConflictOption) *EventUpsertBulk {
	ecb.conflict = opts
	return &EventUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ecb *EventCreateBulk) OnConflictColumns(columns ...string) *EventUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertBulk{
		create: ecb,
	}
}

// EventUpsertBulk is the builder for "upsert"-ing
// a bulk of Event nodes.
type EventUpsertBulk struct {
	create *EventCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *EventUpsertBulk) UpdateNewValues() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *EventUpsertBulk) Ignore() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertBulk) DoNothing() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreateBulk.OnConflict
// documentation for more info.
func (u *EventUpsertBulk) Update(set func(*EventUpsert)) *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetTime sets the "time" field.
func (u *EventUpsertBulk) SetTime(v time.Time) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateTime() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTime()
	})
}

// SetName sets the "name" field.
func (u *EventUpsertBulk) SetName(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateName() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateName()
	})
}

// SetSignature sets the "signature" field.
func (u *EventUpsertBulk) SetSignature(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateSignature() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateSignature()
	})
}

// SetAddress sets the "address" field.
func (u *EventUpsertBulk) SetAddress(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateAddress() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateAddress()
	})
}

// SetBlockNumber sets the "block_number" field.
func (u *EventUpsertBulk) SetBlockNumber(v uint64) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetBlockNumber(v)
	})
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateBlockNumber() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateBlockNumber()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *EventUpsertBulk) SetTxHash(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateTxHash() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTxHash()
	})
}

// SetTxIndex sets the "tx_index" field.
func (u *EventUpsertBulk) SetTxIndex(v uint) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetTxIndex(v)
	})
}

// UpdateTxIndex sets the "tx_index" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateTxIndex() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTxIndex()
	})
}

// SetBlockHash sets the "block_hash" field.
func (u *EventUpsertBulk) SetBlockHash(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetBlockHash(v)
	})
}

// UpdateBlockHash sets the "block_hash" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateBlockHash() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateBlockHash()
	})
}

// SetIndex sets the "index" field.
func (u *EventUpsertBulk) SetIndex(v uint) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateIndex() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateIndex()
	})
}

// Exec executes the query.
func (u *EventUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
