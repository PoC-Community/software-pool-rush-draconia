// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"draconia/ent/character"
	"draconia/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CharacterUpdate is the builder for updating Character entities.
type CharacterUpdate struct {
	config
	hooks    []Hook
	mutation *CharacterMutation
}

// Where appends a list predicates to the CharacterUpdate builder.
func (cu *CharacterUpdate) Where(ps ...predicate.Character) *CharacterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetPrice sets the "price" field.
func (cu *CharacterUpdate) SetPrice(i int64) *CharacterUpdate {
	cu.mutation.ResetPrice()
	cu.mutation.SetPrice(i)
	return cu
}

// AddPrice adds i to the "price" field.
func (cu *CharacterUpdate) AddPrice(i int64) *CharacterUpdate {
	cu.mutation.AddPrice(i)
	return cu
}

// Mutation returns the CharacterMutation object of the builder.
func (cu *CharacterUpdate) Mutation() *CharacterMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CharacterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CharacterMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CharacterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CharacterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CharacterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CharacterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   character.Table,
			Columns: character.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: character.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Price(); ok {
		_spec.SetField(character.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedPrice(); ok {
		_spec.AddField(character.FieldPrice, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CharacterUpdateOne is the builder for updating a single Character entity.
type CharacterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CharacterMutation
}

// SetPrice sets the "price" field.
func (cuo *CharacterUpdateOne) SetPrice(i int64) *CharacterUpdateOne {
	cuo.mutation.ResetPrice()
	cuo.mutation.SetPrice(i)
	return cuo
}

// AddPrice adds i to the "price" field.
func (cuo *CharacterUpdateOne) AddPrice(i int64) *CharacterUpdateOne {
	cuo.mutation.AddPrice(i)
	return cuo
}

// Mutation returns the CharacterMutation object of the builder.
func (cuo *CharacterUpdateOne) Mutation() *CharacterMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CharacterUpdateOne) Select(field string, fields ...string) *CharacterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Character entity.
func (cuo *CharacterUpdateOne) Save(ctx context.Context) (*Character, error) {
	return withHooks[*Character, CharacterMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CharacterUpdateOne) SaveX(ctx context.Context) *Character {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CharacterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CharacterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CharacterUpdateOne) sqlSave(ctx context.Context) (_node *Character, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   character.Table,
			Columns: character.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: character.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Character.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, character.FieldID)
		for _, f := range fields {
			if !character.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != character.FieldID {
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
	if value, ok := cuo.mutation.Price(); ok {
		_spec.SetField(character.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedPrice(); ok {
		_spec.AddField(character.FieldPrice, field.TypeInt64, value)
	}
	_node = &Character{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}