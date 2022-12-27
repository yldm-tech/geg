// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"myetc.lantron.ltd/m/ent/etcsyncrecord"
	"myetc.lantron.ltd/m/ent/predicate"
)

// ETCSyncRecordDelete is the builder for deleting a ETCSyncRecord entity.
type ETCSyncRecordDelete struct {
	config
	hooks    []Hook
	mutation *ETCSyncRecordMutation
}

// Where appends a list predicates to the ETCSyncRecordDelete builder.
func (esrd *ETCSyncRecordDelete) Where(ps ...predicate.ETCSyncRecord) *ETCSyncRecordDelete {
	esrd.mutation.Where(ps...)
	return esrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (esrd *ETCSyncRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(esrd.hooks) == 0 {
		affected, err = esrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ETCSyncRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			esrd.mutation = mutation
			affected, err = esrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(esrd.hooks) - 1; i >= 0; i-- {
			if esrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = esrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, esrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (esrd *ETCSyncRecordDelete) ExecX(ctx context.Context) int {
	n, err := esrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (esrd *ETCSyncRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: etcsyncrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: etcsyncrecord.FieldID,
			},
		},
	}
	if ps := esrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, esrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ETCSyncRecordDeleteOne is the builder for deleting a single ETCSyncRecord entity.
type ETCSyncRecordDeleteOne struct {
	esrd *ETCSyncRecordDelete
}

// Exec executes the deletion query.
func (esrdo *ETCSyncRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := esrdo.esrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{etcsyncrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (esrdo *ETCSyncRecordDeleteOne) ExecX(ctx context.Context) {
	esrdo.esrd.ExecX(ctx)
}
