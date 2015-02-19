// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

import (
	"strconv"
	"strings"
)

type Parser interface {
	Parse(string) error
	Clear()
	Count() int

	String() string
}

type boolParser []bool

func (b *boolParser) Parse(s string) error {
	t, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*b = append(*b, t)
	return nil
}

func (b *boolParser) Clear() {
	*b = []bool{}
}

func (b boolParser) Count() int {
	return len(b)
}

func (b boolParser) String() string {
	w := []byte{}
	for i, t := range b {
		if i > 0 {
			w = append(w, ' ')
		}
		w = strconv.AppendBool(w, t)
	}
	return string(w)
}

type intParser []int

func (b *intParser) Parse(s string) error {
	t, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*b = append(*b, t)
	return nil
}

func (b *intParser) Clear() {
	*b = []int{}
}

func (b intParser) Count() int {
	return len(b)
}

func (b intParser) String() string {
	w := ""
	for i, t := range b {
		if i > 0 {
			w += " "
		}
		w += strconv.Itoa(t)
	}
	return w
}

type stringParser []string

func (b *stringParser) Parse(s string) error {
	*b = append(*b, s)
	return nil
}

func (b *stringParser) Clear() {
	*b = []string{}
}

func (b stringParser) Count() int {
	return len(b)
}

func (b stringParser) String() string {
	return strings.Join([]string(b), " ")
}
