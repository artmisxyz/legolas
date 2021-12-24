// Code generated by entc, DO NOT EDIT.

package uniswapv3poolmint

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artmisxyz/legolas/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// TickLower applies equality check predicate on the "tick_lower" field. It's identical to TickLowerEQ.
func TickLower(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTickLower), v))
	})
}

// TickUpper applies equality check predicate on the "tick_upper" field. It's identical to TickUpperEQ.
func TickUpper(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTickUpper), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Amount0 applies equality check predicate on the "amount0" field. It's identical to Amount0EQ.
func Amount0(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount0), v))
	})
}

// Amount1 applies equality check predicate on the "amount1" field. It's identical to Amount1EQ.
func Amount1(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount1), v))
	})
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwner), v))
	})
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOwner), v...))
	})
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOwner), v...))
	})
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOwner), v))
	})
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOwner), v))
	})
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOwner), v))
	})
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOwner), v))
	})
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOwner), v))
	})
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOwner), v))
	})
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOwner), v))
	})
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOwner), v))
	})
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOwner), v))
	})
}

// TickLowerEQ applies the EQ predicate on the "tick_lower" field.
func TickLowerEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTickLower), v))
	})
}

// TickLowerNEQ applies the NEQ predicate on the "tick_lower" field.
func TickLowerNEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTickLower), v))
	})
}

// TickLowerIn applies the In predicate on the "tick_lower" field.
func TickLowerIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTickLower), v...))
	})
}

// TickLowerNotIn applies the NotIn predicate on the "tick_lower" field.
func TickLowerNotIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTickLower), v...))
	})
}

// TickLowerGT applies the GT predicate on the "tick_lower" field.
func TickLowerGT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTickLower), v))
	})
}

// TickLowerGTE applies the GTE predicate on the "tick_lower" field.
func TickLowerGTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTickLower), v))
	})
}

// TickLowerLT applies the LT predicate on the "tick_lower" field.
func TickLowerLT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTickLower), v))
	})
}

// TickLowerLTE applies the LTE predicate on the "tick_lower" field.
func TickLowerLTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTickLower), v))
	})
}

// TickLowerContains applies the Contains predicate on the "tick_lower" field.
func TickLowerContains(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTickLower), v))
	})
}

// TickLowerHasPrefix applies the HasPrefix predicate on the "tick_lower" field.
func TickLowerHasPrefix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTickLower), v))
	})
}

// TickLowerHasSuffix applies the HasSuffix predicate on the "tick_lower" field.
func TickLowerHasSuffix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTickLower), v))
	})
}

// TickLowerEqualFold applies the EqualFold predicate on the "tick_lower" field.
func TickLowerEqualFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTickLower), v))
	})
}

// TickLowerContainsFold applies the ContainsFold predicate on the "tick_lower" field.
func TickLowerContainsFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTickLower), v))
	})
}

// TickUpperEQ applies the EQ predicate on the "tick_upper" field.
func TickUpperEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTickUpper), v))
	})
}

// TickUpperNEQ applies the NEQ predicate on the "tick_upper" field.
func TickUpperNEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTickUpper), v))
	})
}

// TickUpperIn applies the In predicate on the "tick_upper" field.
func TickUpperIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTickUpper), v...))
	})
}

// TickUpperNotIn applies the NotIn predicate on the "tick_upper" field.
func TickUpperNotIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTickUpper), v...))
	})
}

// TickUpperGT applies the GT predicate on the "tick_upper" field.
func TickUpperGT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTickUpper), v))
	})
}

// TickUpperGTE applies the GTE predicate on the "tick_upper" field.
func TickUpperGTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTickUpper), v))
	})
}

// TickUpperLT applies the LT predicate on the "tick_upper" field.
func TickUpperLT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTickUpper), v))
	})
}

// TickUpperLTE applies the LTE predicate on the "tick_upper" field.
func TickUpperLTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTickUpper), v))
	})
}

// TickUpperContains applies the Contains predicate on the "tick_upper" field.
func TickUpperContains(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTickUpper), v))
	})
}

// TickUpperHasPrefix applies the HasPrefix predicate on the "tick_upper" field.
func TickUpperHasPrefix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTickUpper), v))
	})
}

// TickUpperHasSuffix applies the HasSuffix predicate on the "tick_upper" field.
func TickUpperHasSuffix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTickUpper), v))
	})
}

// TickUpperEqualFold applies the EqualFold predicate on the "tick_upper" field.
func TickUpperEqualFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTickUpper), v))
	})
}

// TickUpperContainsFold applies the ContainsFold predicate on the "tick_upper" field.
func TickUpperContainsFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTickUpper), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount), v))
	})
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount), v))
	})
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount), v))
	})
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount), v))
	})
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount), v))
	})
}

// Amount0EQ applies the EQ predicate on the "amount0" field.
func Amount0EQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount0), v))
	})
}

// Amount0NEQ applies the NEQ predicate on the "amount0" field.
func Amount0NEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount0), v))
	})
}

// Amount0In applies the In predicate on the "amount0" field.
func Amount0In(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func Amount0NotIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func Amount0GT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount0), v))
	})
}

// Amount0GTE applies the GTE predicate on the "amount0" field.
func Amount0GTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount0), v))
	})
}

// Amount0LT applies the LT predicate on the "amount0" field.
func Amount0LT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount0), v))
	})
}

// Amount0LTE applies the LTE predicate on the "amount0" field.
func Amount0LTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount0), v))
	})
}

// Amount0Contains applies the Contains predicate on the "amount0" field.
func Amount0Contains(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount0), v))
	})
}

// Amount0HasPrefix applies the HasPrefix predicate on the "amount0" field.
func Amount0HasPrefix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount0), v))
	})
}

// Amount0HasSuffix applies the HasSuffix predicate on the "amount0" field.
func Amount0HasSuffix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount0), v))
	})
}

// Amount0EqualFold applies the EqualFold predicate on the "amount0" field.
func Amount0EqualFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount0), v))
	})
}

// Amount0ContainsFold applies the ContainsFold predicate on the "amount0" field.
func Amount0ContainsFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount0), v))
	})
}

// Amount1EQ applies the EQ predicate on the "amount1" field.
func Amount1EQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount1), v))
	})
}

// Amount1NEQ applies the NEQ predicate on the "amount1" field.
func Amount1NEQ(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount1), v))
	})
}

// Amount1In applies the In predicate on the "amount1" field.
func Amount1In(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func Amount1NotIn(vs ...string) predicate.UniswapV3PoolMint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func Amount1GT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount1), v))
	})
}

// Amount1GTE applies the GTE predicate on the "amount1" field.
func Amount1GTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount1), v))
	})
}

// Amount1LT applies the LT predicate on the "amount1" field.
func Amount1LT(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount1), v))
	})
}

// Amount1LTE applies the LTE predicate on the "amount1" field.
func Amount1LTE(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount1), v))
	})
}

// Amount1Contains applies the Contains predicate on the "amount1" field.
func Amount1Contains(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount1), v))
	})
}

// Amount1HasPrefix applies the HasPrefix predicate on the "amount1" field.
func Amount1HasPrefix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount1), v))
	})
}

// Amount1HasSuffix applies the HasSuffix predicate on the "amount1" field.
func Amount1HasSuffix(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount1), v))
	})
}

// Amount1EqualFold applies the EqualFold predicate on the "amount1" field.
func Amount1EqualFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount1), v))
	})
}

// Amount1ContainsFold applies the ContainsFold predicate on the "amount1" field.
func Amount1ContainsFold(v string) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount1), v))
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func And(predicates ...predicate.UniswapV3PoolMint) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UniswapV3PoolMint) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
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
func Not(p predicate.UniswapV3PoolMint) predicate.UniswapV3PoolMint {
	return predicate.UniswapV3PoolMint(func(s *sql.Selector) {
		p(s.Not())
	})
}
