// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

import (
	"fmt"
)

type Value struct {
	Name  string
	Alias string

	Usage string

	Parser  Parser
	Default Parser
	Repeat  bool

	checks []Check
}

type Option interface {
	Apply(*Value)
}

type Check func() error

type Alias struct {
	Alias string
}

func (o Alias) Apply(v *Value) {
	v.Alias = o.Alias
}

type Usage struct {
	Usage string
}

func (o Usage) Apply(v *Value) {
	v.Usage = o.Usage
}

type Repeat struct {
	Repeat bool
}

func (o Repeat) Apply(v *Value) {
	v.Repeat = o.Repeat
}

type ErrRepeatMin struct {
	Name string
	Min  int
	Got  int
}

func (e ErrRepeatMin) Error() string {
	return fmt.Sprintf("expected %s to appear at least %d times; got %d", e.Name, e.Min, e.Got)
}

type RepeatMin struct {
	RepeatMin int
}

func (o RepeatMin) Apply(v *Value) {
	v.checks = append(v.checks, func() error {
		if v.Parser.Count() < o.RepeatMin {
			return &ErrRepeatMin{v.Name, o.RepeatMin, v.Parser.Count()}
		}
		return nil
	})
}

type ErrRepeatMax struct {
	Name string
	Max  int
	Got  int
}

func (e ErrRepeatMax) Error() string {
	return fmt.Sprintf("expected %s to appear at least %d times; got %d", e.Name, e.Max, e.Got)
}

type RepeatMax struct {
	RepeatMax int
}

func (o RepeatMax) Apply(v *Value) {
	v.checks = append(v.checks, func() error {
		if v.Parser.Count() > o.RepeatMax {
			return &ErrRepeatMax{v.Name, o.RepeatMax, v.Parser.Count()}
		}
		return nil
	})
}

type Type func() Value

func Bool() Value {
	return Value{Parser: new(boolParser), Default: &boolParser{false}}
}

type BoolCheck struct {
	Check func(*Value, []bool) error
}

func (o BoolCheck) Apply(v *Value) {
	v.checks = append(v.checks, func() error {
		return o.Check(v, []bool(*v.Parser.(*boolParser)))
	})
}

type BoolDefault struct {
	Default bool
}

func (o BoolDefault) Apply(v *Value) {
	v.Default = &boolParser{o.Default}
}

func Int() Value {
	return Value{Parser: new(intParser), Default: &intParser{0}}
}

type IntDefault struct {
	Default int
}

func (o IntDefault) Apply(v *Value) {
	v.Default = &intParser{o.Default}
}

type IntCheck struct {
	Check func(*Value, []int) error
}

func (o IntCheck) Apply(v *Value) {
	v.checks = append(v.checks, func() error {
		return o.Check(v, []int(*v.Parser.(*intParser)))
	})
}

type ErrIntMin struct {
	Name string
	Min  int
	Got  int
}

func (e ErrIntMin) Error() string {
	return fmt.Sprintf("expected %s to be at least %d; got %d", e.Name, e.Min, e.Got)
}

type IntMin struct {
	Min int
}

func (o IntMin) Apply(v *Value) {
	IntCheck{func(v *Value, a []int) error {
		for _, x := range a {
			if x < o.Min {
				return &ErrIntMin{v.Name, o.Min, x}
			}
		}
		return nil
	}}.Apply(v)
}

type ErrIntMax struct {
	Name string
	Max  int
	Got  int
}

func (e ErrIntMax) Error() string {
	return fmt.Sprintf("expected %s to be at most %d; got %d", e.Name, e.Max, e.Got)
}

type IntMax struct {
	Max int
}

func (o IntMax) Apply(v *Value) {
	IntCheck{func(v *Value, a []int) error {
		for _, x := range a {
			if x > o.Max {
				return &ErrIntMax{v.Name, o.Max, x}
			}
		}
		return nil
	}}.Apply(v)
}

func String() Value {
	return Value{Parser: new(stringParser), Default: &stringParser{""}}
}

type StringDefault struct {
	Default string
}

func (o StringDefault) Apply(v *Value) {
	v.Default = &stringParser{o.Default}
}

type StringCheck struct {
	Check func(*Value, []string) error
}

func (o StringCheck) Apply(v *Value) {
	v.checks = append(v.checks, func() error {
		return o.Check(v, []string(*v.Parser.(*stringParser)))
	})
}

type ErrStringLenMin struct {
	Name string
	Min  int
	Got  int
}

func (e ErrStringLenMin) Error() string {
	return fmt.Sprintf("expected length of %s to be at least %d; got %d", e.Name, e.Min, e.Got)
}

type StringLenMin struct {
	LenMin int
}

func (o StringLenMin) Apply(v *Value) {
	StringCheck{func(v *Value, a []string) error {
		for _, x := range a {
			if len(x) < o.LenMin {
				return &ErrStringLenMin{v.Name, o.LenMin, len(x)}
			}
		}
		return nil
	}}.Apply(v)
}

type ErrStringLenMax struct {
	Name string
	Max  int
	Got  int
}

func (e ErrStringLenMax) Error() string {
	return fmt.Sprintf("expected length of %s to be at most %d; got %d", e.Name, e.Max, e.Got)
}

type StringLenMax struct {
	LenMax int
}

func (o StringLenMax) Apply(v *Value) {
	StringCheck{func(v *Value, a []string) error {
		for _, x := range a {
			if len(x) > o.LenMax {
				return &ErrStringLenMax{v.Name, o.LenMax, len(x)}
			}
		}
		return nil
	}}.Apply(v)
}
