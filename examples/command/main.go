package main

import (
	"fmt"
	"github.com/dmulholl/argo"
)

func main() {
	// Instantiate an ArgParser instance.
	parser := argo.NewParser()
	parser.Helptext = "Usage: example..."
	parser.Version = "1.2.3"

	// Register a command, "boo".
	cmdParser := parser.NewCommand("boo")
	cmdParser.Helptext = "Usage: boo..."
	cmdParser.Callback = boo

	// The command can have its own flags and options.
	cmdParser.NewFlag("foo f")
	cmdParser.NewStringOption("bar b", "fallback")

	// Parse the command line arguments.
	parser.Parse()
	fmt.Println(parser)
}

func boo(cmdName string, cmdParser *argo.ArgParser) {
	fmt.Println("------------- boo! -------------")
	fmt.Println(cmdParser)
	fmt.Println("--------------------------------")
}
