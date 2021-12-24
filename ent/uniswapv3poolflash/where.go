// Code generated by entc, DO NOT EDIT.

package uniswapv3poolflash

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/artmisxyz/legolas/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Sender applies equality check predicate on the "sender" field. It's identical to SenderEQ.
func Sender(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSender), v))
	})
}

// Recipient applies equality check predicate on the "recipient" field. It's identical to RecipientEQ.
func Recipient(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// Amount0 applies equality check predicate on the "amount0" field. It's identical to Amount0EQ.
func Amount0(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount0), v))
	})
}

// Amount1 applies equality check predicate on the "amount1" field. It's identical to Amount1EQ.
func Amount1(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount1), v))
	})
}

// Paid0 applies equality check predicate on the "paid0" field. It's identical to Paid0EQ.
func Paid0(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaid0), v))
	})
}

// Paid1 applies equality check predicate on the "paid1" field. It's identical to Paid1EQ.
func Paid1(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaid1), v))
	})
}

// SenderEQ applies the EQ predicate on the "sender" field.
func SenderEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSender), v))
	})
}

// SenderNEQ applies the NEQ predicate on the "sender" field.
func SenderNEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSender), v))
	})
}

// SenderIn applies the In predicate on the "sender" field.
func SenderIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSender), v...))
	})
}

// SenderNotIn applies the NotIn predicate on the "sender" field.
func SenderNotIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSender), v...))
	})
}

// SenderGT applies the GT predicate on the "sender" field.
func SenderGT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSender), v))
	})
}

// SenderGTE applies the GTE predicate on the "sender" field.
func SenderGTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSender), v))
	})
}

// SenderLT applies the LT predicate on the "sender" field.
func SenderLT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSender), v))
	})
}

// SenderLTE applies the LTE predicate on the "sender" field.
func SenderLTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSender), v))
	})
}

// SenderContains applies the Contains predicate on the "sender" field.
func SenderContains(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSender), v))
	})
}

// SenderHasPrefix applies the HasPrefix predicate on the "sender" field.
func SenderHasPrefix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSender), v))
	})
}

// SenderHasSuffix applies the HasSuffix predicate on the "sender" field.
func SenderHasSuffix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSender), v))
	})
}

// SenderEqualFold applies the EqualFold predicate on the "sender" field.
func SenderEqualFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSender), v))
	})
}

// SenderContainsFold applies the ContainsFold predicate on the "sender" field.
func SenderContainsFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSender), v))
	})
}

// RecipientEQ applies the EQ predicate on the "recipient" field.
func RecipientEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// RecipientNEQ applies the NEQ predicate on the "recipient" field.
func RecipientNEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecipient), v))
	})
}

// RecipientIn applies the In predicate on the "recipient" field.
func RecipientIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecipient), v...))
	})
}

// RecipientNotIn applies the NotIn predicate on the "recipient" field.
func RecipientNotIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecipient), v...))
	})
}

// RecipientGT applies the GT predicate on the "recipient" field.
func RecipientGT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecipient), v))
	})
}

// RecipientGTE applies the GTE predicate on the "recipient" field.
func RecipientGTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecipient), v))
	})
}

// RecipientLT applies the LT predicate on the "recipient" field.
func RecipientLT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecipient), v))
	})
}

// RecipientLTE applies the LTE predicate on the "recipient" field.
func RecipientLTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecipient), v))
	})
}

// RecipientContains applies the Contains predicate on the "recipient" field.
func RecipientContains(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecipient), v))
	})
}

// RecipientHasPrefix applies the HasPrefix predicate on the "recipient" field.
func RecipientHasPrefix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecipient), v))
	})
}

// RecipientHasSuffix applies the HasSuffix predicate on the "recipient" field.
func RecipientHasSuffix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecipient), v))
	})
}

// RecipientEqualFold applies the EqualFold predicate on the "recipient" field.
func RecipientEqualFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecipient), v))
	})
}

// RecipientContainsFold applies the ContainsFold predicate on the "recipient" field.
func RecipientContainsFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecipient), v))
	})
}

// Amount0EQ applies the EQ predicate on the "amount0" field.
func Amount0EQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount0), v))
	})
}

// Amount0NEQ applies the NEQ predicate on the "amount0" field.
func Amount0NEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount0), v))
	})
}

// Amount0In applies the In predicate on the "amount0" field.
func Amount0In(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func Amount0NotIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func Amount0GT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount0), v))
	})
}

// Amount0GTE applies the GTE predicate on the "amount0" field.
func Amount0GTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount0), v))
	})
}

// Amount0LT applies the LT predicate on the "amount0" field.
func Amount0LT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount0), v))
	})
}

// Amount0LTE applies the LTE predicate on the "amount0" field.
func Amount0LTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount0), v))
	})
}

// Amount0Contains applies the Contains predicate on the "amount0" field.
func Amount0Contains(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount0), v))
	})
}

// Amount0HasPrefix applies the HasPrefix predicate on the "amount0" field.
func Amount0HasPrefix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount0), v))
	})
}

// Amount0HasSuffix applies the HasSuffix predicate on the "amount0" field.
func Amount0HasSuffix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount0), v))
	})
}

// Amount0EqualFold applies the EqualFold predicate on the "amount0" field.
func Amount0EqualFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount0), v))
	})
}

// Amount0ContainsFold applies the ContainsFold predicate on the "amount0" field.
func Amount0ContainsFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount0), v))
	})
}

// Amount1EQ applies the EQ predicate on the "amount1" field.
func Amount1EQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount1), v))
	})
}

// Amount1NEQ applies the NEQ predicate on the "amount1" field.
func Amount1NEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount1), v))
	})
}

// Amount1In applies the In predicate on the "amount1" field.
func Amount1In(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func Amount1NotIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func Amount1GT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount1), v))
	})
}

// Amount1GTE applies the GTE predicate on the "amount1" field.
func Amount1GTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount1), v))
	})
}

// Amount1LT applies the LT predicate on the "amount1" field.
func Amount1LT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount1), v))
	})
}

// Amount1LTE applies the LTE predicate on the "amount1" field.
func Amount1LTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount1), v))
	})
}

// Amount1Contains applies the Contains predicate on the "amount1" field.
func Amount1Contains(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount1), v))
	})
}

// Amount1HasPrefix applies the HasPrefix predicate on the "amount1" field.
func Amount1HasPrefix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount1), v))
	})
}

// Amount1HasSuffix applies the HasSuffix predicate on the "amount1" field.
func Amount1HasSuffix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount1), v))
	})
}

// Amount1EqualFold applies the EqualFold predicate on the "amount1" field.
func Amount1EqualFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount1), v))
	})
}

// Amount1ContainsFold applies the ContainsFold predicate on the "amount1" field.
func Amount1ContainsFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount1), v))
	})
}

// Paid0EQ applies the EQ predicate on the "paid0" field.
func Paid0EQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaid0), v))
	})
}

// Paid0NEQ applies the NEQ predicate on the "paid0" field.
func Paid0NEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaid0), v))
	})
}

// Paid0In applies the In predicate on the "paid0" field.
func Paid0In(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaid0), v...))
	})
}

// Paid0NotIn applies the NotIn predicate on the "paid0" field.
func Paid0NotIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaid0), v...))
	})
}

// Paid0GT applies the GT predicate on the "paid0" field.
func Paid0GT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaid0), v))
	})
}

// Paid0GTE applies the GTE predicate on the "paid0" field.
func Paid0GTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaid0), v))
	})
}

// Paid0LT applies the LT predicate on the "paid0" field.
func Paid0LT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaid0), v))
	})
}

// Paid0LTE applies the LTE predicate on the "paid0" field.
func Paid0LTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaid0), v))
	})
}

// Paid0Contains applies the Contains predicate on the "paid0" field.
func Paid0Contains(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaid0), v))
	})
}

// Paid0HasPrefix applies the HasPrefix predicate on the "paid0" field.
func Paid0HasPrefix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaid0), v))
	})
}

// Paid0HasSuffix applies the HasSuffix predicate on the "paid0" field.
func Paid0HasSuffix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaid0), v))
	})
}

// Paid0EqualFold applies the EqualFold predicate on the "paid0" field.
func Paid0EqualFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaid0), v))
	})
}

// Paid0ContainsFold applies the ContainsFold predicate on the "paid0" field.
func Paid0ContainsFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaid0), v))
	})
}

// Paid1EQ applies the EQ predicate on the "paid1" field.
func Paid1EQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaid1), v))
	})
}

// Paid1NEQ applies the NEQ predicate on the "paid1" field.
func Paid1NEQ(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaid1), v))
	})
}

// Paid1In applies the In predicate on the "paid1" field.
func Paid1In(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaid1), v...))
	})
}

// Paid1NotIn applies the NotIn predicate on the "paid1" field.
func Paid1NotIn(vs ...string) predicate.UniswapV3PoolFlash {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaid1), v...))
	})
}

// Paid1GT applies the GT predicate on the "paid1" field.
func Paid1GT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaid1), v))
	})
}

// Paid1GTE applies the GTE predicate on the "paid1" field.
func Paid1GTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaid1), v))
	})
}

// Paid1LT applies the LT predicate on the "paid1" field.
func Paid1LT(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaid1), v))
	})
}

// Paid1LTE applies the LTE predicate on the "paid1" field.
func Paid1LTE(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaid1), v))
	})
}

// Paid1Contains applies the Contains predicate on the "paid1" field.
func Paid1Contains(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaid1), v))
	})
}

// Paid1HasPrefix applies the HasPrefix predicate on the "paid1" field.
func Paid1HasPrefix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaid1), v))
	})
}

// Paid1HasSuffix applies the HasSuffix predicate on the "paid1" field.
func Paid1HasSuffix(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaid1), v))
	})
}

// Paid1EqualFold applies the EqualFold predicate on the "paid1" field.
func Paid1EqualFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaid1), v))
	})
}

// Paid1ContainsFold applies the ContainsFold predicate on the "paid1" field.
func Paid1ContainsFold(v string) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaid1), v))
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func And(predicates ...predicate.UniswapV3PoolFlash) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UniswapV3PoolFlash) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
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
func Not(p predicate.UniswapV3PoolFlash) predicate.UniswapV3PoolFlash {
	return predicate.UniswapV3PoolFlash(func(s *sql.Selector) {
		p(s.Not())
	})
}
