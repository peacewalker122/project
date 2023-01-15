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
	"github.com/peacewalker122/project/db/repository/postgres/ent/account"
	"github.com/peacewalker122/project/db/repository/postgres/ent/predicate"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetOwner sets the "owner" field.
func (au *AccountUpdate) SetOwner(s string) *AccountUpdate {
	au.mutation.SetOwner(s)
	return au
}

// SetIsPrivate sets the "is_private" field.
func (au *AccountUpdate) SetIsPrivate(b bool) *AccountUpdate {
	au.mutation.SetIsPrivate(b)
	return au
}

// SetNillableIsPrivate sets the "is_private" field if the given value is not nil.
func (au *AccountUpdate) SetNillableIsPrivate(b *bool) *AccountUpdate {
	if b != nil {
		au.SetIsPrivate(*b)
	}
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccountUpdate) SetCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetFollower sets the "follower" field.
func (au *AccountUpdate) SetFollower(i int64) *AccountUpdate {
	au.mutation.ResetFollower()
	au.mutation.SetFollower(i)
	return au
}

// SetNillableFollower sets the "follower" field if the given value is not nil.
func (au *AccountUpdate) SetNillableFollower(i *int64) *AccountUpdate {
	if i != nil {
		au.SetFollower(*i)
	}
	return au
}

// AddFollower adds i to the "follower" field.
func (au *AccountUpdate) AddFollower(i int64) *AccountUpdate {
	au.mutation.AddFollower(i)
	return au
}

// SetFollowing sets the "following" field.
func (au *AccountUpdate) SetFollowing(i int64) *AccountUpdate {
	au.mutation.ResetFollowing()
	au.mutation.SetFollowing(i)
	return au
}

// SetNillableFollowing sets the "following" field if the given value is not nil.
func (au *AccountUpdate) SetNillableFollowing(i *int64) *AccountUpdate {
	if i != nil {
		au.SetFollowing(*i)
	}
	return au
}

// AddFollowing adds i to the "following" field.
func (au *AccountUpdate) AddFollowing(i int64) *AccountUpdate {
	au.mutation.AddFollowing(i)
	return au
}

// SetPhotoDir sets the "photo_dir" field.
func (au *AccountUpdate) SetPhotoDir(s string) *AccountUpdate {
	au.mutation.SetPhotoDir(s)
	return au
}

// SetNillablePhotoDir sets the "photo_dir" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePhotoDir(s *string) *AccountUpdate {
	if s != nil {
		au.SetPhotoDir(*s)
	}
	return au
}

// ClearPhotoDir clears the value of the "photo_dir" field.
func (au *AccountUpdate) ClearPhotoDir() *AccountUpdate {
	au.mutation.ClearPhotoDir()
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Owner(); ok {
		if err := account.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "Account.owner": %w`, err)}
		}
	}
	if v, ok := au.mutation.PhotoDir(); ok {
		if err := account.PhotoDirValidator(v); err != nil {
			return &ValidationError{Name: "photo_dir", err: fmt.Errorf(`ent: validator failed for field "Account.photo_dir": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Owner(); ok {
		_spec.SetField(account.FieldOwner, field.TypeString, value)
	}
	if value, ok := au.mutation.IsPrivate(); ok {
		_spec.SetField(account.FieldIsPrivate, field.TypeBool, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Follower(); ok {
		_spec.SetField(account.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedFollower(); ok {
		_spec.AddField(account.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := au.mutation.Following(); ok {
		_spec.SetField(account.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedFollowing(); ok {
		_spec.AddField(account.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := au.mutation.PhotoDir(); ok {
		_spec.SetField(account.FieldPhotoDir, field.TypeString, value)
	}
	if au.mutation.PhotoDirCleared() {
		_spec.ClearField(account.FieldPhotoDir, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetOwner sets the "owner" field.
func (auo *AccountUpdateOne) SetOwner(s string) *AccountUpdateOne {
	auo.mutation.SetOwner(s)
	return auo
}

// SetIsPrivate sets the "is_private" field.
func (auo *AccountUpdateOne) SetIsPrivate(b bool) *AccountUpdateOne {
	auo.mutation.SetIsPrivate(b)
	return auo
}

// SetNillableIsPrivate sets the "is_private" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableIsPrivate(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetIsPrivate(*b)
	}
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccountUpdateOne) SetCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetFollower sets the "follower" field.
func (auo *AccountUpdateOne) SetFollower(i int64) *AccountUpdateOne {
	auo.mutation.ResetFollower()
	auo.mutation.SetFollower(i)
	return auo
}

// SetNillableFollower sets the "follower" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableFollower(i *int64) *AccountUpdateOne {
	if i != nil {
		auo.SetFollower(*i)
	}
	return auo
}

// AddFollower adds i to the "follower" field.
func (auo *AccountUpdateOne) AddFollower(i int64) *AccountUpdateOne {
	auo.mutation.AddFollower(i)
	return auo
}

// SetFollowing sets the "following" field.
func (auo *AccountUpdateOne) SetFollowing(i int64) *AccountUpdateOne {
	auo.mutation.ResetFollowing()
	auo.mutation.SetFollowing(i)
	return auo
}

// SetNillableFollowing sets the "following" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableFollowing(i *int64) *AccountUpdateOne {
	if i != nil {
		auo.SetFollowing(*i)
	}
	return auo
}

// AddFollowing adds i to the "following" field.
func (auo *AccountUpdateOne) AddFollowing(i int64) *AccountUpdateOne {
	auo.mutation.AddFollowing(i)
	return auo
}

// SetPhotoDir sets the "photo_dir" field.
func (auo *AccountUpdateOne) SetPhotoDir(s string) *AccountUpdateOne {
	auo.mutation.SetPhotoDir(s)
	return auo
}

// SetNillablePhotoDir sets the "photo_dir" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePhotoDir(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPhotoDir(*s)
	}
	return auo
}

// ClearPhotoDir clears the value of the "photo_dir" field.
func (auo *AccountUpdateOne) ClearPhotoDir() *AccountUpdateOne {
	auo.mutation.ClearPhotoDir()
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Owner(); ok {
		if err := account.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "Account.owner": %w`, err)}
		}
	}
	if v, ok := auo.mutation.PhotoDir(); ok {
		if err := account.PhotoDirValidator(v); err != nil {
			return &ValidationError{Name: "photo_dir", err: fmt.Errorf(`ent: validator failed for field "Account.photo_dir": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Owner(); ok {
		_spec.SetField(account.FieldOwner, field.TypeString, value)
	}
	if value, ok := auo.mutation.IsPrivate(); ok {
		_spec.SetField(account.FieldIsPrivate, field.TypeBool, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Follower(); ok {
		_spec.SetField(account.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedFollower(); ok {
		_spec.AddField(account.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.Following(); ok {
		_spec.SetField(account.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedFollowing(); ok {
		_spec.AddField(account.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.PhotoDir(); ok {
		_spec.SetField(account.FieldPhotoDir, field.TypeString, value)
	}
	if auo.mutation.PhotoDirCleared() {
		_spec.ClearField(account.FieldPhotoDir, field.TypeString)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
