// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon

type Handler interface {
	Run(*Context)
}

type HandlerFunc func(*Context)

func (h HandlerFunc) Run(c *Context) {
	h(c)
}
