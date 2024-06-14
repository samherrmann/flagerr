package flagerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/samherrmann/flagerr"
)

func TestIsHelpError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "HelpError",
			err:  &flagerr.HelpError{},
			want: true,
		},
		{
			name: "wrapped HelpError",
			err:  fmt.Errorf("error: %w", &flagerr.HelpError{}),
			want: true,
		},
		{
			name: "generic error",
			err:  errors.New("oops"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flagerr.IsHelpError(tt.err); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
