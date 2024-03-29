// Code generated by ent, DO NOT EDIT.

package users

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// HashedPassword applies equality check predicate on the "hashed_password" field. It's identical to HashedPasswordEQ.
func HashedPassword(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedPassword), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullName), v))
	})
}

// PasswordChangedAt applies equality check predicate on the "password_changed_at" field. It's identical to PasswordChangedAtEQ.
func PasswordChangedAt(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordChangedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// HashedPasswordEQ applies the EQ predicate on the "hashed_password" field.
func HashedPasswordEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordNEQ applies the NEQ predicate on the "hashed_password" field.
func HashedPasswordNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordIn applies the In predicate on the "hashed_password" field.
func HashedPasswordIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHashedPassword), v...))
	})
}

// HashedPasswordNotIn applies the NotIn predicate on the "hashed_password" field.
func HashedPasswordNotIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHashedPassword), v...))
	})
}

// HashedPasswordGT applies the GT predicate on the "hashed_password" field.
func HashedPasswordGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordGTE applies the GTE predicate on the "hashed_password" field.
func HashedPasswordGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordLT applies the LT predicate on the "hashed_password" field.
func HashedPasswordLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordLTE applies the LTE predicate on the "hashed_password" field.
func HashedPasswordLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordContains applies the Contains predicate on the "hashed_password" field.
func HashedPasswordContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordHasPrefix applies the HasPrefix predicate on the "hashed_password" field.
func HashedPasswordHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordHasSuffix applies the HasSuffix predicate on the "hashed_password" field.
func HashedPasswordHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordIsNil applies the IsNil predicate on the "hashed_password" field.
func HashedPasswordIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHashedPassword)))
	})
}

// HashedPasswordNotNil applies the NotNil predicate on the "hashed_password" field.
func HashedPasswordNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHashedPassword)))
	})
}

// HashedPasswordEqualFold applies the EqualFold predicate on the "hashed_password" field.
func HashedPasswordEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHashedPassword), v))
	})
}

// HashedPasswordContainsFold applies the ContainsFold predicate on the "hashed_password" field.
func HashedPasswordContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHashedPassword), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullName), v))
	})
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFullName), v))
	})
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFullName), v...))
	})
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFullName), v...))
	})
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFullName), v))
	})
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFullName), v))
	})
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFullName), v))
	})
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFullName), v))
	})
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFullName), v))
	})
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFullName), v))
	})
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFullName), v))
	})
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFullName), v))
	})
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFullName), v))
	})
}

// PasswordChangedAtEQ applies the EQ predicate on the "password_changed_at" field.
func PasswordChangedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPasswordChangedAt), v))
	})
}

// PasswordChangedAtNEQ applies the NEQ predicate on the "password_changed_at" field.
func PasswordChangedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPasswordChangedAt), v))
	})
}

// PasswordChangedAtIn applies the In predicate on the "password_changed_at" field.
func PasswordChangedAtIn(vs ...time.Time) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPasswordChangedAt), v...))
	})
}

// PasswordChangedAtNotIn applies the NotIn predicate on the "password_changed_at" field.
func PasswordChangedAtNotIn(vs ...time.Time) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPasswordChangedAt), v...))
	})
}

// PasswordChangedAtGT applies the GT predicate on the "password_changed_at" field.
func PasswordChangedAtGT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPasswordChangedAt), v))
	})
}

// PasswordChangedAtGTE applies the GTE predicate on the "password_changed_at" field.
func PasswordChangedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPasswordChangedAt), v))
	})
}

// PasswordChangedAtLT applies the LT predicate on the "password_changed_at" field.
func PasswordChangedAtLT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPasswordChangedAt), v))
	})
}

// PasswordChangedAtLTE applies the LTE predicate on the "password_changed_at" field.
func PasswordChangedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPasswordChangedAt), v))
	})
}

// PasswordChangedAtIsNil applies the IsNil predicate on the "password_changed_at" field.
func PasswordChangedAtIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPasswordChangedAt)))
	})
}

// PasswordChangedAtNotNil applies the NotNil predicate on the "password_changed_at" field.
func PasswordChangedAtNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPasswordChangedAt)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Users {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
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
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		p(s.Not())
	})
}
