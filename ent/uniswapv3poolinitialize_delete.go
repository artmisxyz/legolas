// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/legolas/ent/predicate"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolinitialize"
)

// UniswapV3PoolInitializeDelete is the builder for deleting a UniswapV3PoolInitialize entity.
type UniswapV3PoolInitializeDelete struct {
	config
	hooks    []Hook
	mutation *UniswapV3PoolInitializeMutation
}

// Where appends a list predicates to the UniswapV3PoolInitializeDelete builder.
func (uvid *UniswapV3PoolInitializeDelete) Where(ps ...predicate.UniswapV3PoolInitialize) *UniswapV3PoolInitializeDelete {
	uvid.mutation.Where(ps...)
	return uvid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uvid *UniswapV3PoolInitializeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uvid.hooks) == 0 {
		affected, err = uvid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3PoolInitializeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uvid.mutation = mutation
			affected, err = uvid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uvid.hooks) - 1; i >= 0; i-- {
			if uvid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvid *UniswapV3PoolInitializeDelete) ExecX(ctx context.Context) int {
	n, err := uvid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uvid *UniswapV3PoolInitializeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: uniswapv3poolinitialize.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolinitialize.FieldID,
			},
		},
	}
	if ps := uvid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uvid.driver, _spec)
}

// UniswapV3PoolInitializeDeleteOne is the builder for deleting a single UniswapV3PoolInitialize entity.
type UniswapV3PoolInitializeDeleteOne struct {
	uvid *UniswapV3PoolInitializeDelete
}

// Exec executes the deletion query.
func (uvido *UniswapV3PoolInitializeDeleteOne) Exec(ctx context.Context) error {
	n, err := uvido.uvid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{uniswapv3poolinitialize.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uvido *UniswapV3PoolInitializeDeleteOne) ExecX(ctx context.Context) {
	uvido.uvid.ExecX(ctx)
}
