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

func testEscrows(t *testing.T) {
	t.Parallel()

	query := Escrows()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEscrowsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
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

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEscrowsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Escrows().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEscrowsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EscrowSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEscrowsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EscrowExists(ctx, tx, o.Txid)
	if err != nil {
		t.Errorf("Unable to check if Escrow exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EscrowExists to return true, but got false.")
	}
}

func testEscrowsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	escrowFound, err := FindEscrow(ctx, tx, o.Txid)
	if err != nil {
		t.Error(err)
	}

	if escrowFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEscrowsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Escrows().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEscrowsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Escrows().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEscrowsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	escrowOne := &Escrow{}
	escrowTwo := &Escrow{}
	if err = randomize.Struct(seed, escrowOne, escrowDBTypes, false, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}
	if err = randomize.Struct(seed, escrowTwo, escrowDBTypes, false, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = escrowOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = escrowTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Escrows().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEscrowsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	escrowOne := &Escrow{}
	escrowTwo := &Escrow{}
	if err = randomize.Struct(seed, escrowOne, escrowDBTypes, false, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}
	if err = randomize.Struct(seed, escrowTwo, escrowDBTypes, false, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = escrowOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = escrowTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func escrowBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func escrowAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Escrow) error {
	*o = Escrow{}
	return nil
}

func testEscrowsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Escrow{}
	o := &Escrow{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, escrowDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Escrow object: %s", err)
	}

	AddEscrowHook(boil.BeforeInsertHook, escrowBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	escrowBeforeInsertHooks = []EscrowHook{}

	AddEscrowHook(boil.AfterInsertHook, escrowAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	escrowAfterInsertHooks = []EscrowHook{}

	AddEscrowHook(boil.AfterSelectHook, escrowAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	escrowAfterSelectHooks = []EscrowHook{}

	AddEscrowHook(boil.BeforeUpdateHook, escrowBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	escrowBeforeUpdateHooks = []EscrowHook{}

	AddEscrowHook(boil.AfterUpdateHook, escrowAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	escrowAfterUpdateHooks = []EscrowHook{}

	AddEscrowHook(boil.BeforeDeleteHook, escrowBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	escrowBeforeDeleteHooks = []EscrowHook{}

	AddEscrowHook(boil.AfterDeleteHook, escrowAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	escrowAfterDeleteHooks = []EscrowHook{}

	AddEscrowHook(boil.BeforeUpsertHook, escrowBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	escrowBeforeUpsertHooks = []EscrowHook{}

	AddEscrowHook(boil.AfterUpsertHook, escrowAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	escrowAfterUpsertHooks = []EscrowHook{}
}

func testEscrowsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEscrowsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(escrowColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEscrowsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
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

func testEscrowsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EscrowSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEscrowsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Escrows().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	escrowDBTypes = map[string]string{`Txid`: `BYTEA`, `InnerMasterKey`: `VARCHAR`, `OuterMasterKey`: `VARCHAR`, `Slots`: `INT`, `PublisherAddr`: `VARCHAR`}
	_             = bytes.MinRead
)

func testEscrowsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(escrowPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(escrowAllColumns) == len(escrowPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEscrowsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(escrowAllColumns) == len(escrowPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Escrow{}
	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, escrowDBTypes, true, escrowPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(escrowAllColumns, escrowPrimaryKeyColumns) {
		fields = escrowAllColumns
	} else {
		fields = strmangle.SetComplement(
			escrowAllColumns,
			escrowPrimaryKeyColumns,
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

	slice := EscrowSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
