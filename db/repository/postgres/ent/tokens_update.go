// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/peacewalker122/project/db/repository/postgres/ent/predicate"
	"github.com/peacewalker122/project/db/repository/postgres/ent/tokens"
)

// TokensUpdate is the builder for updating Tokens entities.
type TokensUpdate struct {
	config
	hooks    []Hook
	mutation *TokensMutation
}

// Where appends a list predicates to the TokensUpdate builder.
func (tu *TokensUpdate) Where(ps ...predicate.Tokens) *TokensUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetEmail sets the "email" field.
func (tu *TokensUpdate) SetEmail(s string) *TokensUpdate {
	tu.mutation.SetEmail(s)
	return tu
}

// SetAccessToken sets the "access_token" field.
func (tu *TokensUpdate) SetAccessToken(s string) *TokensUpdate {
	tu.mutation.SetAccessToken(s)
	return tu
}

// SetRefreshToken sets the "refresh_token" field.
func (tu *TokensUpdate) SetRefreshToken(s string) *TokensUpdate {
	tu.mutation.SetRefreshToken(s)
	return tu
}

// SetTokenType sets the "token_type" field.
func (tu *TokensUpdate) SetTokenType(s string) *TokensUpdate {
	tu.mutation.SetTokenType(s)
	return tu
}

// SetExpiry sets the "expiry" field.
func (tu *TokensUpdate) SetExpiry(t time.Time) *TokensUpdate {
	tu.mutation.SetExpiry(t)
	return tu
}

// SetRaw sets the "raw" field.
func (tu *TokensUpdate) SetRaw(m map[string]interface{}) *TokensUpdate {
	tu.mutation.SetRaw(m)
	return tu
}

// Mutation returns the TokensMutation object of the builder.
func (tu *TokensUpdate) Mutation() *TokensMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokensUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokensMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokensUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokensUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokensUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TokensUpdate) check() error {
	if v, ok := tu.mutation.Email(); ok {
		if err := tokens.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Tokens.email": %w`, err)}
		}
	}
	if v, ok := tu.mutation.AccessToken(); ok {
		if err := tokens.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Tokens.access_token": %w`, err)}
		}
	}
	if v, ok := tu.mutation.TokenType(); ok {
		if err := tokens.TokenTypeValidator(v); err != nil {
			return &ValidationError{Name: "token_type", err: fmt.Errorf(`ent: validator failed for field "Tokens.token_type": %w`, err)}
		}
	}
	return nil
}

func (tu *TokensUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tokens.Table,
			Columns: tokens.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tokens.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Email(); ok {
		_spec.SetField(tokens.FieldEmail, field.TypeString, value)
	}
	if value, ok := tu.mutation.AccessToken(); ok {
		_spec.SetField(tokens.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := tu.mutation.RefreshToken(); ok {
		_spec.SetField(tokens.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := tu.mutation.TokenType(); ok {
		_spec.SetField(tokens.FieldTokenType, field.TypeString, value)
	}
	if value, ok := tu.mutation.Expiry(); ok {
		_spec.SetField(tokens.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Raw(); ok {
		_spec.SetField(tokens.FieldRaw, field.TypeJSON, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TokensUpdateOne is the builder for updating a single Tokens entity.
type TokensUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokensMutation
}

// SetEmail sets the "email" field.
func (tuo *TokensUpdateOne) SetEmail(s string) *TokensUpdateOne {
	tuo.mutation.SetEmail(s)
	return tuo
}

// SetAccessToken sets the "access_token" field.
func (tuo *TokensUpdateOne) SetAccessToken(s string) *TokensUpdateOne {
	tuo.mutation.SetAccessToken(s)
	return tuo
}

// SetRefreshToken sets the "refresh_token" field.
func (tuo *TokensUpdateOne) SetRefreshToken(s string) *TokensUpdateOne {
	tuo.mutation.SetRefreshToken(s)
	return tuo
}

// SetTokenType sets the "token_type" field.
func (tuo *TokensUpdateOne) SetTokenType(s string) *TokensUpdateOne {
	tuo.mutation.SetTokenType(s)
	return tuo
}

// SetExpiry sets the "expiry" field.
func (tuo *TokensUpdateOne) SetExpiry(t time.Time) *TokensUpdateOne {
	tuo.mutation.SetExpiry(t)
	return tuo
}

// SetRaw sets the "raw" field.
func (tuo *TokensUpdateOne) SetRaw(m map[string]interface{}) *TokensUpdateOne {
	tuo.mutation.SetRaw(m)
	return tuo
}

// Mutation returns the TokensMutation object of the builder.
func (tuo *TokensUpdateOne) Mutation() *TokensMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokensUpdateOne) Select(field string, fields ...string) *TokensUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tokens entity.
func (tuo *TokensUpdateOne) Save(ctx context.Context) (*Tokens, error) {
	var (
		err  error
		node *Tokens
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokensMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokensUpdateOne) SaveX(ctx context.Context) *Tokens {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokensUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokensUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TokensUpdateOne) check() error {
	if v, ok := tuo.mutation.Email(); ok {
		if err := tokens.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Tokens.email": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.AccessToken(); ok {
		if err := tokens.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Tokens.access_token": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.TokenType(); ok {
		if err := tokens.TokenTypeValidator(v); err != nil {
			return &ValidationError{Name: "token_type", err: fmt.Errorf(`ent: validator failed for field "Tokens.token_type": %w`, err)}
		}
	}
	return nil
}

func (tuo *TokensUpdateOne) sqlSave(ctx context.Context) (_node *Tokens, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tokens.Table,
			Columns: tokens.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tokens.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tokens.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tokens.FieldID)
		for _, f := range fields {
			if !tokens.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tokens.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Email(); ok {
		_spec.SetField(tokens.FieldEmail, field.TypeString, value)
	}
	if value, ok := tuo.mutation.AccessToken(); ok {
		_spec.SetField(tokens.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := tuo.mutation.RefreshToken(); ok {
		_spec.SetField(tokens.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := tuo.mutation.TokenType(); ok {
		_spec.SetField(tokens.FieldTokenType, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Expiry(); ok {
		_spec.SetField(tokens.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Raw(); ok {
		_spec.SetField(tokens.FieldRaw, field.TypeJSON, value)
	}
	_node = &Tokens{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tokens.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}