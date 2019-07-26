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

	e, err := EscrowExists(ctx, tx, o.ID)
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

	escrowFound, err := FindEscrow(ctx, tx, o.ID)
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

func testEscrowToManyBundles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Escrow
	var b, c Bundle

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bundleDBTypes, false, bundleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bundleDBTypes, false, bundleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.EscrowID = a.ID
	c.EscrowID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Bundles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.EscrowID == b.EscrowID {
			bFound = true
		}
		if v.EscrowID == c.EscrowID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EscrowSlice{&a}
	if err = a.L.LoadBundles(ctx, tx, false, (*[]*Escrow)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Bundles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Bundles = nil
	if err = a.L.LoadBundles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Bundles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testEscrowToManyEscrowCaches(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Escrow
	var b, c EscrowCache

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, escrowDBTypes, true, escrowColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, escrowCacheDBTypes, false, escrowCacheColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, escrowCacheDBTypes, false, escrowCacheColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.EscrowID = a.ID
	c.EscrowID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.EscrowCaches().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.EscrowID == b.EscrowID {
			bFound = true
		}
		if v.EscrowID == c.EscrowID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EscrowSlice{&a}
	if err = a.L.LoadEscrowCaches(ctx, tx, false, (*[]*Escrow)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.EscrowCaches); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.EscrowCaches = nil
	if err = a.L.LoadEscrowCaches(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.EscrowCaches); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testEscrowToManyAddOpBundles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Escrow
	var b, c, d, e Bundle

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, escrowDBTypes, false, strmangle.SetComplement(escrowPrimaryKeyColumns, escrowColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Bundle{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bundleDBTypes, false, strmangle.SetComplement(bundlePrimaryKeyColumns, bundleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Bundle{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBundles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.EscrowID {
			t.Error("foreign key was wrong value", a.ID, first.EscrowID)
		}
		if a.ID != second.EscrowID {
			t.Error("foreign key was wrong value", a.ID, second.EscrowID)
		}

		if first.R.Escrow != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Escrow != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Bundles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Bundles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Bundles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testEscrowToManyAddOpEscrowCaches(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Escrow
	var b, c, d, e EscrowCache

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, escrowDBTypes, false, strmangle.SetComplement(escrowPrimaryKeyColumns, escrowColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*EscrowCache{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, escrowCacheDBTypes, false, strmangle.SetComplement(escrowCachePrimaryKeyColumns, escrowCacheColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*EscrowCache{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddEscrowCaches(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.EscrowID {
			t.Error("foreign key was wrong value", a.ID, first.EscrowID)
		}
		if a.ID != second.EscrowID {
			t.Error("foreign key was wrong value", a.ID, second.EscrowID)
		}

		if first.R.Escrow != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Escrow != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.EscrowCaches[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.EscrowCaches[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.EscrowCaches().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
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
	escrowDBTypes = map[string]string{`ID`: `integer`, `Txid`: `bytea`, `StartBlock`: `integer`, `EndBlock`: `integer`, `State`: `enum.escrow_state('ok','aborted')`, `PublicKey`: `bytea`, `PrivateKey`: `bytea`, `Raw`: `bytea`}
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

func testEscrowsUpsert(t *testing.T) {
	t.Parallel()

	if len(escrowAllColumns) == len(escrowPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Escrow{}
	if err = randomize.Struct(seed, &o, escrowDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Escrow: %s", err)
	}

	count, err := Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, escrowDBTypes, false, escrowPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Escrow struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Escrow: %s", err)
	}

	count, err = Escrows().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
