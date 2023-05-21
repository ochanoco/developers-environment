// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/proxy/ent/predicate"
	"github.com/ochanoco/proxy/ent/servicelog"
)

// ServiceLogUpdate is the builder for updating ServiceLog entities.
type ServiceLogUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceLogMutation
}

// Where appends a list predicates to the ServiceLogUpdate builder.
func (slu *ServiceLogUpdate) Where(ps ...predicate.ServiceLog) *ServiceLogUpdate {
	slu.mutation.Where(ps...)
	return slu
}

// SetHeaders sets the "headers" field.
func (slu *ServiceLogUpdate) SetHeaders(s string) *ServiceLogUpdate {
	slu.mutation.SetHeaders(s)
	return slu
}

// SetBody sets the "body" field.
func (slu *ServiceLogUpdate) SetBody(b []byte) *ServiceLogUpdate {
	slu.mutation.SetBody(b)
	return slu
}

// Mutation returns the ServiceLogMutation object of the builder.
func (slu *ServiceLogUpdate) Mutation() *ServiceLogMutation {
	return slu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (slu *ServiceLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(slu.hooks) == 0 {
		affected, err = slu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			slu.mutation = mutation
			affected, err = slu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(slu.hooks) - 1; i >= 0; i-- {
			if slu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = slu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, slu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (slu *ServiceLogUpdate) SaveX(ctx context.Context) int {
	affected, err := slu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (slu *ServiceLogUpdate) Exec(ctx context.Context) error {
	_, err := slu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slu *ServiceLogUpdate) ExecX(ctx context.Context) {
	if err := slu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (slu *ServiceLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   servicelog.Table,
			Columns: servicelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicelog.FieldID,
			},
		},
	}
	if ps := slu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := slu.mutation.Headers(); ok {
		_spec.SetField(servicelog.FieldHeaders, field.TypeString, value)
	}
	if value, ok := slu.mutation.Body(); ok {
		_spec.SetField(servicelog.FieldBody, field.TypeBytes, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, slu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ServiceLogUpdateOne is the builder for updating a single ServiceLog entity.
type ServiceLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceLogMutation
}

// SetHeaders sets the "headers" field.
func (sluo *ServiceLogUpdateOne) SetHeaders(s string) *ServiceLogUpdateOne {
	sluo.mutation.SetHeaders(s)
	return sluo
}

// SetBody sets the "body" field.
func (sluo *ServiceLogUpdateOne) SetBody(b []byte) *ServiceLogUpdateOne {
	sluo.mutation.SetBody(b)
	return sluo
}

// Mutation returns the ServiceLogMutation object of the builder.
func (sluo *ServiceLogUpdateOne) Mutation() *ServiceLogMutation {
	return sluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sluo *ServiceLogUpdateOne) Select(field string, fields ...string) *ServiceLogUpdateOne {
	sluo.fields = append([]string{field}, fields...)
	return sluo
}

// Save executes the query and returns the updated ServiceLog entity.
func (sluo *ServiceLogUpdateOne) Save(ctx context.Context) (*ServiceLog, error) {
	var (
		err  error
		node *ServiceLog
	)
	if len(sluo.hooks) == 0 {
		node, err = sluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sluo.mutation = mutation
			node, err = sluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sluo.hooks) - 1; i >= 0; i-- {
			if sluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ServiceLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ServiceLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sluo *ServiceLogUpdateOne) SaveX(ctx context.Context) *ServiceLog {
	node, err := sluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sluo *ServiceLogUpdateOne) Exec(ctx context.Context) error {
	_, err := sluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sluo *ServiceLogUpdateOne) ExecX(ctx context.Context) {
	if err := sluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sluo *ServiceLogUpdateOne) sqlSave(ctx context.Context) (_node *ServiceLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   servicelog.Table,
			Columns: servicelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicelog.FieldID,
			},
		},
	}
	id, ok := sluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ServiceLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servicelog.FieldID)
		for _, f := range fields {
			if !servicelog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != servicelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sluo.mutation.Headers(); ok {
		_spec.SetField(servicelog.FieldHeaders, field.TypeString, value)
	}
	if value, ok := sluo.mutation.Body(); ok {
		_spec.SetField(servicelog.FieldBody, field.TypeBytes, value)
	}
	_node = &ServiceLog{config: sluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
