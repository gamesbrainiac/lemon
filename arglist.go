// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

type ArgList []interface{}

func (l ArgList) Bool(i int) bool {
	return l.Bools(i)[0]
}

func (l ArgList) Bools(i int) []bool {
	return []bool(*l[i].(*boolParser))
}

func (l ArgList) Int(i int) int {
	return l.Ints(i)[0]
}

func (l ArgList) Ints(i int) []int {
	return []int(*l[i].(*intParser))
}

func (l ArgList) String(i int) string {
	return l.Strings(i)[0]
}

func (l ArgList) Strings(i int) []string {
	return []string(*l[i].(*stringParser))
}
