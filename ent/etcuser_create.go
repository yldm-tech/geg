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
	"myetc.lantron.ltd/m/ent/etcuser"
)

// ETCUserCreate is the builder for creating a ETCUser entity.
type ETCUserCreate struct {
	config
	mutation *ETCUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEtcUsername sets the "etc_username" field.
func (euc *ETCUserCreate) SetEtcUsername(i int64) *ETCUserCreate {
	euc.mutation.SetEtcUsername(i)
	return euc
}

// SetEtcPassword sets the "etc_password" field.
func (euc *ETCUserCreate) SetEtcPassword(s string) *ETCUserCreate {
	euc.mutation.SetEtcPassword(s)
	return euc
}

// SetPointUsername sets the "point_username" field.
func (euc *ETCUserCreate) SetPointUsername(s string) *ETCUserCreate {
	euc.mutation.SetPointUsername(s)
	return euc
}

// SetPointPassword sets the "point_password" field.
func (euc *ETCUserCreate) SetPointPassword(s string) *ETCUserCreate {
	euc.mutation.SetPointPassword(s)
	return euc
}

// SetCreatedAt sets the "created_at" field.
func (euc *ETCUserCreate) SetCreatedAt(t time.Time) *ETCUserCreate {
	euc.mutation.SetCreatedAt(t)
	return euc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euc *ETCUserCreate) SetNillableCreatedAt(t *time.Time) *ETCUserCreate {
	if t != nil {
		euc.SetCreatedAt(*t)
	}
	return euc
}

// SetUpdatedAt sets the "updated_at" field.
func (euc *ETCUserCreate) SetUpdatedAt(t time.Time) *ETCUserCreate {
	euc.mutation.SetUpdatedAt(t)
	return euc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (euc *ETCUserCreate) SetNillableUpdatedAt(t *time.Time) *ETCUserCreate {
	if t != nil {
		euc.SetUpdatedAt(*t)
	}
	return euc
}

// SetID sets the "id" field.
func (euc *ETCUserCreate) SetID(i int64) *ETCUserCreate {
	euc.mutation.SetID(i)
	return euc
}

// Mutation returns the ETCUserMutation object of the builder.
func (euc *ETCUserCreate) Mutation() *ETCUserMutation {
	return euc.mutation
}

// Save creates the ETCUser in the database.
func (euc *ETCUserCreate) Save(ctx context.Context) (*ETCUser, error) {
	var (
		err  error
		node *ETCUser
	)
	euc.defaults()
	if len(euc.hooks) == 0 {
		if err = euc.check(); err != nil {
			return nil, err
		}
		node, err = euc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ETCUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euc.check(); err != nil {
				return nil, err
			}
			euc.mutation = mutation
			if node, err = euc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(euc.hooks) - 1; i >= 0; i-- {
			if euc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ETCUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ETCUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (euc *ETCUserCreate) SaveX(ctx context.Context) *ETCUser {
	v, err := euc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (euc *ETCUserCreate) Exec(ctx context.Context) error {
	_, err := euc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euc *ETCUserCreate) ExecX(ctx context.Context) {
	if err := euc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euc *ETCUserCreate) defaults() {
	if _, ok := euc.mutation.CreatedAt(); !ok {
		v := etcuser.DefaultCreatedAt
		euc.mutation.SetCreatedAt(v)
	}
	if _, ok := euc.mutation.UpdatedAt(); !ok {
		v := etcuser.DefaultUpdatedAt
		euc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euc *ETCUserCreate) check() error {
	if _, ok := euc.mutation.EtcUsername(); !ok {
		return &ValidationError{Name: "etc_username", err: errors.New(`ent: missing required field "ETCUser.etc_username"`)}
	}
	if _, ok := euc.mutation.EtcPassword(); !ok {
		return &ValidationError{Name: "etc_password", err: errors.New(`ent: missing required field "ETCUser.etc_password"`)}
	}
	if _, ok := euc.mutation.PointUsername(); !ok {
		return &ValidationError{Name: "point_username", err: errors.New(`ent: missing required field "ETCUser.point_username"`)}
	}
	if _, ok := euc.mutation.PointPassword(); !ok {
		return &ValidationError{Name: "point_password", err: errors.New(`ent: missing required field "ETCUser.point_password"`)}
	}
	if _, ok := euc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ETCUser.created_at"`)}
	}
	if _, ok := euc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ETCUser.updated_at"`)}
	}
	return nil
}

func (euc *ETCUserCreate) sqlSave(ctx context.Context) (*ETCUser, error) {
	_node, _spec := euc.createSpec()
	if err := sqlgraph.CreateNode(ctx, euc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (euc *ETCUserCreate) createSpec() (*ETCUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ETCUser{config: euc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: etcuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: etcuser.FieldID,
			},
		}
	)
	_spec.OnConflict = euc.conflict
	if id, ok := euc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := euc.mutation.EtcUsername(); ok {
		_spec.SetField(etcuser.FieldEtcUsername, field.TypeInt64, value)
		_node.EtcUsername = value
	}
	if value, ok := euc.mutation.EtcPassword(); ok {
		_spec.SetField(etcuser.FieldEtcPassword, field.TypeString, value)
		_node.EtcPassword = value
	}
	if value, ok := euc.mutation.PointUsername(); ok {
		_spec.SetField(etcuser.FieldPointUsername, field.TypeString, value)
		_node.PointUsername = value
	}
	if value, ok := euc.mutation.PointPassword(); ok {
		_spec.SetField(etcuser.FieldPointPassword, field.TypeString, value)
		_node.PointPassword = value
	}
	if value, ok := euc.mutation.CreatedAt(); ok {
		_spec.SetField(etcuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := euc.mutation.UpdatedAt(); ok {
		_spec.SetField(etcuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ETCUser.Create().
//		SetEtcUsername(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ETCUserUpsert) {
//			SetEtcUsername(v+v).
//		}).
//		Exec(ctx)
func (euc *ETCUserCreate) OnConflict(opts ...sql.ConflictOption) *ETCUserUpsertOne {
	euc.conflict = opts
	return &ETCUserUpsertOne{
		create: euc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ETCUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (euc *ETCUserCreate) OnConflictColumns(columns ...string) *ETCUserUpsertOne {
	euc.conflict = append(euc.conflict, sql.ConflictColumns(columns...))
	return &ETCUserUpsertOne{
		create: euc,
	}
}

type (
	// ETCUserUpsertOne is the builder for "upsert"-ing
	//  one ETCUser node.
	ETCUserUpsertOne struct {
		create *ETCUserCreate
	}

	// ETCUserUpsert is the "OnConflict" setter.
	ETCUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetEtcUsername sets the "etc_username" field.
func (u *ETCUserUpsert) SetEtcUsername(v int64) *ETCUserUpsert {
	u.Set(etcuser.FieldEtcUsername, v)
	return u
}

// UpdateEtcUsername sets the "etc_username" field to the value that was provided on create.
func (u *ETCUserUpsert) UpdateEtcUsername() *ETCUserUpsert {
	u.SetExcluded(etcuser.FieldEtcUsername)
	return u
}

// AddEtcUsername adds v to the "etc_username" field.
func (u *ETCUserUpsert) AddEtcUsername(v int64) *ETCUserUpsert {
	u.Add(etcuser.FieldEtcUsername, v)
	return u
}

// SetEtcPassword sets the "etc_password" field.
func (u *ETCUserUpsert) SetEtcPassword(v string) *ETCUserUpsert {
	u.Set(etcuser.FieldEtcPassword, v)
	return u
}

// UpdateEtcPassword sets the "etc_password" field to the value that was provided on create.
func (u *ETCUserUpsert) UpdateEtcPassword() *ETCUserUpsert {
	u.SetExcluded(etcuser.FieldEtcPassword)
	return u
}

// SetPointUsername sets the "point_username" field.
func (u *ETCUserUpsert) SetPointUsername(v string) *ETCUserUpsert {
	u.Set(etcuser.FieldPointUsername, v)
	return u
}

// UpdatePointUsername sets the "point_username" field to the value that was provided on create.
func (u *ETCUserUpsert) UpdatePointUsername() *ETCUserUpsert {
	u.SetExcluded(etcuser.FieldPointUsername)
	return u
}

// SetPointPassword sets the "point_password" field.
func (u *ETCUserUpsert) SetPointPassword(v string) *ETCUserUpsert {
	u.Set(etcuser.FieldPointPassword, v)
	return u
}

// UpdatePointPassword sets the "point_password" field to the value that was provided on create.
func (u *ETCUserUpsert) UpdatePointPassword() *ETCUserUpsert {
	u.SetExcluded(etcuser.FieldPointPassword)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ETCUserUpsert) SetCreatedAt(v time.Time) *ETCUserUpsert {
	u.Set(etcuser.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ETCUserUpsert) UpdateCreatedAt() *ETCUserUpsert {
	u.SetExcluded(etcuser.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ETCUserUpsert) SetUpdatedAt(v time.Time) *ETCUserUpsert {
	u.Set(etcuser.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ETCUserUpsert) UpdateUpdatedAt() *ETCUserUpsert {
	u.SetExcluded(etcuser.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ETCUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(etcuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ETCUserUpsertOne) UpdateNewValues() *ETCUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(etcuser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ETCUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ETCUserUpsertOne) Ignore() *ETCUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ETCUserUpsertOne) DoNothing() *ETCUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ETCUserCreate.OnConflict
// documentation for more info.
func (u *ETCUserUpsertOne) Update(set func(*ETCUserUpsert)) *ETCUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ETCUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEtcUsername sets the "etc_username" field.
func (u *ETCUserUpsertOne) SetEtcUsername(v int64) *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetEtcUsername(v)
	})
}

// AddEtcUsername adds v to the "etc_username" field.
func (u *ETCUserUpsertOne) AddEtcUsername(v int64) *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.AddEtcUsername(v)
	})
}

// UpdateEtcUsername sets the "etc_username" field to the value that was provided on create.
func (u *ETCUserUpsertOne) UpdateEtcUsername() *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateEtcUsername()
	})
}

// SetEtcPassword sets the "etc_password" field.
func (u *ETCUserUpsertOne) SetEtcPassword(v string) *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetEtcPassword(v)
	})
}

// UpdateEtcPassword sets the "etc_password" field to the value that was provided on create.
func (u *ETCUserUpsertOne) UpdateEtcPassword() *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateEtcPassword()
	})
}

// SetPointUsername sets the "point_username" field.
func (u *ETCUserUpsertOne) SetPointUsername(v string) *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetPointUsername(v)
	})
}

// UpdatePointUsername sets the "point_username" field to the value that was provided on create.
func (u *ETCUserUpsertOne) UpdatePointUsername() *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdatePointUsername()
	})
}

// SetPointPassword sets the "point_password" field.
func (u *ETCUserUpsertOne) SetPointPassword(v string) *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetPointPassword(v)
	})
}

// UpdatePointPassword sets the "point_password" field to the value that was provided on create.
func (u *ETCUserUpsertOne) UpdatePointPassword() *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdatePointPassword()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ETCUserUpsertOne) SetCreatedAt(v time.Time) *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ETCUserUpsertOne) UpdateCreatedAt() *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ETCUserUpsertOne) SetUpdatedAt(v time.Time) *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ETCUserUpsertOne) UpdateUpdatedAt() *ETCUserUpsertOne {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ETCUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ETCUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ETCUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ETCUserUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ETCUserUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ETCUserCreateBulk is the builder for creating many ETCUser entities in bulk.
type ETCUserCreateBulk struct {
	config
	builders []*ETCUserCreate
	conflict []sql.ConflictOption
}

// Save creates the ETCUser entities in the database.
func (eucb *ETCUserCreateBulk) Save(ctx context.Context) ([]*ETCUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eucb.builders))
	nodes := make([]*ETCUser, len(eucb.builders))
	mutators := make([]Mutator, len(eucb.builders))
	for i := range eucb.builders {
		func(i int, root context.Context) {
			builder := eucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ETCUserMutation)
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
					_, err = mutators[i+1].Mutate(root, eucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = eucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, eucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eucb *ETCUserCreateBulk) SaveX(ctx context.Context) []*ETCUser {
	v, err := eucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eucb *ETCUserCreateBulk) Exec(ctx context.Context) error {
	_, err := eucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eucb *ETCUserCreateBulk) ExecX(ctx context.Context) {
	if err := eucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ETCUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ETCUserUpsert) {
//			SetEtcUsername(v+v).
//		}).
//		Exec(ctx)
func (eucb *ETCUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *ETCUserUpsertBulk {
	eucb.conflict = opts
	return &ETCUserUpsertBulk{
		create: eucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ETCUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eucb *ETCUserCreateBulk) OnConflictColumns(columns ...string) *ETCUserUpsertBulk {
	eucb.conflict = append(eucb.conflict, sql.ConflictColumns(columns...))
	return &ETCUserUpsertBulk{
		create: eucb,
	}
}

// ETCUserUpsertBulk is the builder for "upsert"-ing
// a bulk of ETCUser nodes.
type ETCUserUpsertBulk struct {
	create *ETCUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ETCUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(etcuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ETCUserUpsertBulk) UpdateNewValues() *ETCUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(etcuser.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ETCUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ETCUserUpsertBulk) Ignore() *ETCUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ETCUserUpsertBulk) DoNothing() *ETCUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ETCUserCreateBulk.OnConflict
// documentation for more info.
func (u *ETCUserUpsertBulk) Update(set func(*ETCUserUpsert)) *ETCUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ETCUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEtcUsername sets the "etc_username" field.
func (u *ETCUserUpsertBulk) SetEtcUsername(v int64) *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetEtcUsername(v)
	})
}

// AddEtcUsername adds v to the "etc_username" field.
func (u *ETCUserUpsertBulk) AddEtcUsername(v int64) *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.AddEtcUsername(v)
	})
}

// UpdateEtcUsername sets the "etc_username" field to the value that was provided on create.
func (u *ETCUserUpsertBulk) UpdateEtcUsername() *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateEtcUsername()
	})
}

// SetEtcPassword sets the "etc_password" field.
func (u *ETCUserUpsertBulk) SetEtcPassword(v string) *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetEtcPassword(v)
	})
}

// UpdateEtcPassword sets the "etc_password" field to the value that was provided on create.
func (u *ETCUserUpsertBulk) UpdateEtcPassword() *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateEtcPassword()
	})
}

// SetPointUsername sets the "point_username" field.
func (u *ETCUserUpsertBulk) SetPointUsername(v string) *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetPointUsername(v)
	})
}

// UpdatePointUsername sets the "point_username" field to the value that was provided on create.
func (u *ETCUserUpsertBulk) UpdatePointUsername() *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdatePointUsername()
	})
}

// SetPointPassword sets the "point_password" field.
func (u *ETCUserUpsertBulk) SetPointPassword(v string) *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetPointPassword(v)
	})
}

// UpdatePointPassword sets the "point_password" field to the value that was provided on create.
func (u *ETCUserUpsertBulk) UpdatePointPassword() *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdatePointPassword()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ETCUserUpsertBulk) SetCreatedAt(v time.Time) *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ETCUserUpsertBulk) UpdateCreatedAt() *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ETCUserUpsertBulk) SetUpdatedAt(v time.Time) *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ETCUserUpsertBulk) UpdateUpdatedAt() *ETCUserUpsertBulk {
	return u.Update(func(s *ETCUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ETCUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ETCUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ETCUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ETCUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}