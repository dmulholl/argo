package argo

import (
	"testing"
)

/* -------- */
/*  Flags.  */
/* -------- */

func TestFlagEmpty(t *testing.T) {
	parser := NewParser()
	parser.NewFlag("bool")
	parser.ParseArgs([]string{})
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
	parser.ParseArgs([]string{"foo", "bar"})
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
	parser.ParseArgs([]string{"--bool"})
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
	parser.ParseArgs([]string{"-b"})
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
	parser.ParseArgs([]string{"--bool", "--bool", "--bool"})
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
	parser.ParseArgs([]string{"-b", "-b", "-b"})
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
	parser.ParseArgs([]string{})
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
	parser.ParseArgs([]string{"foo", "bar"})
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
	parser.ParseArgs([]string{"--opt", "value"})
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
	parser.ParseArgs([]string{"-o", "value"})
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
	parser.ParseArgs([]string{"--opt", "a", "b", "--opt", "c"})
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
	parser.ParseArgs([]string{"-o", "a", "b", "-o", "c"})
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
	parser.ParseArgs([]string{})
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
	parser.ParseArgs([]string{"foo", "bar"})
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
	parser.ParseArgs([]string{"--opt", "202"})
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
	parser.ParseArgs([]string{"-o", "202"})
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
	parser.ParseArgs([]string{"--opt", "-202"})
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
	parser.ParseArgs([]string{"--opt", "1", "2", "--opt", "3"})
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
	parser.ParseArgs([]string{"-o", "1", "2", "-o", "3"})
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
	parser.ParseArgs([]string{})
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
	parser.ParseArgs([]string{"foo", "bar"})
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
	parser.ParseArgs([]string{"--opt", "2.2"})
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
	parser.ParseArgs([]string{"-o", "2.2"})
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
	parser.ParseArgs([]string{"--opt", "-2.2"})
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
	parser.ParseArgs([]string{"--opt", "1.0", "2.0", "--opt", "3.0"})
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
	parser.ParseArgs([]string{"-o", "1.0", "2.0", "-o", "3.0"})
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
	parser.ParseArgs([]string{})
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
	parser.ParseArgs([]string{
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
	parser.ParseArgs([]string{
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
	parser.ParseArgs([]string{"-bsif", "value", "202", "2.2"})
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
	parser.ParseArgs([]string{})
	if parser.HasArgs() != false {
		t.Fail()
	}
}

func TestPositionalArgs(t *testing.T) {
	parser := NewParser()
	parser.ParseArgs([]string{"foo", "bar"})
	if parser.HasArgs() != true {
		t.Fail()
	}
	if parser.CountArgs() != 2 {
		t.Fail()
	}
	if parser.Arg(0) != "foo" {
		t.Fail()
	}
	if parser.Arg(1) != "bar" {
		t.Fail()
	}
	if parser.Args()[0] != "foo" {
		t.Fail()
	}
	if parser.Args()[1] != "bar" {
		t.Fail()
	}
}

func TestPositionalArgsAsInts(t *testing.T) {
	parser := NewParser()
	parser.ParseArgs([]string{"1", "11"})
	if len(parser.ArgsAsInts()) != 2 {
		t.Fail()
	}
	if parser.ArgsAsInts()[0] != 1 {
		t.Fail()
	}
	if parser.ArgsAsInts()[1] != 11 {
		t.Fail()
	}
}

func TestPositionalArgsAsFloats(t *testing.T) {
	parser := NewParser()
	parser.ParseArgs([]string{"1.1", "11.1"})
	if len(parser.ArgsAsFloats()) != 2 {
		t.Fail()
	}
	if parser.ArgsAsFloats()[0] != 1.1 {
		t.Fail()
	}
	if parser.ArgsAsFloats()[1] != 11.1 {
		t.Fail()
	}
}

/* ----------- */
/*  Commands.  */
/* ----------- */

func TestCommandAbsent(t *testing.T) {
	parser := NewParser()
	parser.NewCommand("cmd", "helptext")
	parser.ParseArgs([]string{})
	if parser.HasCommand() != false {
		t.Fail()
	}
}

func TestCommandPresent(t *testing.T) {
	parser := NewParser()
	cmdParser := parser.NewCommand("cmd", "helptext")
	parser.ParseArgs([]string{"cmd"})
	if parser.HasCommand() != true {
		t.Fail()
	}
	if parser.CommandName() != "cmd" {
		t.Fail()
	}
	if parser.CommandParser() != cmdParser {
		t.Fail()
	}
}

func TestCommandWithOptions(t *testing.T) {
	parser := NewParser()
	cmdParser := parser.NewCommand("cmd", "helptext")
	cmdParser.NewFlag("bool")
	cmdParser.NewStringOption("string", "default")
	cmdParser.NewIntOption("int", 101)
	cmdParser.NewFloatOption("float", 1.1)
	parser.ParseArgs([]string{
		"cmd",
		"foo", "bar",
		"--string", "value",
		"--int", "202",
		"--float", "2.2",
	})
	if parser.HasCommand() != true {
		t.Fail()
	}
	if parser.CommandName() != "cmd" {
		t.Fail()
	}
	if parser.CommandParser() != cmdParser {
		t.Fail()
	}
	if cmdParser.HasArgs() != true {
		t.Fail()
	}
	if cmdParser.CountArgs() != 2 {
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
