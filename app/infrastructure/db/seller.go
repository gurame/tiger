// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Seller is an object representing the database table.
type Seller struct {
	Sellerid        string      `boil:"sellerid" json:"sellerid" toml:"sellerid" yaml:"sellerid"`
	Taxid           string      `boil:"taxid" json:"taxid" toml:"taxid" yaml:"taxid"`
	Name            null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Created         time.Time   `boil:"created" json:"created" toml:"created" yaml:"created"`
	Createdby       string      `boil:"createdby" json:"createdby" toml:"createdby" yaml:"createdby"`
	Lastmodified    null.Time   `boil:"lastmodified" json:"lastmodified,omitempty" toml:"lastmodified" yaml:"lastmodified,omitempty"`
	Lastmodifiedby  null.String `boil:"lastmodifiedby" json:"lastmodifiedby,omitempty" toml:"lastmodifiedby" yaml:"lastmodifiedby,omitempty"`
	Baseentitystate int         `boil:"baseentitystate" json:"baseentitystate" toml:"baseentitystate" yaml:"baseentitystate"`

	R *sellerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sellerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SellerColumns = struct {
	Sellerid        string
	Taxid           string
	Name            string
	Created         string
	Createdby       string
	Lastmodified    string
	Lastmodifiedby  string
	Baseentitystate string
}{
	Sellerid:        "sellerid",
	Taxid:           "taxid",
	Name:            "name",
	Created:         "created",
	Createdby:       "createdby",
	Lastmodified:    "lastmodified",
	Lastmodifiedby:  "lastmodifiedby",
	Baseentitystate: "baseentitystate",
}

var SellerTableColumns = struct {
	Sellerid        string
	Taxid           string
	Name            string
	Created         string
	Createdby       string
	Lastmodified    string
	Lastmodifiedby  string
	Baseentitystate string
}{
	Sellerid:        "seller.sellerid",
	Taxid:           "seller.taxid",
	Name:            "seller.name",
	Created:         "seller.created",
	Createdby:       "seller.createdby",
	Lastmodified:    "seller.lastmodified",
	Lastmodifiedby:  "seller.lastmodifiedby",
	Baseentitystate: "seller.baseentitystate",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var SellerWhere = struct {
	Sellerid        whereHelperstring
	Taxid           whereHelperstring
	Name            whereHelpernull_String
	Created         whereHelpertime_Time
	Createdby       whereHelperstring
	Lastmodified    whereHelpernull_Time
	Lastmodifiedby  whereHelpernull_String
	Baseentitystate whereHelperint
}{
	Sellerid:        whereHelperstring{field: "\"seller\".\"sellerid\""},
	Taxid:           whereHelperstring{field: "\"seller\".\"taxid\""},
	Name:            whereHelpernull_String{field: "\"seller\".\"name\""},
	Created:         whereHelpertime_Time{field: "\"seller\".\"created\""},
	Createdby:       whereHelperstring{field: "\"seller\".\"createdby\""},
	Lastmodified:    whereHelpernull_Time{field: "\"seller\".\"lastmodified\""},
	Lastmodifiedby:  whereHelpernull_String{field: "\"seller\".\"lastmodifiedby\""},
	Baseentitystate: whereHelperint{field: "\"seller\".\"baseentitystate\""},
}

// SellerRels is where relationship names are stored.
var SellerRels = struct {
}{}

// sellerR is where relationships are stored.
type sellerR struct {
}

// NewStruct creates a new relationship struct
func (*sellerR) NewStruct() *sellerR {
	return &sellerR{}
}

// sellerL is where Load methods for each relationship are stored.
type sellerL struct{}

var (
	sellerAllColumns            = []string{"sellerid", "taxid", "name", "created", "createdby", "lastmodified", "lastmodifiedby", "baseentitystate"}
	sellerColumnsWithoutDefault = []string{"sellerid", "taxid", "created", "createdby", "baseentitystate"}
	sellerColumnsWithDefault    = []string{"name", "lastmodified", "lastmodifiedby"}
	sellerPrimaryKeyColumns     = []string{"sellerid"}
	sellerGeneratedColumns      = []string{}
)

type (
	// SellerSlice is an alias for a slice of pointers to Seller.
	// This should almost always be used instead of []Seller.
	SellerSlice []*Seller
	// SellerHook is the signature for custom Seller hook methods
	SellerHook func(context.Context, boil.ContextExecutor, *Seller) error

	sellerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sellerType                 = reflect.TypeOf(&Seller{})
	sellerMapping              = queries.MakeStructMapping(sellerType)
	sellerPrimaryKeyMapping, _ = queries.BindMapping(sellerType, sellerMapping, sellerPrimaryKeyColumns)
	sellerInsertCacheMut       sync.RWMutex
	sellerInsertCache          = make(map[string]insertCache)
	sellerUpdateCacheMut       sync.RWMutex
	sellerUpdateCache          = make(map[string]updateCache)
	sellerUpsertCacheMut       sync.RWMutex
	sellerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sellerAfterSelectHooks []SellerHook

var sellerBeforeInsertHooks []SellerHook
var sellerAfterInsertHooks []SellerHook

var sellerBeforeUpdateHooks []SellerHook
var sellerAfterUpdateHooks []SellerHook

var sellerBeforeDeleteHooks []SellerHook
var sellerAfterDeleteHooks []SellerHook

var sellerBeforeUpsertHooks []SellerHook
var sellerAfterUpsertHooks []SellerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Seller) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Seller) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Seller) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Seller) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Seller) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Seller) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Seller) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Seller) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Seller) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSellerHook registers your hook function for all future operations.
func AddSellerHook(hookPoint boil.HookPoint, sellerHook SellerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sellerAfterSelectHooks = append(sellerAfterSelectHooks, sellerHook)
	case boil.BeforeInsertHook:
		sellerBeforeInsertHooks = append(sellerBeforeInsertHooks, sellerHook)
	case boil.AfterInsertHook:
		sellerAfterInsertHooks = append(sellerAfterInsertHooks, sellerHook)
	case boil.BeforeUpdateHook:
		sellerBeforeUpdateHooks = append(sellerBeforeUpdateHooks, sellerHook)
	case boil.AfterUpdateHook:
		sellerAfterUpdateHooks = append(sellerAfterUpdateHooks, sellerHook)
	case boil.BeforeDeleteHook:
		sellerBeforeDeleteHooks = append(sellerBeforeDeleteHooks, sellerHook)
	case boil.AfterDeleteHook:
		sellerAfterDeleteHooks = append(sellerAfterDeleteHooks, sellerHook)
	case boil.BeforeUpsertHook:
		sellerBeforeUpsertHooks = append(sellerBeforeUpsertHooks, sellerHook)
	case boil.AfterUpsertHook:
		sellerAfterUpsertHooks = append(sellerAfterUpsertHooks, sellerHook)
	}
}

// OneG returns a single seller record from the query using the global executor.
func (q sellerQuery) OneG(ctx context.Context) (*Seller, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single seller record from the query.
func (q sellerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Seller, error) {
	o := &Seller{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for seller")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Seller records from the query using the global executor.
func (q sellerQuery) AllG(ctx context.Context) (SellerSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Seller records from the query.
func (q sellerQuery) All(ctx context.Context, exec boil.ContextExecutor) (SellerSlice, error) {
	var o []*Seller

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to Seller slice")
	}

	if len(sellerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Seller records in the query using the global executor
func (q sellerQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Seller records in the query.
func (q sellerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count seller rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q sellerQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q sellerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if seller exists")
	}

	return count > 0, nil
}

// Sellers retrieves all the records using an executor.
func Sellers(mods ...qm.QueryMod) sellerQuery {
	mods = append(mods, qm.From("\"seller\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"seller\".*"})
	}

	return sellerQuery{q}
}

// FindSellerG retrieves a single record by ID.
func FindSellerG(ctx context.Context, sellerid string, selectCols ...string) (*Seller, error) {
	return FindSeller(ctx, boil.GetContextDB(), sellerid, selectCols...)
}

// FindSeller retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSeller(ctx context.Context, exec boil.ContextExecutor, sellerid string, selectCols ...string) (*Seller, error) {
	sellerObj := &Seller{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"seller\" where \"sellerid\"=$1", sel,
	)

	q := queries.Raw(query, sellerid)

	err := q.Bind(ctx, exec, sellerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from seller")
	}

	if err = sellerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sellerObj, err
	}

	return sellerObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Seller) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Seller) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no seller provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sellerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sellerInsertCacheMut.RLock()
	cache, cached := sellerInsertCache[key]
	sellerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sellerAllColumns,
			sellerColumnsWithDefault,
			sellerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sellerType, sellerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sellerType, sellerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"seller\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"seller\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into seller")
	}

	if !cached {
		sellerInsertCacheMut.Lock()
		sellerInsertCache[key] = cache
		sellerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Seller record using the global executor.
// See Update for more documentation.
func (o *Seller) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Seller.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Seller) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sellerUpdateCacheMut.RLock()
	cache, cached := sellerUpdateCache[key]
	sellerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sellerAllColumns,
			sellerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update seller, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"seller\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, sellerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sellerType, sellerMapping, append(wl, sellerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update seller row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for seller")
	}

	if !cached {
		sellerUpdateCacheMut.Lock()
		sellerUpdateCache[key] = cache
		sellerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q sellerQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q sellerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for seller")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for seller")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SellerSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SellerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"seller\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, sellerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in seller slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all seller")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Seller) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Seller) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no seller provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sellerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	sellerUpsertCacheMut.RLock()
	cache, cached := sellerUpsertCache[key]
	sellerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sellerAllColumns,
			sellerColumnsWithDefault,
			sellerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sellerAllColumns,
			sellerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db: unable to upsert seller, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(sellerPrimaryKeyColumns))
			copy(conflict, sellerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"seller\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(sellerType, sellerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sellerType, sellerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "db: unable to upsert seller")
	}

	if !cached {
		sellerUpsertCacheMut.Lock()
		sellerUpsertCache[key] = cache
		sellerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Seller record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Seller) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Seller record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Seller) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Seller provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sellerPrimaryKeyMapping)
	sql := "DELETE FROM \"seller\" WHERE \"sellerid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from seller")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for seller")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q sellerQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q sellerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no sellerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from seller")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for seller")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SellerSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SellerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sellerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"seller\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sellerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from seller slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for seller")
	}

	if len(sellerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Seller) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("db: no Seller provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Seller) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSeller(ctx, exec, o.Sellerid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SellerSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("db: empty SellerSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SellerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SellerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"seller\".* FROM \"seller\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sellerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in SellerSlice")
	}

	*o = slice

	return nil
}

// SellerExistsG checks if the Seller row exists.
func SellerExistsG(ctx context.Context, sellerid string) (bool, error) {
	return SellerExists(ctx, boil.GetContextDB(), sellerid)
}

// SellerExists checks if the Seller row exists.
func SellerExists(ctx context.Context, exec boil.ContextExecutor, sellerid string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"seller\" where \"sellerid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, sellerid)
	}
	row := exec.QueryRowContext(ctx, sql, sellerid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if seller exists")
	}

	return exists, nil
}
