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
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/post"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/predicate"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetOwner sets the "owner" field.
func (pu *PostUpdate) SetOwner(s string) *PostUpdate {
	pu.mutation.SetOwner(s)
	return pu
}

// SetIsPrivate sets the "is_private" field.
func (pu *PostUpdate) SetIsPrivate(b bool) *PostUpdate {
	pu.mutation.SetIsPrivate(b)
	return pu
}

// SetNillableIsPrivate sets the "is_private" field if the given value is not nil.
func (pu *PostUpdate) SetNillableIsPrivate(b *bool) *PostUpdate {
	if b != nil {
		pu.SetIsPrivate(*b)
	}
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PostUpdate) SetCreatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCreatedAt(t *time.Time) *PostUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetFollower sets the "follower" field.
func (pu *PostUpdate) SetFollower(i int64) *PostUpdate {
	pu.mutation.ResetFollower()
	pu.mutation.SetFollower(i)
	return pu
}

// SetNillableFollower sets the "follower" field if the given value is not nil.
func (pu *PostUpdate) SetNillableFollower(i *int64) *PostUpdate {
	if i != nil {
		pu.SetFollower(*i)
	}
	return pu
}

// AddFollower adds i to the "follower" field.
func (pu *PostUpdate) AddFollower(i int64) *PostUpdate {
	pu.mutation.AddFollower(i)
	return pu
}

// SetFollowing sets the "following" field.
func (pu *PostUpdate) SetFollowing(i int64) *PostUpdate {
	pu.mutation.ResetFollowing()
	pu.mutation.SetFollowing(i)
	return pu
}

// SetNillableFollowing sets the "following" field if the given value is not nil.
func (pu *PostUpdate) SetNillableFollowing(i *int64) *PostUpdate {
	if i != nil {
		pu.SetFollowing(*i)
	}
	return pu
}

// AddFollowing adds i to the "following" field.
func (pu *PostUpdate) AddFollowing(i int64) *PostUpdate {
	pu.mutation.AddFollowing(i)
	return pu
}

// SetPhotoDir sets the "photo_dir" field.
func (pu *PostUpdate) SetPhotoDir(s string) *PostUpdate {
	pu.mutation.SetPhotoDir(s)
	return pu
}

// SetNillablePhotoDir sets the "photo_dir" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePhotoDir(s *string) *PostUpdate {
	if s != nil {
		pu.SetPhotoDir(*s)
	}
	return pu
}

// ClearPhotoDir clears the value of the "photo_dir" field.
func (pu *PostUpdate) ClearPhotoDir() *PostUpdate {
	pu.mutation.ClearPhotoDir()
	return pu
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Owner(); ok {
		if err := post.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "Post.owner": %w`, err)}
		}
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: post.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Owner(); ok {
		_spec.SetField(post.FieldOwner, field.TypeString, value)
	}
	if value, ok := pu.mutation.IsPrivate(); ok {
		_spec.SetField(post.FieldIsPrivate, field.TypeBool, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Follower(); ok {
		_spec.SetField(post.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedFollower(); ok {
		_spec.AddField(post.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Following(); ok {
		_spec.SetField(post.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedFollowing(); ok {
		_spec.AddField(post.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.PhotoDir(); ok {
		_spec.SetField(post.FieldPhotoDir, field.TypeString, value)
	}
	if pu.mutation.PhotoDirCleared() {
		_spec.ClearField(post.FieldPhotoDir, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetOwner sets the "owner" field.
func (puo *PostUpdateOne) SetOwner(s string) *PostUpdateOne {
	puo.mutation.SetOwner(s)
	return puo
}

// SetIsPrivate sets the "is_private" field.
func (puo *PostUpdateOne) SetIsPrivate(b bool) *PostUpdateOne {
	puo.mutation.SetIsPrivate(b)
	return puo
}

// SetNillableIsPrivate sets the "is_private" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableIsPrivate(b *bool) *PostUpdateOne {
	if b != nil {
		puo.SetIsPrivate(*b)
	}
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PostUpdateOne) SetCreatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCreatedAt(t *time.Time) *PostUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetFollower sets the "follower" field.
func (puo *PostUpdateOne) SetFollower(i int64) *PostUpdateOne {
	puo.mutation.ResetFollower()
	puo.mutation.SetFollower(i)
	return puo
}

// SetNillableFollower sets the "follower" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableFollower(i *int64) *PostUpdateOne {
	if i != nil {
		puo.SetFollower(*i)
	}
	return puo
}

// AddFollower adds i to the "follower" field.
func (puo *PostUpdateOne) AddFollower(i int64) *PostUpdateOne {
	puo.mutation.AddFollower(i)
	return puo
}

// SetFollowing sets the "following" field.
func (puo *PostUpdateOne) SetFollowing(i int64) *PostUpdateOne {
	puo.mutation.ResetFollowing()
	puo.mutation.SetFollowing(i)
	return puo
}

// SetNillableFollowing sets the "following" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableFollowing(i *int64) *PostUpdateOne {
	if i != nil {
		puo.SetFollowing(*i)
	}
	return puo
}

// AddFollowing adds i to the "following" field.
func (puo *PostUpdateOne) AddFollowing(i int64) *PostUpdateOne {
	puo.mutation.AddFollowing(i)
	return puo
}

// SetPhotoDir sets the "photo_dir" field.
func (puo *PostUpdateOne) SetPhotoDir(s string) *PostUpdateOne {
	puo.mutation.SetPhotoDir(s)
	return puo
}

// SetNillablePhotoDir sets the "photo_dir" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePhotoDir(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetPhotoDir(*s)
	}
	return puo
}

// ClearPhotoDir clears the value of the "photo_dir" field.
func (puo *PostUpdateOne) ClearPhotoDir() *PostUpdateOne {
	puo.mutation.ClearPhotoDir()
	return puo
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Post)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Owner(); ok {
		if err := post.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "Post.owner": %w`, err)}
		}
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: post.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Owner(); ok {
		_spec.SetField(post.FieldOwner, field.TypeString, value)
	}
	if value, ok := puo.mutation.IsPrivate(); ok {
		_spec.SetField(post.FieldIsPrivate, field.TypeBool, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Follower(); ok {
		_spec.SetField(post.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedFollower(); ok {
		_spec.AddField(post.FieldFollower, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Following(); ok {
		_spec.SetField(post.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedFollowing(); ok {
		_spec.AddField(post.FieldFollowing, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.PhotoDir(); ok {
		_spec.SetField(post.FieldPhotoDir, field.TypeString, value)
	}
	if puo.mutation.PhotoDirCleared() {
		_spec.ClearField(post.FieldPhotoDir, field.TypeString)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
