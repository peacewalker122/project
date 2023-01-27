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
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/likefeature"
)

// LikeFeatureCreate is the builder for creating a LikeFeature entity.
type LikeFeatureCreate struct {
	config
	mutation *LikeFeatureMutation
	hooks    []Hook
}

// SetFromAccountID sets the "from_account_id" field.
func (lfc *LikeFeatureCreate) SetFromAccountID(i int64) *LikeFeatureCreate {
	lfc.mutation.SetFromAccountID(i)
	return lfc
}

// SetIsLike sets the "is_like" field.
func (lfc *LikeFeatureCreate) SetIsLike(b bool) *LikeFeatureCreate {
	lfc.mutation.SetIsLike(b)
	return lfc
}

// SetNillableIsLike sets the "is_like" field if the given value is not nil.
func (lfc *LikeFeatureCreate) SetNillableIsLike(b *bool) *LikeFeatureCreate {
	if b != nil {
		lfc.SetIsLike(*b)
	}
	return lfc
}

// SetPostID sets the "post_id" field.
func (lfc *LikeFeatureCreate) SetPostID(u uuid.UUID) *LikeFeatureCreate {
	lfc.mutation.SetPostID(u)
	return lfc
}

// SetCreatedAt sets the "created_at" field.
func (lfc *LikeFeatureCreate) SetCreatedAt(t time.Time) *LikeFeatureCreate {
	lfc.mutation.SetCreatedAt(t)
	return lfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lfc *LikeFeatureCreate) SetNillableCreatedAt(t *time.Time) *LikeFeatureCreate {
	if t != nil {
		lfc.SetCreatedAt(*t)
	}
	return lfc
}

// Mutation returns the LikeFeatureMutation object of the builder.
func (lfc *LikeFeatureCreate) Mutation() *LikeFeatureMutation {
	return lfc.mutation
}

// Save creates the LikeFeature in the database.
func (lfc *LikeFeatureCreate) Save(ctx context.Context) (*LikeFeature, error) {
	var (
		err  error
		node *LikeFeature
	)
	lfc.defaults()
	if len(lfc.hooks) == 0 {
		if err = lfc.check(); err != nil {
			return nil, err
		}
		node, err = lfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeFeatureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lfc.check(); err != nil {
				return nil, err
			}
			lfc.mutation = mutation
			if node, err = lfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lfc.hooks) - 1; i >= 0; i-- {
			if lfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lfc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lfc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LikeFeature)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LikeFeatureMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lfc *LikeFeatureCreate) SaveX(ctx context.Context) *LikeFeature {
	v, err := lfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lfc *LikeFeatureCreate) Exec(ctx context.Context) error {
	_, err := lfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lfc *LikeFeatureCreate) ExecX(ctx context.Context) {
	if err := lfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lfc *LikeFeatureCreate) defaults() {
	if _, ok := lfc.mutation.IsLike(); !ok {
		v := likefeature.DefaultIsLike
		lfc.mutation.SetIsLike(v)
	}
	if _, ok := lfc.mutation.CreatedAt(); !ok {
		v := likefeature.DefaultCreatedAt
		lfc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lfc *LikeFeatureCreate) check() error {
	if _, ok := lfc.mutation.FromAccountID(); !ok {
		return &ValidationError{Name: "from_account_id", err: errors.New(`ent: missing required field "LikeFeature.from_account_id"`)}
	}
	if _, ok := lfc.mutation.IsLike(); !ok {
		return &ValidationError{Name: "is_like", err: errors.New(`ent: missing required field "LikeFeature.is_like"`)}
	}
	if _, ok := lfc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "LikeFeature.post_id"`)}
	}
	if _, ok := lfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LikeFeature.created_at"`)}
	}
	return nil
}

func (lfc *LikeFeatureCreate) sqlSave(ctx context.Context) (*LikeFeature, error) {
	_node, _spec := lfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lfc *LikeFeatureCreate) createSpec() (*LikeFeature, *sqlgraph.CreateSpec) {
	var (
		_node = &LikeFeature{config: lfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: likefeature.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: likefeature.FieldID,
			},
		}
	)
	if value, ok := lfc.mutation.FromAccountID(); ok {
		_spec.SetField(likefeature.FieldFromAccountID, field.TypeInt64, value)
		_node.FromAccountID = value
	}
	if value, ok := lfc.mutation.IsLike(); ok {
		_spec.SetField(likefeature.FieldIsLike, field.TypeBool, value)
		_node.IsLike = value
	}
	if value, ok := lfc.mutation.PostID(); ok {
		_spec.SetField(likefeature.FieldPostID, field.TypeUUID, value)
		_node.PostID = value
	}
	if value, ok := lfc.mutation.CreatedAt(); ok {
		_spec.SetField(likefeature.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// LikeFeatureCreateBulk is the builder for creating many LikeFeature entities in bulk.
type LikeFeatureCreateBulk struct {
	config
	builders []*LikeFeatureCreate
}

// Save creates the LikeFeature entities in the database.
func (lfcb *LikeFeatureCreateBulk) Save(ctx context.Context) ([]*LikeFeature, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lfcb.builders))
	nodes := make([]*LikeFeature, len(lfcb.builders))
	mutators := make([]Mutator, len(lfcb.builders))
	for i := range lfcb.builders {
		func(i int, root context.Context) {
			builder := lfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikeFeatureMutation)
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
					_, err = mutators[i+1].Mutate(root, lfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lfcb *LikeFeatureCreateBulk) SaveX(ctx context.Context) []*LikeFeature {
	v, err := lfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lfcb *LikeFeatureCreateBulk) Exec(ctx context.Context) error {
	_, err := lfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lfcb *LikeFeatureCreateBulk) ExecX(ctx context.Context) {
	if err := lfcb.Exec(ctx); err != nil {
		panic(err)
	}
}