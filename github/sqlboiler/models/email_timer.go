// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// EmailTimer is an object representing the database table.
type EmailTimer struct {
	DataID    uint64    `boil:"data_id" json:"data_id" toml:"data_id" yaml:"data_id"`
	Title     string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Content   string    `boil:"content" json:"content" toml:"content" yaml:"content"`
	StartTime int       `boil:"start_time" json:"start_time" toml:"start_time" yaml:"start_time"`
	EndTime   int       `boil:"end_time" json:"end_time" toml:"end_time" yaml:"end_time"`
	EmailType bool      `boil:"email_type" json:"email_type" toml:"email_type" yaml:"email_type"`
	Attaches  string    `boil:"attaches" json:"attaches" toml:"attaches" yaml:"attaches"`
	IsSend    bool      `boil:"is_send" json:"is_send" toml:"is_send" yaml:"is_send"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *emailTimerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L emailTimerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EmailTimerColumns = struct {
	DataID    string
	Title     string
	Content   string
	StartTime string
	EndTime   string
	EmailType string
	Attaches  string
	IsSend    string
	CreatedAt string
}{
	DataID:    "data_id",
	Title:     "title",
	Content:   "content",
	StartTime: "start_time",
	EndTime:   "end_time",
	EmailType: "email_type",
	Attaches:  "attaches",
	IsSend:    "is_send",
	CreatedAt: "created_at",
}

// EmailTimerRels is where relationship names are stored.
var EmailTimerRels = struct {
}{}

// emailTimerR is where relationships are stored.
type emailTimerR struct {
}

// NewStruct creates a new relationship struct
func (*emailTimerR) NewStruct() *emailTimerR {
	return &emailTimerR{}
}

// emailTimerL is where Load methods for each relationship are stored.
type emailTimerL struct{}

var (
	emailTimerColumns               = []string{"data_id", "title", "content", "start_time", "end_time", "email_type", "attaches", "is_send", "created_at"}
	emailTimerColumnsWithoutDefault = []string{"title", "content", "attaches"}
	emailTimerColumnsWithDefault    = []string{"data_id", "start_time", "end_time", "email_type", "is_send", "created_at"}
	emailTimerPrimaryKeyColumns     = []string{"data_id"}
)

type (
	// EmailTimerSlice is an alias for a slice of pointers to EmailTimer.
	// This should generally be used opposed to []EmailTimer.
	EmailTimerSlice []*EmailTimer

	emailTimerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	emailTimerType                 = reflect.TypeOf(&EmailTimer{})
	emailTimerMapping              = queries.MakeStructMapping(emailTimerType)
	emailTimerPrimaryKeyMapping, _ = queries.BindMapping(emailTimerType, emailTimerMapping, emailTimerPrimaryKeyColumns)
	emailTimerInsertCacheMut       sync.RWMutex
	emailTimerInsertCache          = make(map[string]insertCache)
	emailTimerUpdateCacheMut       sync.RWMutex
	emailTimerUpdateCache          = make(map[string]updateCache)
	emailTimerUpsertCacheMut       sync.RWMutex
	emailTimerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

// One returns a single emailTimer record from the query.
func (q emailTimerQuery) One(exec boil.Executor) (*EmailTimer, error) {
	o := &EmailTimer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for email_timer")
	}

	return o, nil
}

// All returns all EmailTimer records from the query.
func (q emailTimerQuery) All(exec boil.Executor) (EmailTimerSlice, error) {
	var o []*EmailTimer

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EmailTimer slice")
	}

	return o, nil
}

// Count returns the count of all EmailTimer records in the query.
func (q emailTimerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count email_timer rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q emailTimerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if email_timer exists")
	}

	return count > 0, nil
}

// EmailTimers retrieves all the records using an executor.
func EmailTimers(mods ...qm.QueryMod) emailTimerQuery {
	mods = append(mods, qm.From("`email_timer`"))
	return emailTimerQuery{NewQuery(mods...)}
}

// FindEmailTimer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEmailTimer(exec boil.Executor, dataID uint64, selectCols ...string) (*EmailTimer, error) {
	emailTimerObj := &EmailTimer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `email_timer` where `data_id`=?", sel,
	)

	q := queries.Raw(query, dataID)

	err := q.Bind(nil, exec, emailTimerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from email_timer")
	}

	return emailTimerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EmailTimer) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no email_timer provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(emailTimerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	emailTimerInsertCacheMut.RLock()
	cache, cached := emailTimerInsertCache[key]
	emailTimerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			emailTimerColumns,
			emailTimerColumnsWithDefault,
			emailTimerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(emailTimerType, emailTimerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(emailTimerType, emailTimerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `email_timer` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `email_timer` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `email_timer` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, emailTimerPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into email_timer")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.DataID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == emailTimerMapping["DataID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.DataID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for email_timer")
	}

CacheNoHooks:
	if !cached {
		emailTimerInsertCacheMut.Lock()
		emailTimerInsertCache[key] = cache
		emailTimerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the EmailTimer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EmailTimer) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	emailTimerUpdateCacheMut.RLock()
	cache, cached := emailTimerUpdateCache[key]
	emailTimerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			emailTimerColumns,
			emailTimerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update email_timer, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `email_timer` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, emailTimerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(emailTimerType, emailTimerMapping, append(wl, emailTimerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update email_timer row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for email_timer")
	}

	if !cached {
		emailTimerUpdateCacheMut.Lock()
		emailTimerUpdateCache[key] = cache
		emailTimerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q emailTimerQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for email_timer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for email_timer")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EmailTimerSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailTimerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `email_timer` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailTimerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in emailTimer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all emailTimer")
	}
	return rowsAff, nil
}

var mySQLEmailTimerUniqueColumns = []string{
	"data_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EmailTimer) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no email_timer provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(emailTimerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEmailTimerUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	emailTimerUpsertCacheMut.RLock()
	cache, cached := emailTimerUpsertCache[key]
	emailTimerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			emailTimerColumns,
			emailTimerColumnsWithDefault,
			emailTimerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			emailTimerColumns,
			emailTimerPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert email_timer, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "email_timer", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `email_timer` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(emailTimerType, emailTimerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(emailTimerType, emailTimerMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for email_timer")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.DataID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == emailTimerMapping["data_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(emailTimerType, emailTimerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for email_timer")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for email_timer")
	}

CacheNoHooks:
	if !cached {
		emailTimerUpsertCacheMut.Lock()
		emailTimerUpsertCache[key] = cache
		emailTimerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single EmailTimer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EmailTimer) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EmailTimer provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), emailTimerPrimaryKeyMapping)
	sql := "DELETE FROM `email_timer` WHERE `data_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from email_timer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for email_timer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q emailTimerQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no emailTimerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from email_timer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for email_timer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EmailTimerSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EmailTimer slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailTimerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `email_timer` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailTimerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from emailTimer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for email_timer")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EmailTimer) Reload(exec boil.Executor) error {
	ret, err := FindEmailTimer(exec, o.DataID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EmailTimerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EmailTimerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailTimerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `email_timer`.* FROM `email_timer` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailTimerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EmailTimerSlice")
	}

	*o = slice

	return nil
}

// EmailTimerExists checks if the EmailTimer row exists.
func EmailTimerExists(exec boil.Executor, dataID uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `email_timer` where `data_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, dataID)
	}

	row := exec.QueryRow(sql, dataID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if email_timer exists")
	}

	return exists, nil
}
