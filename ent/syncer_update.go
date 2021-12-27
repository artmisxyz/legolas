// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/legolas/ent/predicate"
	"github.com/artmisxyz/legolas/ent/syncer"
)

// SyncerUpdate is the builder for updating Syncer entities.
type SyncerUpdate struct {
	config
	hooks    []Hook
	mutation *SyncerMutation
}

// Where appends a list predicates to the SyncerUpdate builder.
func (su *SyncerUpdate) Where(ps ...predicate.Syncer) *SyncerUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SyncerUpdate) SetName(s string) *SyncerUpdate {
	su.mutation.SetName(s)
	return su
}

// SetStart sets the "start" field.
func (su *SyncerUpdate) SetStart(u uint64) *SyncerUpdate {
	su.mutation.ResetStart()
	su.mutation.SetStart(u)
	return su
}

// AddStart adds u to the "start" field.
func (su *SyncerUpdate) AddStart(u uint64) *SyncerUpdate {
	su.mutation.AddStart(u)
	return su
}

// SetFinish sets the "finish" field.
func (su *SyncerUpdate) SetFinish(u uint64) *SyncerUpdate {
	su.mutation.ResetFinish()
	su.mutation.SetFinish(u)
	return su
}

// AddFinish adds u to the "finish" field.
func (su *SyncerUpdate) AddFinish(u uint64) *SyncerUpdate {
	su.mutation.AddFinish(u)
	return su
}

// SetCurrent sets the "current" field.
func (su *SyncerUpdate) SetCurrent(u uint64) *SyncerUpdate {
	su.mutation.ResetCurrent()
	su.mutation.SetCurrent(u)
	return su
}

// AddCurrent adds u to the "current" field.
func (su *SyncerUpdate) AddCurrent(u uint64) *SyncerUpdate {
	su.mutation.AddCurrent(u)
	return su
}

// Mutation returns the SyncerMutation object of the builder.
func (su *SyncerUpdate) Mutation() *SyncerMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SyncerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SyncerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SyncerUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SyncerUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SyncerUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SyncerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   syncer.Table,
			Columns: syncer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: syncer.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syncer.FieldName,
		})
	}
	if value, ok := su.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldStart,
		})
	}
	if value, ok := su.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldStart,
		})
	}
	if value, ok := su.mutation.Finish(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldFinish,
		})
	}
	if value, ok := su.mutation.AddedFinish(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldFinish,
		})
	}
	if value, ok := su.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldCurrent,
		})
	}
	if value, ok := su.mutation.AddedCurrent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldCurrent,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syncer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SyncerUpdateOne is the builder for updating a single Syncer entity.
type SyncerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SyncerMutation
}

// SetName sets the "name" field.
func (suo *SyncerUpdateOne) SetName(s string) *SyncerUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetStart sets the "start" field.
func (suo *SyncerUpdateOne) SetStart(u uint64) *SyncerUpdateOne {
	suo.mutation.ResetStart()
	suo.mutation.SetStart(u)
	return suo
}

// AddStart adds u to the "start" field.
func (suo *SyncerUpdateOne) AddStart(u uint64) *SyncerUpdateOne {
	suo.mutation.AddStart(u)
	return suo
}

// SetFinish sets the "finish" field.
func (suo *SyncerUpdateOne) SetFinish(u uint64) *SyncerUpdateOne {
	suo.mutation.ResetFinish()
	suo.mutation.SetFinish(u)
	return suo
}

// AddFinish adds u to the "finish" field.
func (suo *SyncerUpdateOne) AddFinish(u uint64) *SyncerUpdateOne {
	suo.mutation.AddFinish(u)
	return suo
}

// SetCurrent sets the "current" field.
func (suo *SyncerUpdateOne) SetCurrent(u uint64) *SyncerUpdateOne {
	suo.mutation.ResetCurrent()
	suo.mutation.SetCurrent(u)
	return suo
}

// AddCurrent adds u to the "current" field.
func (suo *SyncerUpdateOne) AddCurrent(u uint64) *SyncerUpdateOne {
	suo.mutation.AddCurrent(u)
	return suo
}

// Mutation returns the SyncerMutation object of the builder.
func (suo *SyncerUpdateOne) Mutation() *SyncerMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SyncerUpdateOne) Select(field string, fields ...string) *SyncerUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Syncer entity.
func (suo *SyncerUpdateOne) Save(ctx context.Context) (*Syncer, error) {
	var (
		err  error
		node *Syncer
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SyncerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SyncerUpdateOne) SaveX(ctx context.Context) *Syncer {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SyncerUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SyncerUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SyncerUpdateOne) sqlSave(ctx context.Context) (_node *Syncer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   syncer.Table,
			Columns: syncer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: syncer.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Syncer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, syncer.FieldID)
		for _, f := range fields {
			if !syncer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != syncer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syncer.FieldName,
		})
	}
	if value, ok := suo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldStart,
		})
	}
	if value, ok := suo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldStart,
		})
	}
	if value, ok := suo.mutation.Finish(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldFinish,
		})
	}
	if value, ok := suo.mutation.AddedFinish(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldFinish,
		})
	}
	if value, ok := suo.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldCurrent,
		})
	}
	if value, ok := suo.mutation.AddedCurrent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: syncer.FieldCurrent,
		})
	}
	_node = &Syncer{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syncer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}