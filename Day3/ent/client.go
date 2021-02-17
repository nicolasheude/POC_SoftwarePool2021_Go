// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"SofwareGoDay3/ent/migrate"

	"SofwareGoDay3/ent/contact"
	"SofwareGoDay3/ent/developper"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Contact is the client for interacting with the Contact builders.
	Contact *ContactClient
	// Developper is the client for interacting with the Developper builders.
	Developper *DevelopperClient
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
	c.Contact = NewContactClient(c.config)
	c.Developper = NewDevelopperClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Contact:    NewContactClient(cfg),
		Developper: NewDevelopperClient(cfg),
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		Contact:    NewContactClient(cfg),
		Developper: NewDevelopperClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Contact.
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
	c.Contact.Use(hooks...)
	c.Developper.Use(hooks...)
}

// ContactClient is a client for the Contact schema.
type ContactClient struct {
	config
}

// NewContactClient returns a client for the Contact from the given config.
func NewContactClient(c config) *ContactClient {
	return &ContactClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contact.Hooks(f(g(h())))`.
func (c *ContactClient) Use(hooks ...Hook) {
	c.hooks.Contact = append(c.hooks.Contact, hooks...)
}

// Create returns a create builder for Contact.
func (c *ContactClient) Create() *ContactCreate {
	mutation := newContactMutation(c.config, OpCreate)
	return &ContactCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contact entities.
func (c *ContactClient) CreateBulk(builders ...*ContactCreate) *ContactCreateBulk {
	return &ContactCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contact.
func (c *ContactClient) Update() *ContactUpdate {
	mutation := newContactMutation(c.config, OpUpdate)
	return &ContactUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContactClient) UpdateOne(co *Contact) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContact(co))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContactClient) UpdateOneID(id int) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContactID(id))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contact.
func (c *ContactClient) Delete() *ContactDelete {
	mutation := newContactMutation(c.config, OpDelete)
	return &ContactDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ContactClient) DeleteOne(co *Contact) *ContactDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ContactClient) DeleteOneID(id int) *ContactDeleteOne {
	builder := c.Delete().Where(contact.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContactDeleteOne{builder}
}

// Query returns a query builder for Contact.
func (c *ContactClient) Query() *ContactQuery {
	return &ContactQuery{config: c.config}
}

// Get returns a Contact entity by its id.
func (c *ContactClient) Get(ctx context.Context, id int) (*Contact, error) {
	return c.Query().Where(contact.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContactClient) GetX(ctx context.Context, id int) *Contact {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Contact.
func (c *ContactClient) QueryOwner(co *Contact) *DevelopperQuery {
	query := &DevelopperQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contact.Table, contact.FieldID, id),
			sqlgraph.To(developper.Table, developper.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, contact.OwnerTable, contact.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ContactClient) Hooks() []Hook {
	return c.hooks.Contact
}

// DevelopperClient is a client for the Developper schema.
type DevelopperClient struct {
	config
}

// NewDevelopperClient returns a client for the Developper from the given config.
func NewDevelopperClient(c config) *DevelopperClient {
	return &DevelopperClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `developper.Hooks(f(g(h())))`.
func (c *DevelopperClient) Use(hooks ...Hook) {
	c.hooks.Developper = append(c.hooks.Developper, hooks...)
}

// Create returns a create builder for Developper.
func (c *DevelopperClient) Create() *DevelopperCreate {
	mutation := newDevelopperMutation(c.config, OpCreate)
	return &DevelopperCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Developper entities.
func (c *DevelopperClient) CreateBulk(builders ...*DevelopperCreate) *DevelopperCreateBulk {
	return &DevelopperCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Developper.
func (c *DevelopperClient) Update() *DevelopperUpdate {
	mutation := newDevelopperMutation(c.config, OpUpdate)
	return &DevelopperUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DevelopperClient) UpdateOne(d *Developper) *DevelopperUpdateOne {
	mutation := newDevelopperMutation(c.config, OpUpdateOne, withDevelopper(d))
	return &DevelopperUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DevelopperClient) UpdateOneID(id int) *DevelopperUpdateOne {
	mutation := newDevelopperMutation(c.config, OpUpdateOne, withDevelopperID(id))
	return &DevelopperUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Developper.
func (c *DevelopperClient) Delete() *DevelopperDelete {
	mutation := newDevelopperMutation(c.config, OpDelete)
	return &DevelopperDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DevelopperClient) DeleteOne(d *Developper) *DevelopperDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DevelopperClient) DeleteOneID(id int) *DevelopperDeleteOne {
	builder := c.Delete().Where(developper.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DevelopperDeleteOne{builder}
}

// Query returns a query builder for Developper.
func (c *DevelopperClient) Query() *DevelopperQuery {
	return &DevelopperQuery{config: c.config}
}

// Get returns a Developper entity by its id.
func (c *DevelopperClient) Get(ctx context.Context, id int) (*Developper, error) {
	return c.Query().Where(developper.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DevelopperClient) GetX(ctx context.Context, id int) *Developper {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryContact queries the contact edge of a Developper.
func (c *DevelopperClient) QueryContact(d *Developper) *ContactQuery {
	query := &ContactQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(developper.Table, developper.FieldID, id),
			sqlgraph.To(contact.Table, contact.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, developper.ContactTable, developper.ContactColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DevelopperClient) Hooks() []Hook {
	return c.hooks.Developper
}