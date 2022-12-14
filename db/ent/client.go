// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/peacewalker122/project/db/ent/migrate"

	"github.com/peacewalker122/project/db/ent/account"
	"github.com/peacewalker122/project/db/ent/accountnotifs"
	"github.com/peacewalker122/project/db/ent/notifread"
	"github.com/peacewalker122/project/db/ent/tokens"
	"github.com/peacewalker122/project/db/ent/users"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// AccountNotifs is the client for interacting with the AccountNotifs builders.
	AccountNotifs *AccountNotifsClient
	// NotifRead is the client for interacting with the NotifRead builders.
	NotifRead *NotifReadClient
	// Tokens is the client for interacting with the Tokens builders.
	Tokens *TokensClient
	// Users is the client for interacting with the Users builders.
	Users *UsersClient
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
	c.Account = NewAccountClient(c.config)
	c.AccountNotifs = NewAccountNotifsClient(c.config)
	c.NotifRead = NewNotifReadClient(c.config)
	c.Tokens = NewTokensClient(c.config)
	c.Users = NewUsersClient(c.config)
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
		ctx:           ctx,
		config:        cfg,
		Account:       NewAccountClient(cfg),
		AccountNotifs: NewAccountNotifsClient(cfg),
		NotifRead:     NewNotifReadClient(cfg),
		Tokens:        NewTokensClient(cfg),
		Users:         NewUsersClient(cfg),
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
		ctx:           ctx,
		config:        cfg,
		Account:       NewAccountClient(cfg),
		AccountNotifs: NewAccountNotifsClient(cfg),
		NotifRead:     NewNotifReadClient(cfg),
		Tokens:        NewTokensClient(cfg),
		Users:         NewUsersClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
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
	c.Account.Use(hooks...)
	c.AccountNotifs.Use(hooks...)
	c.NotifRead.Use(hooks...)
	c.Tokens.Use(hooks...)
	c.Users.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id int) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id int) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id int) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id int) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// AccountNotifsClient is a client for the AccountNotifs schema.
type AccountNotifsClient struct {
	config
}

// NewAccountNotifsClient returns a client for the AccountNotifs from the given config.
func NewAccountNotifsClient(c config) *AccountNotifsClient {
	return &AccountNotifsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accountnotifs.Hooks(f(g(h())))`.
func (c *AccountNotifsClient) Use(hooks ...Hook) {
	c.hooks.AccountNotifs = append(c.hooks.AccountNotifs, hooks...)
}

// Create returns a builder for creating a AccountNotifs entity.
func (c *AccountNotifsClient) Create() *AccountNotifsCreate {
	mutation := newAccountNotifsMutation(c.config, OpCreate)
	return &AccountNotifsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AccountNotifs entities.
func (c *AccountNotifsClient) CreateBulk(builders ...*AccountNotifsCreate) *AccountNotifsCreateBulk {
	return &AccountNotifsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AccountNotifs.
func (c *AccountNotifsClient) Update() *AccountNotifsUpdate {
	mutation := newAccountNotifsMutation(c.config, OpUpdate)
	return &AccountNotifsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountNotifsClient) UpdateOne(an *AccountNotifs) *AccountNotifsUpdateOne {
	mutation := newAccountNotifsMutation(c.config, OpUpdateOne, withAccountNotifs(an))
	return &AccountNotifsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountNotifsClient) UpdateOneID(id uuid.UUID) *AccountNotifsUpdateOne {
	mutation := newAccountNotifsMutation(c.config, OpUpdateOne, withAccountNotifsID(id))
	return &AccountNotifsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AccountNotifs.
func (c *AccountNotifsClient) Delete() *AccountNotifsDelete {
	mutation := newAccountNotifsMutation(c.config, OpDelete)
	return &AccountNotifsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountNotifsClient) DeleteOne(an *AccountNotifs) *AccountNotifsDeleteOne {
	return c.DeleteOneID(an.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountNotifsClient) DeleteOneID(id uuid.UUID) *AccountNotifsDeleteOne {
	builder := c.Delete().Where(accountnotifs.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountNotifsDeleteOne{builder}
}

// Query returns a query builder for AccountNotifs.
func (c *AccountNotifsClient) Query() *AccountNotifsQuery {
	return &AccountNotifsQuery{
		config: c.config,
	}
}

// Get returns a AccountNotifs entity by its id.
func (c *AccountNotifsClient) Get(ctx context.Context, id uuid.UUID) (*AccountNotifs, error) {
	return c.Query().Where(accountnotifs.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountNotifsClient) GetX(ctx context.Context, id uuid.UUID) *AccountNotifs {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountNotifsClient) Hooks() []Hook {
	return c.hooks.AccountNotifs
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

// TokensClient is a client for the Tokens schema.
type TokensClient struct {
	config
}

// NewTokensClient returns a client for the Tokens from the given config.
func NewTokensClient(c config) *TokensClient {
	return &TokensClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tokens.Hooks(f(g(h())))`.
func (c *TokensClient) Use(hooks ...Hook) {
	c.hooks.Tokens = append(c.hooks.Tokens, hooks...)
}

// Create returns a builder for creating a Tokens entity.
func (c *TokensClient) Create() *TokensCreate {
	mutation := newTokensMutation(c.config, OpCreate)
	return &TokensCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tokens entities.
func (c *TokensClient) CreateBulk(builders ...*TokensCreate) *TokensCreateBulk {
	return &TokensCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tokens.
func (c *TokensClient) Update() *TokensUpdate {
	mutation := newTokensMutation(c.config, OpUpdate)
	return &TokensUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokensClient) UpdateOne(t *Tokens) *TokensUpdateOne {
	mutation := newTokensMutation(c.config, OpUpdateOne, withTokens(t))
	return &TokensUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokensClient) UpdateOneID(id uuid.UUID) *TokensUpdateOne {
	mutation := newTokensMutation(c.config, OpUpdateOne, withTokensID(id))
	return &TokensUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tokens.
func (c *TokensClient) Delete() *TokensDelete {
	mutation := newTokensMutation(c.config, OpDelete)
	return &TokensDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokensClient) DeleteOne(t *Tokens) *TokensDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TokensClient) DeleteOneID(id uuid.UUID) *TokensDeleteOne {
	builder := c.Delete().Where(tokens.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokensDeleteOne{builder}
}

// Query returns a query builder for Tokens.
func (c *TokensClient) Query() *TokensQuery {
	return &TokensQuery{
		config: c.config,
	}
}

// Get returns a Tokens entity by its id.
func (c *TokensClient) Get(ctx context.Context, id uuid.UUID) (*Tokens, error) {
	return c.Query().Where(tokens.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokensClient) GetX(ctx context.Context, id uuid.UUID) *Tokens {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TokensClient) Hooks() []Hook {
	return c.hooks.Tokens
}

// UsersClient is a client for the Users schema.
type UsersClient struct {
	config
}

// NewUsersClient returns a client for the Users from the given config.
func NewUsersClient(c config) *UsersClient {
	return &UsersClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `users.Hooks(f(g(h())))`.
func (c *UsersClient) Use(hooks ...Hook) {
	c.hooks.Users = append(c.hooks.Users, hooks...)
}

// Create returns a builder for creating a Users entity.
func (c *UsersClient) Create() *UsersCreate {
	mutation := newUsersMutation(c.config, OpCreate)
	return &UsersCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Users entities.
func (c *UsersClient) CreateBulk(builders ...*UsersCreate) *UsersCreateBulk {
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Users.
func (c *UsersClient) Update() *UsersUpdate {
	mutation := newUsersMutation(c.config, OpUpdate)
	return &UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersClient) UpdateOne(u *Users) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsers(u))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersClient) UpdateOneID(id uuid.UUID) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsersID(id))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Users.
func (c *UsersClient) Delete() *UsersDelete {
	mutation := newUsersMutation(c.config, OpDelete)
	return &UsersDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersClient) DeleteOne(u *Users) *UsersDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UsersClient) DeleteOneID(id uuid.UUID) *UsersDeleteOne {
	builder := c.Delete().Where(users.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersDeleteOne{builder}
}

// Query returns a query builder for Users.
func (c *UsersClient) Query() *UsersQuery {
	return &UsersQuery{
		config: c.config,
	}
}

// Get returns a Users entity by its id.
func (c *UsersClient) Get(ctx context.Context, id uuid.UUID) (*Users, error) {
	return c.Query().Where(users.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersClient) GetX(ctx context.Context, id uuid.UUID) *Users {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UsersClient) Hooks() []Hook {
	return c.hooks.Users
}
