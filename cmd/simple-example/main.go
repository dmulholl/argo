package main

import (
	"fmt"
	"os"

	"github.com/dmulholl/argo"
)

func main() {
	// Instantiate an ArgParser instance.
	parser := argo.NewParser()
	parser.Helptext = "Usage: example..."
	parser.Version = "1.2.3"

	// Register a flag and a string-valued option.
	parser.NewFlag("foo f")
	parser.NewStringOption("bar b", "fallback")

	// Parse the command line arguments.
	if err := parser.ParseOsArgs(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(parser)
}
