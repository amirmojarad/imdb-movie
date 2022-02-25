// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-movie/ent/movie"
	"imdb-movie/ent/predicate"
	"imdb-movie/ent/user"
	"sync"
	"time"

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
	TypeMovie = "Movie"
	TypeUser  = "User"
)

// MovieMutation represents an operation that mutates the Movie nodes in the graph.
type MovieMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	year          *string
	rated         *float32
	addrated      *float32
	realease_date *time.Time
	genre         *string
	language      *string
	poster        *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Movie, error)
	predicates    []predicate.Movie
}

var _ ent.Mutation = (*MovieMutation)(nil)

// movieOption allows management of the mutation configuration using functional options.
type movieOption func(*MovieMutation)

// newMovieMutation creates new mutation for the Movie entity.
func newMovieMutation(c config, op Op, opts ...movieOption) *MovieMutation {
	m := &MovieMutation{
		config:        c,
		op:            op,
		typ:           TypeMovie,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMovieID sets the ID field of the mutation.
func withMovieID(id int) movieOption {
	return func(m *MovieMutation) {
		var (
			err   error
			once  sync.Once
			value *Movie
		)
		m.oldValue = func(ctx context.Context) (*Movie, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Movie.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMovie sets the old Movie of the mutation.
func withMovie(node *Movie) movieOption {
	return func(m *MovieMutation) {
		m.oldValue = func(context.Context) (*Movie, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MovieMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MovieMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MovieMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MovieMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Movie.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *MovieMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *MovieMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *MovieMutation) ResetTitle() {
	m.title = nil
}

// SetYear sets the "year" field.
func (m *MovieMutation) SetYear(s string) {
	m.year = &s
}

// Year returns the value of the "year" field in the mutation.
func (m *MovieMutation) Year() (r string, exists bool) {
	v := m.year
	if v == nil {
		return
	}
	return *v, true
}

// OldYear returns the old "year" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldYear(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldYear is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldYear requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldYear: %w", err)
	}
	return oldValue.Year, nil
}

// ResetYear resets all changes to the "year" field.
func (m *MovieMutation) ResetYear() {
	m.year = nil
}

// SetRated sets the "rated" field.
func (m *MovieMutation) SetRated(f float32) {
	m.rated = &f
	m.addrated = nil
}

// Rated returns the value of the "rated" field in the mutation.
func (m *MovieMutation) Rated() (r float32, exists bool) {
	v := m.rated
	if v == nil {
		return
	}
	return *v, true
}

// OldRated returns the old "rated" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldRated(ctx context.Context) (v float32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRated: %w", err)
	}
	return oldValue.Rated, nil
}

// AddRated adds f to the "rated" field.
func (m *MovieMutation) AddRated(f float32) {
	if m.addrated != nil {
		*m.addrated += f
	} else {
		m.addrated = &f
	}
}

// AddedRated returns the value that was added to the "rated" field in this mutation.
func (m *MovieMutation) AddedRated() (r float32, exists bool) {
	v := m.addrated
	if v == nil {
		return
	}
	return *v, true
}

// ResetRated resets all changes to the "rated" field.
func (m *MovieMutation) ResetRated() {
	m.rated = nil
	m.addrated = nil
}

// SetRealeaseDate sets the "realease_date" field.
func (m *MovieMutation) SetRealeaseDate(t time.Time) {
	m.realease_date = &t
}

// RealeaseDate returns the value of the "realease_date" field in the mutation.
func (m *MovieMutation) RealeaseDate() (r time.Time, exists bool) {
	v := m.realease_date
	if v == nil {
		return
	}
	return *v, true
}

// OldRealeaseDate returns the old "realease_date" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldRealeaseDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRealeaseDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRealeaseDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRealeaseDate: %w", err)
	}
	return oldValue.RealeaseDate, nil
}

// ResetRealeaseDate resets all changes to the "realease_date" field.
func (m *MovieMutation) ResetRealeaseDate() {
	m.realease_date = nil
}

// SetGenre sets the "genre" field.
func (m *MovieMutation) SetGenre(s string) {
	m.genre = &s
}

// Genre returns the value of the "genre" field in the mutation.
func (m *MovieMutation) Genre() (r string, exists bool) {
	v := m.genre
	if v == nil {
		return
	}
	return *v, true
}

// OldGenre returns the old "genre" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldGenre(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGenre is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGenre requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGenre: %w", err)
	}
	return oldValue.Genre, nil
}

// ResetGenre resets all changes to the "genre" field.
func (m *MovieMutation) ResetGenre() {
	m.genre = nil
}

// SetLanguage sets the "language" field.
func (m *MovieMutation) SetLanguage(s string) {
	m.language = &s
}

// Language returns the value of the "language" field in the mutation.
func (m *MovieMutation) Language() (r string, exists bool) {
	v := m.language
	if v == nil {
		return
	}
	return *v, true
}

// OldLanguage returns the old "language" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldLanguage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLanguage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLanguage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLanguage: %w", err)
	}
	return oldValue.Language, nil
}

// ResetLanguage resets all changes to the "language" field.
func (m *MovieMutation) ResetLanguage() {
	m.language = nil
}

// SetPoster sets the "poster" field.
func (m *MovieMutation) SetPoster(s string) {
	m.poster = &s
}

// Poster returns the value of the "poster" field in the mutation.
func (m *MovieMutation) Poster() (r string, exists bool) {
	v := m.poster
	if v == nil {
		return
	}
	return *v, true
}

// OldPoster returns the old "poster" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldPoster(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPoster is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPoster requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPoster: %w", err)
	}
	return oldValue.Poster, nil
}

// ResetPoster resets all changes to the "poster" field.
func (m *MovieMutation) ResetPoster() {
	m.poster = nil
}

// Where appends a list predicates to the MovieMutation builder.
func (m *MovieMutation) Where(ps ...predicate.Movie) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MovieMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Movie).
func (m *MovieMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MovieMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.title != nil {
		fields = append(fields, movie.FieldTitle)
	}
	if m.year != nil {
		fields = append(fields, movie.FieldYear)
	}
	if m.rated != nil {
		fields = append(fields, movie.FieldRated)
	}
	if m.realease_date != nil {
		fields = append(fields, movie.FieldRealeaseDate)
	}
	if m.genre != nil {
		fields = append(fields, movie.FieldGenre)
	}
	if m.language != nil {
		fields = append(fields, movie.FieldLanguage)
	}
	if m.poster != nil {
		fields = append(fields, movie.FieldPoster)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MovieMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case movie.FieldTitle:
		return m.Title()
	case movie.FieldYear:
		return m.Year()
	case movie.FieldRated:
		return m.Rated()
	case movie.FieldRealeaseDate:
		return m.RealeaseDate()
	case movie.FieldGenre:
		return m.Genre()
	case movie.FieldLanguage:
		return m.Language()
	case movie.FieldPoster:
		return m.Poster()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MovieMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case movie.FieldTitle:
		return m.OldTitle(ctx)
	case movie.FieldYear:
		return m.OldYear(ctx)
	case movie.FieldRated:
		return m.OldRated(ctx)
	case movie.FieldRealeaseDate:
		return m.OldRealeaseDate(ctx)
	case movie.FieldGenre:
		return m.OldGenre(ctx)
	case movie.FieldLanguage:
		return m.OldLanguage(ctx)
	case movie.FieldPoster:
		return m.OldPoster(ctx)
	}
	return nil, fmt.Errorf("unknown Movie field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MovieMutation) SetField(name string, value ent.Value) error {
	switch name {
	case movie.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case movie.FieldYear:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetYear(v)
		return nil
	case movie.FieldRated:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRated(v)
		return nil
	case movie.FieldRealeaseDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRealeaseDate(v)
		return nil
	case movie.FieldGenre:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGenre(v)
		return nil
	case movie.FieldLanguage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLanguage(v)
		return nil
	case movie.FieldPoster:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPoster(v)
		return nil
	}
	return fmt.Errorf("unknown Movie field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MovieMutation) AddedFields() []string {
	var fields []string
	if m.addrated != nil {
		fields = append(fields, movie.FieldRated)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MovieMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case movie.FieldRated:
		return m.AddedRated()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MovieMutation) AddField(name string, value ent.Value) error {
	switch name {
	case movie.FieldRated:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRated(v)
		return nil
	}
	return fmt.Errorf("unknown Movie numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MovieMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MovieMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MovieMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Movie nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MovieMutation) ResetField(name string) error {
	switch name {
	case movie.FieldTitle:
		m.ResetTitle()
		return nil
	case movie.FieldYear:
		m.ResetYear()
		return nil
	case movie.FieldRated:
		m.ResetRated()
		return nil
	case movie.FieldRealeaseDate:
		m.ResetRealeaseDate()
		return nil
	case movie.FieldGenre:
		m.ResetGenre()
		return nil
	case movie.FieldLanguage:
		m.ResetLanguage()
		return nil
	case movie.FieldPoster:
		m.ResetPoster()
		return nil
	}
	return fmt.Errorf("unknown Movie field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MovieMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MovieMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MovieMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MovieMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MovieMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MovieMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MovieMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Movie unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MovieMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Movie edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	username      *string
	email         *string
	full_name     *string
	password      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetFullName sets the "full_name" field.
func (m *UserMutation) SetFullName(s string) {
	m.full_name = &s
}

// FullName returns the value of the "full_name" field in the mutation.
func (m *UserMutation) FullName() (r string, exists bool) {
	v := m.full_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFullName returns the old "full_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldFullName(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFullName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFullName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFullName: %w", err)
	}
	return oldValue.FullName, nil
}

// ClearFullName clears the value of the "full_name" field.
func (m *UserMutation) ClearFullName() {
	m.full_name = nil
	m.clearedFields[user.FieldFullName] = struct{}{}
}

// FullNameCleared returns if the "full_name" field was cleared in this mutation.
func (m *UserMutation) FullNameCleared() bool {
	_, ok := m.clearedFields[user.FieldFullName]
	return ok
}

// ResetFullName resets all changes to the "full_name" field.
func (m *UserMutation) ResetFullName() {
	m.full_name = nil
	delete(m.clearedFields, user.FieldFullName)
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.full_name != nil {
		fields = append(fields, user.FieldFullName)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	case user.FieldEmail:
		return m.Email()
	case user.FieldFullName:
		return m.FullName()
	case user.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldFullName:
		return m.OldFullName(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldFullName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFullName(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldFullName) {
		fields = append(fields, user.FieldFullName)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldFullName:
		m.ClearFullName()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldFullName:
		m.ResetFullName()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
