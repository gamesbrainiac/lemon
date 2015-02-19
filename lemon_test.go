// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

import (
	"testing"
)

func TestNotFound(t *testing.T) {
	ok := false
	l := Lemon{}
	l.NewAction("foo", "").HandlerFunc(func(*Context) {})
	l.NotFoundHandler = HandlerFunc(func(c *Context) {
		if len(c.Flags) != 0 {
			t.Errorf("Expected 0 flags, got %d", len(c.Flags))
		}
		if len(c.Args) != 2 {
			t.Errorf("Expected 2 arguments, got %d", len(c.Args))
		}
		ok = true
	})

	l.Run([]string{"bar", "a", "b"})

	if !ok {
		t.Error("Didn't reach bar handler")
	}
}
