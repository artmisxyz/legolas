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
	"github.com/artmisxyz/legolas/ent/uniswapv3poolburn"
)

// UniswapV3PoolBurnCreate is the builder for creating a UniswapV3PoolBurn entity.
type UniswapV3PoolBurnCreate struct {
	config
	mutation *UniswapV3PoolBurnMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOwner sets the "owner" field.
func (uvbc *UniswapV3PoolBurnCreate) SetOwner(s string) *UniswapV3PoolBurnCreate {
	uvbc.mutation.SetOwner(s)
	return uvbc
}

// SetTickLower sets the "tick_lower" field.
func (uvbc *UniswapV3PoolBurnCreate) SetTickLower(s string) *UniswapV3PoolBurnCreate {
	uvbc.mutation.SetTickLower(s)
	return uvbc
}

// SetTickUpper sets the "tick_upper" field.
func (uvbc *UniswapV3PoolBurnCreate) SetTickUpper(s string) *UniswapV3PoolBurnCreate {
	uvbc.mutation.SetTickUpper(s)
	return uvbc
}

// SetAmount sets the "amount" field.
func (uvbc *UniswapV3PoolBurnCreate) SetAmount(s string) *UniswapV3PoolBurnCreate {
	uvbc.mutation.SetAmount(s)
	return uvbc
}

// SetAmount0 sets the "amount0" field.
func (uvbc *UniswapV3PoolBurnCreate) SetAmount0(s string) *UniswapV3PoolBurnCreate {
	uvbc.mutation.SetAmount0(s)
	return uvbc
}

// SetAmount1 sets the "amount1" field.
func (uvbc *UniswapV3PoolBurnCreate) SetAmount1(s string) *UniswapV3PoolBurnCreate {
	uvbc.mutation.SetAmount1(s)
	return uvbc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (uvbc *UniswapV3PoolBurnCreate) SetEventID(id int) *UniswapV3PoolBurnCreate {
	uvbc.mutation.SetEventID(id)
	return uvbc
}

// SetEvent sets the "event" edge to the Event entity.
func (uvbc *UniswapV3PoolBurnCreate) SetEvent(e *Event) *UniswapV3PoolBurnCreate {
	return uvbc.SetEventID(e.ID)
}

// Mutation returns the UniswapV3PoolBurnMutation object of the builder.
func (uvbc *UniswapV3PoolBurnCreate) Mutation() *UniswapV3PoolBurnMutation {
	return uvbc.mutation
}

// Save creates the UniswapV3PoolBurn in the database.
func (uvbc *UniswapV3PoolBurnCreate) Save(ctx context.Context) (*UniswapV3PoolBurn, error) {
	var (
		err  error
		node *UniswapV3PoolBurn
	)
	if len(uvbc.hooks) == 0 {
		if err = uvbc.check(); err != nil {
			return nil, err
		}
		node, err = uvbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3PoolBurnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uvbc.check(); err != nil {
				return nil, err
			}
			uvbc.mutation = mutation
			if node, err = uvbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uvbc.hooks) - 1; i >= 0; i-- {
			if uvbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uvbc *UniswapV3PoolBurnCreate) SaveX(ctx context.Context) *UniswapV3PoolBurn {
	v, err := uvbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvbc *UniswapV3PoolBurnCreate) Exec(ctx context.Context) error {
	_, err := uvbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvbc *UniswapV3PoolBurnCreate) ExecX(ctx context.Context) {
	if err := uvbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvbc *UniswapV3PoolBurnCreate) check() error {
	if _, ok := uvbc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "owner"`)}
	}
	if _, ok := uvbc.mutation.TickLower(); !ok {
		return &ValidationError{Name: "tick_lower", err: errors.New(`ent: missing required field "tick_lower"`)}
	}
	if _, ok := uvbc.mutation.TickUpper(); !ok {
		return &ValidationError{Name: "tick_upper", err: errors.New(`ent: missing required field "tick_upper"`)}
	}
	if _, ok := uvbc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if _, ok := uvbc.mutation.Amount0(); !ok {
		return &ValidationError{Name: "amount0", err: errors.New(`ent: missing required field "amount0"`)}
	}
	if _, ok := uvbc.mutation.Amount1(); !ok {
		return &ValidationError{Name: "amount1", err: errors.New(`ent: missing required field "amount1"`)}
	}
	if _, ok := uvbc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New("ent: missing required edge \"event\"")}
	}
	return nil
}

func (uvbc *UniswapV3PoolBurnCreate) sqlSave(ctx context.Context) (*UniswapV3PoolBurn, error) {
	_node, _spec := uvbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uvbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uvbc *UniswapV3PoolBurnCreate) createSpec() (*UniswapV3PoolBurn, *sqlgraph.CreateSpec) {
	var (
		_node = &UniswapV3PoolBurn{config: uvbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: uniswapv3poolburn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolburn.FieldID,
			},
		}
	)
	_spec.OnConflict = uvbc.conflict
	if value, ok := uvbc.mutation.Owner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolburn.FieldOwner,
		})
		_node.Owner = value
	}
	if value, ok := uvbc.mutation.TickLower(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolburn.FieldTickLower,
		})
		_node.TickLower = value
	}
	if value, ok := uvbc.mutation.TickUpper(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolburn.FieldTickUpper,
		})
		_node.TickUpper = value
	}
	if value, ok := uvbc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolburn.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := uvbc.mutation.Amount0(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolburn.FieldAmount0,
		})
		_node.Amount0 = value
	}
	if value, ok := uvbc.mutation.Amount1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: uniswapv3poolburn.FieldAmount1,
		})
		_node.Amount1 = value
	}
	if nodes := uvbc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   uniswapv3poolburn.EventTable,
			Columns: []string{uniswapv3poolburn.EventColumn},
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
//	client.UniswapV3PoolBurn.Create().
//		SetOwner(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UniswapV3PoolBurnUpsert) {
//			SetOwner(v+v).
//		}).
//		Exec(ctx)
//
func (uvbc *UniswapV3PoolBurnCreate) OnConflict(opts ...sql.ConflictOption) *UniswapV3PoolBurnUpsertOne {
	uvbc.conflict = opts
	return &UniswapV3PoolBurnUpsertOne{
		create: uvbc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UniswapV3PoolBurn.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uvbc *UniswapV3PoolBurnCreate) OnConflictColumns(columns ...string) *UniswapV3PoolBurnUpsertOne {
	uvbc.conflict = append(uvbc.conflict, sql.ConflictColumns(columns...))
	return &UniswapV3PoolBurnUpsertOne{
		create: uvbc,
	}
}

type (
	// UniswapV3PoolBurnUpsertOne is the builder for "upsert"-ing
	//  one UniswapV3PoolBurn node.
	UniswapV3PoolBurnUpsertOne struct {
		create *UniswapV3PoolBurnCreate
	}

	// UniswapV3PoolBurnUpsert is the "OnConflict" setter.
	UniswapV3PoolBurnUpsert struct {
		*sql.UpdateSet
	}
)

// SetOwner sets the "owner" field.
func (u *UniswapV3PoolBurnUpsert) SetOwner(v string) *UniswapV3PoolBurnUpsert {
	u.Set(uniswapv3poolburn.FieldOwner, v)
	return u
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsert) UpdateOwner() *UniswapV3PoolBurnUpsert {
	u.SetExcluded(uniswapv3poolburn.FieldOwner)
	return u
}

// SetTickLower sets the "tick_lower" field.
func (u *UniswapV3PoolBurnUpsert) SetTickLower(v string) *UniswapV3PoolBurnUpsert {
	u.Set(uniswapv3poolburn.FieldTickLower, v)
	return u
}

// UpdateTickLower sets the "tick_lower" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsert) UpdateTickLower() *UniswapV3PoolBurnUpsert {
	u.SetExcluded(uniswapv3poolburn.FieldTickLower)
	return u
}

// SetTickUpper sets the "tick_upper" field.
func (u *UniswapV3PoolBurnUpsert) SetTickUpper(v string) *UniswapV3PoolBurnUpsert {
	u.Set(uniswapv3poolburn.FieldTickUpper, v)
	return u
}

// UpdateTickUpper sets the "tick_upper" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsert) UpdateTickUpper() *UniswapV3PoolBurnUpsert {
	u.SetExcluded(uniswapv3poolburn.FieldTickUpper)
	return u
}

// SetAmount sets the "amount" field.
func (u *UniswapV3PoolBurnUpsert) SetAmount(v string) *UniswapV3PoolBurnUpsert {
	u.Set(uniswapv3poolburn.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsert) UpdateAmount() *UniswapV3PoolBurnUpsert {
	u.SetExcluded(uniswapv3poolburn.FieldAmount)
	return u
}

// SetAmount0 sets the "amount0" field.
func (u *UniswapV3PoolBurnUpsert) SetAmount0(v string) *UniswapV3PoolBurnUpsert {
	u.Set(uniswapv3poolburn.FieldAmount0, v)
	return u
}

// UpdateAmount0 sets the "amount0" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsert) UpdateAmount0() *UniswapV3PoolBurnUpsert {
	u.SetExcluded(uniswapv3poolburn.FieldAmount0)
	return u
}

// SetAmount1 sets the "amount1" field.
func (u *UniswapV3PoolBurnUpsert) SetAmount1(v string) *UniswapV3PoolBurnUpsert {
	u.Set(uniswapv3poolburn.FieldAmount1, v)
	return u
}

// UpdateAmount1 sets the "amount1" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsert) UpdateAmount1() *UniswapV3PoolBurnUpsert {
	u.SetExcluded(uniswapv3poolburn.FieldAmount1)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UniswapV3PoolBurn.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *UniswapV3PoolBurnUpsertOne) UpdateNewValues() *UniswapV3PoolBurnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UniswapV3PoolBurn.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UniswapV3PoolBurnUpsertOne) Ignore() *UniswapV3PoolBurnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UniswapV3PoolBurnUpsertOne) DoNothing() *UniswapV3PoolBurnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UniswapV3PoolBurnCreate.OnConflict
// documentation for more info.
func (u *UniswapV3PoolBurnUpsertOne) Update(set func(*UniswapV3PoolBurnUpsert)) *UniswapV3PoolBurnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UniswapV3PoolBurnUpsert{UpdateSet: update})
	}))
	return u
}

// SetOwner sets the "owner" field.
func (u *UniswapV3PoolBurnUpsertOne) SetOwner(v string) *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertOne) UpdateOwner() *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateOwner()
	})
}

// SetTickLower sets the "tick_lower" field.
func (u *UniswapV3PoolBurnUpsertOne) SetTickLower(v string) *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetTickLower(v)
	})
}

// UpdateTickLower sets the "tick_lower" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertOne) UpdateTickLower() *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateTickLower()
	})
}

// SetTickUpper sets the "tick_upper" field.
func (u *UniswapV3PoolBurnUpsertOne) SetTickUpper(v string) *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetTickUpper(v)
	})
}

// UpdateTickUpper sets the "tick_upper" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertOne) UpdateTickUpper() *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateTickUpper()
	})
}

// SetAmount sets the "amount" field.
func (u *UniswapV3PoolBurnUpsertOne) SetAmount(v string) *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertOne) UpdateAmount() *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateAmount()
	})
}

// SetAmount0 sets the "amount0" field.
func (u *UniswapV3PoolBurnUpsertOne) SetAmount0(v string) *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetAmount0(v)
	})
}

// UpdateAmount0 sets the "amount0" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertOne) UpdateAmount0() *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateAmount0()
	})
}

// SetAmount1 sets the "amount1" field.
func (u *UniswapV3PoolBurnUpsertOne) SetAmount1(v string) *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetAmount1(v)
	})
}

// UpdateAmount1 sets the "amount1" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertOne) UpdateAmount1() *UniswapV3PoolBurnUpsertOne {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateAmount1()
	})
}

// Exec executes the query.
func (u *UniswapV3PoolBurnUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UniswapV3PoolBurnCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UniswapV3PoolBurnUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UniswapV3PoolBurnUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UniswapV3PoolBurnUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UniswapV3PoolBurnCreateBulk is the builder for creating many UniswapV3PoolBurn entities in bulk.
type UniswapV3PoolBurnCreateBulk struct {
	config
	builders []*UniswapV3PoolBurnCreate
	conflict []sql.ConflictOption
}

// Save creates the UniswapV3PoolBurn entities in the database.
func (uvbcb *UniswapV3PoolBurnCreateBulk) Save(ctx context.Context) ([]*UniswapV3PoolBurn, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uvbcb.builders))
	nodes := make([]*UniswapV3PoolBurn, len(uvbcb.builders))
	mutators := make([]Mutator, len(uvbcb.builders))
	for i := range uvbcb.builders {
		func(i int, root context.Context) {
			builder := uvbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UniswapV3PoolBurnMutation)
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
					_, err = mutators[i+1].Mutate(root, uvbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uvbcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uvbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uvbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uvbcb *UniswapV3PoolBurnCreateBulk) SaveX(ctx context.Context) []*UniswapV3PoolBurn {
	v, err := uvbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uvbcb *UniswapV3PoolBurnCreateBulk) Exec(ctx context.Context) error {
	_, err := uvbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvbcb *UniswapV3PoolBurnCreateBulk) ExecX(ctx context.Context) {
	if err := uvbcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UniswapV3PoolBurn.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UniswapV3PoolBurnUpsert) {
//			SetOwner(v+v).
//		}).
//		Exec(ctx)
//
func (uvbcb *UniswapV3PoolBurnCreateBulk) OnConflict(opts ...sql.ConflictOption) *UniswapV3PoolBurnUpsertBulk {
	uvbcb.conflict = opts
	return &UniswapV3PoolBurnUpsertBulk{
		create: uvbcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UniswapV3PoolBurn.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uvbcb *UniswapV3PoolBurnCreateBulk) OnConflictColumns(columns ...string) *UniswapV3PoolBurnUpsertBulk {
	uvbcb.conflict = append(uvbcb.conflict, sql.ConflictColumns(columns...))
	return &UniswapV3PoolBurnUpsertBulk{
		create: uvbcb,
	}
}

// UniswapV3PoolBurnUpsertBulk is the builder for "upsert"-ing
// a bulk of UniswapV3PoolBurn nodes.
type UniswapV3PoolBurnUpsertBulk struct {
	create *UniswapV3PoolBurnCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UniswapV3PoolBurn.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *UniswapV3PoolBurnUpsertBulk) UpdateNewValues() *UniswapV3PoolBurnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UniswapV3PoolBurn.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UniswapV3PoolBurnUpsertBulk) Ignore() *UniswapV3PoolBurnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UniswapV3PoolBurnUpsertBulk) DoNothing() *UniswapV3PoolBurnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UniswapV3PoolBurnCreateBulk.OnConflict
// documentation for more info.
func (u *UniswapV3PoolBurnUpsertBulk) Update(set func(*UniswapV3PoolBurnUpsert)) *UniswapV3PoolBurnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UniswapV3PoolBurnUpsert{UpdateSet: update})
	}))
	return u
}

// SetOwner sets the "owner" field.
func (u *UniswapV3PoolBurnUpsertBulk) SetOwner(v string) *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertBulk) UpdateOwner() *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateOwner()
	})
}

// SetTickLower sets the "tick_lower" field.
func (u *UniswapV3PoolBurnUpsertBulk) SetTickLower(v string) *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetTickLower(v)
	})
}

// UpdateTickLower sets the "tick_lower" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertBulk) UpdateTickLower() *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateTickLower()
	})
}

// SetTickUpper sets the "tick_upper" field.
func (u *UniswapV3PoolBurnUpsertBulk) SetTickUpper(v string) *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetTickUpper(v)
	})
}

// UpdateTickUpper sets the "tick_upper" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertBulk) UpdateTickUpper() *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateTickUpper()
	})
}

// SetAmount sets the "amount" field.
func (u *UniswapV3PoolBurnUpsertBulk) SetAmount(v string) *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertBulk) UpdateAmount() *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateAmount()
	})
}

// SetAmount0 sets the "amount0" field.
func (u *UniswapV3PoolBurnUpsertBulk) SetAmount0(v string) *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetAmount0(v)
	})
}

// UpdateAmount0 sets the "amount0" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertBulk) UpdateAmount0() *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateAmount0()
	})
}

// SetAmount1 sets the "amount1" field.
func (u *UniswapV3PoolBurnUpsertBulk) SetAmount1(v string) *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.SetAmount1(v)
	})
}

// UpdateAmount1 sets the "amount1" field to the value that was provided on create.
func (u *UniswapV3PoolBurnUpsertBulk) UpdateAmount1() *UniswapV3PoolBurnUpsertBulk {
	return u.Update(func(s *UniswapV3PoolBurnUpsert) {
		s.UpdateAmount1()
	})
}

// Exec executes the query.
func (u *UniswapV3PoolBurnUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UniswapV3PoolBurnCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UniswapV3PoolBurnCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UniswapV3PoolBurnUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
