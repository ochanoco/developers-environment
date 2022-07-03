// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/ochanoco/ochano.co-auth/proxy/ent/migrate"

	"github.com/ochanoco/ochano.co-auth/proxy/ent/project"
	"github.com/ochanoco/ochano.co-auth/proxy/ent/whitelist"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Project is the client for interacting with the Project builders.
	Project *ProjectClient
	// WhiteList is the client for interacting with the WhiteList builders.
	WhiteList *WhiteListClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Project = NewProjectClient(c.config)
	c.WhiteList = NewWhiteListClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Project:   NewProjectClient(cfg),
		WhiteList: NewWhiteListClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Project:   NewProjectClient(cfg),
		WhiteList: NewWhiteListClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Project.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Project.Use(hooks...)
	c.WhiteList.Use(hooks...)
}

// ProjectClient is a client for the Project schema.
type ProjectClient struct {
	config
}

// NewProjectClient returns a client for the Project from the given config.
func NewProjectClient(c config) *ProjectClient {
	return &ProjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `project.Hooks(f(g(h())))`.
func (c *ProjectClient) Use(hooks ...Hook) {
	c.hooks.Project = append(c.hooks.Project, hooks...)
}

// Create returns a create builder for Project.
func (c *ProjectClient) Create() *ProjectCreate {
	mutation := newProjectMutation(c.config, OpCreate)
	return &ProjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Project entities.
func (c *ProjectClient) CreateBulk(builders ...*ProjectCreate) *ProjectCreateBulk {
	return &ProjectCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Project.
func (c *ProjectClient) Update() *ProjectUpdate {
	mutation := newProjectMutation(c.config, OpUpdate)
	return &ProjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProjectClient) UpdateOne(pr *Project) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProject(pr))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProjectClient) UpdateOneID(id int) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProjectID(id))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Project.
func (c *ProjectClient) Delete() *ProjectDelete {
	mutation := newProjectMutation(c.config, OpDelete)
	return &ProjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProjectClient) DeleteOne(pr *Project) *ProjectDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProjectClient) DeleteOneID(id int) *ProjectDeleteOne {
	builder := c.Delete().Where(project.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProjectDeleteOne{builder}
}

// Query returns a query builder for Project.
func (c *ProjectClient) Query() *ProjectQuery {
	return &ProjectQuery{
		config: c.config,
	}
}

// Get returns a Project entity by its id.
func (c *ProjectClient) Get(ctx context.Context, id int) (*Project, error) {
	return c.Query().Where(project.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProjectClient) GetX(ctx context.Context, id int) *Project {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWhitelists queries the whitelists edge of a Project.
func (c *ProjectClient) QueryWhitelists(pr *Project) *WhiteListQuery {
	query := &WhiteListQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(whitelist.Table, whitelist.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, project.WhitelistsTable, project.WhitelistsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProjectClient) Hooks() []Hook {
	return c.hooks.Project
}

// WhiteListClient is a client for the WhiteList schema.
type WhiteListClient struct {
	config
}

// NewWhiteListClient returns a client for the WhiteList from the given config.
func NewWhiteListClient(c config) *WhiteListClient {
	return &WhiteListClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `whitelist.Hooks(f(g(h())))`.
func (c *WhiteListClient) Use(hooks ...Hook) {
	c.hooks.WhiteList = append(c.hooks.WhiteList, hooks...)
}

// Create returns a create builder for WhiteList.
func (c *WhiteListClient) Create() *WhiteListCreate {
	mutation := newWhiteListMutation(c.config, OpCreate)
	return &WhiteListCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WhiteList entities.
func (c *WhiteListClient) CreateBulk(builders ...*WhiteListCreate) *WhiteListCreateBulk {
	return &WhiteListCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WhiteList.
func (c *WhiteListClient) Update() *WhiteListUpdate {
	mutation := newWhiteListMutation(c.config, OpUpdate)
	return &WhiteListUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WhiteListClient) UpdateOne(wl *WhiteList) *WhiteListUpdateOne {
	mutation := newWhiteListMutation(c.config, OpUpdateOne, withWhiteList(wl))
	return &WhiteListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WhiteListClient) UpdateOneID(id int) *WhiteListUpdateOne {
	mutation := newWhiteListMutation(c.config, OpUpdateOne, withWhiteListID(id))
	return &WhiteListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WhiteList.
func (c *WhiteListClient) Delete() *WhiteListDelete {
	mutation := newWhiteListMutation(c.config, OpDelete)
	return &WhiteListDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WhiteListClient) DeleteOne(wl *WhiteList) *WhiteListDeleteOne {
	return c.DeleteOneID(wl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WhiteListClient) DeleteOneID(id int) *WhiteListDeleteOne {
	builder := c.Delete().Where(whitelist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WhiteListDeleteOne{builder}
}

// Query returns a query builder for WhiteList.
func (c *WhiteListClient) Query() *WhiteListQuery {
	return &WhiteListQuery{
		config: c.config,
	}
}

// Get returns a WhiteList entity by its id.
func (c *WhiteListClient) Get(ctx context.Context, id int) (*WhiteList, error) {
	return c.Query().Where(whitelist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WhiteListClient) GetX(ctx context.Context, id int) *WhiteList {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a WhiteList.
func (c *WhiteListClient) QueryOwner(wl *WhiteList) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(whitelist.Table, whitelist.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, whitelist.OwnerTable, whitelist.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(wl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WhiteListClient) Hooks() []Hook {
	return c.hooks.WhiteList
}
