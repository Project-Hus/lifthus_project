// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hus-auth/ent/community"
	"hus-auth/ent/predicate"
	"hus-auth/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CommunityUpdate is the builder for updating Community entities.
type CommunityUpdate struct {
	config
	hooks    []Hook
	mutation *CommunityMutation
}

// Where appends a list predicates to the CommunityUpdate builder.
func (cu *CommunityUpdate) Where(ps ...predicate.Community) *CommunityUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CommunityUpdate) SetName(s string) *CommunityUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cu *CommunityUpdate) AddUserIDs(ids ...uuid.UUID) *CommunityUpdate {
	cu.mutation.AddUserIDs(ids...)
	return cu
}

// AddUsers adds the "users" edges to the User entity.
func (cu *CommunityUpdate) AddUsers(u ...*User) *CommunityUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserIDs(ids...)
}

// Mutation returns the CommunityMutation object of the builder.
func (cu *CommunityUpdate) Mutation() *CommunityMutation {
	return cu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cu *CommunityUpdate) ClearUsers() *CommunityUpdate {
	cu.mutation.ClearUsers()
	return cu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cu *CommunityUpdate) RemoveUserIDs(ids ...uuid.UUID) *CommunityUpdate {
	cu.mutation.RemoveUserIDs(ids...)
	return cu
}

// RemoveUsers removes "users" edges to User entities.
func (cu *CommunityUpdate) RemoveUsers(u ...*User) *CommunityUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommunityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CommunityMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommunityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommunityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommunityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommunityUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := community.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Community.name": %w`, err)}
		}
	}
	return nil
}

func (cu *CommunityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(community.Table, community.Columns, sqlgraph.NewFieldSpec(community.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(community.FieldName, field.TypeString, value)
	}
	if cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.UsersTable,
			Columns: community.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.UsersTable,
			Columns: community.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.UsersTable,
			Columns: community.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{community.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommunityUpdateOne is the builder for updating a single Community entity.
type CommunityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommunityMutation
}

// SetName sets the "name" field.
func (cuo *CommunityUpdateOne) SetName(s string) *CommunityUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cuo *CommunityUpdateOne) AddUserIDs(ids ...uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.AddUserIDs(ids...)
	return cuo
}

// AddUsers adds the "users" edges to the User entity.
func (cuo *CommunityUpdateOne) AddUsers(u ...*User) *CommunityUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserIDs(ids...)
}

// Mutation returns the CommunityMutation object of the builder.
func (cuo *CommunityUpdateOne) Mutation() *CommunityMutation {
	return cuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cuo *CommunityUpdateOne) ClearUsers() *CommunityUpdateOne {
	cuo.mutation.ClearUsers()
	return cuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cuo *CommunityUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *CommunityUpdateOne {
	cuo.mutation.RemoveUserIDs(ids...)
	return cuo
}

// RemoveUsers removes "users" edges to User entities.
func (cuo *CommunityUpdateOne) RemoveUsers(u ...*User) *CommunityUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the CommunityUpdate builder.
func (cuo *CommunityUpdateOne) Where(ps ...predicate.Community) *CommunityUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommunityUpdateOne) Select(field string, fields ...string) *CommunityUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Community entity.
func (cuo *CommunityUpdateOne) Save(ctx context.Context) (*Community, error) {
	return withHooks[*Community, CommunityMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommunityUpdateOne) SaveX(ctx context.Context) *Community {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommunityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommunityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommunityUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := community.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Community.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CommunityUpdateOne) sqlSave(ctx context.Context) (_node *Community, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(community.Table, community.Columns, sqlgraph.NewFieldSpec(community.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Community.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, community.FieldID)
		for _, f := range fields {
			if !community.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != community.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(community.FieldName, field.TypeString, value)
	}
	if cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.UsersTable,
			Columns: community.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.UsersTable,
			Columns: community.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.UsersTable,
			Columns: community.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Community{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{community.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
