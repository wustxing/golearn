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

func testEmailSends(t *testing.T) {
	t.Parallel()

	query := EmailSends()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEmailSendsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
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

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailSendsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EmailSends().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailSendsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailSendSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmailSendsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EmailSendExists(tx, o.DataID)
	if err != nil {
		t.Errorf("Unable to check if EmailSend exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EmailSendExists to return true, but got false.")
	}
}

func testEmailSendsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	emailSendFound, err := FindEmailSend(tx, o.DataID)
	if err != nil {
		t.Error(err)
	}

	if emailSendFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEmailSendsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EmailSends().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testEmailSendsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EmailSends().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEmailSendsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	emailSendOne := &EmailSend{}
	emailSendTwo := &EmailSend{}
	if err = randomize.Struct(seed, emailSendOne, emailSendDBTypes, false, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}
	if err = randomize.Struct(seed, emailSendTwo, emailSendDBTypes, false, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = emailSendOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailSendTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailSends().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEmailSendsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	emailSendOne := &EmailSend{}
	emailSendTwo := &EmailSend{}
	if err = randomize.Struct(seed, emailSendOne, emailSendDBTypes, false, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}
	if err = randomize.Struct(seed, emailSendTwo, emailSendDBTypes, false, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = emailSendOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = emailSendTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testEmailSendsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailSendsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(emailSendColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmailSendsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
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

func testEmailSendsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmailSendSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testEmailSendsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmailSends().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	emailSendDBTypes = map[string]string{`DataID`: `bigint`, `EmailID`: `bigint`, `EmailType`: `tinyint`, `GetTime`: `int`, `IsAttachReceive`: `tinyint`, `IsDel`: `tinyint`, `IsRead`: `tinyint`, `UpdatedAt`: `timestamp`, `UserID`: `bigint`}
	_                = bytes.MinRead
)

func testEmailSendsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(emailSendPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(emailSendColumns) == len(emailSendPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEmailSendsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(emailSendColumns) == len(emailSendPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmailSend{}
	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, emailSendDBTypes, true, emailSendPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(emailSendColumns, emailSendPrimaryKeyColumns) {
		fields = emailSendColumns
	} else {
		fields = strmangle.SetComplement(
			emailSendColumns,
			emailSendPrimaryKeyColumns,
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

	slice := EmailSendSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEmailSendsUpsert(t *testing.T) {
	t.Parallel()

	if len(emailSendColumns) == len(emailSendPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLEmailSendUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EmailSend{}
	if err = randomize.Struct(seed, &o, emailSendDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailSend: %s", err)
	}

	count, err := EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, emailSendDBTypes, false, emailSendPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmailSend struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmailSend: %s", err)
	}

	count, err = EmailSends().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
