// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

type Context struct {
	Flags FlagSet
	Args  ArgList

	Lemon  *Lemon
	Action *Action

	Error error
}
