// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"example.com/m/v2/ent/predicate"
	"example.com/m/v2/ent/serviceprovider"
	"example.com/m/v2/ent/whitelist"

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
	TypeServiceProvider = "ServiceProvider"
	TypeWhiteList       = "WhiteList"
)

// ServiceProviderMutation represents an operation that mutates the ServiceProvider nodes in the graph.
type ServiceProviderMutation struct {
	config
	op                Op
	typ               string
	id                *int
	host              *string
	destination_ip    *string
	clearedFields     map[string]struct{}
	whitelists        map[int]struct{}
	removedwhitelists map[int]struct{}
	clearedwhitelists bool
	done              bool
	oldValue          func(context.Context) (*ServiceProvider, error)
	predicates        []predicate.ServiceProvider
}

var _ ent.Mutation = (*ServiceProviderMutation)(nil)

// serviceproviderOption allows management of the mutation configuration using functional options.
type serviceproviderOption func(*ServiceProviderMutation)

// newServiceProviderMutation creates new mutation for the ServiceProvider entity.
func newServiceProviderMutation(c config, op Op, opts ...serviceproviderOption) *ServiceProviderMutation {
	m := &ServiceProviderMutation{
		config:        c,
		op:            op,
		typ:           TypeServiceProvider,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withServiceProviderID sets the ID field of the mutation.
func withServiceProviderID(id int) serviceproviderOption {
	return func(m *ServiceProviderMutation) {
		var (
			err   error
			once  sync.Once
			value *ServiceProvider
		)
		m.oldValue = func(ctx context.Context) (*ServiceProvider, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ServiceProvider.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withServiceProvider sets the old ServiceProvider of the mutation.
func withServiceProvider(node *ServiceProvider) serviceproviderOption {
	return func(m *ServiceProviderMutation) {
		m.oldValue = func(context.Context) (*ServiceProvider, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ServiceProviderMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ServiceProviderMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ServiceProviderMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ServiceProviderMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ServiceProvider.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetHost sets the "host" field.
func (m *ServiceProviderMutation) SetHost(s string) {
	m.host = &s
}

// Host returns the value of the "host" field in the mutation.
func (m *ServiceProviderMutation) Host() (r string, exists bool) {
	v := m.host
	if v == nil {
		return
	}
	return *v, true
}

// OldHost returns the old "host" field's value of the ServiceProvider entity.
// If the ServiceProvider object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceProviderMutation) OldHost(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHost is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHost requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHost: %w", err)
	}
	return oldValue.Host, nil
}

// ResetHost resets all changes to the "host" field.
func (m *ServiceProviderMutation) ResetHost() {
	m.host = nil
}

// SetDestinationIP sets the "destination_ip" field.
func (m *ServiceProviderMutation) SetDestinationIP(s string) {
	m.destination_ip = &s
}

// DestinationIP returns the value of the "destination_ip" field in the mutation.
func (m *ServiceProviderMutation) DestinationIP() (r string, exists bool) {
	v := m.destination_ip
	if v == nil {
		return
	}
	return *v, true
}

// OldDestinationIP returns the old "destination_ip" field's value of the ServiceProvider entity.
// If the ServiceProvider object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceProviderMutation) OldDestinationIP(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDestinationIP is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDestinationIP requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDestinationIP: %w", err)
	}
	return oldValue.DestinationIP, nil
}

// ResetDestinationIP resets all changes to the "destination_ip" field.
func (m *ServiceProviderMutation) ResetDestinationIP() {
	m.destination_ip = nil
}

// AddWhitelistIDs adds the "whitelists" edge to the WhiteList entity by ids.
func (m *ServiceProviderMutation) AddWhitelistIDs(ids ...int) {
	if m.whitelists == nil {
		m.whitelists = make(map[int]struct{})
	}
	for i := range ids {
		m.whitelists[ids[i]] = struct{}{}
	}
}

// ClearWhitelists clears the "whitelists" edge to the WhiteList entity.
func (m *ServiceProviderMutation) ClearWhitelists() {
	m.clearedwhitelists = true
}

// WhitelistsCleared reports if the "whitelists" edge to the WhiteList entity was cleared.
func (m *ServiceProviderMutation) WhitelistsCleared() bool {
	return m.clearedwhitelists
}

// RemoveWhitelistIDs removes the "whitelists" edge to the WhiteList entity by IDs.
func (m *ServiceProviderMutation) RemoveWhitelistIDs(ids ...int) {
	if m.removedwhitelists == nil {
		m.removedwhitelists = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.whitelists, ids[i])
		m.removedwhitelists[ids[i]] = struct{}{}
	}
}

// RemovedWhitelists returns the removed IDs of the "whitelists" edge to the WhiteList entity.
func (m *ServiceProviderMutation) RemovedWhitelistsIDs() (ids []int) {
	for id := range m.removedwhitelists {
		ids = append(ids, id)
	}
	return
}

// WhitelistsIDs returns the "whitelists" edge IDs in the mutation.
func (m *ServiceProviderMutation) WhitelistsIDs() (ids []int) {
	for id := range m.whitelists {
		ids = append(ids, id)
	}
	return
}

// ResetWhitelists resets all changes to the "whitelists" edge.
func (m *ServiceProviderMutation) ResetWhitelists() {
	m.whitelists = nil
	m.clearedwhitelists = false
	m.removedwhitelists = nil
}

// Where appends a list predicates to the ServiceProviderMutation builder.
func (m *ServiceProviderMutation) Where(ps ...predicate.ServiceProvider) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ServiceProviderMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (ServiceProvider).
func (m *ServiceProviderMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ServiceProviderMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.host != nil {
		fields = append(fields, serviceprovider.FieldHost)
	}
	if m.destination_ip != nil {
		fields = append(fields, serviceprovider.FieldDestinationIP)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ServiceProviderMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case serviceprovider.FieldHost:
		return m.Host()
	case serviceprovider.FieldDestinationIP:
		return m.DestinationIP()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ServiceProviderMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case serviceprovider.FieldHost:
		return m.OldHost(ctx)
	case serviceprovider.FieldDestinationIP:
		return m.OldDestinationIP(ctx)
	}
	return nil, fmt.Errorf("unknown ServiceProvider field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServiceProviderMutation) SetField(name string, value ent.Value) error {
	switch name {
	case serviceprovider.FieldHost:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHost(v)
		return nil
	case serviceprovider.FieldDestinationIP:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDestinationIP(v)
		return nil
	}
	return fmt.Errorf("unknown ServiceProvider field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ServiceProviderMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ServiceProviderMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServiceProviderMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ServiceProvider numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ServiceProviderMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ServiceProviderMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ServiceProviderMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ServiceProvider nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ServiceProviderMutation) ResetField(name string) error {
	switch name {
	case serviceprovider.FieldHost:
		m.ResetHost()
		return nil
	case serviceprovider.FieldDestinationIP:
		m.ResetDestinationIP()
		return nil
	}
	return fmt.Errorf("unknown ServiceProvider field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ServiceProviderMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.whitelists != nil {
		edges = append(edges, serviceprovider.EdgeWhitelists)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ServiceProviderMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case serviceprovider.EdgeWhitelists:
		ids := make([]ent.Value, 0, len(m.whitelists))
		for id := range m.whitelists {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ServiceProviderMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedwhitelists != nil {
		edges = append(edges, serviceprovider.EdgeWhitelists)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ServiceProviderMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case serviceprovider.EdgeWhitelists:
		ids := make([]ent.Value, 0, len(m.removedwhitelists))
		for id := range m.removedwhitelists {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ServiceProviderMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedwhitelists {
		edges = append(edges, serviceprovider.EdgeWhitelists)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ServiceProviderMutation) EdgeCleared(name string) bool {
	switch name {
	case serviceprovider.EdgeWhitelists:
		return m.clearedwhitelists
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ServiceProviderMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown ServiceProvider unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ServiceProviderMutation) ResetEdge(name string) error {
	switch name {
	case serviceprovider.EdgeWhitelists:
		m.ResetWhitelists()
		return nil
	}
	return fmt.Errorf("unknown ServiceProvider edge %s", name)
}

// WhiteListMutation represents an operation that mutates the WhiteList nodes in the graph.
type WhiteListMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_path         *string
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

// SetPath sets the "path" field.
func (m *WhiteListMutation) SetPath(s string) {
	m._path = &s
}

// Path returns the value of the "path" field in the mutation.
func (m *WhiteListMutation) Path() (r string, exists bool) {
	v := m._path
	if v == nil {
		return
	}
	return *v, true
}

// OldPath returns the old "path" field's value of the WhiteList entity.
// If the WhiteList object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WhiteListMutation) OldPath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPath: %w", err)
	}
	return oldValue.Path, nil
}

// ResetPath resets all changes to the "path" field.
func (m *WhiteListMutation) ResetPath() {
	m._path = nil
}

// SetOwnerID sets the "owner" edge to the ServiceProvider entity by id.
func (m *WhiteListMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the ServiceProvider entity.
func (m *WhiteListMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the ServiceProvider entity was cleared.
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
	if m._path != nil {
		fields = append(fields, whitelist.FieldPath)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WhiteListMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case whitelist.FieldPath:
		return m.Path()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WhiteListMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case whitelist.FieldPath:
		return m.OldPath(ctx)
	}
	return nil, fmt.Errorf("unknown WhiteList field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WhiteListMutation) SetField(name string, value ent.Value) error {
	switch name {
	case whitelist.FieldPath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPath(v)
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
	case whitelist.FieldPath:
		m.ResetPath()
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
