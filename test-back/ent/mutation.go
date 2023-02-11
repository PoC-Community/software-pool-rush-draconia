// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"draconia/ent/character"
	"draconia/ent/predicate"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCharacter = "Character"
)

// CharacterMutation represents an operation that mutates the Character nodes in the graph.
type CharacterMutation struct {
	config
	op            Op
	typ           string
	id            *int
	price         *int64
	addprice      *int64
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Character, error)
	predicates    []predicate.Character
}

var _ ent.Mutation = (*CharacterMutation)(nil)

// characterOption allows management of the mutation configuration using functional options.
type characterOption func(*CharacterMutation)

// newCharacterMutation creates new mutation for the Character entity.
func newCharacterMutation(c config, op Op, opts ...characterOption) *CharacterMutation {
	m := &CharacterMutation{
		config:        c,
		op:            op,
		typ:           TypeCharacter,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCharacterID sets the ID field of the mutation.
func withCharacterID(id int) characterOption {
	return func(m *CharacterMutation) {
		var (
			err   error
			once  sync.Once
			value *Character
		)
		m.oldValue = func(ctx context.Context) (*Character, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Character.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCharacter sets the old Character of the mutation.
func withCharacter(node *Character) characterOption {
	return func(m *CharacterMutation) {
		m.oldValue = func(context.Context) (*Character, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CharacterMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CharacterMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CharacterMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CharacterMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Character.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetPrice sets the "price" field.
func (m *CharacterMutation) SetPrice(i int64) {
	m.price = &i
	m.addprice = nil
}

// Price returns the value of the "price" field in the mutation.
func (m *CharacterMutation) Price() (r int64, exists bool) {
	v := m.price
	if v == nil {
		return
	}
	return *v, true
}

// OldPrice returns the old "price" field's value of the Character entity.
// If the Character object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CharacterMutation) OldPrice(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPrice is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPrice requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrice: %w", err)
	}
	return oldValue.Price, nil
}

// AddPrice adds i to the "price" field.
func (m *CharacterMutation) AddPrice(i int64) {
	if m.addprice != nil {
		*m.addprice += i
	} else {
		m.addprice = &i
	}
}

// AddedPrice returns the value that was added to the "price" field in this mutation.
func (m *CharacterMutation) AddedPrice() (r int64, exists bool) {
	v := m.addprice
	if v == nil {
		return
	}
	return *v, true
}

// ResetPrice resets all changes to the "price" field.
func (m *CharacterMutation) ResetPrice() {
	m.price = nil
	m.addprice = nil
}

// Where appends a list predicates to the CharacterMutation builder.
func (m *CharacterMutation) Where(ps ...predicate.Character) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CharacterMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CharacterMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Character, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CharacterMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CharacterMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Character).
func (m *CharacterMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CharacterMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.price != nil {
		fields = append(fields, character.FieldPrice)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CharacterMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case character.FieldPrice:
		return m.Price()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CharacterMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case character.FieldPrice:
		return m.OldPrice(ctx)
	}
	return nil, fmt.Errorf("unknown Character field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CharacterMutation) SetField(name string, value ent.Value) error {
	switch name {
	case character.FieldPrice:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrice(v)
		return nil
	}
	return fmt.Errorf("unknown Character field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CharacterMutation) AddedFields() []string {
	var fields []string
	if m.addprice != nil {
		fields = append(fields, character.FieldPrice)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CharacterMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case character.FieldPrice:
		return m.AddedPrice()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CharacterMutation) AddField(name string, value ent.Value) error {
	switch name {
	case character.FieldPrice:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPrice(v)
		return nil
	}
	return fmt.Errorf("unknown Character numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CharacterMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CharacterMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CharacterMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Character nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CharacterMutation) ResetField(name string) error {
	switch name {
	case character.FieldPrice:
		m.ResetPrice()
		return nil
	}
	return fmt.Errorf("unknown Character field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CharacterMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CharacterMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CharacterMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CharacterMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CharacterMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CharacterMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CharacterMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Character unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CharacterMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Character edge %s", name)
}