// Code generated by entc, DO NOT EDIT.

package uniswapv3poolinitialize

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artmisxyz/legolas/ent/predicate"
	"github.com/artmisxyz/legolas/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SqrtPriceX96 applies equality check predicate on the "sqrt_price_x96" field. It's identical to SqrtPriceX96EQ.
func SqrtPriceX96(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSqrtPriceX96), v))
	})
}

// Tick applies equality check predicate on the "tick" field. It's identical to TickEQ.
func Tick(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTick), v))
	})
}

// SqrtPriceX96EQ applies the EQ predicate on the "sqrt_price_x96" field.
func SqrtPriceX96EQ(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSqrtPriceX96), v))
	})
}

// SqrtPriceX96NEQ applies the NEQ predicate on the "sqrt_price_x96" field.
func SqrtPriceX96NEQ(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSqrtPriceX96), v))
	})
}

// SqrtPriceX96In applies the In predicate on the "sqrt_price_x96" field.
func SqrtPriceX96In(vs ...*schema.BigInt) predicate.UniswapV3PoolInitialize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSqrtPriceX96), v...))
	})
}

// SqrtPriceX96NotIn applies the NotIn predicate on the "sqrt_price_x96" field.
func SqrtPriceX96NotIn(vs ...*schema.BigInt) predicate.UniswapV3PoolInitialize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSqrtPriceX96), v...))
	})
}

// SqrtPriceX96GT applies the GT predicate on the "sqrt_price_x96" field.
func SqrtPriceX96GT(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSqrtPriceX96), v))
	})
}

// SqrtPriceX96GTE applies the GTE predicate on the "sqrt_price_x96" field.
func SqrtPriceX96GTE(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSqrtPriceX96), v))
	})
}

// SqrtPriceX96LT applies the LT predicate on the "sqrt_price_x96" field.
func SqrtPriceX96LT(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSqrtPriceX96), v))
	})
}

// SqrtPriceX96LTE applies the LTE predicate on the "sqrt_price_x96" field.
func SqrtPriceX96LTE(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSqrtPriceX96), v))
	})
}

// SqrtPriceX96Contains applies the Contains predicate on the "sqrt_price_x96" field.
func SqrtPriceX96Contains(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSqrtPriceX96), vc))
	})
}

// SqrtPriceX96HasPrefix applies the HasPrefix predicate on the "sqrt_price_x96" field.
func SqrtPriceX96HasPrefix(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSqrtPriceX96), vc))
	})
}

// SqrtPriceX96HasSuffix applies the HasSuffix predicate on the "sqrt_price_x96" field.
func SqrtPriceX96HasSuffix(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSqrtPriceX96), vc))
	})
}

// SqrtPriceX96EqualFold applies the EqualFold predicate on the "sqrt_price_x96" field.
func SqrtPriceX96EqualFold(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSqrtPriceX96), vc))
	})
}

// SqrtPriceX96ContainsFold applies the ContainsFold predicate on the "sqrt_price_x96" field.
func SqrtPriceX96ContainsFold(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSqrtPriceX96), vc))
	})
}

// TickEQ applies the EQ predicate on the "tick" field.
func TickEQ(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTick), v))
	})
}

// TickNEQ applies the NEQ predicate on the "tick" field.
func TickNEQ(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTick), v))
	})
}

// TickIn applies the In predicate on the "tick" field.
func TickIn(vs ...*schema.BigInt) predicate.UniswapV3PoolInitialize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTick), v...))
	})
}

// TickNotIn applies the NotIn predicate on the "tick" field.
func TickNotIn(vs ...*schema.BigInt) predicate.UniswapV3PoolInitialize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTick), v...))
	})
}

// TickGT applies the GT predicate on the "tick" field.
func TickGT(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTick), v))
	})
}

// TickGTE applies the GTE predicate on the "tick" field.
func TickGTE(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTick), v))
	})
}

// TickLT applies the LT predicate on the "tick" field.
func TickLT(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTick), v))
	})
}

// TickLTE applies the LTE predicate on the "tick" field.
func TickLTE(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTick), v))
	})
}

// TickContains applies the Contains predicate on the "tick" field.
func TickContains(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTick), vc))
	})
}

// TickHasPrefix applies the HasPrefix predicate on the "tick" field.
func TickHasPrefix(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTick), vc))
	})
}

// TickHasSuffix applies the HasSuffix predicate on the "tick" field.
func TickHasSuffix(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTick), vc))
	})
}

// TickEqualFold applies the EqualFold predicate on the "tick" field.
func TickEqualFold(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTick), vc))
	})
}

// TickContainsFold applies the ContainsFold predicate on the "tick" field.
func TickContainsFold(v *schema.BigInt) predicate.UniswapV3PoolInitialize {
	vc := v.String()
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTick), vc))
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
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
func And(predicates ...predicate.UniswapV3PoolInitialize) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UniswapV3PoolInitialize) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
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
func Not(p predicate.UniswapV3PoolInitialize) predicate.UniswapV3PoolInitialize {
	return predicate.UniswapV3PoolInitialize(func(s *sql.Selector) {
		p(s.Not())
	})
}