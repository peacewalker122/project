// Code generated by ent, DO NOT EDIT.

package qoute_retweet_feature

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FromAccountID applies equality check predicate on the "from_account_id" field. It's identical to FromAccountIDEQ.
func FromAccountID(v int64) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromAccountID), v))
	})
}

// QouteRetweet applies equality check predicate on the "qoute_retweet" field. It's identical to QouteRetweetEQ.
func QouteRetweet(v bool) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQouteRetweet), v))
	})
}

// Qoute applies equality check predicate on the "qoute" field. It's identical to QouteEQ.
func Qoute(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQoute), v))
	})
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v uuid.UUID) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// FromAccountIDEQ applies the EQ predicate on the "from_account_id" field.
func FromAccountIDEQ(v int64) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDNEQ applies the NEQ predicate on the "from_account_id" field.
func FromAccountIDNEQ(v int64) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDIn applies the In predicate on the "from_account_id" field.
func FromAccountIDIn(vs ...int64) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFromAccountID), v...))
	})
}

// FromAccountIDNotIn applies the NotIn predicate on the "from_account_id" field.
func FromAccountIDNotIn(vs ...int64) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFromAccountID), v...))
	})
}

// FromAccountIDGT applies the GT predicate on the "from_account_id" field.
func FromAccountIDGT(v int64) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDGTE applies the GTE predicate on the "from_account_id" field.
func FromAccountIDGTE(v int64) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDLT applies the LT predicate on the "from_account_id" field.
func FromAccountIDLT(v int64) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFromAccountID), v))
	})
}

// FromAccountIDLTE applies the LTE predicate on the "from_account_id" field.
func FromAccountIDLTE(v int64) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFromAccountID), v))
	})
}

// QouteRetweetEQ applies the EQ predicate on the "qoute_retweet" field.
func QouteRetweetEQ(v bool) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQouteRetweet), v))
	})
}

// QouteRetweetNEQ applies the NEQ predicate on the "qoute_retweet" field.
func QouteRetweetNEQ(v bool) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQouteRetweet), v))
	})
}

// QouteEQ applies the EQ predicate on the "qoute" field.
func QouteEQ(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQoute), v))
	})
}

// QouteNEQ applies the NEQ predicate on the "qoute" field.
func QouteNEQ(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQoute), v))
	})
}

// QouteIn applies the In predicate on the "qoute" field.
func QouteIn(vs ...string) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldQoute), v...))
	})
}

// QouteNotIn applies the NotIn predicate on the "qoute" field.
func QouteNotIn(vs ...string) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldQoute), v...))
	})
}

// QouteGT applies the GT predicate on the "qoute" field.
func QouteGT(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQoute), v))
	})
}

// QouteGTE applies the GTE predicate on the "qoute" field.
func QouteGTE(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQoute), v))
	})
}

// QouteLT applies the LT predicate on the "qoute" field.
func QouteLT(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQoute), v))
	})
}

// QouteLTE applies the LTE predicate on the "qoute" field.
func QouteLTE(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQoute), v))
	})
}

// QouteContains applies the Contains predicate on the "qoute" field.
func QouteContains(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldQoute), v))
	})
}

// QouteHasPrefix applies the HasPrefix predicate on the "qoute" field.
func QouteHasPrefix(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldQoute), v))
	})
}

// QouteHasSuffix applies the HasSuffix predicate on the "qoute" field.
func QouteHasSuffix(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldQoute), v))
	})
}

// QouteEqualFold applies the EqualFold predicate on the "qoute" field.
func QouteEqualFold(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldQoute), v))
	})
}

// QouteContainsFold applies the ContainsFold predicate on the "qoute" field.
func QouteContainsFold(v string) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldQoute), v))
	})
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v uuid.UUID) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostID), v))
	})
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v uuid.UUID) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPostID), v))
	})
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...uuid.UUID) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPostID), v...))
	})
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...uuid.UUID) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPostID), v...))
	})
}

// PostIDGT applies the GT predicate on the "post_id" field.
func PostIDGT(v uuid.UUID) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPostID), v))
	})
}

// PostIDGTE applies the GTE predicate on the "post_id" field.
func PostIDGTE(v uuid.UUID) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPostID), v))
	})
}

// PostIDLT applies the LT predicate on the "post_id" field.
func PostIDLT(v uuid.UUID) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPostID), v))
	})
}

// PostIDLTE applies the LTE predicate on the "post_id" field.
func PostIDLTE(v uuid.UUID) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPostID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Qoute_retweet_feature {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Qoute_retweet_feature) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Qoute_retweet_feature) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
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
func Not(p predicate.Qoute_retweet_feature) predicate.Qoute_retweet_feature {
	return predicate.Qoute_retweet_feature(func(s *sql.Selector) {
		p(s.Not())
	})
}