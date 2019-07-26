// +build sqlboiler_test

// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
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

func testTransactionAuditlogs(t *testing.T) {
	t.Parallel()

	query := TransactionAuditlogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTransactionAuditlogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionAuditlogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TransactionAuditlogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionAuditlogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactionAuditlogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionAuditlogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TransactionAuditlogExists(ctx, tx, o.Rowid)
	if err != nil {
		t.Errorf("Unable to check if TransactionAuditlog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TransactionAuditlogExists to return true, but got false.")
	}
}

func testTransactionAuditlogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	transactionAuditlogFound, err := FindTransactionAuditlog(ctx, tx, o.Rowid)
	if err != nil {
		t.Error(err)
	}

	if transactionAuditlogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTransactionAuditlogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TransactionAuditlogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTransactionAuditlogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TransactionAuditlogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTransactionAuditlogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionAuditlogOne := &TransactionAuditlog{}
	transactionAuditlogTwo := &TransactionAuditlog{}
	if err = randomize.Struct(seed, transactionAuditlogOne, transactionAuditlogDBTypes, false, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionAuditlogTwo, transactionAuditlogDBTypes, false, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactionAuditlogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactionAuditlogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TransactionAuditlogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTransactionAuditlogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	transactionAuditlogOne := &TransactionAuditlog{}
	transactionAuditlogTwo := &TransactionAuditlog{}
	if err = randomize.Struct(seed, transactionAuditlogOne, transactionAuditlogDBTypes, false, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionAuditlogTwo, transactionAuditlogDBTypes, false, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactionAuditlogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactionAuditlogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func transactionAuditlogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func transactionAuditlogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TransactionAuditlog) error {
	*o = TransactionAuditlog{}
	return nil
}

func testTransactionAuditlogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TransactionAuditlog{}
	o := &TransactionAuditlog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog object: %s", err)
	}

	AddTransactionAuditlogHook(boil.BeforeInsertHook, transactionAuditlogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogBeforeInsertHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.AfterInsertHook, transactionAuditlogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogAfterInsertHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.AfterSelectHook, transactionAuditlogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogAfterSelectHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.BeforeUpdateHook, transactionAuditlogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogBeforeUpdateHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.AfterUpdateHook, transactionAuditlogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogAfterUpdateHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.BeforeDeleteHook, transactionAuditlogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogBeforeDeleteHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.AfterDeleteHook, transactionAuditlogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogAfterDeleteHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.BeforeUpsertHook, transactionAuditlogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogBeforeUpsertHooks = []TransactionAuditlogHook{}

	AddTransactionAuditlogHook(boil.AfterUpsertHook, transactionAuditlogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	transactionAuditlogAfterUpsertHooks = []TransactionAuditlogHook{}
}

func testTransactionAuditlogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionAuditlogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(transactionAuditlogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionAuditlogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTransactionAuditlogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactionAuditlogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTransactionAuditlogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TransactionAuditlogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	transactionAuditlogDBTypes = map[string]string{`Rowid`: `integer`, `Raw`: `bytea`, `Status`: `enum.transaction_status('pending','mined','rejected')`}
	_                          = bytes.MinRead
)

func testTransactionAuditlogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(transactionAuditlogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(transactionAuditlogAllColumns) == len(transactionAuditlogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTransactionAuditlogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(transactionAuditlogAllColumns) == len(transactionAuditlogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TransactionAuditlog{}
	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactionAuditlogDBTypes, true, transactionAuditlogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(transactionAuditlogAllColumns, transactionAuditlogPrimaryKeyColumns) {
		fields = transactionAuditlogAllColumns
	} else {
		fields = strmangle.SetComplement(
			transactionAuditlogAllColumns,
			transactionAuditlogPrimaryKeyColumns,
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

	slice := TransactionAuditlogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTransactionAuditlogsUpsert(t *testing.T) {
	t.Parallel()

	if len(transactionAuditlogAllColumns) == len(transactionAuditlogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TransactionAuditlog{}
	if err = randomize.Struct(seed, &o, transactionAuditlogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TransactionAuditlog: %s", err)
	}

	count, err := TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, transactionAuditlogDBTypes, false, transactionAuditlogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TransactionAuditlog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TransactionAuditlog: %s", err)
	}

	count, err = TransactionAuditlogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
