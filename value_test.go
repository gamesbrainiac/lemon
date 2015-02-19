// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

import (
	"reflect"
	"testing"
)

func testOption(t *testing.T, u Value, o Option, v Value) {
	o.Apply(&u)
	if !reflect.DeepEqual(u, v) {
		t.Fatalf("Expected %#v, got %#v", v, u)
	}
}

func testCheck(t *testing.T, v Value, o Option, ex error) {
	o.Apply(&v)
	for _, fn := range v.checks {
		err := fn()
		if !reflect.DeepEqual(ex, err) {
			t.Fatalf("Expected %#v, got %#v", ex, err)
		}
	}
}

func TestAlias(t *testing.T) {
	testOption(t, Value{}, Alias{"gopher"}, Value{Alias: "gopher"})
}

func TestUsage(t *testing.T) {
	testOption(t, Value{}, Usage{"gopher"}, Value{Usage: "gopher"})
}

func TestRepeat(t *testing.T) {
	testOption(t, Value{}, Repeat{true}, Value{Repeat: true})
}

func TestRepeatMin(t *testing.T) {
	testCheck(t, Value{Parser: &boolParser{true, true}}, RepeatMin{2}, nil)
	testCheck(t, Value{Parser: &boolParser{true, true, true}}, RepeatMin{2}, nil)
	testCheck(t, Value{Parser: &boolParser{true}}, RepeatMin{2}, &ErrRepeatMin{"", 2, 1})
}

func TestRepeatMax(t *testing.T) {
	testCheck(t, Value{Parser: &boolParser{true, true}}, RepeatMax{2}, nil)
	testCheck(t, Value{Parser: &boolParser{true}}, RepeatMax{2}, nil)
	testCheck(t, Value{Parser: &boolParser{true, true, true}}, RepeatMax{2}, &ErrRepeatMax{"", 2, 3})
}

func TestBoolDefault(t *testing.T) {
	testOption(t, Value{}, BoolDefault{true}, Value{Default: &boolParser{true}})
}

func TestIntDefault(t *testing.T) {
	testOption(t, Value{}, IntDefault{2}, Value{Default: &intParser{2}})
}

func TestIntMin(t *testing.T) {
	testCheck(t, Value{Parser: &intParser{2, 3}}, IntMin{2}, nil)
	testCheck(t, Value{Parser: &intParser{2, 1}}, IntMin{2}, &ErrIntMin{"", 2, 1})
	testCheck(t, Value{Parser: &intParser{2}}, IntMin{2}, nil)
	testCheck(t, Value{Parser: &intParser{1}}, IntMin{2}, &ErrIntMin{"", 2, 1})
}

func TestIntMax(t *testing.T) {
	testCheck(t, Value{Parser: &intParser{2, 1}}, IntMax{2}, nil)
	testCheck(t, Value{Parser: &intParser{2, 3}}, IntMax{2}, &ErrIntMax{"", 2, 3})
	testCheck(t, Value{Parser: &intParser{2}}, IntMax{2}, nil)
	testCheck(t, Value{Parser: &intParser{3}}, IntMax{2}, &ErrIntMax{"", 2, 3})
}

func TestStringDefault(t *testing.T) {
	testOption(t, Value{}, StringDefault{"gopher"}, Value{Default: &stringParser{"gopher"}})
}

func TestStringLenMin(t *testing.T) {
	testCheck(t, Value{Parser: &stringParser{"go", "gopher"}}, StringLenMin{2}, nil)
	testCheck(t, Value{Parser: &stringParser{"go", "g"}}, StringLenMin{2}, &ErrStringLenMin{"", 2, 1})
	testCheck(t, Value{Parser: &stringParser{"go"}}, StringLenMin{2}, nil)
	testCheck(t, Value{Parser: &stringParser{"g"}}, StringLenMin{2}, &ErrStringLenMin{"", 2, 1})
}

func TestStringLenMax(t *testing.T) {
	testCheck(t, Value{Parser: &stringParser{"go", "g"}}, StringLenMax{2}, nil)
	testCheck(t, Value{Parser: &stringParser{"go", "gopher"}}, StringLenMax{2}, &ErrStringLenMax{"", 2, 6})
	testCheck(t, Value{Parser: &stringParser{"go"}}, StringLenMax{2}, nil)
	testCheck(t, Value{Parser: &stringParser{"gopher"}}, StringLenMax{2}, &ErrStringLenMax{"", 2, 6})
}
