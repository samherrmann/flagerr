package flagerr_test

import (
	"testing"

	"github.com/samherrmann/flagerr"
)

func TestFlagSet_Parse(t *testing.T) {
	t.Run("returns *flagerr.HelpError when -help is invoked", func(t *testing.T) {
		err := parseFlagSet("-help")

		if !flagerr.IsHelpError(err) {
			t.Fatalf("want *flagerr.HelpError, got %T", err)
		}
	})

	t.Run("includes usage info in error", func(t *testing.T) {
		err := parseFlagSet("-bar")

		got := err.Error()
		want := "flag provided but not defined: -bar\n" +
			"Usage of my-app:\n" +
			"  -foo string\n" +
			"    	Some description of foo (default \"my-defalt-value\")\n"
		if got != want {
			t.Fatalf("got\n%v\nwant\n%v\n", got, want)
		}
	})
}

func parseFlagSet(args ...string) error {
	flagSet := flagerr.NewFlagSet("my-app")
	flagSet.String("foo", "my-defalt-value", "Some description of foo")
	return flagSet.Parse(args)
}
