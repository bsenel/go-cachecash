// +build sqlboiler_test

// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testMempoolTransactions(t *testing.T) {
	t.Parallel()

	query := MempoolTransactions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMempoolTransactionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
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

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMempoolTransactionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MempoolTransactions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMempoolTransactionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MempoolTransactionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMempoolTransactionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MempoolTransactionExists(ctx, tx, o.Rowid)
	if err != nil {
		t.Errorf("Unable to check if MempoolTransaction exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MempoolTransactionExists to return true, but got false.")
	}
}

func testMempoolTransactionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mempoolTransactionFound, err := FindMempoolTransaction(ctx, tx, o.Rowid)
	if err != nil {
		t.Error(err)
	}

	if mempoolTransactionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMempoolTransactionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MempoolTransactions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMempoolTransactionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MempoolTransactions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMempoolTransactionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mempoolTransactionOne := &MempoolTransaction{}
	mempoolTransactionTwo := &MempoolTransaction{}
	if err = randomize.Struct(seed, mempoolTransactionOne, mempoolTransactionDBTypes, false, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}
	if err = randomize.Struct(seed, mempoolTransactionTwo, mempoolTransactionDBTypes, false, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mempoolTransactionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mempoolTransactionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MempoolTransactions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMempoolTransactionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mempoolTransactionOne := &MempoolTransaction{}
	mempoolTransactionTwo := &MempoolTransaction{}
	if err = randomize.Struct(seed, mempoolTransactionOne, mempoolTransactionDBTypes, false, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}
	if err = randomize.Struct(seed, mempoolTransactionTwo, mempoolTransactionDBTypes, false, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mempoolTransactionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mempoolTransactionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mempoolTransactionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func mempoolTransactionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MempoolTransaction) error {
	*o = MempoolTransaction{}
	return nil
}

func testMempoolTransactionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MempoolTransaction{}
	o := &MempoolTransaction{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction object: %s", err)
	}

	AddMempoolTransactionHook(boil.BeforeInsertHook, mempoolTransactionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionBeforeInsertHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.AfterInsertHook, mempoolTransactionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionAfterInsertHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.AfterSelectHook, mempoolTransactionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionAfterSelectHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.BeforeUpdateHook, mempoolTransactionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionBeforeUpdateHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.AfterUpdateHook, mempoolTransactionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionAfterUpdateHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.BeforeDeleteHook, mempoolTransactionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionBeforeDeleteHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.AfterDeleteHook, mempoolTransactionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionAfterDeleteHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.BeforeUpsertHook, mempoolTransactionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionBeforeUpsertHooks = []MempoolTransactionHook{}

	AddMempoolTransactionHook(boil.AfterUpsertHook, mempoolTransactionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mempoolTransactionAfterUpsertHooks = []MempoolTransactionHook{}
}

func testMempoolTransactionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMempoolTransactionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mempoolTransactionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMempoolTransactionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
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

func testMempoolTransactionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MempoolTransactionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMempoolTransactionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MempoolTransactions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mempoolTransactionDBTypes = map[string]string{`Rowid`: `integer`, `Txid`: `bytea`, `Raw`: `bytea`}
	_                         = bytes.MinRead
)

func testMempoolTransactionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mempoolTransactionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mempoolTransactionAllColumns) == len(mempoolTransactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMempoolTransactionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mempoolTransactionAllColumns) == len(mempoolTransactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MempoolTransaction{}
	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mempoolTransactionDBTypes, true, mempoolTransactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mempoolTransactionAllColumns, mempoolTransactionPrimaryKeyColumns) {
		fields = mempoolTransactionAllColumns
	} else {
		fields = strmangle.SetComplement(
			mempoolTransactionAllColumns,
			mempoolTransactionPrimaryKeyColumns,
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

	slice := MempoolTransactionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMempoolTransactionsUpsert(t *testing.T) {
	t.Parallel()

	if len(mempoolTransactionAllColumns) == len(mempoolTransactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MempoolTransaction{}
	if err = randomize.Struct(seed, &o, mempoolTransactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MempoolTransaction: %s", err)
	}

	count, err := MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mempoolTransactionDBTypes, false, mempoolTransactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MempoolTransaction struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MempoolTransaction: %s", err)
	}

	count, err = MempoolTransactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
