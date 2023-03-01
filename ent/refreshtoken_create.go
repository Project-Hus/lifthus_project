// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hus-auth/ent/refreshtoken"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RefreshTokenCreate is the builder for creating a RefreshToken entity.
type RefreshTokenCreate struct {
	config
	mutation *RefreshTokenMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (rtc *RefreshTokenCreate) SetUID(s string) *RefreshTokenCreate {
	rtc.mutation.SetUID(s)
	return rtc
}

// SetRevoked sets the "revoked" field.
func (rtc *RefreshTokenCreate) SetRevoked(b bool) *RefreshTokenCreate {
	rtc.mutation.SetRevoked(b)
	return rtc
}

// SetNillableRevoked sets the "revoked" field if the given value is not nil.
func (rtc *RefreshTokenCreate) SetNillableRevoked(b *bool) *RefreshTokenCreate {
	if b != nil {
		rtc.SetRevoked(*b)
	}
	return rtc
}

// SetLastUsedAt sets the "last_used_at" field.
func (rtc *RefreshTokenCreate) SetLastUsedAt(t time.Time) *RefreshTokenCreate {
	rtc.mutation.SetLastUsedAt(t)
	return rtc
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (rtc *RefreshTokenCreate) SetNillableLastUsedAt(t *time.Time) *RefreshTokenCreate {
	if t != nil {
		rtc.SetLastUsedAt(*t)
	}
	return rtc
}

// SetCreatedAt sets the "created_at" field.
func (rtc *RefreshTokenCreate) SetCreatedAt(t time.Time) *RefreshTokenCreate {
	rtc.mutation.SetCreatedAt(t)
	return rtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtc *RefreshTokenCreate) SetNillableCreatedAt(t *time.Time) *RefreshTokenCreate {
	if t != nil {
		rtc.SetCreatedAt(*t)
	}
	return rtc
}

// SetID sets the "id" field.
func (rtc *RefreshTokenCreate) SetID(s string) *RefreshTokenCreate {
	rtc.mutation.SetID(s)
	return rtc
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtc *RefreshTokenCreate) Mutation() *RefreshTokenMutation {
	return rtc.mutation
}

// Save creates the RefreshToken in the database.
func (rtc *RefreshTokenCreate) Save(ctx context.Context) (*RefreshToken, error) {
	rtc.defaults()
	return withHooks[*RefreshToken, RefreshTokenMutation](ctx, rtc.sqlSave, rtc.mutation, rtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rtc *RefreshTokenCreate) SaveX(ctx context.Context) *RefreshToken {
	v, err := rtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rtc *RefreshTokenCreate) Exec(ctx context.Context) error {
	_, err := rtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtc *RefreshTokenCreate) ExecX(ctx context.Context) {
	if err := rtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtc *RefreshTokenCreate) defaults() {
	if _, ok := rtc.mutation.Revoked(); !ok {
		v := refreshtoken.DefaultRevoked
		rtc.mutation.SetRevoked(v)
	}
	if _, ok := rtc.mutation.LastUsedAt(); !ok {
		v := refreshtoken.DefaultLastUsedAt
		rtc.mutation.SetLastUsedAt(v)
	}
	if _, ok := rtc.mutation.CreatedAt(); !ok {
		v := refreshtoken.DefaultCreatedAt
		rtc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtc *RefreshTokenCreate) check() error {
	if _, ok := rtc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "RefreshToken.uid"`)}
	}
	if _, ok := rtc.mutation.Revoked(); !ok {
		return &ValidationError{Name: "revoked", err: errors.New(`ent: missing required field "RefreshToken.revoked"`)}
	}
	if _, ok := rtc.mutation.LastUsedAt(); !ok {
		return &ValidationError{Name: "last_used_at", err: errors.New(`ent: missing required field "RefreshToken.last_used_at"`)}
	}
	if _, ok := rtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RefreshToken.created_at"`)}
	}
	return nil
}

func (rtc *RefreshTokenCreate) sqlSave(ctx context.Context) (*RefreshToken, error) {
	if err := rtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RefreshToken.ID type: %T", _spec.ID.Value)
		}
	}
	rtc.mutation.id = &_node.ID
	rtc.mutation.done = true
	return _node, nil
}

func (rtc *RefreshTokenCreate) createSpec() (*RefreshToken, *sqlgraph.CreateSpec) {
	var (
		_node = &RefreshToken{config: rtc.config}
		_spec = sqlgraph.NewCreateSpec(refreshtoken.Table, sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeString))
	)
	if id, ok := rtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rtc.mutation.UID(); ok {
		_spec.SetField(refreshtoken.FieldUID, field.TypeString, value)
		_node.UID = value
	}
	if value, ok := rtc.mutation.Revoked(); ok {
		_spec.SetField(refreshtoken.FieldRevoked, field.TypeBool, value)
		_node.Revoked = value
	}
	if value, ok := rtc.mutation.LastUsedAt(); ok {
		_spec.SetField(refreshtoken.FieldLastUsedAt, field.TypeTime, value)
		_node.LastUsedAt = value
	}
	if value, ok := rtc.mutation.CreatedAt(); ok {
		_spec.SetField(refreshtoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// RefreshTokenCreateBulk is the builder for creating many RefreshToken entities in bulk.
type RefreshTokenCreateBulk struct {
	config
	builders []*RefreshTokenCreate
}

// Save creates the RefreshToken entities in the database.
func (rtcb *RefreshTokenCreateBulk) Save(ctx context.Context) ([]*RefreshToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rtcb.builders))
	nodes := make([]*RefreshToken, len(rtcb.builders))
	mutators := make([]Mutator, len(rtcb.builders))
	for i := range rtcb.builders {
		func(i int, root context.Context) {
			builder := rtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RefreshTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, rtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rtcb *RefreshTokenCreateBulk) SaveX(ctx context.Context) []*RefreshToken {
	v, err := rtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rtcb *RefreshTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := rtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtcb *RefreshTokenCreateBulk) ExecX(ctx context.Context) {
	if err := rtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
