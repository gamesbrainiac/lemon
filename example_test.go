// Copyright 2014 Mahmud Ridwan. All rights reserved.

package lemon_test

import (
	"fmt"
	"strings"

	lmn "github.com/hjr265/lemon"
)

func handleEcho(c *lmn.Context) {
	str := strings.Join(c.Args.Strings(0), c.Flags.String("t"))
	if c.Flags.Bool("r") {
		t := []rune(str)
		for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
			t[i], t[j] = t[j], t[i]
		}
		str = string(t)
	}
	if c.Flags.Bool("e") {
		str += "\n"
	}

	for i := 0; i < c.Flags.Int("n"); i++ {
		fmt.Print(str)
		fmt.Print(c.Flags.String("s"))
	}
}

func Example() {
	l := &lmn.Lemon{}
	l.NewAction("echo", "echos the string(s) to standard output").
		Flag("n", lmn.Int, lmn.IntMin{0}, lmn.IntMax{8}, lmn.IntDefault{1}, lmn.Usage{"repeat"}).
		Flag("r", lmn.Bool, lmn.Usage{"reverse"}).
		Flag("e", lmn.Bool, lmn.BoolDefault{true}, lmn.Usage{"end line"}).
		Flag("t", lmn.String, lmn.StringDefault{" "}, lmn.Usage{"join token"}).
		Flag("s", lmn.String, lmn.StringDefault{""}, lmn.Usage{"suffix token"}).
		Arg(lmn.String, lmn.Alias{"word"}, lmn.Repeat{true}).
		HandlerFunc(handleEcho)

	l.Run([]string{"echo", "-n=2", "-e=false", "-t=_", "-s", " ", "Hello", "world!"})

	// Output: Hello_world! Hello_world!
}
