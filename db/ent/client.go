// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/peacewalker122/project/db/ent/migrate"

	"github.com/peacewalker122/project/db/ent/notif"
	"github.com/peacewalker122/project/db/ent/notifread"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Notif is the client for interacting with the Notif builders.
	Notif *NotifClient
	// NotifRead is the client for interacting with the NotifRead builders.
	NotifRead *NotifReadClient
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
	c.Notif = NewNotifClient(c.config)
	c.NotifRead = NewNotifReadClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		Notif:     NewNotifClient(cfg),
		NotifRead: NewNotifReadClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		Notif:     NewNotifClient(cfg),
		NotifRead: NewNotifReadClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Notif.
//		Query().
//		Count(ctx)
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
	c.Notif.Use(hooks...)
	c.NotifRead.Use(hooks...)
}

// NotifClient is a client for the Notif schema.
type NotifClient struct {
	config
}

// NewNotifClient returns a client for the Notif from the given config.
func NewNotifClient(c config) *NotifClient {
	return &NotifClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notif.Hooks(f(g(h())))`.
func (c *NotifClient) Use(hooks ...Hook) {
	c.hooks.Notif = append(c.hooks.Notif, hooks...)
}

// Create returns a builder for creating a Notif entity.
func (c *NotifClient) Create() *NotifCreate {
	mutation := newNotifMutation(c.config, OpCreate)
	return &NotifCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notif entities.
func (c *NotifClient) CreateBulk(builders ...*NotifCreate) *NotifCreateBulk {
	return &NotifCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notif.
func (c *NotifClient) Update() *NotifUpdate {
	mutation := newNotifMutation(c.config, OpUpdate)
	return &NotifUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotifClient) UpdateOne(n *Notif) *NotifUpdateOne {
	mutation := newNotifMutation(c.config, OpUpdateOne, withNotif(n))
	return &NotifUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotifClient) UpdateOneID(id uuid.UUID) *NotifUpdateOne {
	mutation := newNotifMutation(c.config, OpUpdateOne, withNotifID(id))
	return &NotifUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notif.
func (c *NotifClient) Delete() *NotifDelete {
	mutation := newNotifMutation(c.config, OpDelete)
	return &NotifDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NotifClient) DeleteOne(n *Notif) *NotifDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NotifClient) DeleteOneID(id uuid.UUID) *NotifDeleteOne {
	builder := c.Delete().Where(notif.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotifDeleteOne{builder}
}

// Query returns a query builder for Notif.
func (c *NotifClient) Query() *NotifQuery {
	return &NotifQuery{
		config: c.config,
	}
}

// Get returns a Notif entity by its id.
func (c *NotifClient) Get(ctx context.Context, id uuid.UUID) (*Notif, error) {
	return c.Query().Where(notif.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotifClient) GetX(ctx context.Context, id uuid.UUID) *Notif {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NotifClient) Hooks() []Hook {
	return c.hooks.Notif
}

// NotifReadClient is a client for the NotifRead schema.
type NotifReadClient struct {
	config
}

// NewNotifReadClient returns a client for the NotifRead from the given config.
func NewNotifReadClient(c config) *NotifReadClient {
	return &NotifReadClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notifread.Hooks(f(g(h())))`.
func (c *NotifReadClient) Use(hooks ...Hook) {
	c.hooks.NotifRead = append(c.hooks.NotifRead, hooks...)
}

// Create returns a builder for creating a NotifRead entity.
func (c *NotifReadClient) Create() *NotifReadCreate {
	mutation := newNotifReadMutation(c.config, OpCreate)
	return &NotifReadCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NotifRead entities.
func (c *NotifReadClient) CreateBulk(builders ...*NotifReadCreate) *NotifReadCreateBulk {
	return &NotifReadCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NotifRead.
func (c *NotifReadClient) Update() *NotifReadUpdate {
	mutation := newNotifReadMutation(c.config, OpUpdate)
	return &NotifReadUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NotifReadClient) UpdateOne(nr *NotifRead) *NotifReadUpdateOne {
	mutation := newNotifReadMutation(c.config, OpUpdateOne, withNotifRead(nr))
	return &NotifReadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NotifReadClient) UpdateOneID(id int) *NotifReadUpdateOne {
	mutation := newNotifReadMutation(c.config, OpUpdateOne, withNotifReadID(id))
	return &NotifReadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NotifRead.
func (c *NotifReadClient) Delete() *NotifReadDelete {
	mutation := newNotifReadMutation(c.config, OpDelete)
	return &NotifReadDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NotifReadClient) DeleteOne(nr *NotifRead) *NotifReadDeleteOne {
	return c.DeleteOneID(nr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NotifReadClient) DeleteOneID(id int) *NotifReadDeleteOne {
	builder := c.Delete().Where(notifread.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NotifReadDeleteOne{builder}
}

// Query returns a query builder for NotifRead.
func (c *NotifReadClient) Query() *NotifReadQuery {
	return &NotifReadQuery{
		config: c.config,
	}
}

// Get returns a NotifRead entity by its id.
func (c *NotifReadClient) Get(ctx context.Context, id int) (*NotifRead, error) {
	return c.Query().Where(notifread.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NotifReadClient) GetX(ctx context.Context, id int) *NotifRead {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NotifReadClient) Hooks() []Hook {
	return c.hooks.NotifRead
}
