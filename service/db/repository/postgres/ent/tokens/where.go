// Code generated by ent, DO NOT EDIT.

package tokens

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// RefreshToken applies equality check predicate on the "refresh_token" field. It's identical to RefreshTokenEQ.
func RefreshToken(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefreshToken), v))
	})
}

// TokenType applies equality check predicate on the "token_type" field. It's identical to TokenTypeEQ.
func TokenType(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenType), v))
	})
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccessToken), v))
	})
}

// RefreshTokenEQ applies the EQ predicate on the "refresh_token" field.
func RefreshTokenEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenNEQ applies the NEQ predicate on the "refresh_token" field.
func RefreshTokenNEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenIn applies the In predicate on the "refresh_token" field.
func RefreshTokenIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRefreshToken), v...))
	})
}

// RefreshTokenNotIn applies the NotIn predicate on the "refresh_token" field.
func RefreshTokenNotIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRefreshToken), v...))
	})
}

// RefreshTokenGT applies the GT predicate on the "refresh_token" field.
func RefreshTokenGT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenGTE applies the GTE predicate on the "refresh_token" field.
func RefreshTokenGTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenLT applies the LT predicate on the "refresh_token" field.
func RefreshTokenLT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenLTE applies the LTE predicate on the "refresh_token" field.
func RefreshTokenLTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenContains applies the Contains predicate on the "refresh_token" field.
func RefreshTokenContains(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenHasPrefix applies the HasPrefix predicate on the "refresh_token" field.
func RefreshTokenHasPrefix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenHasSuffix applies the HasSuffix predicate on the "refresh_token" field.
func RefreshTokenHasSuffix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenEqualFold applies the EqualFold predicate on the "refresh_token" field.
func RefreshTokenEqualFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenContainsFold applies the ContainsFold predicate on the "refresh_token" field.
func RefreshTokenContainsFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRefreshToken), v))
	})
}

// TokenTypeEQ applies the EQ predicate on the "token_type" field.
func TokenTypeEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenType), v))
	})
}

// TokenTypeNEQ applies the NEQ predicate on the "token_type" field.
func TokenTypeNEQ(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenType), v))
	})
}

// TokenTypeIn applies the In predicate on the "token_type" field.
func TokenTypeIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTokenType), v...))
	})
}

// TokenTypeNotIn applies the NotIn predicate on the "token_type" field.
func TokenTypeNotIn(vs ...string) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTokenType), v...))
	})
}

// TokenTypeGT applies the GT predicate on the "token_type" field.
func TokenTypeGT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenType), v))
	})
}

// TokenTypeGTE applies the GTE predicate on the "token_type" field.
func TokenTypeGTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenType), v))
	})
}

// TokenTypeLT applies the LT predicate on the "token_type" field.
func TokenTypeLT(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenType), v))
	})
}

// TokenTypeLTE applies the LTE predicate on the "token_type" field.
func TokenTypeLTE(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenType), v))
	})
}

// TokenTypeContains applies the Contains predicate on the "token_type" field.
func TokenTypeContains(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenType), v))
	})
}

// TokenTypeHasPrefix applies the HasPrefix predicate on the "token_type" field.
func TokenTypeHasPrefix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenType), v))
	})
}

// TokenTypeHasSuffix applies the HasSuffix predicate on the "token_type" field.
func TokenTypeHasSuffix(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenType), v))
	})
}

// TokenTypeEqualFold applies the EqualFold predicate on the "token_type" field.
func TokenTypeEqualFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenType), v))
	})
}

// TokenTypeContainsFold applies the ContainsFold predicate on the "token_type" field.
func TokenTypeContainsFold(v string) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenType), v))
	})
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiry), v))
	})
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExpiry), v...))
	})
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.Tokens {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExpiry), v...))
	})
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiry), v))
	})
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiry), v))
	})
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiry), v))
	})
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiry), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tokens) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tokens) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
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
func Not(p predicate.Tokens) predicate.Tokens {
	return predicate.Tokens(func(s *sql.Selector) {
		p(s.Not())
	})
}
