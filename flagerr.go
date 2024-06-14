// Package flagerr is a small wrapper around Go's [flag] package that uses
// opinionated error handling.
package flagerr

import (
	"errors"
)

// IsHelpError returns true if err is of type *HelpError.
func IsHelpError(err error) bool {
	var helpErr *HelpError
	return errors.As(err, &helpErr)
}

// HelpError is the error returned by FlagSet.Parse method if the -help or -h
// flag is invoked and no such flag is defined. The error includes the usage
// information.
type HelpError struct {
	usage string
}

// Error implements the error interface.
func (e *HelpError) Error() string {
	return e.usage
}
