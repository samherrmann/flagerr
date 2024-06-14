package flagerr_test

import (
	"fmt"
	"os"

	"github.com/samherrmann/flagerr"
)

func Example() {
	if err := app(); err != nil {
		logError(err)
	}
}

func app() error {
	flagSet := flagerr.NewFlagSet(os.Args[0])

	// Add flags to flagSet as needed.

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	// Run the rest of the app.
	return nil
}

func logError(err error) {
	// If the -help or -h flag was invoked, then print the error to standard out
	// and exit normally (i.e. with code 0).
	if flagerr.IsHelpError(err) {
		fmt.Println(err)
		return
	}
	// Print all other errors to os.Stdout. and exit with code 1.
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
