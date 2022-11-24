package count_test

import (
	"bytes"
	"count"
	"errors"
	"io"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		exp   int
	}{
		{name: "empty", input: "", exp: 0},
		{name: "single line", input: "AnDrEWW!", exp: 1},
		{name: "multi line", input: "This has a line\nbreak", exp: 2},
		{name: "repeated line breaks", input: "This \nhas a line\n\nbreak\n", exp: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := count.NewCounter(
				count.WithInput(bytes.NewBufferString(tt.input)),
			)

			if err != nil {
				t.Fatal(err)
			}
			got := c.Lines()
			if got != tt.exp {
				t.Errorf("want %d, got %d", tt.exp, got)
			}
		})

	}
}

func TestArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     []string
		exp      int
		expError error
	}{
		{name: "existing file", exp: 5, args: []string{"testdata/five.txt"}},
		{name: "non-existing file", expError: errors.New("no such file"), args: []string{"testdata/doesnotexist.txt"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := count.NewCounter(
				count.FromArgs(tt.args),
			)
			if tt.expError != nil {
				if err == nil {
					t.Errorf("expected error but did not get one")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
			got := c.Lines()

			if got != tt.exp {
				t.Errorf("got %d, want %d", got, tt.exp)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     []string
		exp      int
		expError error
	}{
		{name: "flag on", exp: 5, args: []string{"-w", "testdata/five.txt"}},
		{name: "flag off", exp: 5, args: []string{"testdata/five.txt"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := count.NewCounter(
				count.FromArgs(tt.args),
			)
			if tt.expError != nil {
				if err == nil {
					t.Errorf("expected error but did not get one")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
			got := c.Lines()

			if got != tt.exp {
				t.Errorf("got %d, want %d", got, tt.exp)
			}
		})
	}
}

func TestFromArgsBogusFlag(t *testing.T) {
	t.Parallel()

	args := []string{"-bogus"}
	_, err := count.NewCounter(
		count.WithOutput(io.Discard),
		count.FromArgs(args),
	)

	if err == nil {
		t.Fatal("want error, got nil")
	}
}
