// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEmailTimers(t *testing.T) {
	t.Parallel()

	query := EmailTimers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEmailTimersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailTimersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EmailTimers().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailTimersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailTimerSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailTimersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EmailTimerExists(tx, o.DataID)
	if err != nil {
		t.Errorf("Unable to check if EmailTimer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EmailTimerExists to return true, but got false.")
	}
}

func testEmailTimersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	emailTimerFound, err := FindEmailTimer(tx, o.DataID)
	if err != nil {
		t.Error(err)
	}

	if emailTimerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEmailTimersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EmailTimers().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testEmailTimersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EmailTimers().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEmailTimersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	emailTimerOne := &EmailTimer{}
	emailTimerTwo := &EmailTimer{}
	if err = randomize.Struct(seed, emailTimerOne, emailTimerDBTypes, false, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}
	if err = randomize.Struct(seed, emailTimerTwo, emailTimerDBTypes, false, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = emailTimerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailTimerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailTimers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEmailTimersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	emailTimerOne := &EmailTimer{}
	emailTimerTwo := &EmailTimer{}
	if err = randomize.Struct(seed, emailTimerOne, emailTimerDBTypes, false, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}
	if err = randomize.Struct(seed, emailTimerTwo, emailTimerDBTypes, false, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = emailTimerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailTimerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testEmailTimersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailTimersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(emailTimerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailTimersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testEmailTimersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailTimerSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testEmailTimersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailTimers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	emailTimerDBTypes = map[string]string{`Attaches`: `text`, `Content`: `text`, `CreatedAt`: `timestamp`, `DataID`: `bigint`, `EmailType`: `tinyint`, `EndTime`: `int`, `IsSend`: `tinyint`, `StartTime`: `int`, `Title`: `varchar`}
	_                 = bytes.MinRead
)

func testEmailTimersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(emailTimerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(emailTimerColumns) == len(emailTimerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEmailTimersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(emailTimerColumns) == len(emailTimerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailTimer{}
	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailTimerDBTypes, true, emailTimerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(emailTimerColumns, emailTimerPrimaryKeyColumns) {
		fields = emailTimerColumns
	} else {
		fields = strmangle.SetComplement(
			emailTimerColumns,
			emailTimerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := EmailTimerSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEmailTimersUpsert(t *testing.T) {
	t.Parallel()

	if len(emailTimerColumns) == len(emailTimerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLEmailTimerUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EmailTimer{}
	if err = randomize.Struct(seed, &o, emailTimerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailTimer: %s", err)
	}

	count, err := EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, emailTimerDBTypes, false, emailTimerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailTimer struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailTimer: %s", err)
	}

	count, err = EmailTimers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}