// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

type FlagSet map[string]interface{}

func (s FlagSet) Bool(name string) bool {
	return s.Bools(name)[0]
}

func (s FlagSet) Bools(name string) []bool {
	return []bool(*s[name].(*boolParser))
}

func (s FlagSet) Int(name string) int {
	return s.Ints(name)[0]
}

func (s FlagSet) Ints(name string) []int {
	return []int(*s[name].(*intParser))
}

func (s FlagSet) String(name string) string {
	return s.Strings(name)[0]
}

func (s FlagSet) Strings(name string) []string {
	return []string(*s[name].(*stringParser))
}
