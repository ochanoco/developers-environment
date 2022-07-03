// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ochanoco/ochano.co-auth/proxy/ent/predicate"
	"github.com/ochanoco/ochano.co-auth/proxy/ent/project"
	"github.com/ochanoco/ochano.co-auth/proxy/ent/whitelist"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProject   = "Project"
	TypeWhiteList = "WhiteList"
)

// ProjectMutation represents an operation that mutates the Project nodes in the graph.
type ProjectMutation struct {
	config
	op                Op
	typ               string
	id                *int
	name              *string
	domain            *string
	destination       *string
	line_id           *string
	clearedFields     map[string]struct{}
	whitelists        map[int]struct{}
	removedwhitelists map[int]struct{}
	clearedwhitelists bool
	done              bool
	oldValue          func(context.Context) (*Project, error)
	predicates        []predicate.Project
}

var _ ent.Mutation = (*ProjectMutation)(nil)

// projectOption allows management of the mutation configuration using functional options.
type projectOption func(*ProjectMutation)

// newProjectMutation creates new mutation for the Project entity.
func newProjectMutation(c config, op Op, opts ...projectOption) *ProjectMutation {
	m := &ProjectMutation{
		config:        c,
		op:            op,
		typ:           TypeProject,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProjectID sets the ID field of the mutation.
func withProjectID(id int) projectOption {
	return func(m *ProjectMutation) {
		var (
			err   error
			once  sync.Once
			value *Project
		)
		m.oldValue = func(ctx context.Context) (*Project, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Project.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProject sets the old Project of the mutation.
func withProject(node *Project) projectOption {
	return func(m *ProjectMutation) {
		m.oldValue = func(context.Context) (*Project, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProjectMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProjectMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ProjectMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ProjectMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Project.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *ProjectMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ProjectMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ProjectMutation) ResetName() {
	m.name = nil
}

// SetDomain sets the "domain" field.
func (m *ProjectMutation) SetDomain(s string) {
	m.domain = &s
}

// Domain returns the value of the "domain" field in the mutation.
func (m *ProjectMutation) Domain() (r string, exists bool) {
	v := m.domain
	if v == nil {
		return
	}
	return *v, true
}

// OldDomain returns the old "domain" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldDomain(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDomain is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDomain requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDomain: %w", err)
	}
	return oldValue.Domain, nil
}

// ResetDomain resets all changes to the "domain" field.
func (m *ProjectMutation) ResetDomain() {
	m.domain = nil
}

// SetDestination sets the "destination" field.
func (m *ProjectMutation) SetDestination(s string) {
	m.destination = &s
}

// Destination returns the value of the "destination" field in the mutation.
func (m *ProjectMutation) Destination() (r string, exists bool) {
	v := m.destination
	if v == nil {
		return
	}
	return *v, true
}

// OldDestination returns the old "destination" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldDestination(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDestination is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDestination requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDestination: %w", err)
	}
	return oldValue.Destination, nil
}

// ResetDestination resets all changes to the "destination" field.
func (m *ProjectMutation) ResetDestination() {
	m.destination = nil
}

// SetLineID sets the "line_id" field.
func (m *ProjectMutation) SetLineID(s string) {
	m.line_id = &s
}

// LineID returns the value of the "line_id" field in the mutation.
func (m *ProjectMutation) LineID() (r string, exists bool) {
	v := m.line_id
	if v == nil {
		return
	}
	return *v, true
}

// OldLineID returns the old "line_id" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldLineID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLineID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLineID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLineID: %w", err)
	}
	return oldValue.LineID, nil
}

// ResetLineID resets all changes to the "line_id" field.
func (m *ProjectMutation) ResetLineID() {
	m.line_id = nil
}

// AddWhitelistIDs adds the "whitelists" edge to the WhiteList entity by ids.
func (m *ProjectMutation) AddWhitelistIDs(ids ...int) {
	if m.whitelists == nil {
		m.whitelists = make(map[int]struct{})
	}
	for i := range ids {
		m.whitelists[ids[i]] = struct{}{}
	}
}

// ClearWhitelists clears the "whitelists" edge to the WhiteList entity.
func (m *ProjectMutation) ClearWhitelists() {
	m.clearedwhitelists = true
}

// WhitelistsCleared reports if the "whitelists" edge to the WhiteList entity was cleared.
func (m *ProjectMutation) WhitelistsCleared() bool {
	return m.clearedwhitelists
}

// RemoveWhitelistIDs removes the "whitelists" edge to the WhiteList entity by IDs.
func (m *ProjectMutation) RemoveWhitelistIDs(ids ...int) {
	if m.removedwhitelists == nil {
		m.removedwhitelists = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.whitelists, ids[i])
		m.removedwhitelists[ids[i]] = struct{}{}
	}
}

// RemovedWhitelists returns the removed IDs of the "whitelists" edge to the WhiteList entity.
func (m *ProjectMutation) RemovedWhitelistsIDs() (ids []int) {
	for id := range m.removedwhitelists {
		ids = append(ids, id)
	}
	return
}

// WhitelistsIDs returns the "whitelists" edge IDs in the mutation.
func (m *ProjectMutation) WhitelistsIDs() (ids []int) {
	for id := range m.whitelists {
		ids = append(ids, id)
	}
	return
}

// ResetWhitelists resets all changes to the "whitelists" edge.
func (m *ProjectMutation) ResetWhitelists() {
	m.whitelists = nil
	m.clearedwhitelists = false
	m.removedwhitelists = nil
}

// Where appends a list predicates to the ProjectMutation builder.
func (m *ProjectMutation) Where(ps ...predicate.Project) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ProjectMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Project).
func (m *ProjectMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProjectMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, project.FieldName)
	}
	if m.domain != nil {
		fields = append(fields, project.FieldDomain)
	}
	if m.destination != nil {
		fields = append(fields, project.FieldDestination)
	}
	if m.line_id != nil {
		fields = append(fields, project.FieldLineID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProjectMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case project.FieldName:
		return m.Name()
	case project.FieldDomain:
		return m.Domain()
	case project.FieldDestination:
		return m.Destination()
	case project.FieldLineID:
		return m.LineID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProjectMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case project.FieldName:
		return m.OldName(ctx)
	case project.FieldDomain:
		return m.OldDomain(ctx)
	case project.FieldDestination:
		return m.OldDestination(ctx)
	case project.FieldLineID:
		return m.OldLineID(ctx)
	}
	return nil, fmt.Errorf("unknown Project field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProjectMutation) SetField(name string, value ent.Value) error {
	switch name {
	case project.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case project.FieldDomain:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDomain(v)
		return nil
	case project.FieldDestination:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDestination(v)
		return nil
	case project.FieldLineID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLineID(v)
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProjectMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProjectMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProjectMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Project numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProjectMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProjectMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProjectMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Project nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProjectMutation) ResetField(name string) error {
	switch name {
	case project.FieldName:
		m.ResetName()
		return nil
	case project.FieldDomain:
		m.ResetDomain()
		return nil
	case project.FieldDestination:
		m.ResetDestination()
		return nil
	case project.FieldLineID:
		m.ResetLineID()
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProjectMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.whitelists != nil {
		edges = append(edges, project.EdgeWhitelists)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProjectMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case project.EdgeWhitelists:
		ids := make([]ent.Value, 0, len(m.whitelists))
		for id := range m.whitelists {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProjectMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedwhitelists != nil {
		edges = append(edges, project.EdgeWhitelists)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProjectMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case project.EdgeWhitelists:
		ids := make([]ent.Value, 0, len(m.removedwhitelists))
		for id := range m.removedwhitelists {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProjectMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedwhitelists {
		edges = append(edges, project.EdgeWhitelists)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProjectMutation) EdgeCleared(name string) bool {
	switch name {
	case project.EdgeWhitelists:
		return m.clearedwhitelists
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProjectMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Project unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProjectMutation) ResetEdge(name string) error {
	switch name {
	case project.EdgeWhitelists:
		m.ResetWhitelists()
		return nil
	}
	return fmt.Errorf("unknown Project edge %s", name)
}

// WhiteListMutation represents an operation that mutates the WhiteList nodes in the graph.
type WhiteListMutation struct {
	config
	op            Op
	typ           string
	id            *int
	url           *string
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*WhiteList, error)
	predicates    []predicate.WhiteList
}

var _ ent.Mutation = (*WhiteListMutation)(nil)

// whitelistOption allows management of the mutation configuration using functional options.
type whitelistOption func(*WhiteListMutation)

// newWhiteListMutation creates new mutation for the WhiteList entity.
func newWhiteListMutation(c config, op Op, opts ...whitelistOption) *WhiteListMutation {
	m := &WhiteListMutation{
		config:        c,
		op:            op,
		typ:           TypeWhiteList,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWhiteListID sets the ID field of the mutation.
func withWhiteListID(id int) whitelistOption {
	return func(m *WhiteListMutation) {
		var (
			err   error
			once  sync.Once
			value *WhiteList
		)
		m.oldValue = func(ctx context.Context) (*WhiteList, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().WhiteList.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWhiteList sets the old WhiteList of the mutation.
func withWhiteList(node *WhiteList) whitelistOption {
	return func(m *WhiteListMutation) {
		m.oldValue = func(context.Context) (*WhiteList, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WhiteListMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WhiteListMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WhiteListMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *WhiteListMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().WhiteList.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetURL sets the "url" field.
func (m *WhiteListMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *WhiteListMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the WhiteList entity.
// If the WhiteList object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WhiteListMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *WhiteListMutation) ResetURL() {
	m.url = nil
}

// SetOwnerID sets the "owner" edge to the Project entity by id.
func (m *WhiteListMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the Project entity.
func (m *WhiteListMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the Project entity was cleared.
func (m *WhiteListMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *WhiteListMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *WhiteListMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *WhiteListMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the WhiteListMutation builder.
func (m *WhiteListMutation) Where(ps ...predicate.WhiteList) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *WhiteListMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (WhiteList).
func (m *WhiteListMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WhiteListMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.url != nil {
		fields = append(fields, whitelist.FieldURL)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WhiteListMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case whitelist.FieldURL:
		return m.URL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WhiteListMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case whitelist.FieldURL:
		return m.OldURL(ctx)
	}
	return nil, fmt.Errorf("unknown WhiteList field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WhiteListMutation) SetField(name string, value ent.Value) error {
	switch name {
	case whitelist.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	}
	return fmt.Errorf("unknown WhiteList field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WhiteListMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WhiteListMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WhiteListMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown WhiteList numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WhiteListMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WhiteListMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WhiteListMutation) ClearField(name string) error {
	return fmt.Errorf("unknown WhiteList nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WhiteListMutation) ResetField(name string) error {
	switch name {
	case whitelist.FieldURL:
		m.ResetURL()
		return nil
	}
	return fmt.Errorf("unknown WhiteList field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WhiteListMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, whitelist.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WhiteListMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case whitelist.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WhiteListMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WhiteListMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WhiteListMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, whitelist.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WhiteListMutation) EdgeCleared(name string) bool {
	switch name {
	case whitelist.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WhiteListMutation) ClearEdge(name string) error {
	switch name {
	case whitelist.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown WhiteList unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WhiteListMutation) ResetEdge(name string) error {
	switch name {
	case whitelist.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown WhiteList edge %s", name)
}
