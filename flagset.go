package flagerr

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
)

// NewFlagSet returns a *FlagSet that does not call os.Exit and does not print
// any messages to os.Stdout or os.Stderr.
func NewFlagSet(name string) *FlagSet {
	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)

	// Output for usage and error messages.
	output := bytes.NewBufferString("")
	flagSet.SetOutput(output)

	return &FlagSet{flagSet: flagSet, output: output}
}

// The purpose of this alias is to embed flag.FlagSet in FlagSet while keeping
// it private.
type flagSet = flag.FlagSet

// FlagSet represents a set of defined flags.
type FlagSet struct {
	*flagSet
	output *bytes.Buffer
}

// UsageError returns an error that includes the given message and the usage
// information.
func (fs *FlagSet) UsageError(msg string) error {
	fs.Usage()
	return fmt.Errorf("%s\n%s", msg, fs.output.String())
}

// Parse parses flag definitions from the argument list, which should not
// include the command name (e.g. os.Args[1:]). If the -help or -h flag is
// invoked, then *HelpError is returned which includes the usage information.
// This method does not call os.Exit and does not print any messages to
// os.Stdout or os.Stderr.
func (fs *FlagSet) Parse(args []string) error {
	if err := fs.flagSet.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return &HelpError{usage: fs.output.String()}
		}
		return errors.New(fs.output.String())
	}
	return nil
}
