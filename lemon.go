// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

var ErrNotFound = errors.New("action not found")

type Lemon struct {
	NotFoundHandler Handler

	actions map[string]*Action

	output io.Writer
}

func (l *Lemon) NewAction(name, hint string) *Action {
	a := &Action{
		name:  name,
		hint:  hint,
		lemon: l,
	}
	if l.actions == nil {
		l.actions = map[string]*Action{}
	}
	l.actions[name] = a
	return a
}

func (l *Lemon) Run(args []string) error {
	if args == nil {
		args = os.Args[1:]
	}
	if len(args) == 0 {
		args = []string{""}
	}

	a, ok := l.actions[args[0]]
	if !ok {
		if l.NotFoundHandler != nil {
			c := &Context{}
			c.Flags = FlagSet{}
			c.Args = ArgList{}
			c.Lemon = l
			for _, arg := range args[1:] {
				c.Args = append(c.Args, &stringParser{arg})
			}
			l.NotFoundHandler.Run(c)

		} else {
			l.PrintDefaults()
		}

		return ErrNotFound
	}

	return a.Run(args[1:])
}

// VisitAll visits the actions in lexicographical order of names, calling fn for each.
func (l *Lemon) VisitAll(fn func(*Action)) {
	names := []string{}
	for name, a := range l.actions {
		if name == a.name {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	for _, name := range names {
		fn(l.actions[name])
	}
}

// PrintDefaults prints to standard error the list of available actions.
func (l *Lemon) PrintDefaults() {
	w := l.Output()
	fmt.Fprintf(w, "Usage: %s action [action-specific options]\n\n", filepath.Base(os.Args[0]))
	l.VisitAll(func(a *Action) {
		fmt.Fprintf(w, "  %s: %s\n", a.name, a.hint)
	})
	fmt.Fprint(w, "\n")
}

// Output returns the destination for usage and error messages.
func (l *Lemon) Output() io.Writer {
	if l.output == nil {
		return os.Stderr
	}
	return l.output
}

// SetOutput sets the destination for usage and error messages. If output is nil, os.Stderr is used.
func (l *Lemon) SetOutput(output io.Writer) {
	l.output = output
}
