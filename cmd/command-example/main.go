package main

import (
	"fmt"
	"os"

	"github.com/dmulholl/argo"
)

func cmdBooHandler(cmdName string, cmdParser *argo.ArgParser) error {
	fmt.Println("------------- boo! -------------")
	fmt.Println(cmdParser)
	fmt.Println("--------------------------------")
	return nil
}

func main() {
	// Instantiate an ArgParser instance.
	parser := argo.NewParser()
	parser.Helptext = "Usage: example..."
	parser.Version = "1.2.3"

	// Register a command, "boo".
	cmdParser := parser.NewCommand("boo")
	cmdParser.Helptext = "Usage: example boo..."
	cmdParser.Callback = cmdBooHandler

	// The command can have its own flags and options.
	cmdParser.NewFlag("foo f")
	cmdParser.NewStringOption("bar b", "fallback")

	// Parse the command line arguments.
	if err := parser.ParseOsArgs(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(parser)
}
