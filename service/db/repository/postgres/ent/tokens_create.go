// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/tokens"
)

// TokensCreate is the builder for creating a Tokens entity.
type TokensCreate struct {
	config
	mutation *TokensMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (tc *TokensCreate) SetEmail(s string) *TokensCreate {
	tc.mutation.SetEmail(s)
	return tc
}

// SetAccessToken sets the "access_token" field.
func (tc *TokensCreate) SetAccessToken(s string) *TokensCreate {
	tc.mutation.SetAccessToken(s)
	return tc
}

// SetRefreshToken sets the "refresh_token" field.
func (tc *TokensCreate) SetRefreshToken(s string) *TokensCreate {
	tc.mutation.SetRefreshToken(s)
	return tc
}

// SetTokenType sets the "token_type" field.
func (tc *TokensCreate) SetTokenType(s string) *TokensCreate {
	tc.mutation.SetTokenType(s)
	return tc
}

// SetExpiry sets the "expiry" field.
func (tc *TokensCreate) SetExpiry(t time.Time) *TokensCreate {
	tc.mutation.SetExpiry(t)
	return tc
}

// SetRaw sets the "raw" field.
func (tc *TokensCreate) SetRaw(m map[string]interface{}) *TokensCreate {
	tc.mutation.SetRaw(m)
	return tc
}

// SetID sets the "id" field.
func (tc *TokensCreate) SetID(u uuid.UUID) *TokensCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TokensCreate) SetNillableID(u *uuid.UUID) *TokensCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// Mutation returns the TokensMutation object of the builder.
func (tc *TokensCreate) Mutation() *TokensMutation {
	return tc.mutation
}

// Save creates the Tokens in the database.
func (tc *TokensCreate) Save(ctx context.Context) (*Tokens, error) {
	var (
		err  error
		node *Tokens
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokensMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tokens)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TokensMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokensCreate) SaveX(ctx context.Context) *Tokens {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokensCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokensCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokensCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := tokens.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokensCreate) check() error {
	if _, ok := tc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Tokens.email"`)}
	}
	if v, ok := tc.mutation.Email(); ok {
		if err := tokens.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Tokens.email": %w`, err)}
		}
	}
	if _, ok := tc.mutation.AccessToken(); !ok {
		return &ValidationError{Name: "access_token", err: errors.New(`ent: missing required field "Tokens.access_token"`)}
	}
	if v, ok := tc.mutation.AccessToken(); ok {
		if err := tokens.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Tokens.access_token": %w`, err)}
		}
	}
	if _, ok := tc.mutation.RefreshToken(); !ok {
		return &ValidationError{Name: "refresh_token", err: errors.New(`ent: missing required field "Tokens.refresh_token"`)}
	}
	if _, ok := tc.mutation.TokenType(); !ok {
		return &ValidationError{Name: "token_type", err: errors.New(`ent: missing required field "Tokens.token_type"`)}
	}
	if v, ok := tc.mutation.TokenType(); ok {
		if err := tokens.TokenTypeValidator(v); err != nil {
			return &ValidationError{Name: "token_type", err: fmt.Errorf(`ent: validator failed for field "Tokens.token_type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`ent: missing required field "Tokens.expiry"`)}
	}
	if _, ok := tc.mutation.Raw(); !ok {
		return &ValidationError{Name: "raw", err: errors.New(`ent: missing required field "Tokens.raw"`)}
	}
	return nil
}

func (tc *TokensCreate) sqlSave(ctx context.Context) (*Tokens, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TokensCreate) createSpec() (*Tokens, *sqlgraph.CreateSpec) {
	var (
		_node = &Tokens{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tokens.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tokens.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Email(); ok {
		_spec.SetField(tokens.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := tc.mutation.AccessToken(); ok {
		_spec.SetField(tokens.FieldAccessToken, field.TypeString, value)
		_node.AccessToken = value
	}
	if value, ok := tc.mutation.RefreshToken(); ok {
		_spec.SetField(tokens.FieldRefreshToken, field.TypeString, value)
		_node.RefreshToken = value
	}
	if value, ok := tc.mutation.TokenType(); ok {
		_spec.SetField(tokens.FieldTokenType, field.TypeString, value)
		_node.TokenType = value
	}
	if value, ok := tc.mutation.Expiry(); ok {
		_spec.SetField(tokens.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	if value, ok := tc.mutation.Raw(); ok {
		_spec.SetField(tokens.FieldRaw, field.TypeJSON, value)
		_node.Raw = value
	}
	return _node, _spec
}

// TokensCreateBulk is the builder for creating many Tokens entities in bulk.
type TokensCreateBulk struct {
	config
	builders []*TokensCreate
}

// Save creates the Tokens entities in the database.
func (tcb *TokensCreateBulk) Save(ctx context.Context) ([]*Tokens, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tokens, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokensMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokensCreateBulk) SaveX(ctx context.Context) []*Tokens {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokensCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokensCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
