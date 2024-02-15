package argo

import "testing"

/* -------- */
/*  Flags.  */
/* -------- */

func TestFlagEmpty(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool")
	parser.Parse([]string{"ignored"})
	if parser.Found("bool") != false {
		t.Fail()
	}
	if parser.Count("bool") != 0 {
		t.Fail()
	}
}

func TestFlagMissing(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool")
	parser.Parse([]string{"ignored", "foo", "bar"})
	if parser.Found("bool") != false {
		t.Fail()
	}
	if parser.Count("bool") != 0 {
		t.Fail()
	}
}

func TestFlagLongform(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool")
	parser.Parse([]string{"ignored", "--bool"})
	if parser.Found("bool") != true {
		t.Fail()
	}
	if parser.Count("bool") != 1 {
		t.Fail()
	}
}

func TestFlagShortform(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool b")
	parser.Parse([]string{"ignored", "-b"})
	if parser.Found("bool") != true {
		t.Fail()
	}
	if parser.Count("bool") != 1 {
		t.Fail()
	}
}

func TestBoolMultiLongform(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool")
	parser.Parse([]string{"ignored", "--bool", "--bool", "--bool"})
	if parser.Found("bool") != true {
		t.Fail()
	}
	if parser.Count("bool") != 3 {
		t.Fail()
	}
}

func TestBoolMultiShortform(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool b")
	parser.Parse([]string{"ignored", "-b", "-b", "-b"})
	if parser.Found("bool") != true {
		t.Fail()
	}
	if parser.Count("bool") != 3 {
		t.Fail()
	}
}

/* ----------------- */
/*  String options.  */
/* ----------------- */

func TestStringOptionEmpty(t *testing.T) {
	parser := NewParser()
	parser.NewStringOption("opt", "default")
	parser.Parse([]string{"ignored"})
	if parser.StringValue("opt") != "default" {
		t.Fail()
	}
	if parser.Found("opt") != false {
		t.Fail()
	}
	if parser.Count("opt") != 0 {
		t.Fail()
	}
}

func TestStringOptionMissing(t *testing.T) {
	parser := NewParser()
	parser.NewStringOption("opt", "default")
	parser.Parse([]string{"ignored", "foo", "bar"})
	if parser.StringValue("opt") != "default" {
		t.Fail()
	}
	if parser.Found("opt") != false {
		t.Fail()
	}
	if parser.Count("opt") != 0 {
		t.Fail()
	}
}

func TestStringOptionLongform(t *testing.T) {
	parser := NewParser()
	parser.NewStringOption("opt", "default")
	parser.Parse([]string{"ignored", "--opt", "value"})
	if parser.StringValue("opt") != "value" {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestStringOptionShortform(t *testing.T) {
	parser := NewParser()
	parser.NewStringOption("opt o", "default")
	parser.Parse([]string{"ignored", "-o", "value"})
	if parser.StringValue("opt") != "value" {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestStringListLongform(t *testing.T) {
	parser := NewParser()
	parser.NewStringOption("opt", "default")
	parser.Parse([]string{"ignored", "--opt", "a", "b", "--opt", "c"})
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 2 {
		t.Fail()
	}
	if len(parser.StringValues("opt")) != 2 {
		t.Fail()
	}
	if parser.StringValues("opt")[0] != "a" {
		t.Fail()
	}
	if parser.StringValues("opt")[1] != "c" {
		t.Fail()
	}
}

func TestStringListShortform(t *testing.T) {
	parser := NewParser()
	parser.NewStringOption("opt o", "default")
	parser.Parse([]string{"ignored", "-o", "a", "b", "-o", "c"})
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 2 {
		t.Fail()
	}
	if len(parser.StringValues("opt")) != 2 {
		t.Fail()
	}
	if parser.StringValues("opt")[0] != "a" {
		t.Fail()
	}
	if parser.StringValues("opt")[1] != "c" {
		t.Fail()
	}
}

/* ------------------ */
/*  Integer options.  */
/* ------------------ */

func TestIntOptionEmpty(t *testing.T) {
	parser := NewParser()
	parser.NewIntOption("opt", 101)
	parser.Parse([]string{"ignored"})
	if parser.IntValue("opt") != 101 {
		t.Fail()
	}
	if parser.Found("opt") != false {
		t.Fail()
	}
	if parser.Count("opt") != 0 {
		t.Fail()
	}
}

func TestIntOptionMissing(t *testing.T) {
	parser := NewParser()
	parser.NewIntOption("opt", 101)
	parser.Parse([]string{"ignored", "foo", "bar"})
	if parser.IntValue("opt") != 101 {
		t.Fail()
	}
	if parser.Found("opt") != false {
		t.Fail()
	}
	if parser.Count("opt") != 0 {
		t.Fail()
	}
}

func TestIntOptionLongform(t *testing.T) {
	parser := NewParser()
	parser.NewIntOption("opt", 101)
	parser.Parse([]string{"ignored", "--opt", "202"})
	if parser.IntValue("opt") != 202 {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestIntOptionShortform(t *testing.T) {
	parser := NewParser()
	parser.NewIntOption("opt o", 101)
	parser.Parse([]string{"ignored", "-o", "202"})
	if parser.IntValue("opt") != 202 {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestIntOptionNegative(t *testing.T) {
	parser := NewParser()
	parser.NewIntOption("opt", 101)
	parser.Parse([]string{"ignored", "--opt", "-202"})
	if parser.IntValue("opt") != -202 {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestIntListLongform(t *testing.T) {
	parser := NewParser()
	parser.NewIntOption("opt", 101)
	parser.Parse([]string{"ignored", "--opt", "1", "2", "--opt", "3"})
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 2 {
		t.Fail()
	}
	if len(parser.IntValues("opt")) != 2 {
		t.Fail()
	}
	if parser.IntValues("opt")[0] != 1 {
		t.Fail()
	}
	if parser.IntValues("opt")[1] != 3 {
		t.Fail()
	}
}

func TestIntListShortform(t *testing.T) {
	parser := NewParser()
	parser.NewIntOption("opt o", 101)
	parser.Parse([]string{"ignored", "-o", "1", "2", "-o", "3"})
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 2 {
		t.Fail()
	}
	if len(parser.IntValues("opt")) != 2 {
		t.Fail()
	}
	if parser.IntValues("opt")[0] != 1 {
		t.Fail()
	}
	if parser.IntValues("opt")[1] != 3 {
		t.Fail()
	}
}

/* ---------------- */
/*  Float options.  */
/* ---------------- */

func TestFloatOptionEmpty(t *testing.T) {
	parser := NewParser()
	parser.NewFloatOption("opt", 1.1)
	parser.Parse([]string{"ignored"})
	if parser.FloatValue("opt") != 1.1 {
		t.Fail()
	}
	if parser.Found("opt") != false {
		t.Fail()
	}
	if parser.Count("opt") != 0 {
		t.Fail()
	}
}

func TestFloatOptionMissing(t *testing.T) {
	parser := NewParser()
	parser.NewFloatOption("opt", 1.1)
	parser.Parse([]string{"ignored", "foo", "bar"})
	if parser.FloatValue("opt") != 1.1 {
		t.Fail()
	}
	if parser.Found("opt") != false {
		t.Fail()
	}
	if parser.Count("opt") != 0 {
		t.Fail()
	}
}

func TestFloatOptionLongform(t *testing.T) {
	parser := NewParser()
	parser.NewFloatOption("opt", 1.1)
	parser.Parse([]string{"ignored", "--opt", "2.2"})
	if parser.FloatValue("opt") != 2.2 {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestFloatOptionShortform(t *testing.T) {
	parser := NewParser()
	parser.NewFloatOption("opt o", 1.1)
	parser.Parse([]string{"ignored", "-o", "2.2"})
	if parser.FloatValue("opt") != 2.2 {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestFloatOptionNegative(t *testing.T) {
	parser := NewParser()
	parser.NewFloatOption("opt", 1.1)
	parser.Parse([]string{"ignored", "--opt", "-2.2"})
	if parser.FloatValue("opt") != -2.2 {
		t.Fail()
	}
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 1 {
		t.Fail()
	}
}

func TestFloatListLongform(t *testing.T) {
	parser := NewParser()
	parser.NewFloatOption("opt", 0.0)
	parser.Parse([]string{"ignored", "--opt", "1.0", "2.0", "--opt", "3.0"})
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 2 {
		t.Fail()
	}
	if parser.FloatValues("opt")[0] != 1.0 {
		t.Fail()
	}
	if parser.FloatValues("opt")[1] != 3.0 {
		t.Fail()
	}
}

func TestFloatListShortform(t *testing.T) {
	parser := NewParser()
	parser.NewFloatOption("opt o", 0.0)
	parser.Parse([]string{"ignored", "-o", "1.0", "2.0", "-o", "3.0"})
	if parser.Found("opt") != true {
		t.Fail()
	}
	if parser.Count("opt") != 2 {
		t.Fail()
	}
	if parser.FloatValues("opt")[0] != 1.0 {
		t.Fail()
	}
	if parser.FloatValues("opt")[1] != 3.0 {
		t.Fail()
	}
}

/* -------------------------------- */
/*  Multiple option types at once.  */
/* -------------------------------- */

func TestMultiOptionsEmpty(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool1")
	parser.NewFlag("bool2 b")
	parser.NewStringOption("string1", "default1")
	parser.NewStringOption("string2 s", "default2")
	parser.NewIntOption("int1", 101)
	parser.NewIntOption("int2 i", 202)
	parser.NewFloatOption("float1", 1.1)
	parser.NewFloatOption("float2 f", 2.2)
	parser.Parse([]string{"ignored"})
	if parser.Found("bool1") != false {
		t.Fail()
	}
	if parser.Found("bool2") != false {
		t.Fail()
	}
	if parser.StringValue("string1") != "default1" {
		t.Fail()
	}
	if parser.StringValue("string2") != "default2" {
		t.Fail()
	}
	if parser.IntValue("int1") != 101 {
		t.Fail()
	}
	if parser.IntValue("int2") != 202 {
		t.Fail()
	}
	if parser.FloatValue("float1") != 1.1 {
		t.Fail()
	}
	if parser.FloatValue("float2") != 2.2 {
		t.Fail()
	}
}

func TestMultiOptionsLongform(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool1")
	parser.NewFlag("bool2 b")
	parser.NewStringOption("string1", "default1")
	parser.NewStringOption("string2 s", "default2")
	parser.NewIntOption("int1", 101)
	parser.NewIntOption("int2 i", 202)
	parser.NewFloatOption("float1", 1.1)
	parser.NewFloatOption("float2 f", 2.2)
	parser.Parse([]string{
		"ignored",
		"--bool1",
		"--bool2",
		"--string1", "value1",
		"--string2", "value2",
		"--int1", "303",
		"--int2", "404",
		"--float1", "3.3",
		"--float2", "4.4",
	})
	if parser.Found("bool1") != true {
		t.Fail()
	}
	if parser.Found("bool2") != true {
		t.Fail()
	}
	if parser.StringValue("string1") != "value1" {
		t.Fail()
	}
	if parser.StringValue("string2") != "value2" {
		t.Fail()
	}
	if parser.IntValue("int1") != 303 {
		t.Fail()
	}
	if parser.IntValue("int2") != 404 {
		t.Fail()
	}
	if parser.FloatValue("float1") != 3.3 {
		t.Fail()
	}
	if parser.FloatValue("float2") != 4.4 {
		t.Fail()
	}
}

func TestMultiOptionsShortform(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool1")
	parser.NewFlag("bool2 b")
	parser.NewStringOption("string1", "default1")
	parser.NewStringOption("string2 s", "default2")
	parser.NewIntOption("int1", 101)
	parser.NewIntOption("int2 i", 202)
	parser.NewFloatOption("float1", 1.1)
	parser.NewFloatOption("float2 f", 2.2)
	parser.Parse([]string{
		"ignored",
		"--bool1",
		"-b",
		"--string1", "value1",
		"-s", "value2",
		"--int1", "303",
		"-i", "404",
		"--float1", "3.3",
		"-f", "4.4",
	})
	if parser.Found("bool1") != true {
		t.Fail()
	}
	if parser.Found("bool2") != true {
		t.Fail()
	}
	if parser.StringValue("string1") != "value1" {
		t.Fail()
	}
	if parser.StringValue("string2") != "value2" {
		t.Fail()
	}
	if parser.IntValue("int1") != 303 {
		t.Fail()
	}
	if parser.IntValue("int2") != 404 {
		t.Fail()
	}
	if parser.FloatValue("float1") != 3.3 {
		t.Fail()
	}
	if parser.FloatValue("float2") != 4.4 {
		t.Fail()
	}
}

/* ------------------------------- */
/*  Condensed short-form options.  */
/* ------------------------------- */

func TestCondensedOptions(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool b")
	parser.NewStringOption("string s", "default")
	parser.NewIntOption("int i", 101)
	parser.NewFloatOption("float f", 1.1)
	parser.Parse([]string{"ignored", "-bsif", "value", "202", "2.2"})
	if parser.Found("bool") != true {
		t.Fail()
	}
	if parser.StringValue("string") != "value" {
		t.Fail()
	}
	if parser.IntValue("int") != 202 {
		t.Fail()
	}
	if parser.FloatValue("float") != 2.2 {
		t.Fail()
	}
}

/* ----------------------- */
/*  Positional arguments.  */
/* ----------------------- */

func TestPositionalArgsEmpty(t *testing.T) {
	parser := NewParser()
	parser.Parse([]string{"ignored"})
	if len(parser.Args) != 0 {
		t.Fail()
	}
}

func TestPositionalArgs(t *testing.T) {
	parser := NewParser()
	parser.Parse([]string{"ignored", "foo", "bar"})
	if len(parser.Args) != 2 {
		t.Fail()
	}
	if parser.Args[0] != "foo" {
		t.Fail()
	}
	if parser.Args[1] != "bar" {
		t.Fail()
	}
}

func TestPositionalArgsAsInts(t *testing.T) {
	parser := NewParser()
	parser.Parse([]string{"ignored", "123", "456"})
	argsAsInts, err := parser.ArgsAsInts()
	if err != nil {
		t.Fail()
	}
	if len(argsAsInts) != 2 {
		t.Fail()
	}
	if argsAsInts[0] != 123 {
		t.Fail()
	}
	if argsAsInts[1] != 456 {
		t.Fail()
	}
}

func TestPositionalArgsAsFloats(t *testing.T) {
	parser := NewParser()
	parser.Parse([]string{"ignored", "1.0", "123.456"})
	argsAsFloats, err := parser.ArgsAsFloats()
	if err != nil {
		t.Fail()
	}
	if len(argsAsFloats) != 2 {
		t.Fail()
	}
	if argsAsFloats[0] != 1.0 {
		t.Fail()
	}
	if argsAsFloats[1] != 123.456 {
		t.Fail()
	}
}

/* ----------- */
/*  Commands.  */
/* ----------- */

func TestCommandAbsent(t *testing.T) {
	parser := NewParser()
	parser.NewCommand("cmd")
	parser.Parse([]string{"ignored"})
	if parser.FoundCommandName != "" {
		t.Fail()
	}
	if parser.FoundCommandParser != nil {
		t.Fail()
	}
}

func TestCommandPresent(t *testing.T) {
	parser := NewParser()
	cmdParser := parser.NewCommand("cmd")
	parser.Parse([]string{"ignored", "cmd"})
	if parser.FoundCommandName != "cmd" {
		t.Fail()
	}
	if parser.FoundCommandParser != cmdParser {
		t.Fail()
	}
}

func TestCommandWithOptions(t *testing.T) {
	parser := NewParser()
	cmdParser := parser.NewCommand("cmd")
	cmdParser.NewFlag("bool")
	cmdParser.NewStringOption("string", "default")
	cmdParser.NewIntOption("int", 101)
	cmdParser.NewFloatOption("float", 1.1)
	parser.Parse([]string{
		"ignored",
		"cmd",
		"foo", "bar",
		"--string", "value",
		"--int", "202",
		"--float", "2.2",
	})
	if parser.FoundCommandName != "cmd" {
		t.Fail()
	}
	if parser.FoundCommandParser != cmdParser {
		t.Fail()
	}
	if len(cmdParser.Args) != 2 {
		t.Fail()
	}
	if cmdParser.StringValue("string") != "value" {
		t.Fail()
	}
	if cmdParser.IntValue("int") != 202 {
		t.Fail()
	}
	if cmdParser.FloatValue("float") != 2.2 {
		t.Fail()
	}
}
