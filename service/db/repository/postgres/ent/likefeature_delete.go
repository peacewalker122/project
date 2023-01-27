// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/likefeature"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/predicate"
)

// LikeFeatureDelete is the builder for deleting a LikeFeature entity.
type LikeFeatureDelete struct {
	config
	hooks    []Hook
	mutation *LikeFeatureMutation
}

// Where appends a list predicates to the LikeFeatureDelete builder.
func (lfd *LikeFeatureDelete) Where(ps ...predicate.LikeFeature) *LikeFeatureDelete {
	lfd.mutation.Where(ps...)
	return lfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lfd *LikeFeatureDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lfd.hooks) == 0 {
		affected, err = lfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeFeatureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lfd.mutation = mutation
			affected, err = lfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lfd.hooks) - 1; i >= 0; i-- {
			if lfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lfd *LikeFeatureDelete) ExecX(ctx context.Context) int {
	n, err := lfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lfd *LikeFeatureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: likefeature.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: likefeature.FieldID,
			},
		},
	}
	if ps := lfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// LikeFeatureDeleteOne is the builder for deleting a single LikeFeature entity.
type LikeFeatureDeleteOne struct {
	lfd *LikeFeatureDelete
}

// Exec executes the deletion query.
func (lfdo *LikeFeatureDeleteOne) Exec(ctx context.Context) error {
	n, err := lfdo.lfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{likefeature.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lfdo *LikeFeatureDeleteOne) ExecX(ctx context.Context) {
	lfdo.lfd.ExecX(ctx)
}