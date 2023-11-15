// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"rent-a-car/ent/car"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	mutation *CarMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCarName sets the "car_name" field.
func (cc *CarCreate) SetCarName(s string) *CarCreate {
	cc.mutation.SetCarName(s)
	return cc
}

// SetMilesPerGallon sets the "miles_per_gallon" field.
func (cc *CarCreate) SetMilesPerGallon(s string) *CarCreate {
	cc.mutation.SetMilesPerGallon(s)
	return cc
}

// SetCylinders sets the "cylinders" field.
func (cc *CarCreate) SetCylinders(s string) *CarCreate {
	cc.mutation.SetCylinders(s)
	return cc
}

// SetPower sets the "power" field.
func (cc *CarCreate) SetPower(s string) *CarCreate {
	cc.mutation.SetPower(s)
	return cc
}

// SetYear sets the "year" field.
func (cc *CarCreate) SetYear(s string) *CarCreate {
	cc.mutation.SetYear(s)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CarCreate) SetCreatedAt(t time.Time) *CarCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CarCreate) SetNillableCreatedAt(t *time.Time) *CarCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CarCreate) SetUpdatedAt(t time.Time) *CarCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CarCreate) SetNillableUpdatedAt(t *time.Time) *CarCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// Mutation returns the CarMutation object of the builder.
func (cc *CarCreate) Mutation() *CarMutation {
	return cc.mutation
}

// Save creates the Car in the database.
func (cc *CarCreate) Save(ctx context.Context) (*Car, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CarCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CarCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CarCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := car.DefaultCreatedAt
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := car.DefaultUpdatedAt
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CarCreate) check() error {
	if _, ok := cc.mutation.CarName(); !ok {
		return &ValidationError{Name: "car_name", err: errors.New(`ent: missing required field "Car.car_name"`)}
	}
	if _, ok := cc.mutation.MilesPerGallon(); !ok {
		return &ValidationError{Name: "miles_per_gallon", err: errors.New(`ent: missing required field "Car.miles_per_gallon"`)}
	}
	if _, ok := cc.mutation.Cylinders(); !ok {
		return &ValidationError{Name: "cylinders", err: errors.New(`ent: missing required field "Car.cylinders"`)}
	}
	if _, ok := cc.mutation.Power(); !ok {
		return &ValidationError{Name: "power", err: errors.New(`ent: missing required field "Car.power"`)}
	}
	if _, ok := cc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Car.year"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Car.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Car.updated_at"`)}
	}
	return nil
}

func (cc *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CarCreate) createSpec() (*Car, *sqlgraph.CreateSpec) {
	var (
		_node = &Car{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(car.Table, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.CarName(); ok {
		_spec.SetField(car.FieldCarName, field.TypeString, value)
		_node.CarName = value
	}
	if value, ok := cc.mutation.MilesPerGallon(); ok {
		_spec.SetField(car.FieldMilesPerGallon, field.TypeString, value)
		_node.MilesPerGallon = value
	}
	if value, ok := cc.mutation.Cylinders(); ok {
		_spec.SetField(car.FieldCylinders, field.TypeString, value)
		_node.Cylinders = value
	}
	if value, ok := cc.mutation.Power(); ok {
		_spec.SetField(car.FieldPower, field.TypeString, value)
		_node.Power = value
	}
	if value, ok := cc.mutation.Year(); ok {
		_spec.SetField(car.FieldYear, field.TypeString, value)
		_node.Year = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(car.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(car.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Car.Create().
//		SetCarName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CarUpsert) {
//			SetCarName(v+v).
//		}).
//		Exec(ctx)
func (cc *CarCreate) OnConflict(opts ...sql.ConflictOption) *CarUpsertOne {
	cc.conflict = opts
	return &CarUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CarCreate) OnConflictColumns(columns ...string) *CarUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CarUpsertOne{
		create: cc,
	}
}

type (
	// CarUpsertOne is the builder for "upsert"-ing
	//  one Car node.
	CarUpsertOne struct {
		create *CarCreate
	}

	// CarUpsert is the "OnConflict" setter.
	CarUpsert struct {
		*sql.UpdateSet
	}
)

// SetCarName sets the "car_name" field.
func (u *CarUpsert) SetCarName(v string) *CarUpsert {
	u.Set(car.FieldCarName, v)
	return u
}

// UpdateCarName sets the "car_name" field to the value that was provided on create.
func (u *CarUpsert) UpdateCarName() *CarUpsert {
	u.SetExcluded(car.FieldCarName)
	return u
}

// SetMilesPerGallon sets the "miles_per_gallon" field.
func (u *CarUpsert) SetMilesPerGallon(v string) *CarUpsert {
	u.Set(car.FieldMilesPerGallon, v)
	return u
}

// UpdateMilesPerGallon sets the "miles_per_gallon" field to the value that was provided on create.
func (u *CarUpsert) UpdateMilesPerGallon() *CarUpsert {
	u.SetExcluded(car.FieldMilesPerGallon)
	return u
}

// SetCylinders sets the "cylinders" field.
func (u *CarUpsert) SetCylinders(v string) *CarUpsert {
	u.Set(car.FieldCylinders, v)
	return u
}

// UpdateCylinders sets the "cylinders" field to the value that was provided on create.
func (u *CarUpsert) UpdateCylinders() *CarUpsert {
	u.SetExcluded(car.FieldCylinders)
	return u
}

// SetPower sets the "power" field.
func (u *CarUpsert) SetPower(v string) *CarUpsert {
	u.Set(car.FieldPower, v)
	return u
}

// UpdatePower sets the "power" field to the value that was provided on create.
func (u *CarUpsert) UpdatePower() *CarUpsert {
	u.SetExcluded(car.FieldPower)
	return u
}

// SetYear sets the "year" field.
func (u *CarUpsert) SetYear(v string) *CarUpsert {
	u.Set(car.FieldYear, v)
	return u
}

// UpdateYear sets the "year" field to the value that was provided on create.
func (u *CarUpsert) UpdateYear() *CarUpsert {
	u.SetExcluded(car.FieldYear)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CarUpsert) SetCreatedAt(v time.Time) *CarUpsert {
	u.Set(car.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CarUpsert) UpdateCreatedAt() *CarUpsert {
	u.SetExcluded(car.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CarUpsert) SetUpdatedAt(v time.Time) *CarUpsert {
	u.Set(car.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CarUpsert) UpdateUpdatedAt() *CarUpsert {
	u.SetExcluded(car.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CarUpsertOne) UpdateNewValues() *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Car.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CarUpsertOne) Ignore() *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CarUpsertOne) DoNothing() *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CarCreate.OnConflict
// documentation for more info.
func (u *CarUpsertOne) Update(set func(*CarUpsert)) *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CarUpsert{UpdateSet: update})
	}))
	return u
}

// SetCarName sets the "car_name" field.
func (u *CarUpsertOne) SetCarName(v string) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetCarName(v)
	})
}

// UpdateCarName sets the "car_name" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateCarName() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateCarName()
	})
}

// SetMilesPerGallon sets the "miles_per_gallon" field.
func (u *CarUpsertOne) SetMilesPerGallon(v string) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetMilesPerGallon(v)
	})
}

// UpdateMilesPerGallon sets the "miles_per_gallon" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateMilesPerGallon() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateMilesPerGallon()
	})
}

// SetCylinders sets the "cylinders" field.
func (u *CarUpsertOne) SetCylinders(v string) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetCylinders(v)
	})
}

// UpdateCylinders sets the "cylinders" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateCylinders() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateCylinders()
	})
}

// SetPower sets the "power" field.
func (u *CarUpsertOne) SetPower(v string) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetPower(v)
	})
}

// UpdatePower sets the "power" field to the value that was provided on create.
func (u *CarUpsertOne) UpdatePower() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdatePower()
	})
}

// SetYear sets the "year" field.
func (u *CarUpsertOne) SetYear(v string) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetYear(v)
	})
}

// UpdateYear sets the "year" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateYear() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateYear()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CarUpsertOne) SetCreatedAt(v time.Time) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateCreatedAt() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CarUpsertOne) SetUpdatedAt(v time.Time) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateUpdatedAt() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *CarUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CarCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CarUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CarUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CarUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CarCreateBulk is the builder for creating many Car entities in bulk.
type CarCreateBulk struct {
	config
	err      error
	builders []*CarCreate
	conflict []sql.ConflictOption
}

// Save creates the Car entities in the database.
func (ccb *CarCreateBulk) Save(ctx context.Context) ([]*Car, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Car, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CarCreateBulk) SaveX(ctx context.Context) []*Car {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CarCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CarCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Car.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CarUpsert) {
//			SetCarName(v+v).
//		}).
//		Exec(ctx)
func (ccb *CarCreateBulk) OnConflict(opts ...sql.ConflictOption) *CarUpsertBulk {
	ccb.conflict = opts
	return &CarUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CarCreateBulk) OnConflictColumns(columns ...string) *CarUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CarUpsertBulk{
		create: ccb,
	}
}

// CarUpsertBulk is the builder for "upsert"-ing
// a bulk of Car nodes.
type CarUpsertBulk struct {
	create *CarCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CarUpsertBulk) UpdateNewValues() *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CarUpsertBulk) Ignore() *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CarUpsertBulk) DoNothing() *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CarCreateBulk.OnConflict
// documentation for more info.
func (u *CarUpsertBulk) Update(set func(*CarUpsert)) *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CarUpsert{UpdateSet: update})
	}))
	return u
}

// SetCarName sets the "car_name" field.
func (u *CarUpsertBulk) SetCarName(v string) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetCarName(v)
	})
}

// UpdateCarName sets the "car_name" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateCarName() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateCarName()
	})
}

// SetMilesPerGallon sets the "miles_per_gallon" field.
func (u *CarUpsertBulk) SetMilesPerGallon(v string) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetMilesPerGallon(v)
	})
}

// UpdateMilesPerGallon sets the "miles_per_gallon" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateMilesPerGallon() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateMilesPerGallon()
	})
}

// SetCylinders sets the "cylinders" field.
func (u *CarUpsertBulk) SetCylinders(v string) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetCylinders(v)
	})
}

// UpdateCylinders sets the "cylinders" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateCylinders() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateCylinders()
	})
}

// SetPower sets the "power" field.
func (u *CarUpsertBulk) SetPower(v string) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetPower(v)
	})
}

// UpdatePower sets the "power" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdatePower() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdatePower()
	})
}

// SetYear sets the "year" field.
func (u *CarUpsertBulk) SetYear(v string) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetYear(v)
	})
}

// UpdateYear sets the "year" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateYear() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateYear()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CarUpsertBulk) SetCreatedAt(v time.Time) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateCreatedAt() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CarUpsertBulk) SetUpdatedAt(v time.Time) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateUpdatedAt() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *CarUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CarCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CarCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CarUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
