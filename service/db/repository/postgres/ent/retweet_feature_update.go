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
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/predicate"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/retweet_feature"
)

// RetweetFeatureUpdate is the builder for updating Retweet_feature entities.
type RetweetFeatureUpdate struct {
	config
	hooks    []Hook
	mutation *RetweetFeatureMutation
}

// Where appends a list predicates to the RetweetFeatureUpdate builder.
func (rfu *RetweetFeatureUpdate) Where(ps ...predicate.Retweet_feature) *RetweetFeatureUpdate {
	rfu.mutation.Where(ps...)
	return rfu
}

// SetFromAccountID sets the "from_account_id" field.
func (rfu *RetweetFeatureUpdate) SetFromAccountID(i int64) *RetweetFeatureUpdate {
	rfu.mutation.ResetFromAccountID()
	rfu.mutation.SetFromAccountID(i)
	return rfu
}

// AddFromAccountID adds i to the "from_account_id" field.
func (rfu *RetweetFeatureUpdate) AddFromAccountID(i int64) *RetweetFeatureUpdate {
	rfu.mutation.AddFromAccountID(i)
	return rfu
}

// SetRetweet sets the "retweet" field.
func (rfu *RetweetFeatureUpdate) SetRetweet(b bool) *RetweetFeatureUpdate {
	rfu.mutation.SetRetweet(b)
	return rfu
}

// SetNillableRetweet sets the "retweet" field if the given value is not nil.
func (rfu *RetweetFeatureUpdate) SetNillableRetweet(b *bool) *RetweetFeatureUpdate {
	if b != nil {
		rfu.SetRetweet(*b)
	}
	return rfu
}

// SetPostID sets the "post_id" field.
func (rfu *RetweetFeatureUpdate) SetPostID(u uuid.UUID) *RetweetFeatureUpdate {
	rfu.mutation.SetPostID(u)
	return rfu
}

// SetCreatedAt sets the "created_at" field.
func (rfu *RetweetFeatureUpdate) SetCreatedAt(t time.Time) *RetweetFeatureUpdate {
	rfu.mutation.SetCreatedAt(t)
	return rfu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rfu *RetweetFeatureUpdate) SetNillableCreatedAt(t *time.Time) *RetweetFeatureUpdate {
	if t != nil {
		rfu.SetCreatedAt(*t)
	}
	return rfu
}

// Mutation returns the RetweetFeatureMutation object of the builder.
func (rfu *RetweetFeatureUpdate) Mutation() *RetweetFeatureMutation {
	return rfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rfu *RetweetFeatureUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rfu.hooks) == 0 {
		affected, err = rfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RetweetFeatureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rfu.mutation = mutation
			affected, err = rfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rfu.hooks) - 1; i >= 0; i-- {
			if rfu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rfu *RetweetFeatureUpdate) SaveX(ctx context.Context) int {
	affected, err := rfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rfu *RetweetFeatureUpdate) Exec(ctx context.Context) error {
	_, err := rfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfu *RetweetFeatureUpdate) ExecX(ctx context.Context) {
	if err := rfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rfu *RetweetFeatureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   retweet_feature.Table,
			Columns: retweet_feature.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: retweet_feature.FieldID,
			},
		},
	}
	if ps := rfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rfu.mutation.FromAccountID(); ok {
		_spec.SetField(retweet_feature.FieldFromAccountID, field.TypeInt64, value)
	}
	if value, ok := rfu.mutation.AddedFromAccountID(); ok {
		_spec.AddField(retweet_feature.FieldFromAccountID, field.TypeInt64, value)
	}
	if value, ok := rfu.mutation.Retweet(); ok {
		_spec.SetField(retweet_feature.FieldRetweet, field.TypeBool, value)
	}
	if value, ok := rfu.mutation.PostID(); ok {
		_spec.SetField(retweet_feature.FieldPostID, field.TypeUUID, value)
	}
	if value, ok := rfu.mutation.CreatedAt(); ok {
		_spec.SetField(retweet_feature.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{retweet_feature.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RetweetFeatureUpdateOne is the builder for updating a single Retweet_feature entity.
type RetweetFeatureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RetweetFeatureMutation
}

// SetFromAccountID sets the "from_account_id" field.
func (rfuo *RetweetFeatureUpdateOne) SetFromAccountID(i int64) *RetweetFeatureUpdateOne {
	rfuo.mutation.ResetFromAccountID()
	rfuo.mutation.SetFromAccountID(i)
	return rfuo
}

// AddFromAccountID adds i to the "from_account_id" field.
func (rfuo *RetweetFeatureUpdateOne) AddFromAccountID(i int64) *RetweetFeatureUpdateOne {
	rfuo.mutation.AddFromAccountID(i)
	return rfuo
}

// SetRetweet sets the "retweet" field.
func (rfuo *RetweetFeatureUpdateOne) SetRetweet(b bool) *RetweetFeatureUpdateOne {
	rfuo.mutation.SetRetweet(b)
	return rfuo
}

// SetNillableRetweet sets the "retweet" field if the given value is not nil.
func (rfuo *RetweetFeatureUpdateOne) SetNillableRetweet(b *bool) *RetweetFeatureUpdateOne {
	if b != nil {
		rfuo.SetRetweet(*b)
	}
	return rfuo
}

// SetPostID sets the "post_id" field.
func (rfuo *RetweetFeatureUpdateOne) SetPostID(u uuid.UUID) *RetweetFeatureUpdateOne {
	rfuo.mutation.SetPostID(u)
	return rfuo
}

// SetCreatedAt sets the "created_at" field.
func (rfuo *RetweetFeatureUpdateOne) SetCreatedAt(t time.Time) *RetweetFeatureUpdateOne {
	rfuo.mutation.SetCreatedAt(t)
	return rfuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rfuo *RetweetFeatureUpdateOne) SetNillableCreatedAt(t *time.Time) *RetweetFeatureUpdateOne {
	if t != nil {
		rfuo.SetCreatedAt(*t)
	}
	return rfuo
}

// Mutation returns the RetweetFeatureMutation object of the builder.
func (rfuo *RetweetFeatureUpdateOne) Mutation() *RetweetFeatureMutation {
	return rfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rfuo *RetweetFeatureUpdateOne) Select(field string, fields ...string) *RetweetFeatureUpdateOne {
	rfuo.fields = append([]string{field}, fields...)
	return rfuo
}

// Save executes the query and returns the updated Retweet_feature entity.
func (rfuo *RetweetFeatureUpdateOne) Save(ctx context.Context) (*Retweet_feature, error) {
	var (
		err  error
		node *Retweet_feature
	)
	if len(rfuo.hooks) == 0 {
		node, err = rfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RetweetFeatureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rfuo.mutation = mutation
			node, err = rfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rfuo.hooks) - 1; i >= 0; i-- {
			if rfuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rfuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rfuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Retweet_feature)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RetweetFeatureMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rfuo *RetweetFeatureUpdateOne) SaveX(ctx context.Context) *Retweet_feature {
	node, err := rfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rfuo *RetweetFeatureUpdateOne) Exec(ctx context.Context) error {
	_, err := rfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfuo *RetweetFeatureUpdateOne) ExecX(ctx context.Context) {
	if err := rfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rfuo *RetweetFeatureUpdateOne) sqlSave(ctx context.Context) (_node *Retweet_feature, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   retweet_feature.Table,
			Columns: retweet_feature.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: retweet_feature.FieldID,
			},
		},
	}
	id, ok := rfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Retweet_feature.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, retweet_feature.FieldID)
		for _, f := range fields {
			if !retweet_feature.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != retweet_feature.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rfuo.mutation.FromAccountID(); ok {
		_spec.SetField(retweet_feature.FieldFromAccountID, field.TypeInt64, value)
	}
	if value, ok := rfuo.mutation.AddedFromAccountID(); ok {
		_spec.AddField(retweet_feature.FieldFromAccountID, field.TypeInt64, value)
	}
	if value, ok := rfuo.mutation.Retweet(); ok {
		_spec.SetField(retweet_feature.FieldRetweet, field.TypeBool, value)
	}
	if value, ok := rfuo.mutation.PostID(); ok {
		_spec.SetField(retweet_feature.FieldPostID, field.TypeUUID, value)
	}
	if value, ok := rfuo.mutation.CreatedAt(); ok {
		_spec.SetField(retweet_feature.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Retweet_feature{config: rfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{retweet_feature.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
