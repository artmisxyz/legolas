// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artmisxyz/legolas/ent/predicate"
	"github.com/artmisxyz/legolas/ent/uniswapv3poolmint"
)

// UniswapV3PoolMintDelete is the builder for deleting a UniswapV3PoolMint entity.
type UniswapV3PoolMintDelete struct {
	config
	hooks    []Hook
	mutation *UniswapV3PoolMintMutation
}

// Where appends a list predicates to the UniswapV3PoolMintDelete builder.
func (uvmd *UniswapV3PoolMintDelete) Where(ps ...predicate.UniswapV3PoolMint) *UniswapV3PoolMintDelete {
	uvmd.mutation.Where(ps...)
	return uvmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uvmd *UniswapV3PoolMintDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uvmd.hooks) == 0 {
		affected, err = uvmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UniswapV3PoolMintMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uvmd.mutation = mutation
			affected, err = uvmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uvmd.hooks) - 1; i >= 0; i-- {
			if uvmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uvmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uvmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvmd *UniswapV3PoolMintDelete) ExecX(ctx context.Context) int {
	n, err := uvmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uvmd *UniswapV3PoolMintDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: uniswapv3poolmint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: uniswapv3poolmint.FieldID,
			},
		},
	}
	if ps := uvmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uvmd.driver, _spec)
}

// UniswapV3PoolMintDeleteOne is the builder for deleting a single UniswapV3PoolMint entity.
type UniswapV3PoolMintDeleteOne struct {
	uvmd *UniswapV3PoolMintDelete
}

// Exec executes the deletion query.
func (uvmdo *UniswapV3PoolMintDeleteOne) Exec(ctx context.Context) error {
	n, err := uvmdo.uvmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{uniswapv3poolmint.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uvmdo *UniswapV3PoolMintDeleteOne) ExecX(ctx context.Context) {
	uvmdo.uvmd.ExecX(ctx)
}
