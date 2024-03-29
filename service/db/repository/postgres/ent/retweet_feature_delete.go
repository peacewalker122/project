// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/predicate"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/retweet_feature"
)

// RetweetFeatureDelete is the builder for deleting a Retweet_feature entity.
type RetweetFeatureDelete struct {
	config
	hooks    []Hook
	mutation *RetweetFeatureMutation
}

// Where appends a list predicates to the RetweetFeatureDelete builder.
func (rfd *RetweetFeatureDelete) Where(ps ...predicate.Retweet_feature) *RetweetFeatureDelete {
	rfd.mutation.Where(ps...)
	return rfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rfd *RetweetFeatureDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rfd.hooks) == 0 {
		affected, err = rfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RetweetFeatureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rfd.mutation = mutation
			affected, err = rfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rfd.hooks) - 1; i >= 0; i-- {
			if rfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfd *RetweetFeatureDelete) ExecX(ctx context.Context) int {
	n, err := rfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rfd *RetweetFeatureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: retweet_feature.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: retweet_feature.FieldID,
			},
		},
	}
	if ps := rfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// RetweetFeatureDeleteOne is the builder for deleting a single Retweet_feature entity.
type RetweetFeatureDeleteOne struct {
	rfd *RetweetFeatureDelete
}

// Exec executes the deletion query.
func (rfdo *RetweetFeatureDeleteOne) Exec(ctx context.Context) error {
	n, err := rfdo.rfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{retweet_feature.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rfdo *RetweetFeatureDeleteOne) ExecX(ctx context.Context) {
	rfdo.rfd.ExecX(ctx)
}
