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
	"github.com/peacewalker122/project/db/ent/notif"
)

// NotifCreate is the builder for creating a Notif entity.
type NotifCreate struct {
	config
	mutation *NotifMutation
	hooks    []Hook
}

// SetNotifID sets the "notif_id" field.
func (nc *NotifCreate) SetNotifID(u uuid.UUID) *NotifCreate {
	nc.mutation.SetNotifID(u)
	return nc
}

// SetNillableNotifID sets the "notif_id" field if the given value is not nil.
func (nc *NotifCreate) SetNillableNotifID(u *uuid.UUID) *NotifCreate {
	if u != nil {
		nc.SetNotifID(*u)
	}
	return nc
}

// SetAccountID sets the "account_id" field.
func (nc *NotifCreate) SetAccountID(i int64) *NotifCreate {
	nc.mutation.SetAccountID(i)
	return nc
}

// SetNotifType sets the "notif_type" field.
func (nc *NotifCreate) SetNotifType(s string) *NotifCreate {
	nc.mutation.SetNotifType(s)
	return nc
}

// SetNotifTitle sets the "notif_title" field.
func (nc *NotifCreate) SetNotifTitle(s string) *NotifCreate {
	nc.mutation.SetNotifTitle(s)
	return nc
}

// SetNotifContent sets the "notif_content" field.
func (nc *NotifCreate) SetNotifContent(s string) *NotifCreate {
	nc.mutation.SetNotifContent(s)
	return nc
}

// SetNotifTime sets the "notif_time" field.
func (nc *NotifCreate) SetNotifTime(t time.Time) *NotifCreate {
	nc.mutation.SetNotifTime(t)
	return nc
}

// SetNillableNotifTime sets the "notif_time" field if the given value is not nil.
func (nc *NotifCreate) SetNillableNotifTime(t *time.Time) *NotifCreate {
	if t != nil {
		nc.SetNotifTime(*t)
	}
	return nc
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotifCreate) SetCreatedAt(t time.Time) *NotifCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NotifCreate) SetNillableCreatedAt(t *time.Time) *NotifCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// Mutation returns the NotifMutation object of the builder.
func (nc *NotifCreate) Mutation() *NotifMutation {
	return nc.mutation
}

// Save creates the Notif in the database.
func (nc *NotifCreate) Save(ctx context.Context) (*Notif, error) {
	var (
		err  error
		node *Notif
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotifMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Notif)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NotifMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotifCreate) SaveX(ctx context.Context) *Notif {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotifCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotifCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotifCreate) defaults() {
	if _, ok := nc.mutation.NotifID(); !ok {
		v := notif.DefaultNotifID()
		nc.mutation.SetNotifID(v)
	}
	if _, ok := nc.mutation.NotifTime(); !ok {
		v := notif.DefaultNotifTime()
		nc.mutation.SetNotifTime(v)
	}
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := notif.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotifCreate) check() error {
	if _, ok := nc.mutation.NotifID(); !ok {
		return &ValidationError{Name: "notif_id", err: errors.New(`ent: missing required field "Notif.notif_id"`)}
	}
	if _, ok := nc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Notif.account_id"`)}
	}
	if _, ok := nc.mutation.NotifType(); !ok {
		return &ValidationError{Name: "notif_type", err: errors.New(`ent: missing required field "Notif.notif_type"`)}
	}
	if v, ok := nc.mutation.NotifType(); ok {
		if err := notif.NotifTypeValidator(v); err != nil {
			return &ValidationError{Name: "notif_type", err: fmt.Errorf(`ent: validator failed for field "Notif.notif_type": %w`, err)}
		}
	}
	if _, ok := nc.mutation.NotifTitle(); !ok {
		return &ValidationError{Name: "notif_title", err: errors.New(`ent: missing required field "Notif.notif_title"`)}
	}
	if v, ok := nc.mutation.NotifTitle(); ok {
		if err := notif.NotifTitleValidator(v); err != nil {
			return &ValidationError{Name: "notif_title", err: fmt.Errorf(`ent: validator failed for field "Notif.notif_title": %w`, err)}
		}
	}
	if _, ok := nc.mutation.NotifContent(); !ok {
		return &ValidationError{Name: "notif_content", err: errors.New(`ent: missing required field "Notif.notif_content"`)}
	}
	if v, ok := nc.mutation.NotifContent(); ok {
		if err := notif.NotifContentValidator(v); err != nil {
			return &ValidationError{Name: "notif_content", err: fmt.Errorf(`ent: validator failed for field "Notif.notif_content": %w`, err)}
		}
	}
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Notif.created_at"`)}
	}
	return nil
}

func (nc *NotifCreate) sqlSave(ctx context.Context) (*Notif, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nc *NotifCreate) createSpec() (*Notif, *sqlgraph.CreateSpec) {
	var (
		_node = &Notif{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: notif.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notif.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.NotifID(); ok {
		_spec.SetField(notif.FieldNotifID, field.TypeUUID, value)
		_node.NotifID = value
	}
	if value, ok := nc.mutation.AccountID(); ok {
		_spec.SetField(notif.FieldAccountID, field.TypeInt64, value)
		_node.AccountID = value
	}
	if value, ok := nc.mutation.NotifType(); ok {
		_spec.SetField(notif.FieldNotifType, field.TypeString, value)
		_node.NotifType = value
	}
	if value, ok := nc.mutation.NotifTitle(); ok {
		_spec.SetField(notif.FieldNotifTitle, field.TypeString, value)
		_node.NotifTitle = value
	}
	if value, ok := nc.mutation.NotifContent(); ok {
		_spec.SetField(notif.FieldNotifContent, field.TypeString, value)
		_node.NotifContent = value
	}
	if value, ok := nc.mutation.NotifTime(); ok {
		_spec.SetField(notif.FieldNotifTime, field.TypeTime, value)
		_node.NotifTime = &value
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(notif.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// NotifCreateBulk is the builder for creating many Notif entities in bulk.
type NotifCreateBulk struct {
	config
	builders []*NotifCreate
}

// Save creates the Notif entities in the database.
func (ncb *NotifCreateBulk) Save(ctx context.Context) ([]*Notif, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notif, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotifCreateBulk) SaveX(ctx context.Context) []*Notif {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotifCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotifCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
