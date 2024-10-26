// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/kevindamm/q-party/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kevindamm/q-party/ent/category"
	"github.com/kevindamm/q-party/ent/challenge"
	"github.com/kevindamm/q-party/ent/challengegroup"
	"github.com/kevindamm/q-party/ent/episode"
	"github.com/kevindamm/q-party/ent/episoderound"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Challenge is the client for interacting with the Challenge builders.
	Challenge *ChallengeClient
	// ChallengeGroup is the client for interacting with the ChallengeGroup builders.
	ChallengeGroup *ChallengeGroupClient
	// Episode is the client for interacting with the Episode builders.
	Episode *EpisodeClient
	// EpisodeRound is the client for interacting with the EpisodeRound builders.
	EpisodeRound *EpisodeRoundClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Category = NewCategoryClient(c.config)
	c.Challenge = NewChallengeClient(c.config)
	c.ChallengeGroup = NewChallengeGroupClient(c.config)
	c.Episode = NewEpisodeClient(c.config)
	c.EpisodeRound = NewEpisodeRoundClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Category:       NewCategoryClient(cfg),
		Challenge:      NewChallengeClient(cfg),
		ChallengeGroup: NewChallengeGroupClient(cfg),
		Episode:        NewEpisodeClient(cfg),
		EpisodeRound:   NewEpisodeRoundClient(cfg),
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
		ctx:            ctx,
		config:         cfg,
		Category:       NewCategoryClient(cfg),
		Challenge:      NewChallengeClient(cfg),
		ChallengeGroup: NewChallengeGroupClient(cfg),
		Episode:        NewEpisodeClient(cfg),
		EpisodeRound:   NewEpisodeRoundClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Category.
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
	c.Category.Use(hooks...)
	c.Challenge.Use(hooks...)
	c.ChallengeGroup.Use(hooks...)
	c.Episode.Use(hooks...)
	c.EpisodeRound.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Category.Intercept(interceptors...)
	c.Challenge.Intercept(interceptors...)
	c.ChallengeGroup.Intercept(interceptors...)
	c.Episode.Intercept(interceptors...)
	c.EpisodeRound.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CategoryMutation:
		return c.Category.mutate(ctx, m)
	case *ChallengeMutation:
		return c.Challenge.mutate(ctx, m)
	case *ChallengeGroupMutation:
		return c.ChallengeGroup.mutate(ctx, m)
	case *EpisodeMutation:
		return c.Episode.mutate(ctx, m)
	case *EpisodeRoundMutation:
		return c.EpisodeRound.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `category.Intercept(f(g(h())))`.
func (c *CategoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Category = append(c.inters.Category, interceptors...)
}

// Create returns a builder for creating a Category entity.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CategoryClient) MapCreateBulk(slice any, setFunc func(*CategoryCreate, int)) *CategoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CategoryCreateBulk{err: fmt.Errorf("calling to CategoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CategoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id int) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CategoryClient) DeleteOneID(id int) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCategory},
		inters: c.Interceptors(),
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id int) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id int) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChallenges queries the challenges edge of a Category.
func (c *CategoryClient) QueryChallenges(ca *Category) *ChallengeGroupQuery {
	query := (&ChallengeGroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(challengegroup.Table, challengegroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, category.ChallengesTable, category.ChallengesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// Interceptors returns the client interceptors.
func (c *CategoryClient) Interceptors() []Interceptor {
	return c.inters.Category
}

func (c *CategoryClient) mutate(ctx context.Context, m *CategoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Category mutation op: %q", m.Op())
	}
}

// ChallengeClient is a client for the Challenge schema.
type ChallengeClient struct {
	config
}

// NewChallengeClient returns a client for the Challenge from the given config.
func NewChallengeClient(c config) *ChallengeClient {
	return &ChallengeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `challenge.Hooks(f(g(h())))`.
func (c *ChallengeClient) Use(hooks ...Hook) {
	c.hooks.Challenge = append(c.hooks.Challenge, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `challenge.Intercept(f(g(h())))`.
func (c *ChallengeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Challenge = append(c.inters.Challenge, interceptors...)
}

// Create returns a builder for creating a Challenge entity.
func (c *ChallengeClient) Create() *ChallengeCreate {
	mutation := newChallengeMutation(c.config, OpCreate)
	return &ChallengeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Challenge entities.
func (c *ChallengeClient) CreateBulk(builders ...*ChallengeCreate) *ChallengeCreateBulk {
	return &ChallengeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChallengeClient) MapCreateBulk(slice any, setFunc func(*ChallengeCreate, int)) *ChallengeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChallengeCreateBulk{err: fmt.Errorf("calling to ChallengeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChallengeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChallengeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Challenge.
func (c *ChallengeClient) Update() *ChallengeUpdate {
	mutation := newChallengeMutation(c.config, OpUpdate)
	return &ChallengeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChallengeClient) UpdateOne(ch *Challenge) *ChallengeUpdateOne {
	mutation := newChallengeMutation(c.config, OpUpdateOne, withChallenge(ch))
	return &ChallengeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChallengeClient) UpdateOneID(id int) *ChallengeUpdateOne {
	mutation := newChallengeMutation(c.config, OpUpdateOne, withChallengeID(id))
	return &ChallengeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Challenge.
func (c *ChallengeClient) Delete() *ChallengeDelete {
	mutation := newChallengeMutation(c.config, OpDelete)
	return &ChallengeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChallengeClient) DeleteOne(ch *Challenge) *ChallengeDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChallengeClient) DeleteOneID(id int) *ChallengeDeleteOne {
	builder := c.Delete().Where(challenge.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChallengeDeleteOne{builder}
}

// Query returns a query builder for Challenge.
func (c *ChallengeClient) Query() *ChallengeQuery {
	return &ChallengeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChallenge},
		inters: c.Interceptors(),
	}
}

// Get returns a Challenge entity by its id.
func (c *ChallengeClient) Get(ctx context.Context, id int) (*Challenge, error) {
	return c.Query().Where(challenge.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChallengeClient) GetX(ctx context.Context, id int) *Challenge {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryColumn queries the column edge of a Challenge.
func (c *ChallengeClient) QueryColumn(ch *Challenge) *ChallengeGroupQuery {
	query := (&ChallengeGroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(challenge.Table, challenge.FieldID, id),
			sqlgraph.To(challengegroup.Table, challengegroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, challenge.ColumnTable, challenge.ColumnPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChallengeClient) Hooks() []Hook {
	return c.hooks.Challenge
}

// Interceptors returns the client interceptors.
func (c *ChallengeClient) Interceptors() []Interceptor {
	return c.inters.Challenge
}

func (c *ChallengeClient) mutate(ctx context.Context, m *ChallengeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChallengeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChallengeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChallengeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChallengeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Challenge mutation op: %q", m.Op())
	}
}

// ChallengeGroupClient is a client for the ChallengeGroup schema.
type ChallengeGroupClient struct {
	config
}

// NewChallengeGroupClient returns a client for the ChallengeGroup from the given config.
func NewChallengeGroupClient(c config) *ChallengeGroupClient {
	return &ChallengeGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `challengegroup.Hooks(f(g(h())))`.
func (c *ChallengeGroupClient) Use(hooks ...Hook) {
	c.hooks.ChallengeGroup = append(c.hooks.ChallengeGroup, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `challengegroup.Intercept(f(g(h())))`.
func (c *ChallengeGroupClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChallengeGroup = append(c.inters.ChallengeGroup, interceptors...)
}

// Create returns a builder for creating a ChallengeGroup entity.
func (c *ChallengeGroupClient) Create() *ChallengeGroupCreate {
	mutation := newChallengeGroupMutation(c.config, OpCreate)
	return &ChallengeGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChallengeGroup entities.
func (c *ChallengeGroupClient) CreateBulk(builders ...*ChallengeGroupCreate) *ChallengeGroupCreateBulk {
	return &ChallengeGroupCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChallengeGroupClient) MapCreateBulk(slice any, setFunc func(*ChallengeGroupCreate, int)) *ChallengeGroupCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChallengeGroupCreateBulk{err: fmt.Errorf("calling to ChallengeGroupClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChallengeGroupCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChallengeGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChallengeGroup.
func (c *ChallengeGroupClient) Update() *ChallengeGroupUpdate {
	mutation := newChallengeGroupMutation(c.config, OpUpdate)
	return &ChallengeGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChallengeGroupClient) UpdateOne(cg *ChallengeGroup) *ChallengeGroupUpdateOne {
	mutation := newChallengeGroupMutation(c.config, OpUpdateOne, withChallengeGroup(cg))
	return &ChallengeGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChallengeGroupClient) UpdateOneID(id int) *ChallengeGroupUpdateOne {
	mutation := newChallengeGroupMutation(c.config, OpUpdateOne, withChallengeGroupID(id))
	return &ChallengeGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChallengeGroup.
func (c *ChallengeGroupClient) Delete() *ChallengeGroupDelete {
	mutation := newChallengeGroupMutation(c.config, OpDelete)
	return &ChallengeGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChallengeGroupClient) DeleteOne(cg *ChallengeGroup) *ChallengeGroupDeleteOne {
	return c.DeleteOneID(cg.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChallengeGroupClient) DeleteOneID(id int) *ChallengeGroupDeleteOne {
	builder := c.Delete().Where(challengegroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChallengeGroupDeleteOne{builder}
}

// Query returns a query builder for ChallengeGroup.
func (c *ChallengeGroupClient) Query() *ChallengeGroupQuery {
	return &ChallengeGroupQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChallengeGroup},
		inters: c.Interceptors(),
	}
}

// Get returns a ChallengeGroup entity by its id.
func (c *ChallengeGroupClient) Get(ctx context.Context, id int) (*ChallengeGroup, error) {
	return c.Query().Where(challengegroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChallengeGroupClient) GetX(ctx context.Context, id int) *ChallengeGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCategory queries the category edge of a ChallengeGroup.
func (c *ChallengeGroupClient) QueryCategory(cg *ChallengeGroup) *CategoryQuery {
	query := (&CategoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(challengegroup.Table, challengegroup.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, challengegroup.CategoryTable, challengegroup.CategoryPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChallenges queries the challenges edge of a ChallengeGroup.
func (c *ChallengeGroupClient) QueryChallenges(cg *ChallengeGroup) *ChallengeQuery {
	query := (&ChallengeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(challengegroup.Table, challengegroup.FieldID, id),
			sqlgraph.To(challenge.Table, challenge.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, challengegroup.ChallengesTable, challengegroup.ChallengesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEpisodeRound queries the episode_round edge of a ChallengeGroup.
func (c *ChallengeGroupClient) QueryEpisodeRound(cg *ChallengeGroup) *EpisodeRoundQuery {
	query := (&EpisodeRoundClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(challengegroup.Table, challengegroup.FieldID, id),
			sqlgraph.To(episoderound.Table, episoderound.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, challengegroup.EpisodeRoundTable, challengegroup.EpisodeRoundPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChallengeGroupClient) Hooks() []Hook {
	return c.hooks.ChallengeGroup
}

// Interceptors returns the client interceptors.
func (c *ChallengeGroupClient) Interceptors() []Interceptor {
	return c.inters.ChallengeGroup
}

func (c *ChallengeGroupClient) mutate(ctx context.Context, m *ChallengeGroupMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChallengeGroupCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChallengeGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChallengeGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChallengeGroupDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ChallengeGroup mutation op: %q", m.Op())
	}
}

// EpisodeClient is a client for the Episode schema.
type EpisodeClient struct {
	config
}

// NewEpisodeClient returns a client for the Episode from the given config.
func NewEpisodeClient(c config) *EpisodeClient {
	return &EpisodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episode.Hooks(f(g(h())))`.
func (c *EpisodeClient) Use(hooks ...Hook) {
	c.hooks.Episode = append(c.hooks.Episode, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `episode.Intercept(f(g(h())))`.
func (c *EpisodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Episode = append(c.inters.Episode, interceptors...)
}

// Create returns a builder for creating a Episode entity.
func (c *EpisodeClient) Create() *EpisodeCreate {
	mutation := newEpisodeMutation(c.config, OpCreate)
	return &EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Episode entities.
func (c *EpisodeClient) CreateBulk(builders ...*EpisodeCreate) *EpisodeCreateBulk {
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EpisodeClient) MapCreateBulk(slice any, setFunc func(*EpisodeCreate, int)) *EpisodeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EpisodeCreateBulk{err: fmt.Errorf("calling to EpisodeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EpisodeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Episode.
func (c *EpisodeClient) Update() *EpisodeUpdate {
	mutation := newEpisodeMutation(c.config, OpUpdate)
	return &EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeClient) UpdateOne(e *Episode) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisode(e))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeClient) UpdateOneID(id int) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisodeID(id))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Episode.
func (c *EpisodeClient) Delete() *EpisodeDelete {
	mutation := newEpisodeMutation(c.config, OpDelete)
	return &EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EpisodeClient) DeleteOne(e *Episode) *EpisodeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EpisodeClient) DeleteOneID(id int) *EpisodeDeleteOne {
	builder := c.Delete().Where(episode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeDeleteOne{builder}
}

// Query returns a query builder for Episode.
func (c *EpisodeClient) Query() *EpisodeQuery {
	return &EpisodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEpisode},
		inters: c.Interceptors(),
	}
}

// Get returns a Episode entity by its id.
func (c *EpisodeClient) Get(ctx context.Context, id int) (*Episode, error) {
	return c.Query().Where(episode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeClient) GetX(ctx context.Context, id int) *Episode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRounds queries the rounds edge of a Episode.
func (c *EpisodeClient) QueryRounds(e *Episode) *EpisodeRoundQuery {
	query := (&EpisodeRoundClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(episoderound.Table, episoderound.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, episode.RoundsTable, episode.RoundsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EpisodeClient) Hooks() []Hook {
	return c.hooks.Episode
}

// Interceptors returns the client interceptors.
func (c *EpisodeClient) Interceptors() []Interceptor {
	return c.inters.Episode
}

func (c *EpisodeClient) mutate(ctx context.Context, m *EpisodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Episode mutation op: %q", m.Op())
	}
}

// EpisodeRoundClient is a client for the EpisodeRound schema.
type EpisodeRoundClient struct {
	config
}

// NewEpisodeRoundClient returns a client for the EpisodeRound from the given config.
func NewEpisodeRoundClient(c config) *EpisodeRoundClient {
	return &EpisodeRoundClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episoderound.Hooks(f(g(h())))`.
func (c *EpisodeRoundClient) Use(hooks ...Hook) {
	c.hooks.EpisodeRound = append(c.hooks.EpisodeRound, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `episoderound.Intercept(f(g(h())))`.
func (c *EpisodeRoundClient) Intercept(interceptors ...Interceptor) {
	c.inters.EpisodeRound = append(c.inters.EpisodeRound, interceptors...)
}

// Create returns a builder for creating a EpisodeRound entity.
func (c *EpisodeRoundClient) Create() *EpisodeRoundCreate {
	mutation := newEpisodeRoundMutation(c.config, OpCreate)
	return &EpisodeRoundCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EpisodeRound entities.
func (c *EpisodeRoundClient) CreateBulk(builders ...*EpisodeRoundCreate) *EpisodeRoundCreateBulk {
	return &EpisodeRoundCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EpisodeRoundClient) MapCreateBulk(slice any, setFunc func(*EpisodeRoundCreate, int)) *EpisodeRoundCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EpisodeRoundCreateBulk{err: fmt.Errorf("calling to EpisodeRoundClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EpisodeRoundCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EpisodeRoundCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EpisodeRound.
func (c *EpisodeRoundClient) Update() *EpisodeRoundUpdate {
	mutation := newEpisodeRoundMutation(c.config, OpUpdate)
	return &EpisodeRoundUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeRoundClient) UpdateOne(er *EpisodeRound) *EpisodeRoundUpdateOne {
	mutation := newEpisodeRoundMutation(c.config, OpUpdateOne, withEpisodeRound(er))
	return &EpisodeRoundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeRoundClient) UpdateOneID(id int) *EpisodeRoundUpdateOne {
	mutation := newEpisodeRoundMutation(c.config, OpUpdateOne, withEpisodeRoundID(id))
	return &EpisodeRoundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EpisodeRound.
func (c *EpisodeRoundClient) Delete() *EpisodeRoundDelete {
	mutation := newEpisodeRoundMutation(c.config, OpDelete)
	return &EpisodeRoundDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EpisodeRoundClient) DeleteOne(er *EpisodeRound) *EpisodeRoundDeleteOne {
	return c.DeleteOneID(er.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EpisodeRoundClient) DeleteOneID(id int) *EpisodeRoundDeleteOne {
	builder := c.Delete().Where(episoderound.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeRoundDeleteOne{builder}
}

// Query returns a query builder for EpisodeRound.
func (c *EpisodeRoundClient) Query() *EpisodeRoundQuery {
	return &EpisodeRoundQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEpisodeRound},
		inters: c.Interceptors(),
	}
}

// Get returns a EpisodeRound entity by its id.
func (c *EpisodeRoundClient) Get(ctx context.Context, id int) (*EpisodeRound, error) {
	return c.Query().Where(episoderound.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeRoundClient) GetX(ctx context.Context, id int) *EpisodeRound {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryColumns queries the columns edge of a EpisodeRound.
func (c *EpisodeRoundClient) QueryColumns(er *EpisodeRound) *ChallengeGroupQuery {
	query := (&ChallengeGroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := er.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episoderound.Table, episoderound.FieldID, id),
			sqlgraph.To(challengegroup.Table, challengegroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, episoderound.ColumnsTable, episoderound.ColumnsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(er.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEpisode queries the episode edge of a EpisodeRound.
func (c *EpisodeRoundClient) QueryEpisode(er *EpisodeRound) *EpisodeQuery {
	query := (&EpisodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := er.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episoderound.Table, episoderound.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, episoderound.EpisodeTable, episoderound.EpisodePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(er.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EpisodeRoundClient) Hooks() []Hook {
	return c.hooks.EpisodeRound
}

// Interceptors returns the client interceptors.
func (c *EpisodeRoundClient) Interceptors() []Interceptor {
	return c.inters.EpisodeRound
}

func (c *EpisodeRoundClient) mutate(ctx context.Context, m *EpisodeRoundMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EpisodeRoundCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EpisodeRoundUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EpisodeRoundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EpisodeRoundDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EpisodeRound mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Category, Challenge, ChallengeGroup, Episode, EpisodeRound []ent.Hook
	}
	inters struct {
		Category, Challenge, ChallengeGroup, Episode, EpisodeRound []ent.Interceptor
	}
)
