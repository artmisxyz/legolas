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
	"github.com/artmisxyz/legolas/ent/uniswapv3collect"
)

// UniswapV3CollectCreate is the builder for creating a UniswapV3Collect entity.
type UniswapV3CollectCreate struct {
	config
	mutation *UniswapV3CollectMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTokenID sets the "token_id" field.
func (uvc *UniswapV3CollectCreate) SetTokenID(s string) *UniswapV3CollectCreate {
	uvc.mutation.SetTokenID(s)
	return uvc
}

// SetRecipient sets the "recipient" field.
func (uvc *UniswapV3CollectCreate) SetRecipient(s string) *UniswapV3CollectCreate {
	uvc.mutation.SetRecipient(s)
	return uvc
}

// SetAmount0 sets the "amount0" field.
func (uvc *UniswapV3CollectCreate) SetAmount0(s string) *UniswapV3CollectCreate {
	uvc.mutation.SetAmount0(s)
	return uvc
}

// SetAmount1 sets the "amount1" field.
func (uvc *UniswapV3CollectCreate) SetAmount1(s string) *UniswapV3CollectCreate {
	uvc.mutation.SetAmount1(s)
	return uvc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvc *UniswapV3CollectCreate) SetEventID(id int) *UniswapV3CollectCreate {
	uvc.mutation.SetEventID(id)
	return uvc
}

// SetEvent sets the "event" edge to the Event entity.
func (uvc *UniswapV3CollectCreate) SetEvent(e *Event) *UniswapV3CollectCreate {
	return uvc.SetEventID(e.ID)
}

// Mutation returns the UniswapV3CollectMutation object of the builder.
func (uvc *UniswapV3CollectCreate) Mutation() *UniswapV3CollectMutation {
	return uvc.mutation
}

// Save creates the UniswapV3Collect in the database.
func (uvc *UniswapV3CollectCreate) Save(ctx context.Context) (*UniswapV3Collect, error) {
	var (
		err  error
		node *UniswapV3Collect
	)
	if len(uvc.hooks) == 0 {
		if err = uvc.check(); err != nil {
			return nil, err
		}
		node, err = uvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3CollectMutation)
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
func (uvc *UniswapV3CollectCreate) SaveX(ctx context.Context) *UniswapV3Collect {
	v, err := uvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvc *UniswapV3CollectCreate) Exec(ctx context.Context) error {
	_, err := uvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvc *UniswapV3CollectCreate) ExecX(ctx context.Context) {
	if err := uvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvc *UniswapV3CollectCreate) check() error {
	if _, ok := uvc.mutation.TokenID(); !ok {
		return &ValidationError{Name: "token_id", err: errors.New(`ent: missing required field "token_id"`)}
	}
	if _, ok := uvc.mutation.Recipient(); !ok {
		return &ValidationError{Name: "recipient", err: errors.New(`ent: missing required field "recipient"`)}
	}
	if _, ok := uvc.mutation.Amount0(); !ok {
		return &ValidationError{Name: "amount0", err: errors.New(`ent: missing required field "amount0"`)}
	}
	if _, ok := uvc.mutation.Amount1(); !ok {
		return &ValidationError{Name: "amount1", err: errors.New(`ent: missing required field "amount1"`)}
	}
	if _, ok := uvc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New("ent: missing required edge \"event\"")}
	}
	return nil
}

func (uvc *UniswapV3CollectCreate) sqlSave(ctx context.Context) (*UniswapV3Collect, error) {
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

func (uvc *UniswapV3CollectCreate) createSpec() (*UniswapV3Collect, *sqlgraph.CreateSpec) {
	var (
		_node = &UniswapV3Collect{config: uvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: uniswapv3collect.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3collect.FieldID,
			},
		}
	)
	_spec.OnConflict = uvc.conflict
	if value, ok := uvc.mutation.TokenID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3collect.FieldTokenID,
		})
		_node.TokenID = value
	}
	if value, ok := uvc.mutation.Recipient(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3collect.FieldRecipient,
		})
		_node.Recipient = value
	}
	if value, ok := uvc.mutation.Amount0(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3collect.FieldAmount0,
		})
		_node.Amount0 = value
	}
	if value, ok := uvc.mutation.Amount1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3collect.FieldAmount1,
		})
		_node.Amount1 = value
	}
	if nodes := uvc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3collect.EventTable,
			Columns: []string{uniswapv3collect.EventColumn},
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UniswapV3Collect.Create().
//		SetTokenID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UniswapV3CollectUpsert) {
//			SetTokenID(v+v).
//		}).
//		Exec(ctx)
//
func (uvc *UniswapV3CollectCreate) OnConflict(opts ...sql.ConflictOption) *UniswapV3CollectUpsertOne {
	uvc.conflict = opts
	return &UniswapV3CollectUpsertOne{
		create: uvc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UniswapV3Collect.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uvc *UniswapV3CollectCreate) OnConflictColumns(columns ...string) *UniswapV3CollectUpsertOne {
	uvc.conflict = append(uvc.conflict, sql.ConflictColumns(columns...))
	return &UniswapV3CollectUpsertOne{
		create: uvc,
	}
}

type (
	// UniswapV3CollectUpsertOne is the builder for "upsert"-ing
	//  one UniswapV3Collect node.
	UniswapV3CollectUpsertOne struct {
		create *UniswapV3CollectCreate
	}

	// UniswapV3CollectUpsert is the "OnConflict" setter.
	UniswapV3CollectUpsert struct {
		*sql.UpdateSet
	}
)

// SetTokenID sets the "token_id" field.
func (u *UniswapV3CollectUpsert) SetTokenID(v string) *UniswapV3CollectUpsert {
	u.Set(uniswapv3collect.FieldTokenID, v)
	return u
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *UniswapV3CollectUpsert) UpdateTokenID() *UniswapV3CollectUpsert {
	u.SetExcluded(uniswapv3collect.FieldTokenID)
	return u
}

// SetRecipient sets the "recipient" field.
func (u *UniswapV3CollectUpsert) SetRecipient(v string) *UniswapV3CollectUpsert {
	u.Set(uniswapv3collect.FieldRecipient, v)
	return u
}

// UpdateRecipient sets the "recipient" field to the value that was provided on create.
func (u *UniswapV3CollectUpsert) UpdateRecipient() *UniswapV3CollectUpsert {
	u.SetExcluded(uniswapv3collect.FieldRecipient)
	return u
}

// SetAmount0 sets the "amount0" field.
func (u *UniswapV3CollectUpsert) SetAmount0(v string) *UniswapV3CollectUpsert {
	u.Set(uniswapv3collect.FieldAmount0, v)
	return u
}

// UpdateAmount0 sets the "amount0" field to the value that was provided on create.
func (u *UniswapV3CollectUpsert) UpdateAmount0() *UniswapV3CollectUpsert {
	u.SetExcluded(uniswapv3collect.FieldAmount0)
	return u
}

// SetAmount1 sets the "amount1" field.
func (u *UniswapV3CollectUpsert) SetAmount1(v string) *UniswapV3CollectUpsert {
	u.Set(uniswapv3collect.FieldAmount1, v)
	return u
}

// UpdateAmount1 sets the "amount1" field to the value that was provided on create.
func (u *UniswapV3CollectUpsert) UpdateAmount1() *UniswapV3CollectUpsert {
	u.SetExcluded(uniswapv3collect.FieldAmount1)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UniswapV3Collect.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *UniswapV3CollectUpsertOne) UpdateNewValues() *UniswapV3CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UniswapV3Collect.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UniswapV3CollectUpsertOne) Ignore() *UniswapV3CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UniswapV3CollectUpsertOne) DoNothing() *UniswapV3CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UniswapV3CollectCreate.OnConflict
// documentation for more info.
func (u *UniswapV3CollectUpsertOne) Update(set func(*UniswapV3CollectUpsert)) *UniswapV3CollectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UniswapV3CollectUpsert{UpdateSet: update})
	}))
	return u
}

// SetTokenID sets the "token_id" field.
func (u *UniswapV3CollectUpsertOne) SetTokenID(v string) *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetTokenID(v)
	})
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertOne) UpdateTokenID() *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateTokenID()
	})
}

// SetRecipient sets the "recipient" field.
func (u *UniswapV3CollectUpsertOne) SetRecipient(v string) *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetRecipient(v)
	})
}

// UpdateRecipient sets the "recipient" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertOne) UpdateRecipient() *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateRecipient()
	})
}

// SetAmount0 sets the "amount0" field.
func (u *UniswapV3CollectUpsertOne) SetAmount0(v string) *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetAmount0(v)
	})
}

// UpdateAmount0 sets the "amount0" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertOne) UpdateAmount0() *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateAmount0()
	})
}

// SetAmount1 sets the "amount1" field.
func (u *UniswapV3CollectUpsertOne) SetAmount1(v string) *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetAmount1(v)
	})
}

// UpdateAmount1 sets the "amount1" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertOne) UpdateAmount1() *UniswapV3CollectUpsertOne {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateAmount1()
	})
}

// Exec executes the query.
func (u *UniswapV3CollectUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UniswapV3CollectCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UniswapV3CollectUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UniswapV3CollectUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UniswapV3CollectUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UniswapV3CollectCreateBulk is the builder for creating many UniswapV3Collect entities in bulk.
type UniswapV3CollectCreateBulk struct {
	config
	builders []*UniswapV3CollectCreate
	conflict []sql.ConflictOption
}

// Save creates the UniswapV3Collect entities in the database.
func (uvcb *UniswapV3CollectCreateBulk) Save(ctx context.Context) ([]*UniswapV3Collect, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uvcb.builders))
	nodes := make([]*UniswapV3Collect, len(uvcb.builders))
	mutators := make([]Mutator, len(uvcb.builders))
	for i := range uvcb.builders {
		func(i int, root context.Context) {
			builder := uvcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UniswapV3CollectMutation)
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
					spec.OnConflict = uvcb.conflict
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
func (uvcb *UniswapV3CollectCreateBulk) SaveX(ctx context.Context) []*UniswapV3Collect {
	v, err := uvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvcb *UniswapV3CollectCreateBulk) Exec(ctx context.Context) error {
	_, err := uvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvcb *UniswapV3CollectCreateBulk) ExecX(ctx context.Context) {
	if err := uvcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UniswapV3Collect.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UniswapV3CollectUpsert) {
//			SetTokenID(v+v).
//		}).
//		Exec(ctx)
//
func (uvcb *UniswapV3CollectCreateBulk) OnConflict(opts ...sql.ConflictOption) *UniswapV3CollectUpsertBulk {
	uvcb.conflict = opts
	return &UniswapV3CollectUpsertBulk{
		create: uvcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UniswapV3Collect.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uvcb *UniswapV3CollectCreateBulk) OnConflictColumns(columns ...string) *UniswapV3CollectUpsertBulk {
	uvcb.conflict = append(uvcb.conflict, sql.ConflictColumns(columns...))
	return &UniswapV3CollectUpsertBulk{
		create: uvcb,
	}
}

// UniswapV3CollectUpsertBulk is the builder for "upsert"-ing
// a bulk of UniswapV3Collect nodes.
type UniswapV3CollectUpsertBulk struct {
	create *UniswapV3CollectCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UniswapV3Collect.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *UniswapV3CollectUpsertBulk) UpdateNewValues() *UniswapV3CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UniswapV3Collect.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UniswapV3CollectUpsertBulk) Ignore() *UniswapV3CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UniswapV3CollectUpsertBulk) DoNothing() *UniswapV3CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UniswapV3CollectCreateBulk.OnConflict
// documentation for more info.
func (u *UniswapV3CollectUpsertBulk) Update(set func(*UniswapV3CollectUpsert)) *UniswapV3CollectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UniswapV3CollectUpsert{UpdateSet: update})
	}))
	return u
}

// SetTokenID sets the "token_id" field.
func (u *UniswapV3CollectUpsertBulk) SetTokenID(v string) *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetTokenID(v)
	})
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertBulk) UpdateTokenID() *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateTokenID()
	})
}

// SetRecipient sets the "recipient" field.
func (u *UniswapV3CollectUpsertBulk) SetRecipient(v string) *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetRecipient(v)
	})
}

// UpdateRecipient sets the "recipient" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertBulk) UpdateRecipient() *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateRecipient()
	})
}

// SetAmount0 sets the "amount0" field.
func (u *UniswapV3CollectUpsertBulk) SetAmount0(v string) *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetAmount0(v)
	})
}

// UpdateAmount0 sets the "amount0" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertBulk) UpdateAmount0() *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateAmount0()
	})
}

// SetAmount1 sets the "amount1" field.
func (u *UniswapV3CollectUpsertBulk) SetAmount1(v string) *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.SetAmount1(v)
	})
}

// UpdateAmount1 sets the "amount1" field to the value that was provided on create.
func (u *UniswapV3CollectUpsertBulk) UpdateAmount1() *UniswapV3CollectUpsertBulk {
	return u.Update(func(s *UniswapV3CollectUpsert) {
		s.UpdateAmount1()
	})
}

// Exec executes the query.
func (u *UniswapV3CollectUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UniswapV3CollectCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UniswapV3CollectCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UniswapV3CollectUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
