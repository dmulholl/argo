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

/* --------- */
/*  Options  */
/* --------- */

// [kind] is one of "flag", "string", "int", or "float".
type option struct {
	kind           string
	count          int
	stringValues   []string
	intValues      []int
	floatValues    []float64
	stringFallback string
	intFallback    int
	floatFallback  float64
}

// This method attempts to set the option's value by parsing a string argument.
func (opt *option) trySetValue(arg string) error {
	switch opt.kind {
	case "string":
		opt.stringValues = append(opt.stringValues, arg)
		return nil

	case "int":
		value, err := strconv.ParseInt(arg, 0, 0)
		if err != nil {
			return fmt.Errorf("cannot parse '%s' as an integer", arg)
		}
		opt.intValues = append(opt.intValues, int(value))
		return nil

	case "float":
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return fmt.Errorf("cannot parse '%s' as a float", arg)
		}
		opt.floatValues = append(opt.floatValues, value)
		return nil

	default:
		panic(fmt.Sprintf("argo: invalid option type: %s", opt.kind))
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
	// The parser's helptext string.
	//
	// Specifying a helptext string for a parser activates an automatic --help flag that prints the
	// parser's helptext and exits. (Also activates an automatic -h shortcut unless registered by
	// another flag/option.)
	Helptext string

	// The parser's version string.
	//
	// Specifying a version string for a parser activates an automatic --version flag that prints the
	// parser's version and exits. (Also activates an automatic -v shortcut unless registered
	// by another flag/option.)
	Version string

	// The parser's callback function.
	//
	// This field is only valid for command subparsers. If the command subparser's registered
	// command is found by the parent parser, and if this field is not nil, the specified callback
	// function will be called automatically. It will be passed the command name and the command's
	// ArgParser instance as arguments.
	Callback func(string, *ArgParser) error

	// If true, enables an automatic 'help' command that prints helptext for subcommands.
	//
	// Defaults to false but gets toggled automatically to true whenever a command is registered.
	// Set this value to false to disable the automatic 'help' command.
	EnableHelpCommand bool

	// After parsing, stores the parser's positional arguments.
	Args []string

	// After parsing, if the parser has found a command, stores the command's name.
	FoundCommandName string

	// After parsing, if the parser has found a command, stores the command's ArgParser instance.
	FoundCommandParser *ArgParser

	// Stores option instances indexed by option name.
	options map[string]*option

	// Stores command parsers indexed by command name.
	commands map[string]*ArgParser
}

// NewParser initializes a new ArgParser instance.
func NewParser() *ArgParser {
	return &ArgParser{
		options:  make(map[string]*option),
		commands: make(map[string]*ArgParser),
		Args:     make([]string, 0),
	}
}

/* ------------------------------ */
/*  ArgParser: register options.  */
/* ------------------------------ */

// NewFlag registers a new flag, i.e. a valueless option that is either present (found) or absent
// (not found). You can check for the presence of a flag using the parser's Found() or Count()
// methods.
//
// The name parameter accepts an unlimited number of space-separated aliases and single-character
// shortcuts.
func (parser *ArgParser) NewFlag(name string) {
	opt := &option{}
	opt.kind = "flag"
	for _, alias := range strings.Split(name, " ") {
		parser.options[alias] = opt
	}
}

// NewStringOption registers a new string-valued option.
//
// The name parameter accepts an unlimited number of space-separated aliases and single-character
// shortcuts. The fallback parameter specifies the option's default value.
func (parser *ArgParser) NewStringOption(name string, fallback string) {
	opt := &option{}
	opt.kind = "string"
	opt.stringFallback = fallback
	for _, alias := range strings.Split(name, " ") {
		parser.options[alias] = opt
	}
}

// NewIntOption registers a new integer-valued option, i.e. the option's value will be parsed
// as an int.
//
// The name parameter accepts an unlimited number of space-separated aliases and single-character
// shortcuts. The fallback parameter specifies the option's default value.
func (parser *ArgParser) NewIntOption(name string, fallback int) {
	opt := &option{}
	opt.kind = "int"
	opt.intFallback = fallback
	for _, alias := range strings.Split(name, " ") {
		parser.options[alias] = opt
	}
}

// NewFloatOption registers a new float-valued option, i.e. the option's value will be parsed
// as a float64.
//
// The name parameter accepts an unlimited number of space-separated aliases and single-character
// shortcuts. The fallback parameter specifies the option's default value.
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
	panic(fmt.Sprintf("argo: '%s' is not a registered flag or option name", name))
}

// Count returns the number of times the specified flag or option was found.
// Any of the flag/option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) Count(name string) int {
	return parser.getOpt(name).count
}

// Found returns true if the specified flag or option was found.
// Any of the flag/option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) Found(name string) bool {
	return parser.getOpt(name).count > 0
}

// StringValue returns the value of the specified string-valued option.
// Any of the option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) StringValue(name string) string {
	opt := parser.getOpt(name)
	if len(opt.stringValues) > 0 {
		return opt.stringValues[len(opt.stringValues)-1]
	} else {
		return opt.stringFallback
	}
}

// IntValue returns the value of the specified integer-valued option.
// Any of the option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) IntValue(name string) int {
	opt := parser.getOpt(name)
	if len(opt.intValues) > 0 {
		return opt.intValues[len(opt.intValues)-1]
	} else {
		return opt.intFallback
	}
}

// FloatValue returns the value of the specified float-valued option.
// Any of the option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) FloatValue(name string) float64 {
	opt := parser.getOpt(name)
	if len(opt.floatValues) > 0 {
		return opt.floatValues[len(opt.floatValues)-1]
	} else {
		return opt.floatFallback
	}
}

// StringValues returns the specified string-valued option's list of values.
// Any of the option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) StringValues(name string) []string {
	return parser.getOpt(name).stringValues
}

// IntValues returns the specified integer-valued option's list of values.
// Any of the option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) IntValues(name string) []int {
	return parser.getOpt(name).intValues
}

// FloatValues returns the specified float-valued option's list of values.
// Any of the option's registered aliases or shortcuts can be used as the name parameter.
//
// Panics if name is not a registered flag or option name.
func (parser *ArgParser) FloatValues(name string) []float64 {
	return parser.getOpt(name).floatValues
}

/* ---------------------------------- */
/*  ArgParser: positional arguments.  */
/* ---------------------------------- */

// ArgsAsInts attempts to return the parser's positional arguments as a slice of integers.
// Returns an error if any of the arguments cannot be parsed as an integer.
func (parser *ArgParser) ArgsAsInts() ([]int, error) {
	values := make([]int, 0)
	for _, arg := range parser.Args {
		value, err := strconv.ParseInt(arg, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("cannot parse '%v' as an integer", arg)
		}
		values = append(values, int(value))
	}
	return values, nil
}

// ArgsAsFloats attempts to return the parser's positional arguments as a slice of floats.
// Returns an error if any of the arguments cannot be parsed as a float.
func (parser *ArgParser) ArgsAsFloats() ([]float64, error) {
	values := make([]float64, 0)
	for _, arg := range parser.Args {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse '%v' as a float", arg)
		}
		values = append(values, value)
	}
	return values, nil
}

/* ---------------------- */
/*  ArgParser: commands.  */
/* ---------------------- */

// NewCommand registers a new command. The name parameter accepts an unlimited number of space-
// separated aliases for the command. Returns the new command's ArgParser instance.
func (parser *ArgParser) NewCommand(name string) *ArgParser {
	parser.EnableHelpCommand = true
	cmdParser := NewParser()
	for _, alias := range strings.Split(name, " ") {
		parser.commands[alias] = cmdParser
	}
	return cmdParser
}

/* ----------------------------- */
/*  ArgParser: parse arguments.  */
/* ----------------------------- */

// Parses a stream of string arguments.
func (parser *ArgParser) parseStream(stream *argstream) error {
	for stream.hasNext() {
		arg := stream.next()

		// If we encounter a -- argument, turn off option-parsing.
		if arg == "--" {
			for stream.hasNext() {
				parser.Args = append(parser.Args, stream.next())
			}
			return nil
		}

		// Is the argument a long-form option or flag?
		if strings.HasPrefix(arg, "--") {
			if err := parser.parseLongOption(arg[2:], stream); err != nil {
				return err
			}
			continue
		}

		// Is the argument a short-form option or flag?
		if strings.HasPrefix(arg, "-") {
			if arg == "-" || unicode.IsDigit([]rune(arg)[1]) {
				parser.Args = append(parser.Args, arg)
			} else {
				if err := parser.parseShortOption(arg[1:], stream); err != nil {
					return err
				}
			}
			continue
		}

		// Is the argument a registered command?
		if len(parser.Args) == 0 {
			if cmdParser, found := parser.commands[arg]; found {
				parser.FoundCommandName = arg
				parser.FoundCommandParser = cmdParser

				if err := cmdParser.parseStream(stream); err != nil {
					return err
				}

				if cmdParser.Callback != nil {
					return cmdParser.Callback(arg, cmdParser)
				}

				break
			}
		}

		// Is the argument the automatic 'help' command?
		if len(parser.Args) == 0 && parser.EnableHelpCommand && arg == "help" {
			if stream.hasNext() {
				name := stream.next()
				if cmdParser, ok := parser.commands[name]; ok {
					cmdParser.exitWithHelptext()
				}
				return fmt.Errorf("help: '%v' is not a recognised command name", name)
			}
			return fmt.Errorf("help: missing argument for the help command")
		}

		// If we get here, we have a positional argument.
		parser.Args = append(parser.Args, arg)
	}

	return nil
}

// Parse parses a slice of string arguments. The arguments will be treated as if they came directly
// from os.Args, i.e. the first argument will be treated as the application's path and will be ignored.
func (parser *ArgParser) Parse(args []string) error {
	return parser.parseStream(newArgStream(args[1:]))
}

// ParseOsArgs parses the application's command line arguments.
// This is a shortcut for calling Parse(os.Args).
func (parser *ArgParser) ParseOsArgs() error {
	return parser.Parse(os.Args)
}

// Parse a long-form option, i.e. an option beginning with a double dash.
func (parser *ArgParser) parseLongOption(arg string, stream *argstream) error {
	// Do we have an option of the form --name=value?
	if strings.Contains(arg, "=") {
		return parser.parseEqualsOption("--", arg)
	}

	// Is the argument a registered flag or option name?
	if opt, found := parser.options[arg]; found {
		opt.count += 1
		if opt.kind == "flag" {
			return nil
		}
		if stream.hasNext() {
			return opt.trySetValue(stream.next())
		}
		return fmt.Errorf("missing argument for option --%v", arg)
	}

	// Is the argument an automatic --help flag?
	if arg == "help" && parser.Helptext != "" {
		parser.exitWithHelptext()
	}

	// Is the argument an automatic --version flag?
	if arg == "version" && parser.Version != "" {
		parser.exitWithVersion()
	}

	// The argument is not a recognised flag or option name.
	return fmt.Errorf("--%v is not a recognised flag or option name", arg)
}

// Parse a short-form option, i.e. an option beginning with a single dash.
func (parser *ArgParser) parseShortOption(arg string, stream *argstream) error {
	// Do we have an option of the form -n=value?
	if strings.Contains(arg, "=") {
		return parser.parseEqualsOption("-", arg)
	}

	// We examine each character individually to support condensed options with trailing arguments,
	// e.g. -abc foo bar. If we don't recognise the character as a registered flag or option name,
	// we check for an automatic -h or -v flag before returning an error.
	for _, char := range arg {
		name := string(char)

		if opt, found := parser.options[name]; found {
			opt.count += 1
			if opt.kind == "flag" {
				continue
			}
			if stream.hasNext() {
				if err := opt.trySetValue(stream.next()); err != nil {
					return err
				}
				continue
			}
			if len([]rune(arg)) > 1 {
				return fmt.Errorf("missing argument for option '%v' in -%v", name, arg)
			}
			return fmt.Errorf("missing argument for option -%v", arg)
		}

		if name == "h" && parser.Helptext != "" {
			parser.exitWithHelptext()
		}

		if name == "v" && parser.Version != "" {
			parser.exitWithVersion()
		}

		if len([]rune(arg)) > 1 {
			return fmt.Errorf("'%v' in -%v is not a recognised flag or option name", name, arg)
		}

		return fmt.Errorf("-%v is not a recognised flag or option name", name)
	}

	return nil
}

// Parse an option of the form --name=value or -n=value.
func (parser *ArgParser) parseEqualsOption(prefix string, arg string) error {
	if !strings.Contains(arg, "=") {
		panic(fmt.Sprintf("argo: invalid call to parseEqualsOption with prefix '%s' and arg '%s'", prefix, arg))
	}

	split := strings.SplitN(arg, "=", 2)
	name := split[0]
	value := split[1]

	// Do we have the name of a registered option?
	opt, found := parser.options[name]
	if !found {
		return fmt.Errorf("%s%s is not a recognised option name", prefix, name)
	}
	opt.count += 1

	// Boolean flags should never be followed by an equals sign.
	if opt.kind == "flag" {
		return fmt.Errorf("invalid value assignment for flag %s%s", prefix, name)
	}

	// Check that a value has been supplied.
	if value == "" {
		return fmt.Errorf("missing value for %s%s", prefix, name)
	}

	// Try to parse the argument as a value of the appropriate type.
	return opt.trySetValue(value)
}

// -------------------------------------------------------------------------
// ArgParser: utilities.
// -------------------------------------------------------------------------

// exitWithHelptext prints the parser's help text, then exits.
func (parser *ArgParser) exitWithHelptext() {
	fmt.Println(strings.TrimSpace(parser.Helptext))
	os.Exit(0)
}

// exitWithVersion prints the parser's version string, then exits.
func (parser *ArgParser) exitWithVersion() {
	fmt.Println(strings.TrimSpace(parser.Version))
	os.Exit(0)
}

// String returns a string representation of the parser instance for debugging.
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
			lines = append(lines, fmt.Sprintf("  %s [%s]: %s", name, opt.kind, values))
		}
	} else {
		lines = append(lines, "  [none]")
	}

	lines = append(lines, "\nArguments:")
	if len(parser.Args) > 0 {
		for _, arg := range parser.Args {
			lines = append(lines, fmt.Sprintf("  %v", arg))
		}
	} else {
		lines = append(lines, "  [none]")
	}

	lines = append(lines, "\nCommand:")
	if parser.FoundCommandName != "" {
		lines = append(lines, fmt.Sprintf("  %v", parser.FoundCommandName))
	} else {
		lines = append(lines, "  [none]")
	}

	return strings.Join(lines, "\n")
}
