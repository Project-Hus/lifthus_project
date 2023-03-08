// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hus-auth/ent/hussession"
	"hus-auth/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HusSessionCreate is the builder for creating a HusSession entity.
type HusSessionCreate struct {
	config
	mutation *HusSessionMutation
	hooks    []Hook
}

// SetExpiredAt sets the "expired_at" field.
func (hsc *HusSessionCreate) SetExpiredAt(t time.Time) *HusSessionCreate {
	hsc.mutation.SetExpiredAt(t)
	return hsc
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (hsc *HusSessionCreate) SetNillableExpiredAt(t *time.Time) *HusSessionCreate {
	if t != nil {
		hsc.SetExpiredAt(*t)
	}
	return hsc
}

// SetCreatedAt sets the "created_at" field.
func (hsc *HusSessionCreate) SetCreatedAt(t time.Time) *HusSessionCreate {
	hsc.mutation.SetCreatedAt(t)
	return hsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hsc *HusSessionCreate) SetNillableCreatedAt(t *time.Time) *HusSessionCreate {
	if t != nil {
		hsc.SetCreatedAt(*t)
	}
	return hsc
}

// SetID sets the "id" field.
func (hsc *HusSessionCreate) SetID(u uuid.UUID) *HusSessionCreate {
	hsc.mutation.SetID(u)
	return hsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hsc *HusSessionCreate) SetNillableID(u *uuid.UUID) *HusSessionCreate {
	if u != nil {
		hsc.SetID(*u)
	}
	return hsc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (hsc *HusSessionCreate) SetOwnerID(id uuid.UUID) *HusSessionCreate {
	hsc.mutation.SetOwnerID(id)
	return hsc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (hsc *HusSessionCreate) SetNillableOwnerID(id *uuid.UUID) *HusSessionCreate {
	if id != nil {
		hsc = hsc.SetOwnerID(*id)
	}
	return hsc
}

// SetOwner sets the "owner" edge to the User entity.
func (hsc *HusSessionCreate) SetOwner(u *User) *HusSessionCreate {
	return hsc.SetOwnerID(u.ID)
}

// Mutation returns the HusSessionMutation object of the builder.
func (hsc *HusSessionCreate) Mutation() *HusSessionMutation {
	return hsc.mutation
}

// Save creates the HusSession in the database.
func (hsc *HusSessionCreate) Save(ctx context.Context) (*HusSession, error) {
	hsc.defaults()
	return withHooks[*HusSession, HusSessionMutation](ctx, hsc.sqlSave, hsc.mutation, hsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hsc *HusSessionCreate) SaveX(ctx context.Context) *HusSession {
	v, err := hsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hsc *HusSessionCreate) Exec(ctx context.Context) error {
	_, err := hsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hsc *HusSessionCreate) ExecX(ctx context.Context) {
	if err := hsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hsc *HusSessionCreate) defaults() {
	if _, ok := hsc.mutation.CreatedAt(); !ok {
		v := hussession.DefaultCreatedAt()
		hsc.mutation.SetCreatedAt(v)
	}
	if _, ok := hsc.mutation.ID(); !ok {
		v := hussession.DefaultID()
		hsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hsc *HusSessionCreate) check() error {
	if _, ok := hsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HusSession.created_at"`)}
	}
	return nil
}

func (hsc *HusSessionCreate) sqlSave(ctx context.Context) (*HusSession, error) {
	if err := hsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	hsc.mutation.id = &_node.ID
	hsc.mutation.done = true
	return _node, nil
}

func (hsc *HusSessionCreate) createSpec() (*HusSession, *sqlgraph.CreateSpec) {
	var (
		_node = &HusSession{config: hsc.config}
		_spec = sqlgraph.NewCreateSpec(hussession.Table, sqlgraph.NewFieldSpec(hussession.FieldID, field.TypeUUID))
	)
	if id, ok := hsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hsc.mutation.ExpiredAt(); ok {
		_spec.SetField(hussession.FieldExpiredAt, field.TypeTime, value)
		_node.ExpiredAt = &value
	}
	if value, ok := hsc.mutation.CreatedAt(); ok {
		_spec.SetField(hussession.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := hsc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hussession.OwnerTable,
			Columns: []string{hussession.OwnerColumn},
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
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HusSessionCreateBulk is the builder for creating many HusSession entities in bulk.
type HusSessionCreateBulk struct {
	config
	builders []*HusSessionCreate
}

// Save creates the HusSession entities in the database.
func (hscb *HusSessionCreateBulk) Save(ctx context.Context) ([]*HusSession, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hscb.builders))
	nodes := make([]*HusSession, len(hscb.builders))
	mutators := make([]Mutator, len(hscb.builders))
	for i := range hscb.builders {
		func(i int, root context.Context) {
			builder := hscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HusSessionMutation)
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
					_, err = mutators[i+1].Mutate(root, hscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, hscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hscb *HusSessionCreateBulk) SaveX(ctx context.Context) []*HusSession {
	v, err := hscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hscb *HusSessionCreateBulk) Exec(ctx context.Context) error {
	_, err := hscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hscb *HusSessionCreateBulk) ExecX(ctx context.Context) {
	if err := hscb.Exec(ctx); err != nil {
		panic(err)
	}
}
