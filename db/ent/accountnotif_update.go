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
	"github.com/peacewalker122/project/db/ent/accountnotif"
	"github.com/peacewalker122/project/db/ent/predicate"
)

// AccountNotifUpdate is the builder for updating AccountNotif entities.
type AccountNotifUpdate struct {
	config
	hooks    []Hook
	mutation *AccountNotifMutation
}

// Where appends a list predicates to the AccountNotifUpdate builder.
func (anu *AccountNotifUpdate) Where(ps ...predicate.AccountNotif) *AccountNotifUpdate {
	anu.mutation.Where(ps...)
	return anu
}

// SetAccountID sets the "account_id" field.
func (anu *AccountNotifUpdate) SetAccountID(i int64) *AccountNotifUpdate {
	anu.mutation.ResetAccountID()
	anu.mutation.SetAccountID(i)
	return anu
}

// AddAccountID adds i to the "account_id" field.
func (anu *AccountNotifUpdate) AddAccountID(i int64) *AccountNotifUpdate {
	anu.mutation.AddAccountID(i)
	return anu
}

// SetNotifType sets the "notif_type" field.
func (anu *AccountNotifUpdate) SetNotifType(s string) *AccountNotifUpdate {
	anu.mutation.SetNotifType(s)
	return anu
}

// SetNotifTitle sets the "notif_title" field.
func (anu *AccountNotifUpdate) SetNotifTitle(s string) *AccountNotifUpdate {
	anu.mutation.SetNotifTitle(s)
	return anu
}

// SetNillableNotifTitle sets the "notif_title" field if the given value is not nil.
func (anu *AccountNotifUpdate) SetNillableNotifTitle(s *string) *AccountNotifUpdate {
	if s != nil {
		anu.SetNotifTitle(*s)
	}
	return anu
}

// ClearNotifTitle clears the value of the "notif_title" field.
func (anu *AccountNotifUpdate) ClearNotifTitle() *AccountNotifUpdate {
	anu.mutation.ClearNotifTitle()
	return anu
}

// SetNotifContent sets the "notif_content" field.
func (anu *AccountNotifUpdate) SetNotifContent(s string) *AccountNotifUpdate {
	anu.mutation.SetNotifContent(s)
	return anu
}

// SetNillableNotifContent sets the "notif_content" field if the given value is not nil.
func (anu *AccountNotifUpdate) SetNillableNotifContent(s *string) *AccountNotifUpdate {
	if s != nil {
		anu.SetNotifContent(*s)
	}
	return anu
}

// ClearNotifContent clears the value of the "notif_content" field.
func (anu *AccountNotifUpdate) ClearNotifContent() *AccountNotifUpdate {
	anu.mutation.ClearNotifContent()
	return anu
}

// SetNotifTime sets the "notif_time" field.
func (anu *AccountNotifUpdate) SetNotifTime(t time.Time) *AccountNotifUpdate {
	anu.mutation.SetNotifTime(t)
	return anu
}

// SetNillableNotifTime sets the "notif_time" field if the given value is not nil.
func (anu *AccountNotifUpdate) SetNillableNotifTime(t *time.Time) *AccountNotifUpdate {
	if t != nil {
		anu.SetNotifTime(*t)
	}
	return anu
}

// ClearNotifTime clears the value of the "notif_time" field.
func (anu *AccountNotifUpdate) ClearNotifTime() *AccountNotifUpdate {
	anu.mutation.ClearNotifTime()
	return anu
}

// SetCreatedAt sets the "created_at" field.
func (anu *AccountNotifUpdate) SetCreatedAt(t time.Time) *AccountNotifUpdate {
	anu.mutation.SetCreatedAt(t)
	return anu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (anu *AccountNotifUpdate) SetNillableCreatedAt(t *time.Time) *AccountNotifUpdate {
	if t != nil {
		anu.SetCreatedAt(*t)
	}
	return anu
}

// Mutation returns the AccountNotifMutation object of the builder.
func (anu *AccountNotifUpdate) Mutation() *AccountNotifMutation {
	return anu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (anu *AccountNotifUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(anu.hooks) == 0 {
		if err = anu.check(); err != nil {
			return 0, err
		}
		affected, err = anu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountNotifMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = anu.check(); err != nil {
				return 0, err
			}
			anu.mutation = mutation
			affected, err = anu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(anu.hooks) - 1; i >= 0; i-- {
			if anu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = anu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, anu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (anu *AccountNotifUpdate) SaveX(ctx context.Context) int {
	affected, err := anu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (anu *AccountNotifUpdate) Exec(ctx context.Context) error {
	_, err := anu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (anu *AccountNotifUpdate) ExecX(ctx context.Context) {
	if err := anu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (anu *AccountNotifUpdate) check() error {
	if v, ok := anu.mutation.NotifType(); ok {
		if err := accountnotif.NotifTypeValidator(v); err != nil {
			return &ValidationError{Name: "notif_type", err: fmt.Errorf(`ent: validator failed for field "AccountNotif.notif_type": %w`, err)}
		}
	}
	if v, ok := anu.mutation.NotifTitle(); ok {
		if err := accountnotif.NotifTitleValidator(v); err != nil {
			return &ValidationError{Name: "notif_title", err: fmt.Errorf(`ent: validator failed for field "AccountNotif.notif_title": %w`, err)}
		}
	}
	if v, ok := anu.mutation.NotifContent(); ok {
		if err := accountnotif.NotifContentValidator(v); err != nil {
			return &ValidationError{Name: "notif_content", err: fmt.Errorf(`ent: validator failed for field "AccountNotif.notif_content": %w`, err)}
		}
	}
	return nil
}

func (anu *AccountNotifUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountnotif.Table,
			Columns: accountnotif.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountnotif.FieldID,
			},
		},
	}
	if ps := anu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := anu.mutation.AccountID(); ok {
		_spec.SetField(accountnotif.FieldAccountID, field.TypeInt64, value)
	}
	if value, ok := anu.mutation.AddedAccountID(); ok {
		_spec.AddField(accountnotif.FieldAccountID, field.TypeInt64, value)
	}
	if value, ok := anu.mutation.NotifType(); ok {
		_spec.SetField(accountnotif.FieldNotifType, field.TypeString, value)
	}
	if value, ok := anu.mutation.NotifTitle(); ok {
		_spec.SetField(accountnotif.FieldNotifTitle, field.TypeString, value)
	}
	if anu.mutation.NotifTitleCleared() {
		_spec.ClearField(accountnotif.FieldNotifTitle, field.TypeString)
	}
	if value, ok := anu.mutation.NotifContent(); ok {
		_spec.SetField(accountnotif.FieldNotifContent, field.TypeString, value)
	}
	if anu.mutation.NotifContentCleared() {
		_spec.ClearField(accountnotif.FieldNotifContent, field.TypeString)
	}
	if value, ok := anu.mutation.NotifTime(); ok {
		_spec.SetField(accountnotif.FieldNotifTime, field.TypeTime, value)
	}
	if anu.mutation.NotifTimeCleared() {
		_spec.ClearField(accountnotif.FieldNotifTime, field.TypeTime)
	}
	if value, ok := anu.mutation.CreatedAt(); ok {
		_spec.SetField(accountnotif.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, anu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountnotif.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountNotifUpdateOne is the builder for updating a single AccountNotif entity.
type AccountNotifUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountNotifMutation
}

// SetAccountID sets the "account_id" field.
func (anuo *AccountNotifUpdateOne) SetAccountID(i int64) *AccountNotifUpdateOne {
	anuo.mutation.ResetAccountID()
	anuo.mutation.SetAccountID(i)
	return anuo
}

// AddAccountID adds i to the "account_id" field.
func (anuo *AccountNotifUpdateOne) AddAccountID(i int64) *AccountNotifUpdateOne {
	anuo.mutation.AddAccountID(i)
	return anuo
}

// SetNotifType sets the "notif_type" field.
func (anuo *AccountNotifUpdateOne) SetNotifType(s string) *AccountNotifUpdateOne {
	anuo.mutation.SetNotifType(s)
	return anuo
}

// SetNotifTitle sets the "notif_title" field.
func (anuo *AccountNotifUpdateOne) SetNotifTitle(s string) *AccountNotifUpdateOne {
	anuo.mutation.SetNotifTitle(s)
	return anuo
}

// SetNillableNotifTitle sets the "notif_title" field if the given value is not nil.
func (anuo *AccountNotifUpdateOne) SetNillableNotifTitle(s *string) *AccountNotifUpdateOne {
	if s != nil {
		anuo.SetNotifTitle(*s)
	}
	return anuo
}

// ClearNotifTitle clears the value of the "notif_title" field.
func (anuo *AccountNotifUpdateOne) ClearNotifTitle() *AccountNotifUpdateOne {
	anuo.mutation.ClearNotifTitle()
	return anuo
}

// SetNotifContent sets the "notif_content" field.
func (anuo *AccountNotifUpdateOne) SetNotifContent(s string) *AccountNotifUpdateOne {
	anuo.mutation.SetNotifContent(s)
	return anuo
}

// SetNillableNotifContent sets the "notif_content" field if the given value is not nil.
func (anuo *AccountNotifUpdateOne) SetNillableNotifContent(s *string) *AccountNotifUpdateOne {
	if s != nil {
		anuo.SetNotifContent(*s)
	}
	return anuo
}

// ClearNotifContent clears the value of the "notif_content" field.
func (anuo *AccountNotifUpdateOne) ClearNotifContent() *AccountNotifUpdateOne {
	anuo.mutation.ClearNotifContent()
	return anuo
}

// SetNotifTime sets the "notif_time" field.
func (anuo *AccountNotifUpdateOne) SetNotifTime(t time.Time) *AccountNotifUpdateOne {
	anuo.mutation.SetNotifTime(t)
	return anuo
}

// SetNillableNotifTime sets the "notif_time" field if the given value is not nil.
func (anuo *AccountNotifUpdateOne) SetNillableNotifTime(t *time.Time) *AccountNotifUpdateOne {
	if t != nil {
		anuo.SetNotifTime(*t)
	}
	return anuo
}

// ClearNotifTime clears the value of the "notif_time" field.
func (anuo *AccountNotifUpdateOne) ClearNotifTime() *AccountNotifUpdateOne {
	anuo.mutation.ClearNotifTime()
	return anuo
}

// SetCreatedAt sets the "created_at" field.
func (anuo *AccountNotifUpdateOne) SetCreatedAt(t time.Time) *AccountNotifUpdateOne {
	anuo.mutation.SetCreatedAt(t)
	return anuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (anuo *AccountNotifUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountNotifUpdateOne {
	if t != nil {
		anuo.SetCreatedAt(*t)
	}
	return anuo
}

// Mutation returns the AccountNotifMutation object of the builder.
func (anuo *AccountNotifUpdateOne) Mutation() *AccountNotifMutation {
	return anuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (anuo *AccountNotifUpdateOne) Select(field string, fields ...string) *AccountNotifUpdateOne {
	anuo.fields = append([]string{field}, fields...)
	return anuo
}

// Save executes the query and returns the updated AccountNotif entity.
func (anuo *AccountNotifUpdateOne) Save(ctx context.Context) (*AccountNotif, error) {
	var (
		err  error
		node *AccountNotif
	)
	if len(anuo.hooks) == 0 {
		if err = anuo.check(); err != nil {
			return nil, err
		}
		node, err = anuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountNotifMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = anuo.check(); err != nil {
				return nil, err
			}
			anuo.mutation = mutation
			node, err = anuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(anuo.hooks) - 1; i >= 0; i-- {
			if anuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = anuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, anuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AccountNotif)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountNotifMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (anuo *AccountNotifUpdateOne) SaveX(ctx context.Context) *AccountNotif {
	node, err := anuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (anuo *AccountNotifUpdateOne) Exec(ctx context.Context) error {
	_, err := anuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (anuo *AccountNotifUpdateOne) ExecX(ctx context.Context) {
	if err := anuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (anuo *AccountNotifUpdateOne) check() error {
	if v, ok := anuo.mutation.NotifType(); ok {
		if err := accountnotif.NotifTypeValidator(v); err != nil {
			return &ValidationError{Name: "notif_type", err: fmt.Errorf(`ent: validator failed for field "AccountNotif.notif_type": %w`, err)}
		}
	}
	if v, ok := anuo.mutation.NotifTitle(); ok {
		if err := accountnotif.NotifTitleValidator(v); err != nil {
			return &ValidationError{Name: "notif_title", err: fmt.Errorf(`ent: validator failed for field "AccountNotif.notif_title": %w`, err)}
		}
	}
	if v, ok := anuo.mutation.NotifContent(); ok {
		if err := accountnotif.NotifContentValidator(v); err != nil {
			return &ValidationError{Name: "notif_content", err: fmt.Errorf(`ent: validator failed for field "AccountNotif.notif_content": %w`, err)}
		}
	}
	return nil
}

func (anuo *AccountNotifUpdateOne) sqlSave(ctx context.Context) (_node *AccountNotif, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountnotif.Table,
			Columns: accountnotif.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountnotif.FieldID,
			},
		},
	}
	id, ok := anuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccountNotif.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := anuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountnotif.FieldID)
		for _, f := range fields {
			if !accountnotif.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accountnotif.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := anuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := anuo.mutation.AccountID(); ok {
		_spec.SetField(accountnotif.FieldAccountID, field.TypeInt64, value)
	}
	if value, ok := anuo.mutation.AddedAccountID(); ok {
		_spec.AddField(accountnotif.FieldAccountID, field.TypeInt64, value)
	}
	if value, ok := anuo.mutation.NotifType(); ok {
		_spec.SetField(accountnotif.FieldNotifType, field.TypeString, value)
	}
	if value, ok := anuo.mutation.NotifTitle(); ok {
		_spec.SetField(accountnotif.FieldNotifTitle, field.TypeString, value)
	}
	if anuo.mutation.NotifTitleCleared() {
		_spec.ClearField(accountnotif.FieldNotifTitle, field.TypeString)
	}
	if value, ok := anuo.mutation.NotifContent(); ok {
		_spec.SetField(accountnotif.FieldNotifContent, field.TypeString, value)
	}
	if anuo.mutation.NotifContentCleared() {
		_spec.ClearField(accountnotif.FieldNotifContent, field.TypeString)
	}
	if value, ok := anuo.mutation.NotifTime(); ok {
		_spec.SetField(accountnotif.FieldNotifTime, field.TypeTime, value)
	}
	if anuo.mutation.NotifTimeCleared() {
		_spec.ClearField(accountnotif.FieldNotifTime, field.TypeTime)
	}
	if value, ok := anuo.mutation.CreatedAt(); ok {
		_spec.SetField(accountnotif.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &AccountNotif{config: anuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, anuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountnotif.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
