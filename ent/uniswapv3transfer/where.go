// Code generated by entc, DO NOT EDIT.

package uniswapv3transfer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artmisxyz/blockinspector/ent/predicate"
	"github.com/artmisxyz/blockinspector/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TokenID applies equality check predicate on the "token_id" field. It's identical to TokenIDEQ.
func TokenID(v *schema.BigInt) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenID), v))
	})
}

// From applies equality check predicate on the "from" field. It's identical to FromEQ.
func From(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// To applies equality check predicate on the "to" field. It's identical to ToEQ.
func To(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// TokenIDEQ applies the EQ predicate on the "token_id" field.
func TokenIDEQ(v *schema.BigInt) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenID), v))
	})
}

// TokenIDNEQ applies the NEQ predicate on the "token_id" field.
func TokenIDNEQ(v *schema.BigInt) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenID), v))
	})
}

// TokenIDIn applies the In predicate on the "token_id" field.
func TokenIDIn(vs ...*schema.BigInt) predicate.UniswapV3Transfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
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
func TokenIDNotIn(vs ...*schema.BigInt) predicate.UniswapV3Transfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
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
func TokenIDGT(v *schema.BigInt) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenID), v))
	})
}

// TokenIDGTE applies the GTE predicate on the "token_id" field.
func TokenIDGTE(v *schema.BigInt) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenID), v))
	})
}

// TokenIDLT applies the LT predicate on the "token_id" field.
func TokenIDLT(v *schema.BigInt) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenID), v))
	})
}

// TokenIDLTE applies the LTE predicate on the "token_id" field.
func TokenIDLTE(v *schema.BigInt) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenID), v))
	})
}

// TokenIDContains applies the Contains predicate on the "token_id" field.
func TokenIDContains(v *schema.BigInt) predicate.UniswapV3Transfer {
	vc := v.String()
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenID), vc))
	})
}

// TokenIDHasPrefix applies the HasPrefix predicate on the "token_id" field.
func TokenIDHasPrefix(v *schema.BigInt) predicate.UniswapV3Transfer {
	vc := v.String()
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenID), vc))
	})
}

// TokenIDHasSuffix applies the HasSuffix predicate on the "token_id" field.
func TokenIDHasSuffix(v *schema.BigInt) predicate.UniswapV3Transfer {
	vc := v.String()
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenID), vc))
	})
}

// TokenIDEqualFold applies the EqualFold predicate on the "token_id" field.
func TokenIDEqualFold(v *schema.BigInt) predicate.UniswapV3Transfer {
	vc := v.String()
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenID), vc))
	})
}

// TokenIDContainsFold applies the ContainsFold predicate on the "token_id" field.
func TokenIDContainsFold(v *schema.BigInt) predicate.UniswapV3Transfer {
	vc := v.String()
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenID), vc))
	})
}

// FromEQ applies the EQ predicate on the "from" field.
func FromEQ(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// FromNEQ applies the NEQ predicate on the "from" field.
func FromNEQ(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFrom), v))
	})
}

// FromIn applies the In predicate on the "from" field.
func FromIn(vs ...string) predicate.UniswapV3Transfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFrom), v...))
	})
}

// FromNotIn applies the NotIn predicate on the "from" field.
func FromNotIn(vs ...string) predicate.UniswapV3Transfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFrom), v...))
	})
}

// FromGT applies the GT predicate on the "from" field.
func FromGT(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFrom), v))
	})
}

// FromGTE applies the GTE predicate on the "from" field.
func FromGTE(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFrom), v))
	})
}

// FromLT applies the LT predicate on the "from" field.
func FromLT(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFrom), v))
	})
}

// FromLTE applies the LTE predicate on the "from" field.
func FromLTE(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFrom), v))
	})
}

// FromContains applies the Contains predicate on the "from" field.
func FromContains(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFrom), v))
	})
}

// FromHasPrefix applies the HasPrefix predicate on the "from" field.
func FromHasPrefix(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFrom), v))
	})
}

// FromHasSuffix applies the HasSuffix predicate on the "from" field.
func FromHasSuffix(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFrom), v))
	})
}

// FromEqualFold applies the EqualFold predicate on the "from" field.
func FromEqualFold(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFrom), v))
	})
}

// FromContainsFold applies the ContainsFold predicate on the "from" field.
func FromContainsFold(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFrom), v))
	})
}

// ToEQ applies the EQ predicate on the "to" field.
func ToEQ(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// ToNEQ applies the NEQ predicate on the "to" field.
func ToNEQ(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTo), v))
	})
}

// ToIn applies the In predicate on the "to" field.
func ToIn(vs ...string) predicate.UniswapV3Transfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTo), v...))
	})
}

// ToNotIn applies the NotIn predicate on the "to" field.
func ToNotIn(vs ...string) predicate.UniswapV3Transfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTo), v...))
	})
}

// ToGT applies the GT predicate on the "to" field.
func ToGT(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTo), v))
	})
}

// ToGTE applies the GTE predicate on the "to" field.
func ToGTE(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTo), v))
	})
}

// ToLT applies the LT predicate on the "to" field.
func ToLT(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTo), v))
	})
}

// ToLTE applies the LTE predicate on the "to" field.
func ToLTE(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTo), v))
	})
}

// ToContains applies the Contains predicate on the "to" field.
func ToContains(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTo), v))
	})
}

// ToHasPrefix applies the HasPrefix predicate on the "to" field.
func ToHasPrefix(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTo), v))
	})
}

// ToHasSuffix applies the HasSuffix predicate on the "to" field.
func ToHasSuffix(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTo), v))
	})
}

// ToEqualFold applies the EqualFold predicate on the "to" field.
func ToEqualFold(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTo), v))
	})
}

// ToContainsFold applies the ContainsFold predicate on the "to" field.
func ToContainsFold(v string) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTo), v))
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
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
func And(predicates ...predicate.UniswapV3Transfer) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UniswapV3Transfer) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
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
func Not(p predicate.UniswapV3Transfer) predicate.UniswapV3Transfer {
	return predicate.UniswapV3Transfer(func(s *sql.Selector) {
		p(s.Not())
	})
}
