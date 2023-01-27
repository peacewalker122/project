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
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/qoute_retweet_feature"
)

// QouteRetweetFeatureCreate is the builder for creating a Qoute_retweet_feature entity.
type QouteRetweetFeatureCreate struct {
	config
	mutation *QouteRetweetFeatureMutation
	hooks    []Hook
}

// SetFromAccountID sets the "from_account_id" field.
func (qrfc *QouteRetweetFeatureCreate) SetFromAccountID(i int64) *QouteRetweetFeatureCreate {
	qrfc.mutation.SetFromAccountID(i)
	return qrfc
}

// SetQouteRetweet sets the "qoute_retweet" field.
func (qrfc *QouteRetweetFeatureCreate) SetQouteRetweet(b bool) *QouteRetweetFeatureCreate {
	qrfc.mutation.SetQouteRetweet(b)
	return qrfc
}

// SetNillableQouteRetweet sets the "qoute_retweet" field if the given value is not nil.
func (qrfc *QouteRetweetFeatureCreate) SetNillableQouteRetweet(b *bool) *QouteRetweetFeatureCreate {
	if b != nil {
		qrfc.SetQouteRetweet(*b)
	}
	return qrfc
}

// SetQoute sets the "qoute" field.
func (qrfc *QouteRetweetFeatureCreate) SetQoute(s string) *QouteRetweetFeatureCreate {
	qrfc.mutation.SetQoute(s)
	return qrfc
}

// SetPostID sets the "post_id" field.
func (qrfc *QouteRetweetFeatureCreate) SetPostID(u uuid.UUID) *QouteRetweetFeatureCreate {
	qrfc.mutation.SetPostID(u)
	return qrfc
}

// SetCreatedAt sets the "created_at" field.
func (qrfc *QouteRetweetFeatureCreate) SetCreatedAt(t time.Time) *QouteRetweetFeatureCreate {
	qrfc.mutation.SetCreatedAt(t)
	return qrfc
}

// Mutation returns the QouteRetweetFeatureMutation object of the builder.
func (qrfc *QouteRetweetFeatureCreate) Mutation() *QouteRetweetFeatureMutation {
	return qrfc.mutation
}

// Save creates the Qoute_retweet_feature in the database.
func (qrfc *QouteRetweetFeatureCreate) Save(ctx context.Context) (*Qoute_retweet_feature, error) {
	var (
		err  error
		node *Qoute_retweet_feature
	)
	qrfc.defaults()
	if len(qrfc.hooks) == 0 {
		if err = qrfc.check(); err != nil {
			return nil, err
		}
		node, err = qrfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QouteRetweetFeatureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qrfc.check(); err != nil {
				return nil, err
			}
			qrfc.mutation = mutation
			if node, err = qrfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(qrfc.hooks) - 1; i >= 0; i-- {
			if qrfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qrfc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, qrfc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Qoute_retweet_feature)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from QouteRetweetFeatureMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qrfc *QouteRetweetFeatureCreate) SaveX(ctx context.Context) *Qoute_retweet_feature {
	v, err := qrfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qrfc *QouteRetweetFeatureCreate) Exec(ctx context.Context) error {
	_, err := qrfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qrfc *QouteRetweetFeatureCreate) ExecX(ctx context.Context) {
	if err := qrfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qrfc *QouteRetweetFeatureCreate) defaults() {
	if _, ok := qrfc.mutation.QouteRetweet(); !ok {
		v := qoute_retweet_feature.DefaultQouteRetweet
		qrfc.mutation.SetQouteRetweet(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qrfc *QouteRetweetFeatureCreate) check() error {
	if _, ok := qrfc.mutation.FromAccountID(); !ok {
		return &ValidationError{Name: "from_account_id", err: errors.New(`ent: missing required field "Qoute_retweet_feature.from_account_id"`)}
	}
	if _, ok := qrfc.mutation.QouteRetweet(); !ok {
		return &ValidationError{Name: "qoute_retweet", err: errors.New(`ent: missing required field "Qoute_retweet_feature.qoute_retweet"`)}
	}
	if _, ok := qrfc.mutation.Qoute(); !ok {
		return &ValidationError{Name: "qoute", err: errors.New(`ent: missing required field "Qoute_retweet_feature.qoute"`)}
	}
	if v, ok := qrfc.mutation.Qoute(); ok {
		if err := qoute_retweet_feature.QouteValidator(v); err != nil {
			return &ValidationError{Name: "qoute", err: fmt.Errorf(`ent: validator failed for field "Qoute_retweet_feature.qoute": %w`, err)}
		}
	}
	if _, ok := qrfc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "Qoute_retweet_feature.post_id"`)}
	}
	if _, ok := qrfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Qoute_retweet_feature.created_at"`)}
	}
	return nil
}

func (qrfc *QouteRetweetFeatureCreate) sqlSave(ctx context.Context) (*Qoute_retweet_feature, error) {
	_node, _spec := qrfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qrfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (qrfc *QouteRetweetFeatureCreate) createSpec() (*Qoute_retweet_feature, *sqlgraph.CreateSpec) {
	var (
		_node = &Qoute_retweet_feature{config: qrfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: qoute_retweet_feature.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: qoute_retweet_feature.FieldID,
			},
		}
	)
	if value, ok := qrfc.mutation.FromAccountID(); ok {
		_spec.SetField(qoute_retweet_feature.FieldFromAccountID, field.TypeInt64, value)
		_node.FromAccountID = value
	}
	if value, ok := qrfc.mutation.QouteRetweet(); ok {
		_spec.SetField(qoute_retweet_feature.FieldQouteRetweet, field.TypeBool, value)
		_node.QouteRetweet = value
	}
	if value, ok := qrfc.mutation.Qoute(); ok {
		_spec.SetField(qoute_retweet_feature.FieldQoute, field.TypeString, value)
		_node.Qoute = value
	}
	if value, ok := qrfc.mutation.PostID(); ok {
		_spec.SetField(qoute_retweet_feature.FieldPostID, field.TypeUUID, value)
		_node.PostID = value
	}
	if value, ok := qrfc.mutation.CreatedAt(); ok {
		_spec.SetField(qoute_retweet_feature.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// QouteRetweetFeatureCreateBulk is the builder for creating many Qoute_retweet_feature entities in bulk.
type QouteRetweetFeatureCreateBulk struct {
	config
	builders []*QouteRetweetFeatureCreate
}

// Save creates the Qoute_retweet_feature entities in the database.
func (qrfcb *QouteRetweetFeatureCreateBulk) Save(ctx context.Context) ([]*Qoute_retweet_feature, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qrfcb.builders))
	nodes := make([]*Qoute_retweet_feature, len(qrfcb.builders))
	mutators := make([]Mutator, len(qrfcb.builders))
	for i := range qrfcb.builders {
		func(i int, root context.Context) {
			builder := qrfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QouteRetweetFeatureMutation)
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
					_, err = mutators[i+1].Mutate(root, qrfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qrfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qrfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qrfcb *QouteRetweetFeatureCreateBulk) SaveX(ctx context.Context) []*Qoute_retweet_feature {
	v, err := qrfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qrfcb *QouteRetweetFeatureCreateBulk) Exec(ctx context.Context) error {
	_, err := qrfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qrfcb *QouteRetweetFeatureCreateBulk) ExecX(ctx context.Context) {
	if err := qrfcb.Exec(ctx); err != nil {
		panic(err)
	}
}