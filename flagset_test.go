package flagerr_test

import (
	"testing"

	"github.com/samherrmann/flagerr"
)

func TestFlagSet_Parse(t *testing.T) {
	t.Run("returns *flagerr.HelpError when -help is invoked", func(t *testing.T) {
		fs := newFlagSet()
		err := fs.Parse([]string{"-help"})

		if !flagerr.IsHelpError(err) {
			t.Fatalf("want *flagerr.HelpError, got %T", err)
		}
	})

	t.Run("includes usage info in error", func(t *testing.T) {
		fs := newFlagSet()
		err := fs.Parse([]string{"-bar"})

		got := err.Error()
		want := "flag provided but not defined: -bar\n" +
			"Usage of my-app:\n" +
			"  -foo string\n" +
			"    	Some description of foo (default \"my-default-value\")\n"
		if got != want {
			t.Fatalf("\ngot:\n%q\nwant:\n%q\n", got, want)
		}
	})
}

func TestFlagSet_UsageError(t *testing.T) {
	fs := newFlagSet()
	if err := fs.Parse(nil); err != nil {
		t.Fatal(err)
	}

	err := fs.UsageError("oops")

	got := err.Error()
	want := "oops\n" +
		"Usage of my-app:\n" +
		"  -foo string\n" +
		"    	Some description of foo (default \"my-default-value\")\n"

	if got != want {
		t.Fatalf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}

func newFlagSet() *flagerr.FlagSet {
	flagSet := flagerr.NewFlagSet("my-app")
	flagSet.String("foo", "my-default-value", "Some description of foo")
	return flagSet
}
