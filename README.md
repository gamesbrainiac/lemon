# Lemon

Lemon is a simple package for building command line applications in Go.

## Installation

Install Lemon using the "go get" command:

```
$ go get github.com/hjr265/lemon
```

The only dependency is the Go distribution itself.

## Example

``` go
package main

import (
	"fmt"
	"strings"

	lmn "github.com/hjr265/lemon"
)

func handleEcho(c *lmn.Context) {
	fmt.Print(strings.Join(c.Args.Strings(0), " "))
	if !c.Flags.Bool("n") {
		fmt.Println()
	}
}

func main() {
	l := &lmn.Lemon{}
	l.NewAction("echo", "echos the string(s) to standard output").
		Flag("n", lmn.Bool, "suppress trailing newline").
		Arg(lmn.String, lmn.Repeat{true}).
		HandlerFunc(handleEcho)

	l.Run(nil)
}
```

```
$ go run main.go -h
Usage: main action [action-specific options]

  echo: echos the string(s) to standard output

```

```
$ go run main.go echo -h
Usage: main echo [flags] [arg0..]

  echo: echos the string(s) to standard output

  -n=false: suppress trailing newline

```

```
$ go run main.go echo Hello world!
Hello, world!
```

## Documentation

- [Reference](http://godoc.org/github.com/hjr265/lemon)

## Contributing

Contributions are welcome.

## License

Lemon is available under the [BSD (3-Clause) License](http://opensource.org/licenses/BSD-3-Clause)

## Inspiration

This project is inspired by the very existence of the awesome package [github.com/codegangsta/cli](https://github.com/codegangsta/cli) and intuitiveness of [github.com/gorilla/mux](https://github.com/gorilla/mux).
