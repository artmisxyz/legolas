// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/blockinspector/ent/event"
	"github.com/artmisxyz/blockinspector/ent/predicate"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3collect"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3decreaseliqudity"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3increaseliqudity"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3poolcreated"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3transfer"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventUpdate) SetName(s string) *EventUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetSignature sets the "signature" field.
func (eu *EventUpdate) SetSignature(s string) *EventUpdate {
	eu.mutation.SetSignature(s)
	return eu
}

// SetAddress sets the "address" field.
func (eu *EventUpdate) SetAddress(s string) *EventUpdate {
	eu.mutation.SetAddress(s)
	return eu
}

// SetBlockNumber sets the "block_number" field.
func (eu *EventUpdate) SetBlockNumber(u uint64) *EventUpdate {
	eu.mutation.ResetBlockNumber()
	eu.mutation.SetBlockNumber(u)
	return eu
}

// AddBlockNumber adds u to the "block_number" field.
func (eu *EventUpdate) AddBlockNumber(u uint64) *EventUpdate {
	eu.mutation.AddBlockNumber(u)
	return eu
}

// SetTxHash sets the "tx_hash" field.
func (eu *EventUpdate) SetTxHash(s string) *EventUpdate {
	eu.mutation.SetTxHash(s)
	return eu
}

// SetTxIndex sets the "tx_index" field.
func (eu *EventUpdate) SetTxIndex(u uint) *EventUpdate {
	eu.mutation.ResetTxIndex()
	eu.mutation.SetTxIndex(u)
	return eu
}

// AddTxIndex adds u to the "tx_index" field.
func (eu *EventUpdate) AddTxIndex(u uint) *EventUpdate {
	eu.mutation.AddTxIndex(u)
	return eu
}

// SetBlockHash sets the "block_hash" field.
func (eu *EventUpdate) SetBlockHash(s string) *EventUpdate {
	eu.mutation.SetBlockHash(s)
	return eu
}

// SetIndex sets the "index" field.
func (eu *EventUpdate) SetIndex(u uint) *EventUpdate {
	eu.mutation.ResetIndex()
	eu.mutation.SetIndex(u)
	return eu
}

// AddIndex adds u to the "index" field.
func (eu *EventUpdate) AddIndex(u uint) *EventUpdate {
	eu.mutation.AddIndex(u)
	return eu
}

// SetHash sets the "hash" field.
func (eu *EventUpdate) SetHash(s string) *EventUpdate {
	eu.mutation.SetHash(s)
	return eu
}

// SetIncreaseLiquidityID sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity by ID.
func (eu *EventUpdate) SetIncreaseLiquidityID(id int) *EventUpdate {
	eu.mutation.SetIncreaseLiquidityID(id)
	return eu
}

// SetNillableIncreaseLiquidityID sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableIncreaseLiquidityID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetIncreaseLiquidityID(*id)
	}
	return eu
}

// SetIncreaseLiquidity sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity.
func (eu *EventUpdate) SetIncreaseLiquidity(u *UniswapV3IncreaseLiqudity) *EventUpdate {
	return eu.SetIncreaseLiquidityID(u.ID)
}

// SetDecreaseLiquidityID sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity by ID.
func (eu *EventUpdate) SetDecreaseLiquidityID(id int) *EventUpdate {
	eu.mutation.SetDecreaseLiquidityID(id)
	return eu
}

// SetNillableDecreaseLiquidityID sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableDecreaseLiquidityID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetDecreaseLiquidityID(*id)
	}
	return eu
}

// SetDecreaseLiquidity sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity.
func (eu *EventUpdate) SetDecreaseLiquidity(u *UniswapV3DecreaseLiqudity) *EventUpdate {
	return eu.SetDecreaseLiquidityID(u.ID)
}

// SetCollectID sets the "collect" edge to the UniswapV3Collect entity by ID.
func (eu *EventUpdate) SetCollectID(id int) *EventUpdate {
	eu.mutation.SetCollectID(id)
	return eu
}

// SetNillableCollectID sets the "collect" edge to the UniswapV3Collect entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableCollectID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetCollectID(*id)
	}
	return eu
}

// SetCollect sets the "collect" edge to the UniswapV3Collect entity.
func (eu *EventUpdate) SetCollect(u *UniswapV3Collect) *EventUpdate {
	return eu.SetCollectID(u.ID)
}

// SetTransferID sets the "transfer" edge to the UniswapV3Transfer entity by ID.
func (eu *EventUpdate) SetTransferID(id int) *EventUpdate {
	eu.mutation.SetTransferID(id)
	return eu
}

// SetNillableTransferID sets the "transfer" edge to the UniswapV3Transfer entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableTransferID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetTransferID(*id)
	}
	return eu
}

// SetTransfer sets the "transfer" edge to the UniswapV3Transfer entity.
func (eu *EventUpdate) SetTransfer(u *UniswapV3Transfer) *EventUpdate {
	return eu.SetTransferID(u.ID)
}

// SetPoolCreatedID sets the "pool_created" edge to the UniswapV3PoolCreated entity by ID.
func (eu *EventUpdate) SetPoolCreatedID(id int) *EventUpdate {
	eu.mutation.SetPoolCreatedID(id)
	return eu
}

// SetNillablePoolCreatedID sets the "pool_created" edge to the UniswapV3PoolCreated entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillablePoolCreatedID(id *int) *EventUpdate {
	if id != nil {
		eu = eu.SetPoolCreatedID(*id)
	}
	return eu
}

// SetPoolCreated sets the "pool_created" edge to the UniswapV3PoolCreated entity.
func (eu *EventUpdate) SetPoolCreated(u *UniswapV3PoolCreated) *EventUpdate {
	return eu.SetPoolCreatedID(u.ID)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearIncreaseLiquidity clears the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity.
func (eu *EventUpdate) ClearIncreaseLiquidity() *EventUpdate {
	eu.mutation.ClearIncreaseLiquidity()
	return eu
}

// ClearDecreaseLiquidity clears the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity.
func (eu *EventUpdate) ClearDecreaseLiquidity() *EventUpdate {
	eu.mutation.ClearDecreaseLiquidity()
	return eu
}

// ClearCollect clears the "collect" edge to the UniswapV3Collect entity.
func (eu *EventUpdate) ClearCollect() *EventUpdate {
	eu.mutation.ClearCollect()
	return eu
}

// ClearTransfer clears the "transfer" edge to the UniswapV3Transfer entity.
func (eu *EventUpdate) ClearTransfer() *EventUpdate {
	eu.mutation.ClearTransfer()
	return eu
}

// ClearPoolCreated clears the "pool_created" edge to the UniswapV3PoolCreated entity.
func (eu *EventUpdate) ClearPoolCreated() *EventUpdate {
	eu.mutation.ClearPoolCreated()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := eu.mutation.Signature(); ok {
		if err := event.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf("ent: validator failed for field \"signature\": %w", err)}
		}
	}
	if v, ok := eu.mutation.Address(); ok {
		if err := event.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := eu.mutation.TxHash(); ok {
		if err := event.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf("ent: validator failed for field \"tx_hash\": %w", err)}
		}
	}
	if v, ok := eu.mutation.BlockHash(); ok {
		if err := event.BlockHashValidator(v); err != nil {
			return &ValidationError{Name: "block_hash", err: fmt.Errorf("ent: validator failed for field \"block_hash\": %w", err)}
		}
	}
	if v, ok := eu.mutation.Hash(); ok {
		if err := event.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	return nil
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := eu.mutation.Signature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSignature,
		})
	}
	if value, ok := eu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldAddress,
		})
	}
	if value, ok := eu.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldBlockNumber,
		})
	}
	if value, ok := eu.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldBlockNumber,
		})
	}
	if value, ok := eu.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldTxHash,
		})
	}
	if value, ok := eu.mutation.TxIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldTxIndex,
		})
	}
	if value, ok := eu.mutation.AddedTxIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldTxIndex,
		})
	}
	if value, ok := eu.mutation.BlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldBlockHash,
		})
	}
	if value, ok := eu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldIndex,
		})
	}
	if value, ok := eu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldIndex,
		})
	}
	if value, ok := eu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldHash,
		})
	}
	if eu.mutation.IncreaseLiquidityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.IncreaseLiquidityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.DecreaseLiquidityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.DecreaseLiquidityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.CollectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.CollectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.TransferCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.TransferIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.PoolCreatedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.PoolCreatedIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetName sets the "name" field.
func (euo *EventUpdateOne) SetName(s string) *EventUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetSignature sets the "signature" field.
func (euo *EventUpdateOne) SetSignature(s string) *EventUpdateOne {
	euo.mutation.SetSignature(s)
	return euo
}

// SetAddress sets the "address" field.
func (euo *EventUpdateOne) SetAddress(s string) *EventUpdateOne {
	euo.mutation.SetAddress(s)
	return euo
}

// SetBlockNumber sets the "block_number" field.
func (euo *EventUpdateOne) SetBlockNumber(u uint64) *EventUpdateOne {
	euo.mutation.ResetBlockNumber()
	euo.mutation.SetBlockNumber(u)
	return euo
}

// AddBlockNumber adds u to the "block_number" field.
func (euo *EventUpdateOne) AddBlockNumber(u uint64) *EventUpdateOne {
	euo.mutation.AddBlockNumber(u)
	return euo
}

// SetTxHash sets the "tx_hash" field.
func (euo *EventUpdateOne) SetTxHash(s string) *EventUpdateOne {
	euo.mutation.SetTxHash(s)
	return euo
}

// SetTxIndex sets the "tx_index" field.
func (euo *EventUpdateOne) SetTxIndex(u uint) *EventUpdateOne {
	euo.mutation.ResetTxIndex()
	euo.mutation.SetTxIndex(u)
	return euo
}

// AddTxIndex adds u to the "tx_index" field.
func (euo *EventUpdateOne) AddTxIndex(u uint) *EventUpdateOne {
	euo.mutation.AddTxIndex(u)
	return euo
}

// SetBlockHash sets the "block_hash" field.
func (euo *EventUpdateOne) SetBlockHash(s string) *EventUpdateOne {
	euo.mutation.SetBlockHash(s)
	return euo
}

// SetIndex sets the "index" field.
func (euo *EventUpdateOne) SetIndex(u uint) *EventUpdateOne {
	euo.mutation.ResetIndex()
	euo.mutation.SetIndex(u)
	return euo
}

// AddIndex adds u to the "index" field.
func (euo *EventUpdateOne) AddIndex(u uint) *EventUpdateOne {
	euo.mutation.AddIndex(u)
	return euo
}

// SetHash sets the "hash" field.
func (euo *EventUpdateOne) SetHash(s string) *EventUpdateOne {
	euo.mutation.SetHash(s)
	return euo
}

// SetIncreaseLiquidityID sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity by ID.
func (euo *EventUpdateOne) SetIncreaseLiquidityID(id int) *EventUpdateOne {
	euo.mutation.SetIncreaseLiquidityID(id)
	return euo
}

// SetNillableIncreaseLiquidityID sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableIncreaseLiquidityID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetIncreaseLiquidityID(*id)
	}
	return euo
}

// SetIncreaseLiquidity sets the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity.
func (euo *EventUpdateOne) SetIncreaseLiquidity(u *UniswapV3IncreaseLiqudity) *EventUpdateOne {
	return euo.SetIncreaseLiquidityID(u.ID)
}

// SetDecreaseLiquidityID sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity by ID.
func (euo *EventUpdateOne) SetDecreaseLiquidityID(id int) *EventUpdateOne {
	euo.mutation.SetDecreaseLiquidityID(id)
	return euo
}

// SetNillableDecreaseLiquidityID sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableDecreaseLiquidityID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetDecreaseLiquidityID(*id)
	}
	return euo
}

// SetDecreaseLiquidity sets the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity.
func (euo *EventUpdateOne) SetDecreaseLiquidity(u *UniswapV3DecreaseLiqudity) *EventUpdateOne {
	return euo.SetDecreaseLiquidityID(u.ID)
}

// SetCollectID sets the "collect" edge to the UniswapV3Collect entity by ID.
func (euo *EventUpdateOne) SetCollectID(id int) *EventUpdateOne {
	euo.mutation.SetCollectID(id)
	return euo
}

// SetNillableCollectID sets the "collect" edge to the UniswapV3Collect entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCollectID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetCollectID(*id)
	}
	return euo
}

// SetCollect sets the "collect" edge to the UniswapV3Collect entity.
func (euo *EventUpdateOne) SetCollect(u *UniswapV3Collect) *EventUpdateOne {
	return euo.SetCollectID(u.ID)
}

// SetTransferID sets the "transfer" edge to the UniswapV3Transfer entity by ID.
func (euo *EventUpdateOne) SetTransferID(id int) *EventUpdateOne {
	euo.mutation.SetTransferID(id)
	return euo
}

// SetNillableTransferID sets the "transfer" edge to the UniswapV3Transfer entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableTransferID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetTransferID(*id)
	}
	return euo
}

// SetTransfer sets the "transfer" edge to the UniswapV3Transfer entity.
func (euo *EventUpdateOne) SetTransfer(u *UniswapV3Transfer) *EventUpdateOne {
	return euo.SetTransferID(u.ID)
}

// SetPoolCreatedID sets the "pool_created" edge to the UniswapV3PoolCreated entity by ID.
func (euo *EventUpdateOne) SetPoolCreatedID(id int) *EventUpdateOne {
	euo.mutation.SetPoolCreatedID(id)
	return euo
}

// SetNillablePoolCreatedID sets the "pool_created" edge to the UniswapV3PoolCreated entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillablePoolCreatedID(id *int) *EventUpdateOne {
	if id != nil {
		euo = euo.SetPoolCreatedID(*id)
	}
	return euo
}

// SetPoolCreated sets the "pool_created" edge to the UniswapV3PoolCreated entity.
func (euo *EventUpdateOne) SetPoolCreated(u *UniswapV3PoolCreated) *EventUpdateOne {
	return euo.SetPoolCreatedID(u.ID)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearIncreaseLiquidity clears the "increase_liquidity" edge to the UniswapV3IncreaseLiqudity entity.
func (euo *EventUpdateOne) ClearIncreaseLiquidity() *EventUpdateOne {
	euo.mutation.ClearIncreaseLiquidity()
	return euo
}

// ClearDecreaseLiquidity clears the "decrease_liquidity" edge to the UniswapV3DecreaseLiqudity entity.
func (euo *EventUpdateOne) ClearDecreaseLiquidity() *EventUpdateOne {
	euo.mutation.ClearDecreaseLiquidity()
	return euo
}

// ClearCollect clears the "collect" edge to the UniswapV3Collect entity.
func (euo *EventUpdateOne) ClearCollect() *EventUpdateOne {
	euo.mutation.ClearCollect()
	return euo
}

// ClearTransfer clears the "transfer" edge to the UniswapV3Transfer entity.
func (euo *EventUpdateOne) ClearTransfer() *EventUpdateOne {
	euo.mutation.ClearTransfer()
	return euo
}

// ClearPoolCreated clears the "pool_created" edge to the UniswapV3PoolCreated entity.
func (euo *EventUpdateOne) ClearPoolCreated() *EventUpdateOne {
	euo.mutation.ClearPoolCreated()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := euo.mutation.Signature(); ok {
		if err := event.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf("ent: validator failed for field \"signature\": %w", err)}
		}
	}
	if v, ok := euo.mutation.Address(); ok {
		if err := event.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := euo.mutation.TxHash(); ok {
		if err := event.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf("ent: validator failed for field \"tx_hash\": %w", err)}
		}
	}
	if v, ok := euo.mutation.BlockHash(); ok {
		if err := event.BlockHashValidator(v); err != nil {
			return &ValidationError{Name: "block_hash", err: fmt.Errorf("ent: validator failed for field \"block_hash\": %w", err)}
		}
	}
	if v, ok := euo.mutation.Hash(); ok {
		if err := event.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	return nil
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Event.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := euo.mutation.Signature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSignature,
		})
	}
	if value, ok := euo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldAddress,
		})
	}
	if value, ok := euo.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldBlockNumber,
		})
	}
	if value, ok := euo.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldBlockNumber,
		})
	}
	if value, ok := euo.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldTxHash,
		})
	}
	if value, ok := euo.mutation.TxIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldTxIndex,
		})
	}
	if value, ok := euo.mutation.AddedTxIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldTxIndex,
		})
	}
	if value, ok := euo.mutation.BlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldBlockHash,
		})
	}
	if value, ok := euo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldIndex,
		})
	}
	if value, ok := euo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: event.FieldIndex,
		})
	}
	if value, ok := euo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldHash,
		})
	}
	if euo.mutation.IncreaseLiquidityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.IncreaseLiquidityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.DecreaseLiquidityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.DecreaseLiquidityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.CollectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.CollectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.TransferCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.TransferIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.PoolCreatedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.PoolCreatedIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
