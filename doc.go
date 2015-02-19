// Package lemon is a simple package for building command line applications in Go.
//
//   func main() {
//     l := &lmn.Lemon{}
//     l.NewAction("hello", "prints \"Hello, world!\"").
//       HandlerFunc(func(c *lmn.Context) { fmt.Printf("Hello, world!\n") })
//
//     l.Run(nil)
//   }
package lemon
