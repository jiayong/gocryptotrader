// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package postgres

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/randomize"
	"github.com/thrasher-corp/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testScriptExecutions(t *testing.T) {
	t.Parallel()

	query := ScriptExecutions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testScriptExecutionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
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

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScriptExecutionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ScriptExecutions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScriptExecutionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScriptExecutionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScriptExecutionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ScriptExecutionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ScriptExecution exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ScriptExecutionExists to return true, but got false.")
	}
}

func testScriptExecutionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	scriptExecutionFound, err := FindScriptExecution(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if scriptExecutionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testScriptExecutionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ScriptExecutions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testScriptExecutionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ScriptExecutions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testScriptExecutionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	scriptExecutionOne := &ScriptExecution{}
	scriptExecutionTwo := &ScriptExecution{}
	if err = randomize.Struct(seed, scriptExecutionOne, scriptExecutionDBTypes, false, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}
	if err = randomize.Struct(seed, scriptExecutionTwo, scriptExecutionDBTypes, false, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scriptExecutionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scriptExecutionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ScriptExecutions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testScriptExecutionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	scriptExecutionOne := &ScriptExecution{}
	scriptExecutionTwo := &ScriptExecution{}
	if err = randomize.Struct(seed, scriptExecutionOne, scriptExecutionDBTypes, false, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}
	if err = randomize.Struct(seed, scriptExecutionTwo, scriptExecutionDBTypes, false, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scriptExecutionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scriptExecutionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func scriptExecutionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func scriptExecutionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ScriptExecution) error {
	*o = ScriptExecution{}
	return nil
}

func testScriptExecutionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ScriptExecution{}
	o := &ScriptExecution{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ScriptExecution object: %s", err)
	}

	AddScriptExecutionHook(boil.BeforeInsertHook, scriptExecutionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	scriptExecutionBeforeInsertHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.AfterInsertHook, scriptExecutionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	scriptExecutionAfterInsertHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.AfterSelectHook, scriptExecutionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	scriptExecutionAfterSelectHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.BeforeUpdateHook, scriptExecutionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	scriptExecutionBeforeUpdateHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.AfterUpdateHook, scriptExecutionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	scriptExecutionAfterUpdateHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.BeforeDeleteHook, scriptExecutionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	scriptExecutionBeforeDeleteHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.AfterDeleteHook, scriptExecutionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	scriptExecutionAfterDeleteHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.BeforeUpsertHook, scriptExecutionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	scriptExecutionBeforeUpsertHooks = []ScriptExecutionHook{}

	AddScriptExecutionHook(boil.AfterUpsertHook, scriptExecutionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	scriptExecutionAfterUpsertHooks = []ScriptExecutionHook{}
}

func testScriptExecutionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testScriptExecutionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(scriptExecutionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testScriptExecutionToOneScriptUsingScript(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ScriptExecution
	var foreign Script

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, scriptDBTypes, false, scriptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Script struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ScriptID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Script().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ScriptExecutionSlice{&local}
	if err = local.L.LoadScript(ctx, tx, false, (*[]*ScriptExecution)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Script == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Script = nil
	if err = local.L.LoadScript(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Script == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testScriptExecutionToOneSetOpScriptUsingScript(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ScriptExecution
	var b, c Script

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, scriptExecutionDBTypes, false, strmangle.SetComplement(scriptExecutionPrimaryKeyColumns, scriptExecutionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, scriptDBTypes, false, strmangle.SetComplement(scriptPrimaryKeyColumns, scriptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, scriptDBTypes, false, strmangle.SetComplement(scriptPrimaryKeyColumns, scriptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Script{&b, &c} {
		err = a.SetScript(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Script != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ScriptExecutions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ScriptID, x.ID) {
			t.Error("foreign key was wrong value", a.ScriptID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ScriptID))
		reflect.Indirect(reflect.ValueOf(&a.ScriptID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ScriptID, x.ID) {
			t.Error("foreign key was wrong value", a.ScriptID, x.ID)
		}
	}
}

func testScriptExecutionToOneRemoveOpScriptUsingScript(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ScriptExecution
	var b Script

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, scriptExecutionDBTypes, false, strmangle.SetComplement(scriptExecutionPrimaryKeyColumns, scriptExecutionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, scriptDBTypes, false, strmangle.SetComplement(scriptPrimaryKeyColumns, scriptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetScript(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveScript(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Script().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Script != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ScriptID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ScriptExecutions) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testScriptExecutionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
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

func testScriptExecutionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScriptExecutionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testScriptExecutionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ScriptExecutions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	scriptExecutionDBTypes = map[string]string{`ID`: `uuid`, `ScriptID`: `uuid`, `ExecutionType`: `character varying`, `ExecutionStatus`: `character varying`, `ExecutionTime`: `timestamp without time zone`}
	_                      = bytes.MinRead
)

func testScriptExecutionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(scriptExecutionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(scriptExecutionAllColumns) == len(scriptExecutionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testScriptExecutionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(scriptExecutionAllColumns) == len(scriptExecutionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ScriptExecution{}
	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scriptExecutionDBTypes, true, scriptExecutionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(scriptExecutionAllColumns, scriptExecutionPrimaryKeyColumns) {
		fields = scriptExecutionAllColumns
	} else {
		fields = strmangle.SetComplement(
			scriptExecutionAllColumns,
			scriptExecutionPrimaryKeyColumns,
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

	slice := ScriptExecutionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testScriptExecutionsUpsert(t *testing.T) {
	t.Parallel()

	if len(scriptExecutionAllColumns) == len(scriptExecutionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ScriptExecution{}
	if err = randomize.Struct(seed, &o, scriptExecutionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ScriptExecution: %s", err)
	}

	count, err := ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, scriptExecutionDBTypes, false, scriptExecutionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ScriptExecution struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ScriptExecution: %s", err)
	}

	count, err = ScriptExecutions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
