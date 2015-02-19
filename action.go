// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Action struct {
	name string
	hint string

	flags map[string]Value
	args  []Value

	handler Handler

	lemon *Lemon
}

// GetName returns the name for the action.
func (a *Action) GetName() string {
	return a.name
}

// GetHint returns the hint for the action.
func (a *Action) GetHint() string {
	return a.hint
}

// Alias defines an alias for the action.
func (a *Action) Alias(name string) *Action {
	a.lemon.actions[name] = a
	return a
}

// Flag defines a flag using the given name and value-type.
func (a *Action) Flag(name string, t Type, options ...Option) *Action {
	if a.flags == nil {
		a.flags = map[string]Value{}
	}

	f := t()
	for _, option := range options {
		option.Apply(&f)
	}
	a.flags[name] = f
	return a
}

// Arg defines a positional argument using the given value-type.
func (a *Action) Arg(t Type, options ...Option) *Action {
	f := t()
	for _, option := range options {
		option.Apply(&f)
	}
	a.args = append(a.args, f)
	return a
}

// Handler sets a handler for the action.
func (a *Action) Handler(h Handler) {
	a.handler = h
}

// Handler sets a handler function for the action.
func (a *Action) HandlerFunc(fn func(*Context)) {
	a.Handler(HandlerFunc(fn))
}

// Parse parses args to generate a context.
func (a *Action) Parse(args []string) (*Context, error) {
	c := &Context{}

	c.Flags = FlagSet{}
	for name, f := range a.flags {
		c.Flags[name] = f.Default
	}
	c.Args = ArgList{}
	for _, f := range a.args {
		c.Args = append(c.Args, f.Default)
	}

	i := 0
	for i < len(args) {
		arg := args[i]
		if len(arg) == 0 || arg[0] != '-' || len(arg) == 1 {
			break
		}
		i++
		arg = arg[1:]

		if arg[0] == '-' {
			if len(arg) == 1 {
				break

			} else {
				arg = arg[1:]
			}
		}

		w := strings.SplitN(arg, "=", 2)
		f, ok := a.flags[w[0]]
		if !ok {
			return nil, errors.New("unrecognized flag")
		}

		_, isBool := f.Parser.(*boolParser)
		if isBool {
			if w[0] == "h" || w[0] == "help" {
				return nil, errors.New("help requested")
			}

			if len(w) == 1 {
				w = append(w, "true")
			}
			err := f.Parser.Parse(w[1])
			if err != nil {
				return nil, err
			}

		} else {
			if len(w) == 1 && len(args) > i {
				w = append(w, args[i])
				i++
			}
			if len(w) == 1 {
				return nil, errors.New("")
			}
			err := f.Parser.Parse(w[1])
			if err != nil {
				return nil, err
			}
		}

		for _, check := range f.checks {
			err := check()
			if err != nil {
				return nil, err
			}
		}

		c.Flags[w[0]] = f.Parser
	}
	for j := 0; j < len(a.args); j++ {
		f := a.args[j]
		if f.Repeat {
			if i == len(args) {
				continue
			}

			for ; i < len(args)-(len(a.args)-j-1); i++ {
				err := f.Parser.Parse(args[i])
				if err != nil {
					return nil, err
				}
			}

		} else {
			if i == len(args) {
				return nil, errors.New("insufficient arguments")
			}

			err := f.Parser.Parse(args[i])
			i++
			if err != nil {
				return nil, err
			}
		}

		for _, check := range f.checks {
			err := check()
			if err != nil {
				return nil, err
			}
		}

		c.Args[j] = f.Parser
	}

	return c, nil
}

// Run runs action a with the given context c.
func (a *Action) Run(args []string) error {
	c, err := a.Parse(args)
	if err != nil {
		a.PrintDefaults()
		return err
	}
	c.Lemon = a.lemon
	c.Action = a

	a.handler.Run(c)
	return nil
}

// PrintDefaults prints to standard error the default values of all flags relevant to this action.
func (a *Action) PrintDefaults() {
	w := a.lemon.Output()
	fmt.Fprintf(w, "Usage: %s %s", filepath.Base(os.Args[0]), a.name)
	if len(a.flags) > 0 {
		fmt.Fprint(w, " [flags]")
	}
	for i, f := range a.args {
		a := f.Alias
		if a == "" {
			a = fmt.Sprintf("arg%d", i)
		}
		if f.Repeat {
			fmt.Fprintf(w, " [%s..]", a)
		} else {
			fmt.Fprintf(w, " %s", a)
		}
	}
	fmt.Fprint(w, "\n\n")
	fmt.Fprintf(w, "  %s: %s\n\n", a.name, a.hint)
	if len(a.flags) > 0 {
		names := []string{}
		for name := range a.flags {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			f := a.flags[name]
			fmt.Fprintf(w, "  -%s=%s: %s\n", name, f.Default, f.Usage)
		}
		fmt.Fprint(w, "\n")
	}
}
