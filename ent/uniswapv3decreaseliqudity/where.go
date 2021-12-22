// Code generated by entc, DO NOT EDIT.

package uniswapv3decreaseliqudity

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artmisxyz/legolas/ent/predicate"
	"github.com/artmisxyz/legolas/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TokenID applies equality check predicate on the "token_id" field. It's identical to TokenIDEQ.
func TokenID(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenID), v))
	})
}

// Liquidity applies equality check predicate on the "liquidity" field. It's identical to LiquidityEQ.
func Liquidity(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLiquidity), v))
	})
}

// Amount0 applies equality check predicate on the "amount0" field. It's identical to Amount0EQ.
func Amount0(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount0), v))
	})
}

// Amount1 applies equality check predicate on the "amount1" field. It's identical to Amount1EQ.
func Amount1(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount1), v))
	})
}

// TokenIDEQ applies the EQ predicate on the "token_id" field.
func TokenIDEQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenID), v))
	})
}

// TokenIDNEQ applies the NEQ predicate on the "token_id" field.
func TokenIDNEQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenID), v))
	})
}

// TokenIDIn applies the In predicate on the "token_id" field.
func TokenIDIn(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTokenID), v...))
	})
}

// TokenIDNotIn applies the NotIn predicate on the "token_id" field.
func TokenIDNotIn(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTokenID), v...))
	})
}

// TokenIDGT applies the GT predicate on the "token_id" field.
func TokenIDGT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenID), v))
	})
}

// TokenIDGTE applies the GTE predicate on the "token_id" field.
func TokenIDGTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenID), v))
	})
}

// TokenIDLT applies the LT predicate on the "token_id" field.
func TokenIDLT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenID), v))
	})
}

// TokenIDLTE applies the LTE predicate on the "token_id" field.
func TokenIDLTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenID), v))
	})
}

// TokenIDContains applies the Contains predicate on the "token_id" field.
func TokenIDContains(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenID), vc))
	})
}

// TokenIDHasPrefix applies the HasPrefix predicate on the "token_id" field.
func TokenIDHasPrefix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenID), vc))
	})
}

// TokenIDHasSuffix applies the HasSuffix predicate on the "token_id" field.
func TokenIDHasSuffix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenID), vc))
	})
}

// TokenIDEqualFold applies the EqualFold predicate on the "token_id" field.
func TokenIDEqualFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenID), vc))
	})
}

// TokenIDContainsFold applies the ContainsFold predicate on the "token_id" field.
func TokenIDContainsFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenID), vc))
	})
}

// LiquidityEQ applies the EQ predicate on the "liquidity" field.
func LiquidityEQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLiquidity), v))
	})
}

// LiquidityNEQ applies the NEQ predicate on the "liquidity" field.
func LiquidityNEQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLiquidity), v))
	})
}

// LiquidityIn applies the In predicate on the "liquidity" field.
func LiquidityIn(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLiquidity), v...))
	})
}

// LiquidityNotIn applies the NotIn predicate on the "liquidity" field.
func LiquidityNotIn(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLiquidity), v...))
	})
}

// LiquidityGT applies the GT predicate on the "liquidity" field.
func LiquidityGT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLiquidity), v))
	})
}

// LiquidityGTE applies the GTE predicate on the "liquidity" field.
func LiquidityGTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLiquidity), v))
	})
}

// LiquidityLT applies the LT predicate on the "liquidity" field.
func LiquidityLT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLiquidity), v))
	})
}

// LiquidityLTE applies the LTE predicate on the "liquidity" field.
func LiquidityLTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLiquidity), v))
	})
}

// LiquidityContains applies the Contains predicate on the "liquidity" field.
func LiquidityContains(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLiquidity), vc))
	})
}

// LiquidityHasPrefix applies the HasPrefix predicate on the "liquidity" field.
func LiquidityHasPrefix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLiquidity), vc))
	})
}

// LiquidityHasSuffix applies the HasSuffix predicate on the "liquidity" field.
func LiquidityHasSuffix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLiquidity), vc))
	})
}

// LiquidityEqualFold applies the EqualFold predicate on the "liquidity" field.
func LiquidityEqualFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLiquidity), vc))
	})
}

// LiquidityContainsFold applies the ContainsFold predicate on the "liquidity" field.
func LiquidityContainsFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLiquidity), vc))
	})
}

// Amount0EQ applies the EQ predicate on the "amount0" field.
func Amount0EQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount0), v))
	})
}

// Amount0NEQ applies the NEQ predicate on the "amount0" field.
func Amount0NEQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount0), v))
	})
}

// Amount0In applies the In predicate on the "amount0" field.
func Amount0In(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount0), v...))
	})
}

// Amount0NotIn applies the NotIn predicate on the "amount0" field.
func Amount0NotIn(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount0), v...))
	})
}

// Amount0GT applies the GT predicate on the "amount0" field.
func Amount0GT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount0), v))
	})
}

// Amount0GTE applies the GTE predicate on the "amount0" field.
func Amount0GTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount0), v))
	})
}

// Amount0LT applies the LT predicate on the "amount0" field.
func Amount0LT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount0), v))
	})
}

// Amount0LTE applies the LTE predicate on the "amount0" field.
func Amount0LTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount0), v))
	})
}

// Amount0Contains applies the Contains predicate on the "amount0" field.
func Amount0Contains(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount0), vc))
	})
}

// Amount0HasPrefix applies the HasPrefix predicate on the "amount0" field.
func Amount0HasPrefix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount0), vc))
	})
}

// Amount0HasSuffix applies the HasSuffix predicate on the "amount0" field.
func Amount0HasSuffix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount0), vc))
	})
}

// Amount0EqualFold applies the EqualFold predicate on the "amount0" field.
func Amount0EqualFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount0), vc))
	})
}

// Amount0ContainsFold applies the ContainsFold predicate on the "amount0" field.
func Amount0ContainsFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount0), vc))
	})
}

// Amount1EQ applies the EQ predicate on the "amount1" field.
func Amount1EQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount1), v))
	})
}

// Amount1NEQ applies the NEQ predicate on the "amount1" field.
func Amount1NEQ(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount1), v))
	})
}

// Amount1In applies the In predicate on the "amount1" field.
func Amount1In(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount1), v...))
	})
}

// Amount1NotIn applies the NotIn predicate on the "amount1" field.
func Amount1NotIn(vs ...*schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount1), v...))
	})
}

// Amount1GT applies the GT predicate on the "amount1" field.
func Amount1GT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount1), v))
	})
}

// Amount1GTE applies the GTE predicate on the "amount1" field.
func Amount1GTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount1), v))
	})
}

// Amount1LT applies the LT predicate on the "amount1" field.
func Amount1LT(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount1), v))
	})
}

// Amount1LTE applies the LTE predicate on the "amount1" field.
func Amount1LTE(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount1), v))
	})
}

// Amount1Contains applies the Contains predicate on the "amount1" field.
func Amount1Contains(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount1), vc))
	})
}

// Amount1HasPrefix applies the HasPrefix predicate on the "amount1" field.
func Amount1HasPrefix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount1), vc))
	})
}

// Amount1HasSuffix applies the HasSuffix predicate on the "amount1" field.
func Amount1HasSuffix(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount1), vc))
	})
}

// Amount1EqualFold applies the EqualFold predicate on the "amount1" field.
func Amount1EqualFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount1), vc))
	})
}

// Amount1ContainsFold applies the ContainsFold predicate on the "amount1" field.
func Amount1ContainsFold(v *schema.BigInt) predicate.UniswapV3DecreaseLiqudity {
	vc := v.String()
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount1), vc))
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UniswapV3DecreaseLiqudity) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UniswapV3DecreaseLiqudity) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UniswapV3DecreaseLiqudity) predicate.UniswapV3DecreaseLiqudity {
	return predicate.UniswapV3DecreaseLiqudity(func(s *sql.Selector) {
		p(s.Not())
	})
}
