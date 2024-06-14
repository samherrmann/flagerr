# flagerr

A small wrapper around Go's [flag.FlagSet](https://pkg.go.dev/flag) that does
not call `os.Exit` and does not print any messages to `os.Stdout` or
`os.Stderr`.

The behavior of the default `flag.FlagSet` is to immediately exit the
application (by calling `os.Exit`) when the user invokes the help flag or when a
parsing error is encountered. Additionally, the default `flag.FlagSet` prints
errors directly to `os.Stderr` before exiting. It's assumed that the default
`flag.FlagSet` was configured this way for simplicity, but this is usually not
the recommended practice for handling errors in Go. The recommended practice is
for functions/methods to return errors and let the caller bubble the error up
the call stack. A common logging facility may then be used at the top of the
stack to collect all errors that the application may encounter.
`flagerr.FlagSet` returned by `flagerr.NewFlagSet` provides the ability to
handle errors using the recommended Go practices. See the
[example](example_test.go) for more details.
