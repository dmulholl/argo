# Argo

A minimalist Go library for parsing command line arguments.


### Features

* Long-form boolean flags with single-character shortcuts: `--flag`, `-f`.

* Long-form string, integer, and floating-point options with
  single-character shortcuts: `--option <arg>`, `-o <arg>`.

* Condensed short-form options: `-abc <arg> <arg>`.

* Automatic `--help` and `--version` flags.

* Support for multivalued options.

* Support for git-style command interfaces with arbitrarily-nested commands.


### Installation

Install the `argo` package:

::: code
    go get github.com/dmulholl/argo/v4

Import the `argo` package:

::: code go
    import "github.com/dmulholl/argo/v4"


### Links

* [Documentation](https://www.dmulholl.com/docs/argo/master/)
* [Simple Example](https://github.com/dmulholl/argo/blob/master/cmd/simple-example/main.go)
* [Command Example](https://github.com/dmulholl/argo/blob/master/cmd/command-example/main.go)
* [API Documentation](https://pkg.go.dev/github.com/dmulholl/argo)


### License

Zero-Clause BSD (0BSD).
