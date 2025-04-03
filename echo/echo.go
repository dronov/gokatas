// Echo prints its command-line arguments. It also shows how to test a command
// (package main) instead of the usual library packages.
//
// Adapted from github.com/adonovan/gopl.io/blob/master/ch11/echo
//
// Level: beginner
// Topics: command, testing
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout // modified during testing

func main() {
	flag.Parse()

	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1) // call os.Exit or log.Fatal only from main if testing
	}
}

func echo(newline bool, sep string, args []string) error {
	output := strings.Join(args, sep)

	if newline {
		output += "\n"
	}

	_, err := fmt.Fprint(out, output)
	return err
}
