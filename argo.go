// Package argo is a library for parsing command line arguments.
package argo

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Prints a message to stderr and exits with a non-zero error code.
func exit(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %v.\n", msg)
	os.Exit(1)
}

/* --------- */
/*  Options  */
/* --------- */

// [kind] is one of "flag", "string", "int", or "float".
type option struct {
	kind           string
	count          int
	stringValues   []string
	intValues      []int64
	floatValues    []float64
	stringFallback string
	intFallback    int64
	floatFallback  float64
}

// This method attempts to set the option's value by parsing a string argument.
func (opt *option) trySetValue(arg string) {
	switch opt.kind {
	case "string":
		opt.stringValues = append(opt.stringValues, arg)

	case "int":
		value, err := strconv.ParseInt(arg, 0, 64)
		if err != nil {
			exit(fmt.Sprintf("cannot parse '%v' as an integer", arg))
		}
		opt.intValues = append(opt.intValues, value)

	case "float":
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			exit(fmt.Sprintf("cannot parse '%v' as a float", arg))
		}
		opt.floatValues = append(opt.floatValues, value)
	}
}

/* ----------- */
/*  ArgStream  */
/* ----------- */

// Makes a slice of string arguments available as a stream.
type argstream struct {
	args   []string
	index  int
	length int
}

// Initialize a new argstream instance.
func newArgStream(args []string) *argstream {
	return &argstream{
		args:   args,
		index:  0,
		length: len(args),
	}
}

// Returns the next argument from the stream.
func (stream *argstream) next() string {
	stream.index += 1
	return stream.args[stream.index-1]
}

// Returns true if the stream contains at least one more element.
func (stream *argstream) hasNext() bool {
	return stream.index < stream.length
}

/* ----------- */
/*  ArgParser  */
/* ----------- */

// An ArgParser instance stores registered options and commands.
type ArgParser struct {
	// Help text for the application or command.
	Helptext string

	// The application's version number.
	Version string

	// Stores option instances indexed by option name.
	options map[string]*option

	// Stores command parsers indexed by command name.
	commands map[string]*ArgParser

	// Stores positional arguments parsed from the input array.
	arguments []string

	// Stores the callback function for a command parser.
	Callback func(string, *ArgParser)

	// Stores the command name, if a command was found while parsing.
	command string

	// If true, enables an automatic 'help' command to print subcommand helptext.
	enableHelpCommand bool
}

// NewParser initializes a new ArgParser instance.
func NewParser() *ArgParser {
	return &ArgParser{
		options:   make(map[string]*option),
		commands:  make(map[string]*ArgParser),
		arguments: make([]string, 0),
	}
}

/* ------------------------------ */
/*  ArgParser: register options.  */
/* ------------------------------ */

// NewFlag registers a new flag. The `name` parameter accepts an unlimited number of
// space-separated aliases and single-character shortcuts.
func (parser *ArgParser) NewFlag(name string) {
	opt := &option{}
	opt.kind = "flag"
	for _, alias := range strings.Split(name, " ") {
		parser.options[alias] = opt
	}
}

// NewStringOption registers a new string-valued option. The `name` parameter accepts an unlimited
// number of space-separated aliases and single-character shortcuts. The `fallback` parameter
// specifies the option's default value.
func (parser *ArgParser) NewStringOption(name string, fallback string) {
	opt := &option{}
	opt.kind = "string"
	opt.stringFallback = fallback
	for _, alias := range strings.Split(name, " ") {
		parser.options[alias] = opt
	}
}

// NewIntOption registers a new integer-valued option. The `name` parameter accepts an unlimited
// number of space-separated aliases and single-character shortcuts. The `fallback` parameter
// specifies the option's default value.
func (parser *ArgParser) NewIntOption(name string, fallback int64) {
	opt := &option{}
	opt.kind = "int"
	opt.intFallback = fallback
	for _, alias := range strings.Split(name, " ") {
		parser.options[alias] = opt
	}
}

// NewFloatOption registers a new float-valued option. The `name` parameter accepts an unlimited
// number of space-separated aliases and single-character shortcuts. The `fallback` parameter
// specifies the option's default value.
func (parser *ArgParser) NewFloatOption(name string, fallback float64) {
	opt := &option{}
	opt.kind = "float"
	opt.floatFallback = fallback
	for _, alias := range strings.Split(name, " ") {
		parser.options[alias] = opt
	}
}

/* ------------------------------------ */
/*  ArgParser: retrieve option values.  */
/* ------------------------------------ */

func (parser *ArgParser) getOpt(name string) *option {
	if opt, found := parser.options[name]; found {
		return opt
	}
	panic(fmt.Sprintf("argo: invalid option name '%s'", name))
}

// Count returns the number of times the specified option was found.
func (parser *ArgParser) Count(name string) int {
	return parser.getOpt(name).count
}

// Found returns true if the specified option was found.
func (parser *ArgParser) Found(name string) bool {
	return parser.getOpt(name).count > 0
}

// StringValue returns the value of the specified string-valued option.
func (parser *ArgParser) StringValue(name string) string {
	opt := parser.getOpt(name)
	if len(opt.stringValues) > 0 {
		return opt.stringValues[len(opt.stringValues)-1]
	} else {
		return opt.stringFallback
	}
}

// IntValue returns the value of the specified integer-valued option.
func (parser *ArgParser) IntValue(name string) int64 {
	opt := parser.getOpt(name)
	if len(opt.intValues) > 0 {
		return opt.intValues[len(opt.intValues)-1]
	} else {
		return opt.intFallback
	}
}

// FloatValue returns the value of the specified float-valued option.
func (parser *ArgParser) FloatValue(name string) float64 {
	opt := parser.getOpt(name)
	if len(opt.floatValues) > 0 {
		return opt.floatValues[len(opt.floatValues)-1]
	} else {
		return opt.floatFallback
	}
}

// StringValues returns the specified string-valued option's list of values.
func (parser *ArgParser) StringValues(name string) []string {
	return parser.getOpt(name).stringValues
}

// IntValues returns the specified integer-valued option's list of values.
func (parser *ArgParser) IntValues(name string) []int64 {
	return parser.getOpt(name).intValues
}

// FloatValues returns the specified float-valued option's list of values.
func (parser *ArgParser) FloatValues(name string) []float64 {
	return parser.getOpt(name).floatValues
}

/* ---------------------------------- */
/*  ArgParser: positional arguments.  */
/* ---------------------------------- */

// HasArgs returns true if the parser has found one or more positional arguments.
func (parser *ArgParser) HasArgs() bool {
	return len(parser.arguments) > 0
}

// CountArgs returns the number of positional arguments.
func (parser *ArgParser) CountArgs() int {
	return len(parser.arguments)
}

// Arg returns the positional argument at the specified index.
func (parser *ArgParser) Arg(index int) string {
	return parser.arguments[index]
}

// Args returns the positional arguments as a slice of strings.
func (parser *ArgParser) Args() []string {
	return parser.arguments
}

// ArgsAsInts attempts to parse and return the positional arguments as a slice of integers. Exits
// with an error message if any of the arguments cannot be parsed as an integer.
func (parser *ArgParser) ArgsAsInts() []int64 {
	values := make([]int64, 0)
	for _, arg := range parser.arguments {
		value, err := strconv.ParseInt(arg, 0, 64)
		if err != nil {
			exit(fmt.Sprintf("cannot parse '%v' as an integer", arg))
		}
		values = append(values, value)
	}
	return values
}

// ArgsAsFloats attempts to parse and return the positional arguments as a slice of floats. Exits
// with an error message if any of the arguments cannot be parsed as a float.
func (parser *ArgParser) ArgsAsFloats() []float64 {
	values := make([]float64, 0)
	for _, arg := range parser.arguments {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			exit(fmt.Sprintf("cannot parse '%v' as a float", arg))
		}
		values = append(values, value)
	}
	return values
}

/* ---------------------- */
/*  ArgParser: commands.  */
/* ---------------------- */

// NewCommand registers a new command. The `name` parameter accepts an unlimited number of space-
// separated aliases for the command. Returns the command's `ArgParser` instance.
func (parser *ArgParser) NewCommand(name string) *ArgParser {
	parser.enableHelpCommand = true
	cmdParser := NewParser()
	for _, alias := range strings.Split(name, " ") {
		parser.commands[alias] = cmdParser
	}
	return cmdParser
}

// HasCommand returns true if the parser has found a command.
func (parser *ArgParser) HasCommand() bool {
	return parser.command != ""
}

// CommandName returns the command name, if the parser has found a command.
func (parser *ArgParser) CommandName() string {
	return parser.command
}

// CommandParser returns the command's parser instance, if the parser has found a command.
func (parser *ArgParser) CommandParser() *ArgParser {
	return parser.commands[parser.command]
}

// This boolean switch toggles support for an automatic `help` command which prints subcommand
// helptext. The value defaults to `false` but gets toggled automatically to true whenever a command
// is registered. You can use this method to disable the feature if required.
func (parser *ArgParser) EnableHelpCommand(enable bool) {
	parser.enableHelpCommand = enable
}

/* ----------------------------- */
/*  ArgParser: parse arguments.  */
/* ----------------------------- */

// Parse a stream of string arguments.
func (parser *ArgParser) parseStream(stream *argstream) {
	isFirstArg := true

	// Loop while we have arguments to process.
	for stream.hasNext() {
		arg := stream.next()

		// If we encounter a -- argument, turn off option-parsing.
		if arg == "--" {
			for stream.hasNext() {
				parser.arguments = append(parser.arguments, stream.next())
			}
			break
		}

		// Is the argument a long-form option or flag?
		if strings.HasPrefix(arg, "--") {
			parser.parseLongOption(arg[2:], stream)
			continue
		}

		// Is the argument a short-form option or flag?
		if strings.HasPrefix(arg, "-") {
			if arg == "-" || unicode.IsDigit([]rune(arg)[1]) {
				parser.arguments = append(parser.arguments, arg)
			} else {
				parser.parseShortOption(arg[1:], stream)
			}
			continue
		}

		// Is the argument a registered command?
		if isFirstArg {
			if cmdParser, found := parser.commands[arg]; found {
				parser.command = arg
				cmdParser.parseStream(stream)
				if cmdParser.Callback != nil {
					cmdParser.Callback(arg, cmdParser)
				}
				break
			}
		}

		// Is the argument the automatic 'help' command?
		if isFirstArg && parser.enableHelpCommand && arg == "help" {
			if stream.hasNext() {
				name := stream.next()
				if cmdParser, ok := parser.commands[name]; ok {
					cmdParser.ExitWithHelptext()
				} else {
					exit(fmt.Sprintf("'%v' is not a recognised command name", name))
				}
			} else {
				exit("the help command requires an argument")
			}
		}

		// If we get here, we have a positional argument.
		parser.arguments = append(parser.arguments, arg)
		isFirstArg = false
	}
}

// ParseArgs parses a slice of string arguments.
func (parser *ArgParser) ParseArgs(args []string) {
	parser.parseStream(newArgStream(args))
}

// Parse parses the application's command line arguments.
func (parser *ArgParser) Parse() {
	parser.ParseArgs(os.Args[1:])
}

// Parse a long-form option, i.e. an option beginning with a double dash.
func (parser *ArgParser) parseLongOption(arg string, stream *argstream) {
	// Do we have an option of the form --name=value?
	if strings.Contains(arg, "=") {
		parser.parseEqualsOption("--", arg)
		return
	}

	// Is the argument a registered option name?
	if opt, found := parser.options[arg]; found {
		opt.count += 1
		if opt.kind == "string" || opt.kind == "int" || opt.kind == "float" {
			if stream.hasNext() {
				opt.trySetValue(stream.next())
			} else {
				exit(fmt.Sprintf("missing argument for --%v", arg))
			}
		}
		return
	}

	// Is the argument an automatic --help flag?
	if arg == "help" && parser.Helptext != "" {
		parser.ExitWithHelptext()
	}

	// Is the argument an automatic --version flag?
	if arg == "version" && parser.Version != "" {
		parser.ExitWithVersion()
	}

	// The argument is not a recognised option name.
	exit(fmt.Sprintf("--%v is not a recognised option name", arg))
}

// Parse a short-form option, i.e. an option beginning with a single dash.
func (parser *ArgParser) parseShortOption(arg string, stream *argstream) {
	// Do we have an option of the form -n=value?
	if strings.Contains(arg, "=") {
		parser.parseEqualsOption("-", arg)
		return
	}

	// We examine each character individually to support condensed options with trailing arguments,
	// e.g. -abc foo bar. If we don't recognise the character as a registered option name, we check
	// for an automatic -h or -v flag before exiting.
	for _, char := range arg {
		name := string(char)
		if opt, found := parser.options[name]; found {
			opt.count += 1
			if opt.kind == "string" || opt.kind == "int" || opt.kind == "float" {
				if stream.hasNext() {
					opt.trySetValue(stream.next())
				} else if len([]rune(arg)) > 1 {
					exit(fmt.Sprintf("missing argument for '%v' in -%v", name, arg))
				} else {
					exit(fmt.Sprintf("missing argument for -%v", arg))
				}
			}
		} else {
			if name == "h" && parser.Helptext != "" {
				parser.ExitWithHelptext()
			} else if name == "v" && parser.Version != "" {
				parser.ExitWithVersion()
			} else if len([]rune(arg)) > 1 {
				exit(fmt.Sprintf("'%v' in -%v is not a recognised option name", name, arg))
			} else {
				exit(fmt.Sprintf("-%v is not a recognised option name", name))
			}
		}
	}
}

// Parse an option of the form --name=value or -n=value.
func (parser *ArgParser) parseEqualsOption(prefix string, arg string) {
	split := strings.SplitN(arg, "=", 2)
	name := split[0]
	value := split[1]

	// Do we have the name of a registered option?
	opt, found := parser.options[name]
	if !found {
		exit(fmt.Sprintf("%s%s is not a recognised option name", prefix, name))
	}
	opt.count += 1

	// Boolean flags should never contain an equals sign.
	if opt.kind == "flag" {
		exit(fmt.Sprintf("invalid format for flag %s%s", prefix, name))
	}

	// Check that a value has been supplied.
	if value == "" {
		exit(fmt.Sprintf("missing value for %s%s", prefix, name))
	}

	// Try to parse the argument as a value of the appropriate type.
	opt.trySetValue(value)
}

// -------------------------------------------------------------------------
// ArgParser: utilities.
// -------------------------------------------------------------------------

// ExitWithHelptext prints the parser's help text, then exits.
func (parser *ArgParser) ExitWithHelptext() {
	fmt.Println(strings.TrimSpace(parser.Helptext))
	os.Exit(0)
}

// ExitWithVersion prints the parser's version string, then exits.
func (parser *ArgParser) ExitWithVersion() {
	fmt.Println(strings.TrimSpace(parser.Version))
	os.Exit(0)
}

// String returns a string representation of the parser instance.
func (parser *ArgParser) String() string {
	lines := make([]string, 0)

	lines = append(lines, "Options:")
	if len(parser.options) > 0 {
		names := make([]string, 0, len(parser.options))
		for name := range parser.options {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			var values string
			opt := parser.options[name]
			switch opt.kind {
			case "flag":
				values = fmt.Sprintf("%v", opt.count)
			case "string":
				values = fmt.Sprintf("(%v) %v", opt.stringFallback, opt.stringValues)
			case "int":
				values = fmt.Sprintf("(%v) %v", opt.intFallback, opt.intValues)
			case "float":
				values = fmt.Sprintf("(%v) %v", opt.floatFallback, opt.floatValues)
			}
			lines = append(lines, fmt.Sprintf("  %v: %v", name, values))
		}
	} else {
		lines = append(lines, "  [none]")
	}

	lines = append(lines, "\nArguments:")
	if len(parser.arguments) > 0 {
		for _, arg := range parser.arguments {
			lines = append(lines, fmt.Sprintf("  %v", arg))
		}
	} else {
		lines = append(lines, "  [none]")
	}

	lines = append(lines, "\nCommand:")
	if parser.HasCommand() {
		lines = append(lines, fmt.Sprintf("  %v", parser.CommandName()))
	} else {
		lines = append(lines, "  [none]")
	}

	return strings.Join(lines, "\n")
}
